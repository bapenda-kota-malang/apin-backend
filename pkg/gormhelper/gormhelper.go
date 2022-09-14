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

// Pageinate based on query parameters
func Paginate(u *url.URL, p *Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := u.Query()
		p.Page, _ = strconv.Atoi(q.Get("page"))
		if p.Page == 0 {
			p.Page = 1
		}

		p.PageSize, _ = strconv.Atoi(q.Get("page_size"))
		switch {
		case p.PageSize > 100:
			p.PageSize = 100
		case p.PageSize <= 0:
			p.PageSize = 10
		}

		offset := (p.Page - 1) * p.PageSize
		return db.Offset(offset).Limit(p.PageSize)
	}
}

// Filters based on query parameters
func Filter(u *url.URL, filter interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		rt := reflect.TypeOf(filter)
		if rt.Kind() != reflect.Struct {
			return db
		}

		q := u.Query()
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
