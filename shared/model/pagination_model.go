package model

type Paging struct {
	Page        int `json:"page"`
	TotalPages  int `json:"totalPages"`
	RowsPerPage int `json:"rowsPerPage"`
	TotalRows   int `json:"totalRows"`
}
