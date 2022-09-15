# Aplikasi Terintegrasi (APIN) - Back End

## Description
Aplikasi yang mengintegrasikan beberapa fungsionalitas terkai pajak menjadi satu seperti BAPENDA, PPAT, AND WP.

## App Structure
Struktur direktori merujuk pada standar komunitas: `golang-standards/project-layout` (lihat references).

### cmd/
Starting poin, cukup mengimport `pkg/core` dan `routes` dari masing-masing handlers. Terdapat 3 aplikasi: 
1. `apin-bapenda`, merupakan starting poin aplikasi backend untuk bapenda
2. `apin-ppat`, merupakan starting poin aplikasi backend untuk ppat
3. `apin-wajibpajak`, merupakan starting poin aplikasi backend untuk wajibpajak

### internal/
Package yang related ke proses bisnis (menyangkut data):
1. `handlers`, menghandle http request (meneruskan request `services`) dan respond (menerjemahkan return dari `service`) 
2. `models`, implementasi model database + data-transfer-object
3. `services`, menghandle proses bisnis.

### pkg/
Package yang tidak related ke bisnis flow, bisa diimpor aplikasi lain kapan saja atau dibuatkan repo sendiri untuk masing-masing pacakge, setara eksternal library. Beberapa pkg utama:
1. `apicore`, package bundle boilerplate menyediakan manajemen konfig sampai aplikasi dapat menjalankan service http
2. `routerhelper`, package untuk membantu pembuatan routing standar (saat dokumen ini dibuat hanya membantu pembuatan routing CRUD)
3. `handlerhelper`, package untuk membantu handler menangani request + response
4. `servicehelper`, pakcage untuk membantu service melakukan aksi yang dianggap standar
4. `servicehelper`, pakcage untuk menangani validasi struct
5. `gormhelper`, package untuk membantu aksi yang berkaitan dengan library `gorm`

## Workflow
1. Tim dapat melewati (cukup tahu) apa yang ada dalam `cmd/` karena direktori tersebut hanya merupakan starting poin.
2. Proses kerja dapat dimulai dari direktori `internal/handlers/$type`, dengan `$type` adalah pilihan yand dibutuhkan:
	1. `bapenda`
	2. `main`
	3. `ppat`
	4. `wajibpajak`
3. Buat routes pada file `routes.go` (kecuali `$type` adalah `main`) dengan scope sesuai kebutuhan. **Note**: routes memanfaatkan handler.
4. Buat (jika belum ada) atau update handler pada direktori `inrwenal/handlers`, pastikan memahami scope bisnis. **Note**: handler memanfaatkan service untuk clean code
5. Buat (jika belum ada) atau update service pada direktori `internal/services`, pastikan memahami scope bisnis.
6. Buat models pada direktori `internal/models` jika diperlukan. 
7. Buat package tambahan untuk proses yang tidak berhubungan dengan flow bisnis pada direktori `pkg` jika diperlukan.

## Usage
Kebutuhan dasar
1. Pastikan di komputer terinstall golang versi 1.18 keatas
2. Pastikan di komputer terinstall postgresql versi 10.x keatas
3. Clone aplikasi ini dari repository ke direktori yang dikehendaki
4. Masuk ke direktori yang telah dibuat dan jalankan perintah `go mod tidy`

Menjalankan aplikasi untuk development / debugging:
1. Masuk ke direktori `cmd/apin-backend`
2. Buat file `config.yml`, dapat dibuat dengan mengcopy file `config.yml.example`.
3. Modifikasi file `config.yml` sesuaikan dengan environment yang ada
4. Eksekusi peringah `go run .`

Menjalankan aplikasi untuk keperluan production:
1. Ikuti langkah 1-3 pada langkah menjalankan aplikasi untuk dev / debugging
2. Eksekusi perintah `go build .`
3. Eksekusi file executable `apin-backend` atau `apin-backend.exe`

## Convention
1. Programming Convetion lihat file `README-CONV-CODE.md`
2. Database Convetion lihat file `README-CONV-DB.md`

## References
1. golang, the programming language (https://go.dev/)
2. chi, the url router management (https://github.com/go-chi/chi atau https://go-chi.io/)
3. gorm, the orm (https://gorm.io/ or https://github.com/go-gorm/gorm atau https://gorm.io/)
4. project-layout, the layout structure (https://github.com/golang-standards/project-layout)
