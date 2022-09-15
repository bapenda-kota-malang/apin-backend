package gormhelper

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	// gi "github.com/juliangruber/go-intersect"

	"gorm.io/gorm"
)

type Pagination struct {
	Page     int
	PageSize int
}

type DateModel struct {
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

// Parsing incoming query to pagination struct and return pagination struct, if error when parsing return error
func ParseQueryPagination(q url.Values) (p Pagination, err error) {
	p.Page, err = strconv.Atoi(q.Get("page"))
	if err != nil {
		return
	}

	if p.Page == 0 {
		p.Page = 1
	}

	p.PageSize, err = strconv.Atoi(q.Get("page_size"))
	if err != nil {
		return
	}

	switch {
	case p.PageSize > 100:
		p.PageSize = 100
	case p.PageSize <= 0:
		p.PageSize = 10
	}

	return
}

// Pageinate based on query parameters
func Paginate(p *Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (p.Page - 1) * p.PageSize
		return db.Offset(offset).Limit(p.PageSize)
	}
}

// Filters based on query parameters
func Filter(q url.Values, filter interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		rt := reflect.TypeOf(filter)
		if rt.Kind() != reflect.Struct {
			return db
		}

		opts := []string{"=", "<", ">", "<=", ">=", "<>"} // the default is without like and between

		for i := 0; i < rt.NumField(); i++ {
			f := rt.Field(i)
			t := strings.Split(f.Tag.Get("json"), ",")[0]
			v := q.Get(t)
			fmt.Println(v)

			if v != "" {
				o := t + "_opt"
				vOpt := q.Get(o)
				if stringInSlice(vOpt, opts) {
					db.Where(f.Name+" "+vOpt+" ?", v)
				} else {
					db.Where(f.Name+" = ?", v)
				}
			}
		}

		return db
	}
}
