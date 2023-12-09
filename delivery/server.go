package delivery

import (
	"database/sql"
	"fmt"

	"enigmacamp.com/livecode-catatan-keuangan/config"
	"enigmacamp.com/livecode-catatan-keuangan/delivery/controller"
	"enigmacamp.com/livecode-catatan-keuangan/repository"
	"enigmacamp.com/livecode-catatan-keuangan/usecase"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	expenseUC usecase.ExpensesUseCase
	engine    *gin.Engine
	host      string
}

func (s *Server) initRoute() {
	rg := s.engine.Group(config.ApiGroup)
	controller.NewExpensesController(s.expenseUC, rg).Route()
}

func (s *Server) Run() {
	s.initRoute()
	if err := s.engine.Run(s.host); err != nil {
		panic(fmt.Errorf("Server not Running on host %s, becauce error %v", s.host, err.Error()))
	}
}

func NewServer() *Server {
	cfg, _ := config.NewConfig()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
	db, err := sql.Open(cfg.Driver, dsn)
	if err != nil {
		panic("connection error")
	}

	expensesRepo := repository.NewExpensesRepository(db)
	expensesUC := usecase.NewExpensesUseCase(expensesRepo)
	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiHost)
	return &Server{
		expenseUC: expensesUC,
		engine:    engine,
		host:      host,
	}

}
