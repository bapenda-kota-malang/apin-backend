package crudbase

type CrudBase interface {
	GetCreateStruct() any
	GetListStruct() any
	GetDetailStruct() any
	GetUpdateStruct() any
	GetFilterStruct() any
}
