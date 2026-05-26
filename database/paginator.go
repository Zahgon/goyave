package database

import (
	"gorm.io/gorm"
)

// Paginator structure containing pagination information and result records.
type Paginator[T any] struct {
	DB *gorm.DB `json:"-"`

	Records *[]T `json:"records"`

	rawQuery          string
	rawQueryVars      []any
	rawCountQuery     string
	rawCountQueryVars []any

	MaxPage     int64 `json:"maxPage"`
	Total       int64 `json:"total"`
	PageSize    int   `json:"pageSize"`
	CurrentPage int   `json:"currentPage"`

	loadedPageInfo bool
}

// PaginatorDTO structure sent to clients as a response.
type PaginatorDTO[T any] struct {
	Records     []T   `json:"records"`
	MaxPage     int64 `json:"maxPage"`
	Total       int64 `json:"total"`
	PageSize    int   `json:"pageSize"`
	CurrentPage int   `json:"currentPage"`
}

func paginateScope(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	_ = "STUB: not implemented"
	return nil
}

// NewPaginator create a new Paginator.
//
// Given DB transaction can contain clauses already, such as WHERE, if you want to
// filter results.
//
//	articles := []model.Article{}
//	tx := db.Where("title LIKE ?", "%"+sqlutil.EscapeLike(search)+"%")
//	paginator := database.NewPaginator(tx, page, pageSize, &articles)
//	err := paginator.Find()
//	if response.WriteDBError(err) {
//		return
//	}
//	response.JSON(http.StatusOK, paginator)
func NewPaginator[T any](db *gorm.DB, page, pageSize int, dest *[]T) *Paginator[T] {
	_ = "STUB: not implemented"
	return nil
}

// Raw set a raw SQL query and count query.
// The Paginator will execute the raw queries instead of automatically creating them.
// The raw query should not contain the "LIMIT" and "OFFSET" clauses, they will be added automatically.
// The count query should return a single number (`COUNT(*)` for example).
func (p *Paginator[T]) Raw(query string, vars []any, countQuery string, countVars []any) *Paginator[T] {
	_ = "STUB: not implemented"
	return nil
}

func (p *Paginator[T]) updatePageInfo(db *gorm.DB) error { _ = "STUB: not implemented"; return nil }

// UpdatePageInfo executes count request to calculate the `Total` and `MaxPage`.
// When calling this function manually, it is advised to use a transaction that is calling
// `Find()` too, to avoid inconsistencies.
func (p *Paginator[T]) UpdatePageInfo() error { _ = "STUB: not implemented"; return nil }

// Find requests page information (total records and max page) if not already fetched using
// `UpdatePageInfo()` and executes the query. The `Paginator` struct is updated automatically,
// as well as the destination slice given in `NewPaginator()`.
//
// The two queries are executed inside a transaction.
func (p *Paginator[T]) Find() error { _ = "STUB: not implemented"; return nil }

// Invalidate previous page info.

func (p *Paginator[T]) rawStatement(tx *gorm.DB) *gorm.DB { _ = "STUB: not implemented"; return nil }
