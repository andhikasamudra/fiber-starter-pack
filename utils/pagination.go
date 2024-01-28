package utils

import (
	"context"
	"fmt"
	"strings"

	"github.com/uptrace/bun"
)

type Paging struct {
	Total       int `json:"total" query:"total"`
	PageSize    int `json:"page_size" query:"page_size"`
	CurrentPage int `json:"current_page" query:"current_page"`
}

type TableListParams struct {
	Sorter string `json:"sorter"`
	Paging
}

func (s *Paging) Norm() {
	if s.PageSize == 0 {
		s.PageSize = 10
	}

	if s.CurrentPage == 0 {
		s.CurrentPage = 1
	}
}

func (s *Paging) GetPageSize() int {
	if s.PageSize == 0 {
		return 10
	}
	return s.PageSize
}

func (s *Paging) GetCurrentPage() int {
	if s.CurrentPage == 0 {
		return 1
	}
	return s.CurrentPage
}

func (s *Paging) GetOffset() int {
	return s.GetPageSize() * (s.GetCurrentPage() - 1)
}

func GetMultipleSorters(sorter string) []string {
	splitSorter := func(r rune) bool {
		return r == ','
	}

	return strings.FieldsFunc(RemoveWhitespace(sorter), splitSorter)
}

func WithMultipleSorters(query *bun.SelectQuery, sorters []string) {
	for _, sorter := range sorters {
		WithSorter(query, sorter)
	}
}

func WithSorter(query *bun.SelectQuery, sorter string) *bun.SelectQuery {
	parsedSorter := strings.Split(sorter, "_")

	fieldName := strings.Join(parsedSorter[:len(parsedSorter)-1], "_")
	order := fmt.Sprintf("%s %s", fieldName, parsedSorter[len(parsedSorter)-1])
	query = query.Order(order)

	return query
}

func WithPaging(ctx context.Context, query *bun.SelectQuery, paging *Paging) (*bun.SelectQuery, error) {
	paging.Norm()
	count, err := query.Count(ctx)
	if err != nil {
		return nil, err
	}
	paging.Total = count
	query = query.Offset(paging.GetOffset())
	query = query.Limit(paging.GetPageSize())

	return query, nil
}

//nolint:ifElseChain
func WithTableListParams(ctx context.Context, query *bun.SelectQuery, tableListParams *TableListParams) (*bun.SelectQuery, error) {
	if tableListParams == nil {
		return query, nil
	}

	multipleSorters := GetMultipleSorters(tableListParams.Sorter)

	if len(multipleSorters) > 1 {
		WithMultipleSorters(query, multipleSorters)
	} else if len(multipleSorters) == 1 {
		WithSorter(query, multipleSorters[0])
	} else {
		WithSorter(query, tableListParams.Sorter)
	}

	if _, err := WithPaging(ctx, query, &tableListParams.Paging); err != nil {
		return nil, err
	}

	return query, nil
}
