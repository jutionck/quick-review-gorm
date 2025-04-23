# Update (Mengubah Data)

## Penjelasan

- `db.Save(&instanceModel)`: Jika instance model sudah punya Primary Key, GORM akan melakukan operasi UPDATE. Ini akan menyimpan semua field dari struct, termasuk yang belum diubah (mengubahnya jadi zero-value jika kosong).
- `db.Model(&instanceModel).Updates(map[string]interface{}{"Kolom1": Nilai1, "Kolom2": Nilai2})` atau `db.Model(&instanceModel).Updates(StructDenganKolomYangDiubah{})`: Metode yang lebih fleksibel untuk hanya mengupdate kolom tertentu. Sangat disarankan daripada Save untuk update parsial.
- Update dengan Where: Bisa mengupdate banyak record sekaligus tanpa mengambilnya terlebih dahulu: `db.Model(&Product{}).Where("price < ?", 50).Updates(Product{Price: 50})`.
