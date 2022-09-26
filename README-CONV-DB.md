1. Penulisan nama tabel menggunakan kombinasi, dengan aturan:
	1. Pada dasarnya menggunakan PascalCase
	2. Underscore digunakan untuk tabel turunan, pemisah nama asli dan tujuan turunan
	3. Tabel turunan yaitu tabel yang diturunkan dari tabel yang lain tetapi tidak digunakan pada flow utama proses bisnis.
	3. Contoh implementasi tabel turunan diantaranya:
		1. Archiving, contoh: `Transaction_2019`, merupakan arsip tabel `Transaction` tahun `2019`.
		2. Trashing untuk data yang dihapus, contoh: `Transaction_deleted`, menyimpan data tabel `transaction` yang telah dihapus.
		3. History untuk menyimpan informasi history, contoh `ProductStock_History`, menyimpan data history tabel `ProductStock`.
2. Penulisan kolom menggunakan kombinasi, dengan aturan:
	1. Pada dasarnya menggunakan `PascalCase`
	2. Penggunaan underscore untuk menandakan sebuah kolom memiliki informasi TC (`Table`, `Column`)
		1. Ditujukan kolom yang menyimpan nilai FK (Foreign-Key)
		2. Dituliskan dengan format: `$Table` + `_` + `$Column`, dimana `$Table` merupakan nama (ambil referensi) tabel dan `$Column` merupakan nama (ambil referensi) kolom.
		3. Contoh panamaan kolom snake_case:
			1. `User_Id` berarti mengambil referensi dari tabel `User` kolom `Id`
			2. `BestProduct_Code` berarti mengambil referensi dari tabel `BestProduct` kolom `Code`
		4. Ketika terdapat lebih dari satu kolom memiliki Informasi TC yang sama, dapat ditambahkan `prefix` untuk mencerminkan kegunaan kolom tersebut
		5. Contoh panamaan kolom snake_case untuk kolom dengan informasi TC lebih dari satu:
			1. `Admin_User_Id` berarti untuk `Admin` mengambil referensi dari tabel `User` kolom `Id`
			2. `Operator_User_Id` berarti untuk `Operator` mengambil referensi dari tabel `User` kolom `Id`
	3. Penggunaan underscore untuk keperluan denormalisasi.
		1. Contoh non-FK, contoh pada tabel transaction dapat ditemukan kolom `Product_Name` dimana `Product` adalah nama table dan `Name` adalah nama kolom