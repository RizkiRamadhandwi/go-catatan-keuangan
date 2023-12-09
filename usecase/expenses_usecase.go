package usecase

import (
	"fmt"
	"time"

	"enigmacamp.com/livecode-catatan-keuangan/entity"
	"enigmacamp.com/livecode-catatan-keuangan/repository"
	"enigmacamp.com/livecode-catatan-keuangan/shared/model"
)

type ExpensesUseCase interface {
	FindAll(page int, size int, startDate string, endDate string) ([]entity.Expenses, model.Paging, error)
	Register(payload entity.Expenses) (entity.Expenses, error)
	FindById(id string) (entity.Expenses, error)
	FindByType(trx string) ([]entity.Expenses, error)
	Update(payload entity.Expenses) (entity.Expenses, error)
	Delete(id string) error
}

type expensesUseCase struct {
	repo repository.ExpensesRepository
}

// FindAll implements ExpensesUseCase.
func (e *expensesUseCase) FindAll(page int, size int, startDate string, endDate string) ([]entity.Expenses, model.Paging, error) {
	return e.repo.List(page, size, startDate, endDate)
}

// FindById implements ExpensesUseCase.
func (e *expensesUseCase) FindById(id string) (entity.Expenses, error) {
	return e.repo.GetById(id)
}

// FindByType implements ExpensesUseCase.
func (e *expensesUseCase) FindByType(trx string) ([]entity.Expenses, error) {
	return e.repo.GetByType(trx)
}

// Register implements ExpensesUseCase.
func (e *expensesUseCase) Register(payload entity.Expenses) (entity.Expenses, error) {
	if payload.Amount <= 0 || payload.Description == "" {
		return entity.Expenses{}, fmt.Errorf("opps, required field")
	}

	payload.Date = time.Now()
	payload.UpdatedAt = time.Now()

	expense, err := e.repo.Create(payload)
	if err != nil {
		return entity.Expenses{}, fmt.Errorf("opps, failed to save data: %v", err.Error())
	}

	return expense, nil
}

// Update implements ExpensesUseCase.
func (e *expensesUseCase) Update(payload entity.Expenses) (entity.Expenses, error) {
	panic("unimplemented")
}

// Delete implements ExpensesUseCase.
func (e *expensesUseCase) Delete(id string) error {
	return e.repo.Delete(id)
}

func NewExpensesUseCase(repo repository.ExpensesRepository) ExpensesUseCase {
	return &expensesUseCase{repo: repo}
}
