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
1. `pkg/core`, bundle boilerplate menyediakan manajemen konfig sampai aplikasi dapat menjalankan service http
2. `pkg/jwt`, pakcage untuk menangani kebutuhan jwt
3. `pkg/structvalidator`, pakcage untuk menangani validasi struct

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

## Misc
Usahakan untuk commit mengikuti konvensi komunitas (https://www.conventionalcommits.org/en/v1.0.0/).

## References
1. golang, the programming language (https://go.dev/)
2. httprouter, the url router management (https://github.com/julienschmidt/httprouter)
3. gorm, the orm (https://gorm.io/ or https://github.com/go-gorm/gorm)
4. project-layout, the layout structure (https://github.com/golang-standards/project-layout)
