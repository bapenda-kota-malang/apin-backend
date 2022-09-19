package gormhelper

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"

	// gi "github.com/juliangruber/go-intersect"

	"gorm.io/gorm"
)

// Filters based on query parameters
func Filter(input interface{}, p *Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		rt := reflect.TypeOf(input)
		if rt.Kind() != reflect.Struct {
			return db
		}

		iV := reflect.ValueOf(input) // input value
		if iV.Kind() != reflect.Struct {
			panic("input must be a struct")
		}

		// allowed option, the default is without like and between
		opts := []string{"=", "<", ">", "<=", ">=", "<>"}

		iT := iV.Type() // input type
		for i := 0; i < iV.NumField(); i++ {
			iTF := iT.Field(i) // input type of the current field
			opt := iTF.Name[len(iTF.Name)-4:]

			// skip option, page, or page_size
			if opt == "_opt" || iTF.Name == "Page" || iTF.Name == "Page_Size" {
				continue
			}

			// check field value
			iVF := iV.Field(i) // input value of the current field
			v := fmt.Sprintf("%v", iVF.Interface())
			if v == "<nil>" { // a bit tricky here, nil is not detected as normal nil, TODO: find out about this
				continue
			}

			// check opt
			vOpt := "="
			o := iV.FieldByName(iTF.Name + "_Opt") // option
			if o.IsValid() && o.Interface() != nil {
				vOpt = o.String()
				continue
			}

			// escape pointer
			for iVF.Kind() == reflect.Ptr {
				iVF = iVF.Elem()
			}

			// add where query
			if stringInSlice(vOpt, opts) {
				db.Where(iTF.Name+" "+vOpt+" ?", iVF.Interface())
			} else {
				db.Where(iTF.Name+" = ?", iVF.Interface())
			}
		}

		// field pagination
		fP := iV.FieldByName("Page")
		fPS := iV.FieldByName("PageSize")
		if fP.IsValid() && fP.Type().Kind() == reflect.Int {
			p.Page = fP.Interface().(int)
			if p.Page == 0 {
				p.Page = 1
			}
		} else {
			p.Page = 1
		}
		if fPS.IsValid() && fPS.Type().Kind() == reflect.Int {
			p.PageSize = fPS.Interface().(int)
			if p.PageSize < 5 && p.PageSize > 100 {
				p.PageSize = 10
			}
		} else {
			p.PageSize = 10
		}
		db.Offset(p.Page).Limit(p.PageSize)

		return db
	}
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
