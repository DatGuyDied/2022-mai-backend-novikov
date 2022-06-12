package domain

import (
	"net/http"
	"strconv"
)

type PaginationOrder string

const (
	PaginationOrderAsc  PaginationOrder = "ASC"
	PaginationOrderDesc PaginationOrder = "DESC"
)

func (po PaginationOrder) IsValid() bool {
	switch po {
	case PaginationOrderAsc, PaginationOrderDesc:
		return true
	default:
		return false
	}
}

const (
	defalutLimit = 10
	defaultPage  = 1
)

type Pagination struct {
	Limit int
	Page  int
	Order PaginationOrder
}

func (p *Pagination) Bind(r *http.Request) {
	q := r.URL.Query()

	p.Limit = parseInt(q.Get("limit"), defalutLimit)
	p.Page = parseInt(q.Get("page"), defaultPage)

	orderStr := q.Get("order")
	if order := PaginationOrder(orderStr); order.IsValid() {
		p.Order = order
	} else {
		p.Order = PaginationOrderAsc
	}
}

func parseInt(str string, fallback int) int {
	if v, err := strconv.Atoi(str); err == nil {
		return v
	}
	return fallback
}
