package config

const (
	InsertExpenses      = `INSERT INTO expenses (date, amount, transaction_type, balance, description, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at`
	GetBalance          = `SELECT balance FROM expenses ORDER BY created_at DESC LIMIT 1`
	FindDate            = `SELECT * FROM expenses WHERE date BETWEEN $1 AND $2 ORDER BY updated_at`
	SelectExpenses      = `SELECT * FROM expenses ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	SelectExpenseById   = `SELECT * FROM expenses WHERE id = $1`
	SelectExpenseByType = `SELECT * FROM expenses WHERE transaction_type = $1`
	SelectCount         = `SELECT COUNT(*) FROM expenses`
	UpdateExpense       = `UPDATE expenses SET amount = $2, transaction_type = $3, description = $4, balance = $5 WHERE id = $1`
	DeleteExpense       = `DELETE FROM expenses WHERE id = $1`
)
