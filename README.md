# Definisi Model

## Penjelasan

Model di GORM direpresentasikan sebagai struct Go. Setiap field dalam struct secara default akan dipetakan menjadi kolom di tabel database.

## Konvensi GORM

- Nama struct (PascalCase) biasanya dipetakan ke nama tabel (snake_case, plural): `User` -> `users`, `Task` -> `tasks`.
- Nama field struct (PascalCase) dipetakan ke nama kolom (snake_case): `CreatedAt` -> `created_at`, `TaskName` -> `task_name`.

## gorm.Model

Ini adalah struct bawaan GORM yang menyediakan field umum seperti ID (primary key auto-increment), `CreatedAt`, `UpdatedAt`, dan `DeletedAt` (untuk soft delete). Sangat disarankan untuk menyertakan ini di sebagian besar model.

## GORM Tags (gorm:"...")

Digunakan untuk mengkustomisasi mapping atau menambahkan constraint pada kolom, contoh:

- `gorm:"column:nama_kolom_kustom"`: Mengubah nama kolom.
- `gorm:"type:varchar(100)"`: Menentukan tipe data SQL.
- `gorm:"primaryKey"`: Menandai sebagai primary key.
- `gorm:"unique"`: Menambahkan constraint unique.
- `gorm:"not null"`: Menambahkan constraint not null.
- `gorm:"default:0"`: Menambahkan nilai default.
