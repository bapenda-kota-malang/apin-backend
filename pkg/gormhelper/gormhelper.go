package gormhelper

import (
	"fmt"
	"reflect"

	// gi "github.com/juliangruber/go-intersect"

	"gorm.io/gorm"
)

// Filters based on query parameters
func Filter(input interface{}) func(db *gorm.DB) *gorm.DB {
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
		opts := []string{"=", "<", ">", "<=", ">=", "<>", "left", "mid", "right"}

		iT := iV.Type() // input type
		for i := 0; i < iV.NumField(); i++ {
			iTF := iT.Field(i) // input type of the current field
			opt := iTF.Name
			if len(iTF.Name) >= 4 {
				opt = iTF.Name[len(iTF.Name)-4:]
			}

			// skip option and pagination related
			if opt == "_Opt" || iTF.Name == "Page" || iTF.Name == "PageSize" || iTF.Name == "NoPagination" {
				continue
			}

			// proceed value
			iVF := iV.Field(i) // input value of the current field
			fmt.Println(iTF.Type.Kind())
			for iVF.Kind() == reflect.Ptr {
				iVF = iVF.Elem()
			}

			// check field value
			lastITF := reflect.TypeOf(iVF)
			if !iVF.IsValid() || iVF.IsZero() || (lastITF.Kind() == reflect.Ptr && iVF.IsNil()) {
				continue
			}

			// check opt
			vOpt := "="
			o := iV.FieldByName(iTF.Name + "_Opt") // option
			if o.IsValid() && o.Interface() != nil {
				if o.Kind() == reflect.Ptr {
					o = o.Elem()
				}
				vOpt = o.Interface().(string)
			}

			// add where query
			if stringInSlice(vOpt, opts[0:6]) {
				db.Where("\""+iTF.Name+"\" "+vOpt+" ?", iVF.Interface()) // F-KING BUG
			} else if vOpt == "left" {
				db.Where("\""+iTF.Name+"\" LIKE ?", fmt.Sprintf("%v%v", iVF.Interface(), "%"))
			} else if vOpt == "mid" {
				db.Where("\""+iTF.Name+"\" LIKE ?", fmt.Sprintf("%v%v%v", "%", iVF.Interface(), "%"))
			} else if vOpt == "right" {
				db.Where("\""+iTF.Name+"\" LIKE ?", fmt.Sprintf("%v%v", "%", iVF.Interface()))
			} else {
				db.Where("\""+iTF.Name+"\" = ?", iVF.Interface()) // F-KING BUG
			}
		}

		return db
	}
}

// Pageinate based on query parameters
func Paginate(input interface{}, p *Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		iV := reflect.ValueOf(input) // input value
		if iV.Kind() != reflect.Struct {
			panic("input must be a struct")
		}
		// field pagination
		fP := iV.FieldByName("Page")
		fPS := iV.FieldByName("PageSize")
		fNP := iV.FieldByName("NoPagination")
		if fNP.IsValid() {
			if fNP.Type().Kind() == reflect.Bool && bool(fNP.Interface().(bool)) {
				return db
			}
		}
		if fP.IsValid() {
			if fP.Type().Kind() == reflect.Int {
				p.Page = fP.Interface().(int)
			} else if fP.Type().Kind() == reflect.Int64 {
				p.Page = int(fP.Interface().(int64))
			} else {
				p.Page = 1
			}
			if p.Page <= 0 {
				p.Page = 1
			}
		} else {
			p.Page = 1
		}
		if fPS.IsValid() {
			if fPS.Type().Kind() == reflect.Int {
				p.PageSize = fPS.Interface().(int)
			} else if fPS.Type().Kind() == reflect.Int64 {
				p.PageSize = int(fPS.Interface().(int64))
			} else {
				p.PageSize = 10
			}
			if p.PageSize <= 0 {
				p.PageSize = 5
			}
		} else {
			p.PageSize = 10
		}
		offset := (p.Page - 1) * p.PageSize
		return db.Offset(offset).Limit(p.PageSize)
	}
}
