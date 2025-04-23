# Read (Membaca Data)

## Penjelasan

### Membaca satu record

- `db.First(&instanceModel, id)`: Mengambil record pertama berdasarkan primary key ID.
- `db.First(&instanceModel, "condition")`: Mengambil record pertama yang cocok dengan kondisi string (kurang disarankan karena rentan SQL injection).
- `db.Where("condition", args...).First(&instanceModel)`: Cara yang lebih aman dan umum menggunakan Where.
- `db.Last(&instanceModel)`: Mengambil record terakhir.

### Membaca banyak record

- `db.Find(&sliceOfModels)`: Mengambil semua record.
- `db.Where("condition", args...).Find(&sliceOfModels)`: Mengambil record yang cocok dengan kondisi.

### Chaining Methods: Metode query GORM bisa dirangkai (chained)

**`db.Where(...).Order(...).Limit(...).Offset(...).Find(...)`.**

- `Where("kolom operator ?", nilai)`: Filtering data (misal: `"price > ?"`, `"name LIKE ?"`, `"id IN ?"`).
- `Order("kolom [asc|desc]")`: Mengurutkan hasil.
- `Limit(n)`: Membatasi jumlah record yang diambil.
- `Offset(n)`: Melewati n record (untuk pagination).
