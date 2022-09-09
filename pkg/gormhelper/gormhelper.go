package gormhelper

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	// gi "github.com/juliangruber/go-intersect"

	"gorm.io/gorm"
)

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := r.URL.Query()
		page, _ := strconv.Atoi(q.Get("page"))
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(q.Get("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func Filter(r *http.Request, filter interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		rt := reflect.TypeOf(filter)
		if rt.Kind() != reflect.Struct {
			return db
		}

		q := r.URL.Query()
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
