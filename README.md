# Pengantar ORM dan GORM

## Penjelasan

Apa itu **ORM (Object-Relational Mapper)?** ORM adalah teknik atau tool yang memetakan (mapping) objek dalam kode program (seperti struct di Go) ke tabel dalam database relasional. Ini memungkinkan Anda berinteraksi dengan database menggunakan sintaks bahasa pemrograman Anda, bukan SQL mentah.

## Mengapa Menggunakan ORM (dan GORM)?

- Mengurangi boilerplate code SQL.
- Meningkatkan produktivitas.
- Memudahkan pengelolaan skema database (migrasi).
- Menyediakan abstraksi untuk berbagai jenis database.

## Apa itu GORM?

GORM adalah ORM populer untuk Go, dikenal karena kemudahan penggunaan, fitur lengkap (`associations`, `hooks`, `eager loading`, `transactions`, dll.), dan komunitas yang aktif.

## Prasyarat Singkat

Sudah memahami dasar Go (struct, pointer, slice, error handling)

# Koneksi Database

## Penjelasan

- GORM memerlukan driver spesifik untuk setiap jenis database (`PostgreSQL`, `MySQL`, `SQLite`, `SQL Server`, dll.) Driver ini yang menerjemahkan perintah GORM ke SQL yang sesuai.
- Koneksi database biasanya menggunakan Data Source Name (DSN) yang berisi informasi seperti `user`, `password`, `host`, `port`, `nama database`.
- Fungsi utama untuk membuka koneksi adalah `gorm.Open()`. - Fungsi ini mengembalikan instance `gorm.DB` yang akan kita gunakan untuk semua operasi database, dan sebuah error.
- Penting untuk selalu memeriksa error setelah `gorm.Open()`.
- Menutup koneksi database saat aplikasi berhenti (`db.Close()` - meskipun ini kurang umum pada aplikasi web modern yang menjaga koneksi tetap hidup).
