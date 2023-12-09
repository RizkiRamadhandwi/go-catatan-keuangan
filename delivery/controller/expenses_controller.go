package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"enigmacamp.com/livecode-catatan-keuangan/config"
	"enigmacamp.com/livecode-catatan-keuangan/entity"
	"enigmacamp.com/livecode-catatan-keuangan/shared/common"
	"enigmacamp.com/livecode-catatan-keuangan/usecase"
	"github.com/gin-gonic/gin"
)

type ExpensesController struct {
	expenseUC usecase.ExpensesUseCase
	rg        *gin.RouterGroup
}

func (e *ExpensesController) createHandler(ctx *gin.Context) {
	var payload entity.Expenses
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	expense, err := e.expenseUC.Register(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(ctx, expense, "Created")
}

func (e *ExpensesController) listHandler(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	size, _ := strconv.Atoi(ctx.Query("size"))

	startDate := ctx.Query("startDate")
	endDate := ctx.Query("endDate")

	expenses, paging, err := e.expenseUC.FindAll(page, size, startDate, endDate)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	var response []interface{}
	for _, expense := range expenses {
		response = append(response, expense)
	}
	common.SendPagedResponse(ctx, response, paging, "OK")
}

func (e *ExpensesController) getById(ctx *gin.Context) {
	id := ctx.Param("id")
	expense, err := e.expenseUC.FindById(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, "expense with ID "+id+" not found")
		return
	}
	common.SendSingleResponse(ctx, expense, "OK")
}

func (e *ExpensesController) getByType(ctx *gin.Context) {
	trx := ctx.Param("type")
	fmt.Println(trx, "kosong")
	expenses, err := e.expenseUC.FindByType(trx)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, "expense with type "+trx+" tidak ada")
		ctx.JSON(http.StatusNotFound, gin.H{"err": err.Error()})
		return
	}
	common.SendSingleResponse(ctx, expenses, "OK")
}

// func (e *ExpensesController) updateHandler(ctx *gin.Context) {
// 	id := ctx.Param("id")
// 	var payload entity.Expenses
// 	if err := ctx.ShouldBindJSON(&payload); err != nil {
// 		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	expense, err := e.expenseUC.Update(id, payload)
// 	if err != nil {
// 		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	common.SendSingleResponse(ctx, expense, "Updated")
// }

func (e *ExpensesController) deleteById(ctx *gin.Context) {
	id := ctx.Param("id")
	err := e.expenseUC.Delete(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, "expense with ID "+id+" not found")
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (e *ExpensesController) Route() {
	e.rg.POST(config.ExpensesPost, e.createHandler)
	e.rg.GET(config.ExpensesList, e.listHandler)
	e.rg.GET(config.ExpensesGetById, e.getById)
	e.rg.GET(config.ExpensesGetByType, e.getByType)
	e.rg.DELETE(config.ExpensesDelete, e.deleteById)
}

func NewExpensesController(expenseUC usecase.ExpensesUseCase, rg *gin.RouterGroup) *ExpensesController {
	return &ExpensesController{
		expenseUC: expenseUC,
		rg:        rg,
	}
}
