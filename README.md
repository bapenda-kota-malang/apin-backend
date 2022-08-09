# Aplikasi Terintegrasi (APIN) - Back End

## Description
Aplikasi yang mengintegrasikan beberapa fungsionalitas menjadi satu seperti BAPENDA, PPAT, AND WP.

## App Structure
Struktur direktori merujuk pada standar komunitas: golang-standards/project-layout (lihat references).

### cmd/
Starting poin, cukup mengimport `pkd/core` dan routes dari masing-masing handlers. Terdapat 3 aplikasi: 
1. `cmd/apin-bapenda`, merupakan starting poin aplikasi backend untuk bapenda
2. `cmd/apin-ppat`, merupakan starting poin aplikasi backend untuk ppat
3. `cmd/apin-wajibpajak`, merupakan starting poin aplikasi backend untuk wajibpajak

### internal/
Package yg related ke proses bisnis:
1. `internal/handlers`, menghandle http request meneruskan ke kode `modules`.
2. `internal/models`, implementasi model database + data-transfer-object
3. `internal/modules`, menghandle proses bisnis.

### pkg/
Package yang tidak related ke bisnis flow, bisa diimpor aplikasi lain kapan saja atau dibuatkan repo sendiri untuk masing-masing pacakge, stara eksternal library.
1. `pkg/core`, package bundle boilerplate menyediakan manajemen konfig sampai aplikasi dapat menjalankan service http
2. `pkg/httproutermod`, package turunan dari httprouter untuk mempermudah menangani request dan memanfaatkan middleware
3. `pkg/jwt`, pakcage untuk menangani kebutuhan jwt
4. `pkg/requester`, pakcage helper untuk http.request, spesifik berfungsi mendapatkan param by key
5. `pkg/structvalidator`, pakcage untuk menangani validasi struct

## Technical Workflow
1. Tim dapat melewati (cukup tahu) apa yang ada dalam `cmd/` karena direktori tersebut merupakan starting poin.
2. Proses kerja dapat dimulai dari direktori `internal/handlers/$handlertype`, dengan `$handlertype` adalah pilihan yand dibutuhkan:
	1. `bapenda`
	2. `main`
	3. `ppat`
	4. `wajibpajak`
3. Buat routes pada file `routes.go` (kecuali untuk `$handlertype`==`main`) dengan scope sesuai kebutuhan. **Note**: routes memanfaatkan handler.
4. Semi-Optional: Buat handler dengan nama direktori dan file sesuai scope saat membuat route jika belum ada, pastikan memahami scope bisnis. **Note**: handler memanfaatkan service untuk clean code
5. Semi-Optional: Buat service pada direktori `internal/services` jika belum ada. pastikan memahami scope bisnis.
6. Optional: Buat models pada direktori `internal/models` jika diperlukan. 
7. Optional: Buat package tambahan pada direktori `pkg` jika diperlukan dan tidak berhubungan dengan flow bisnis.

## Usage
Menjalankan aplikasi untuk development / debugging:
1. Masuk ke direktori `cmd/apin-backend`
2. Buat file `config.yml`, dapat dibuat dengan mengcopy file `config.yml.example`.
3. Modifikasi file `config.yml` sesuaikan dengan environment yang ada
4. Eksekusi peringah `go run .`

Menjalankan aplikasi untuk keperluan production:
1. Ikuti langkah 1-3 pada langkah menjalankan aplikasi untuk dev / debugging
2. Eksekusi perintah `go build .`
3. Eksekusi file executable `apin-backend` atau `apin-backend.exe`

Note: file `config.yml` dapat diletakkan pada beberapa pilihan direktori diantaranya:
1. Direktori yang sama dengan direktori file executable
2. `/etc/apin`
3. `$HOME/apin`

## Convention
1. Programming Convetion lihat file `README-CONV-CODE.md`
2. Database Convetion lihat file `README-CONV-DB.md`

## Git
Usahakan untuk commit mengikuti konvensi komunitas (https://www.conventionalcommits.org/en/v1.0.0/).

Gunakan branch sesuai kebutuhan.


## References
1. golang, the programming language (https://go.dev/)
2. httprouter, the url router management (https://github.com/julienschmidt/httprouter)
3. gorm, the orm (https://gorm.io/ or https://github.com/go-gorm/gorm)
4. project-layout, the layout structure (https://github.com/golang-standards/project-layout)
