# Delete (Menghapus Data)

## Penjelasan:

- Hard Delete: Menghapus record secara permanen dari database. Gunakan `db.Delete(&instanceModel, id) atau db.Where(...).Delete(&Model{})`.
- Soft Delete: Tidak benar-benar menghapus data, melainkan menandainya sebagai "terhapus" (biasanya dengan mengisi kolom `deleted_at`). GORM mendukung ini secara otomatis jika model memiliki field `gorm.DeletedAt`(yang sudah ada di `gorm.Model`). Saat menggunakan `db.Delete()` pada model dengan `DeletedAt`, GORM akan melakukan UPDATE `deleted_at` alih-alih DELETE.
- Mengambil Record yang Di-Soft Delete: Secara default, GORM akan mengecualikan record yang di-soft delete dari hasil query. Gunakan `db.Unscoped().Find(...)` atau `db.Unscoped().Where(...)` untuk menyertakan record yang di-soft delete.
- Menghapus Permanen (dengan Soft Delete aktif): Gunakan `db.Unscoped().Delete(...).`
