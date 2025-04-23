# Migrasi Database

## Penjelasan

Setelah mendefinisikan model, kita perlu membuat tabel yang sesuai di database. Proses ini disebut migrasi.

GORM memiliki fitur `AutoMigrate()` yang secara otomatis membuat tabel, menambahkan kolom yang hilang, dan membuat index berdasarkan definisi model.
Penting: `AutoMigrate` tidak menghapus kolom/tabel yang sudah ada, atau mengubah tipe data kolom yang sudah ada. Untuk skenario produksi yang lebih kompleks, disarankan menggunakan tool migrasi terpisah (seperti `golang-migrate`).
