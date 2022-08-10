1. Menggunakan Bahasa Inggris, kecuali disebutkan eksplisit
2. Penggunaan case (`camelCase`, `snake_case`, `kebab-case`) yang disebutkan tetap dalam koridor tidak melanggar aturan golang.
3. Penulisan nama file menggunakan `camelCase`
4. Penulisan package menggunakan `lowercase`, Bahasa Indonesia untuk istilah terkait proses bisnis
5. Penulisan routing menggunakan `kebab-case`, Bahasa Indonesia untuk istilah terkait proses bisnis
6. Penulisan request dan response header menggunakan `kebab-case` 
7. Penulisan request body (payload) menggunakan menggunakan kombinasi (lihat file: `README-CONV-DB.md`)
7. Penulisan response key menggunakan menggunakan kombinasi (lihat file: `README-CONV-DB.md`)
8. Penulisan variabel menggunakan kombinasi TC (lihat file: `README-CONV-DB.md`)
9. Importing dikelompokkan menjadi 4 bagian dengan pemisah empty break-line dengan kelompok sebagai berikut:
	1. Golang built-in package atau package maintained by golang dev team
	2. Eksternal package
	3. Internal package tapi tidak related langsung dengan current-module
	4. Internal package yang related langsung dengan current-module
10. Contoh import, tanpa komen untuk implementasi kecuali memberi penjelasan :
	```
	package user // current module = user

	// built in package
	import "fmt"
	import "strconv"

	// eksternal package
	import "github.com/goplayground/govalidator"
	import "github.com/goplayground/test"

	// internal package yang tidak related langsung dengan current-module = user
	import "github.com/bapenda-kota-malang/apin/pkg/paramChecker"
	import "github.com/bapenda-kota-malang/apin/internal/models/page"

	// internal package yang related langsung dengan current-module
	import "github.com/bapenda-kota-malang/apin/internal/models/user"
	```
