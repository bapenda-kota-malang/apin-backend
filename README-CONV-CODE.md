# Convention in Executing Task
## Writing Code
1. Menggunakan Bahasa Inggris, kecuali disebutkan eksplisit
2. Penggunaan case (`camelCase`, `snake_case`, `kebab-case`) yang disebutkan tetap dalam koridor tidak melanggar aturan golang.
3. Penulisan nama file iutamakan all `lowercase`, atau menggunakan `camelCase` jika dibutuhkan
4. Penulisan package menggunakan `lowercase`, Bahasa Indonesia untuk istilah terkait proses bisnis
5. Penulisan routing menggunakan `kebab-case`, Bahasa Indonesia untuk istilah terkait proses bisnis
6. Penulisan request dan response header menggunakan `kebab-case` 
7. Penulisan request body (payload) menyesuaikan atribute pada database KECUALI setiap atribute dan kata setelah underscore diawali huruf kecil
8. Penulisan response key menggunakan menggunakan kombinasi (lihat poin sebelumnya)
8. Penulisan variabel menggunakan menggunakan kombinasi (lihat poin sebelumnya), bahasa Indonesia jika itu related ke tabel-kolom
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
## Git
1. Flow
	1. Normal, mengerjakan fitur yang masih dalam pengembangan awal (belum deployed di production) dengan langkah:
		1. Clone dari branch `dev`
		2. Push task yang selesai ke repository
		3. Pull request ke `dev`
	2. Hotfix, mengerjakan fitur yang sudah deployed di production untuk bug-fixing atau refactor dengan langkah
		1. Clone dari branch `main`
		2. Push task yang selesai ke repository
		3. Pull request ke `staging`
2. Branch Name
	1. Menggunakan format `$prefix/$scope`
	2. `$prefix` terdapat beberapa jenis:
		1. `feat`: menambahkan fitur baru
		2. `fix`: mengubah fitur yang ada karena bug
		3. `refactor`: mengubah fitur yang sudah ada selain karena bug
		4. `chore`: menambah dan/atau mengubah sesuatu tidak berkaitan dengan programming 
	3. `$scope` bisa disesuaikan dengan kebutuhan
		1. Tingkat tinggi bisa berupa nama modul secara umum, contoh: `services` karena menyangkut semua service  
		2. Tingkat rendah bisa berupa fitur terkecil yang digunakan, contoh: `regency-name-input`
3. Commit Message
	1. Menggunakan format `$prefix $scope\n\n[$description]`
	2. `$prefix` dan `$scope` dapat dilihat pada poin `Branch Name` sebelumnya
	3. `$description` opsional, berisi detail deskripsi task
