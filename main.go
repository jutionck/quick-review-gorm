package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// --- Bagian 2: Definisi Model ---
type Task struct {
	gorm.Model // Menyertakan ID, CreatedAt, UpdatedAt, DeletedAt
	Task       string
	IsComplete bool
}

// Definisi model User tanpa gorm.Model untuk variasi
type User struct {
	ID    uint   `gorm:"primaryKey"` // Kustomisasi primary key
	Name  string `gorm:"not null"`
	Email string `gorm:"unique;size:100"` // Kolom 'email' unique, tipe varchar(100)
	// Field untuk soft delete akan ditambahkan di branch selanjutnya
	// DeletedAt gorm.DeletedAt `gorm:"index"`
}

func main() {
	// --- Bagian 1: Koneksi Database ---
	log.Println("Attempting to connect to database...")

	// Koneksi ke database SQLite. File 'gorm.db' akan dibuat.
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	log.Println("Database connection successfully opened.")

	// Menutup koneksi saat fungsi main selesai (opsional, tapi praktik baik)
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get underlying DB instance: %v", err)
	}
	defer func() {
		log.Println("Closing database connection.")
		sqlDB.Close()
		log.Println("Database connection closed.")
	}()

	// --- Bagian 3: Auto Migrate ---
	// Kode AutoMigrate akan ditambahkan di branch selanjutnya.

	// --- Bagian 4: Operasi CRUD ---
	// Kode Create, Read, Update, Delete akan ditambahkan di branch selanjutnya.

	log.Println("Application finished.")
}
