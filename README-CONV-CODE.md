1. Penulisan nama file menggunakan camelCase, Bahasa Inggris
2. Penulisan package menggunakan lowercase, Bahasa Inggris untuk istilah non proses binsis, Bahasa Indonesia untuk istilah proses bisnis
3. Penulisan routing menggunakan kebab-case, Bahasa Indonesia kecuali yang tidak berhubungan dengan flow bisnis, contoh: `/user`, `/group`
4. Penulisan request dan response header, menggunakan kombinasi Pascal dan kebab-case
5. Penulisan request payload dan response body menggunakan menggunakan kombinasi sesuai konvensi penulisan database (lihat file: `README-CONV-DB.md`)
6. Importing dikelompokkan menjadi 4 bagian dengan pemisah empty break-line dengan kelompok sebagai berikut:
	1. Golang built-in package atau package maintained by golang dev team
	2. Eksternal package
	3. Internal package tapi tidak related langsung dengan current-module
	4. Internal package yang related langsung dengan current-module
7. Contoh import, tanpa komen untuk implementasi kecuali memberi penjelasan :
	```
	package user // current module = user

	// built in package
	import "fmt"
	import "strconv"

	// eksternal package
	import "github.com/goplayground/govalidator"
	import "github.com/goplayground/test"

	// internal package yang tidak related langsung dengan current-module
	import "github.com/bapenda-kota-malang/apin/internal/libraries/paramChecker"
	import "github.com/bapenda-kota-malang/apin/internal/models/page"

	// internal package yang tidak related langsung dengan current-module
	import "github.com/bapenda-kota-malang/apin/internal/models/user"
	```
