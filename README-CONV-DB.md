1. Penulisan nama tabel menggunakan kombinasi, dengan aturan:
	1. Pada dasarnya menggunakan camelCase
	2. Penggunaan snake_case diterapkan untuk tabel turunan, pemisah nama asli dan tujuan turunan
	3. Tabel turunan yaitu tabel yang diturunkan dari tabel yang lain tetapi tidak digunakan pada flow utama proses bisnis.
	3. Contoh implementasi tabel turunan diantaranya:
		1. Archiving, contoh: `transaction_2019`, merupakan arsip tabel `transaction` tahun `2019`.
		2. Trashing untuk data yang dihapus, contoh: `transaction_deleted`, menyimpan data tabel `transaction` yang telah dihapus.
		3. History untuk menyimpan informasi history, contoh `productStock_history`, menyimpan data history tabel `productStock`.
2. Penulisan kolom menggunakan kombinasi, dengan aturan:
	1. Pada dasarnya menggunakan camelCase
	2. Penggunaan snake_case untuk menandakan sebuah kolom memiliki informasi TC (`table`, `column`), sehingga normalnya penamaan kolom dengan informasi TC ditujukan kolom yang menyimpan nilai FK (Foreign-Key)
	3. Kolom dengan informasi TC dituliskan dengan format: `$table` + `_` + `$column`, dimana `$table` merupakan nama (ambil referensi) tabel dan `$column` merupakan nama (ambil referensi) kolom.
	4. Contoh panamaan kolom snake_case:
		1. `user_id` berarti mengambil referensi dari tabel `user` kolom `id`
		2. `bestProduct_code` berarti mengambil referensi dari tabel `bestProduct` kolom `code`
	5. Ketika terdapat lebih dari satu kolom memiliki Informasi TC yang sama, dapat ditambahkan `prefix` untuk mencerminkan kegunaan kolom tersebut
	6. Contoh panamaan kolom snake_case untuk kolom dengan informasi TC lebih dari satu:
		1. `admin_user_id` berarti untuk `admin` mengambil referensi dari tabel `user` kolom `id`
		2. `operator_user_id` berarti untuk `operator` mengambil referensi dari tabel `user` kolom `id`
	7. Informasi TC dapat juga ditemukan pada kolom non-FK ketika terjadi denormalisasi.
	8. Contoh non-FK, contoh pada tabel transaction dapat ditemukan kolom `product_name` dimana `product` adalah nama table dan `name` adalah nama kolom