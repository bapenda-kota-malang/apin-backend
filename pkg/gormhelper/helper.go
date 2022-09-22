package gormhelper

// just local functions, avoid import if it is simple

// check if stnig availble in an array
func stringInSlice(s string, a []string) bool {
	for _, v := range a {
		if s == v {
			return true
		}
	}
	return false
}

// reflect field value or nil
// func refOperatorVal(iV reflect.Value, iTF reflect.StructField) string {
// 	vOpt := "="
// 	o := iV.FieldByName(iTF.Name + "_Opt") // option
// 	if o.Interface() != nil {
// 		vOpt = o.String()
// 	}
// 	return vOpt
// }
