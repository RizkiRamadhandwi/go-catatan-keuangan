package repository

import (
	"database/sql"
	"log"
	"math"

	"enigmacamp.com/livecode-catatan-keuangan/config"
	"enigmacamp.com/livecode-catatan-keuangan/entity"
	"enigmacamp.com/livecode-catatan-keuangan/shared/model"
)

type ExpensesRepository interface {
	List(page, size int, startDate string, endDate string) ([]entity.Expenses, model.Paging, error)
	Create(payload entity.Expenses) (entity.Expenses, error)
	GetById(id string) (entity.Expenses, error)
	GetByType(trx string) ([]entity.Expenses, error)
	Update(payload entity.Expenses) (entity.Expenses, error)
	Delete(id string) error
}

type expensesRepository struct {
	db *sql.DB
}

// Create implements ExpensesRepository.
func (e *expensesRepository) Create(payload entity.Expenses) (entity.Expenses, error) {
	var lastBalance int
	err := e.db.QueryRow(config.GetBalance).Scan(&lastBalance)
	if err != nil {
		panic(err)
	}

	if lastBalance != 0 && lastBalance >= 0 {
		if payload.TransactionType == "CREDIT" {
			payload.Balance = lastBalance + payload.Amount
		} else if payload.TransactionType == "DEBIT" {
			if lastBalance >= payload.Amount {
				payload.Balance = lastBalance - payload.Amount
			} else {
				log.Println("expensesRepository.QueryRow:", err.Error())
				return entity.Expenses{}, err
			}
		} else {
			log.Println("expensesRepository.QueryRow:", err.Error())
			return entity.Expenses{}, err
		}
	}

	var expense entity.Expenses
	err = e.db.QueryRow(config.InsertExpenses,
		payload.Date, payload.Amount, payload.TransactionType, payload.Balance, payload.Description, payload.UpdatedAt).Scan(
		&expense.ID, &expense.CreatedAt)
	if err != nil {
		log.Println("expensesRepository.QueryRow:", err.Error())
		return entity.Expenses{}, err
	}

	expense.Date = payload.Date
	expense.Amount = payload.Amount
	expense.TransactionType = payload.TransactionType
	expense.Balance = payload.Balance
	expense.Description = payload.Description
	expense.UpdatedAt = payload.UpdatedAt
	return expense, nil

}

// List implements ExpensesRepository.
func (e *expensesRepository) List(page int, size int, startDate string, endDate string) ([]entity.Expenses, model.Paging, error) {
	var expenses []entity.Expenses
	offset := (page - 1) * size
	rows, err := e.db.Query(config.SelectExpenses, size, offset)
	if err != nil {
		log.Println("expensesRepository.Query:", err.Error())
		return nil, model.Paging{}, err
	}

	if startDate != "" && endDate != "" {
		rows, _ = e.db.Query(config.FindDate, startDate, endDate)
	}

	for rows.Next() {
		var expense entity.Expenses
		err := rows.Scan(&expense.ID, &expense.Date, &expense.Amount, &expense.TransactionType, &expense.Balance, &expense.Description, &expense.CreatedAt, &expense.UpdatedAt)

		if err != nil {
			log.Println("expensesRepository.Rows.Next():", err.Error())
			return nil, model.Paging{}, err
		}

		expenses = append(expenses, expense)

	}

	totalRows := 0
	if err := e.db.QueryRow(config.SelectCount).Scan(&totalRows); err != nil {
		return nil, model.Paging{}, err
	}

	paging := model.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   totalRows,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(size))),
	}

	return expenses, paging, nil

}

// GetById implements ExpensesRepository.
func (e *expensesRepository) GetById(id string) (entity.Expenses, error) {
	var expense entity.Expenses

	err := e.db.QueryRow(config.SelectExpenseById, id).Scan(&expense.ID, &expense.Date, &expense.Amount, &expense.TransactionType, &expense.Balance, &expense.Description, &expense.CreatedAt, &expense.UpdatedAt)

	if err != nil {
		log.Println("expensesRepository.Get.QueryRow:", err.Error())
		return entity.Expenses{}, err
	}
	return expense, nil
}

// GetByType implements ExpensesRepository.
func (e *expensesRepository) GetByType(trx string) ([]entity.Expenses, error) {
	var expenses []entity.Expenses
	rows, err := e.db.Query(config.SelectExpenseByType, trx)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var expense entity.Expenses
		err := rows.Scan(&expense.ID, &expense.Date, &expense.Amount, &expense.TransactionType, &expense.Balance, &expense.Description, &expense.CreatedAt, &expense.UpdatedAt)

		if err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}
	return expenses, nil
}

// Update implements ExpensesRepository.
func (e *expensesRepository) Update(payload entity.Expenses) (entity.Expenses, error) {
	panic("unimplemented")
}

// Delete implements ExpensesRepository.
func (e *expensesRepository) Delete(id string) error {
	_, err := e.db.Exec(config.DeleteExpense, id)

	if err != nil {
		log.Println("expensesRepository.DeleteById.Exec:", err.Error())
		return err
	}
	return nil
}

func NewExpensesRepository(db *sql.DB) ExpensesRepository {
	return &expensesRepository{db: db}
}
