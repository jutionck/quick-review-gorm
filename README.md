# Relasi (Overview Singkat)

## Penjelasan

- Database relasional seringkali memiliki hubungan antar tabel (misal: satu User memiliki banyak Post, satu Product dimiliki oleh satu Category). GORM memungkinkan pendefinisian relasi ini dalam struct Go.
- Jenis Relasi: One-to-One, One-to-Many, Many-to-Many.
- Pendefinisian: Ditentukan dengan menambahkan field struct yang merupakan pointer atau slice ke struct model lain, seringkali dikombinasikan dengan tag `gorm:"..."` seperti `foreignKey`, `references`, `many2many`.

## Loading Relasi

- Lazy Loading: Relasi tidak diambil secara otomatis saat query model utama. Perlu query terpisah untuk mengambil data relasi (berpotensi N+1 problem).
- Eager Loading: Mengambil data relasi bersamaan dengan model utama, biasanya menggunakan metode `Preload()`. Ini lebih efisien untuk mengambil banyak data relasi sekaligus.
