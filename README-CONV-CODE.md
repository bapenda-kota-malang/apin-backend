1. Penulisan nama file menggunakan camelCase
2. Penulisan package menggunakan camelCase
2. Penulisan routing menggunakan kebab-case, Bahasa Indonesia kecuali yang tidak berhubungan dengan flow bisnis, contoh: `/user`, `/group`
3. Importing dikelompokkan menjadi 4 bagian dengan pemisah empty break-line dengan kelompok sebagai berikut:
	1. Golang built-in package atau package maintained by golang dev team
	2. Eksternal package
	3. Internal package tapi tidak related langsung dengan current-module
	4. Internal package yang related langsung dengan current-module
4. Contoh import, tanpa komen untuk implementasi kecuali memberi penjelasan :
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
