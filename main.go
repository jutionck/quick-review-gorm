package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// --- Definisi Model (Lengkap untuk Challenge) ---
type Product struct {
	gorm.Model
	Code  string `gorm:"unique;size:10"`
	Price uint
}

type User struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"not null"`
	Email     string         `gorm:"unique"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	// Relasi One-to-Many (User punya banyak Post)
	Posts []Post
}

type Post struct {
	gorm.Model
	Title  string `gorm:"not null"`
	Body   string
	UserID uint // Foreign Key ke User
	// Relasi Many-to-One (Post dimiliki oleh satu User)
	User User
}

func main() {
	// --- Koneksi Database ---
	log.Println("Attempting to connect to database...")
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}
	log.Println("Database connection successfully opened.")

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get underlying DB instance: %v", err)
	}
	defer func() {
		log.Println("Closing database connection.")
		sqlDB.Close()
		log.Println("Database connection closed.")
	}()

	// --- Auto Migrate ---
	log.Println("Running auto migration for Product, User, Post...")
	// Pastikan semua tabel terbuat atau diperbarui
	err = db.AutoMigrate(&Product{}, &User{}, &Post{})
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
	log.Println("Database auto migration finished.")

	// --- CHALLENGE AREA ---
	// Peserta akan menulis kode CRUD di area ini untuk menyelesaikan challenge.
	// Gunakan model User dan Post yang sudah didefinisikan.

	log.Println("\n--- Challenge Starts Here ---")

	// TODO: 1. Buat beberapa User baru (jika belum ada atau ingin data bersih)
	// TODO: 2. Buat beberapa Post baru, hubungkan dengan User menggunakan UserID
	// TODO: 3. Query (Read) semua Post dan tampilkan
	// TODO: 4. Query (Read) semua Post milik User tertentu (gunakan Where atau Preload)
	// TODO: 5. Update judul dari salah satu Post
	// TODO: 6. Soft Delete salah satu Post
	// TODO: 7. Coba cari Post yang di-soft delete (query biasa vs Unscoped)

	log.Println("\n--- Challenge Area Ends Here ---")
	log.Println("Application finished.")
}
