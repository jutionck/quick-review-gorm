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
	log.Println("Running auto migration...")
	// AutoMigrate akan membuat/memperbarui skema tabel berdasarkan model.
	err = db.AutoMigrate(&Task{}, &User{})
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
	log.Println("Database auto migration finished. Tables created/updated.")

	// --- Bagian 4: Operasi CRUD - Create ---
	log.Println("\n--- Creating Data ---")

	// Tambah produk baru
	newTask := Task{Task: "Makan", IsComplete: true}
	result := db.Create(&newTask) // Mengembalikan *gorm.DB

	if result.Error != nil {
		log.Printf("Failed to create task: %v\n", result.Error)
	} else {
		log.Printf("Task created successfully with ID: %d (RowsAffected: %d)\n", newTask.ID, result.RowsAffected)
	}

	// Tambah beberapa user sekaligus
	usersToCreate := []User{
		{Name: "Budi Santoso", Email: "budi@example.com"},
		{Name: "Siti Aminah", Email: "siti@example.com"},
		{Name: "Agus Dharma", Email: "agus@example.com"},
	}
	result = db.Create(&usersToCreate)
	if result.Error != nil {
		log.Printf("Failed to create users: %v\n", result.Error)
	} else {
		log.Printf("Users created successfully. (RowsAffected: %d)\n", result.RowsAffected)
		// ID dari setiap user di slice juga akan terisi setelah operasi create
		// log.Printf("Created Users (with IDs): %+v\n", usersToCreate) // Bisa dicetak
	}

	log.Println("Application finished.")
}
