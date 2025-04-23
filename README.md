# Create (Menambah Data)

## Penjelasan

- Untuk menambah data baru, buat instance struct model, isi field-nya, lalu panggil `db.Create(&instanceModel)`.
- GORM akan mengisi field `ID`, `CreatedAt`, dan `UpdatedAt` (jika menggunakan `gorm.Model`) setelah data berhasil disimpan.
- Kamu bisa menyimpan banyak record sekaligus dengan melewatkan slice of models ke `db.Create()`.
