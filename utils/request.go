package utils

import (
	"gorm.io/gorm"
)

type PaginationRequest struct {
	Paginate  bool  `query:"paginate"`
	Page      int   `query:"page"`
	PerPage   int   `query:"per_page"`
	TotalPage int64 `query:"-"`
	TotalData int64 `query:"-"`
}

func Paginate(db *gorm.DB, pr *PaginationRequest) *gorm.DB {
	if pr.Page <= 0 {
		pr.Page = 1
	}

	if pr.PerPage <= 0 {
		pr.PerPage = 10
	}

	offset := (pr.Page - 1) * pr.PerPage

	return db.Count(&pr.TotalData).Offset(offset).Limit(pr.PerPage)
}
