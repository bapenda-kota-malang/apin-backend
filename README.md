# Aplikasi Terintegrasi (APIN) - Back End

## Description
Aplikasi yang mengintegrasikan beberapa fungsionalitas menjadi satu seperti BAPENDA, PPAT, AND WP.

## Structure
Struktur direktori merujuk pada standar komunitas: golang-standards/project-layout (lihat references). Penjelasan tambahan (sub) dari yang ada di golang-standards/project-layout:
1. `cmd/apin-backend`, merupakan starting poin aplikasi, cukup mengimport core dan routes.
2. `internal/core`, merupakan core utama yang membundle kode boiler-plate aplikasi. Untuk informasi lebih detail silahkan lihat readme di direktori tersebut
3. `internal/modules`, berisi kode yang menghandle proses bisnis.
4. `internal/models`, berisi implementasi model database beserta data -transfer-object

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
1. Programming Convetion lihat di sini:
2. Database Convetion lihat di sini:

## References
1. golang, the programming language (https://go.dev/)
2. httprouter, the url router management (https://github.com/julienschmidt/httprouter)
3. gorm, the orm (https://gorm.io/ or https://github.com/go-gorm/gorm)
4. project-layout, the layout structure (https://github.com/golang-standards/project-layout)
