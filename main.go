package main

import (
	"errors"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// --- Bagian 2: Definisi Model ---
type Product struct {
	gorm.Model        // Sudah termasuk DeletedAt
	Code       string `gorm:"unique;size:10"`
	Price      uint
}

type User struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"not null"`
	Email     string         `gorm:"unique"`
	DeletedAt gorm.DeletedAt `gorm:"index"` // Ditambahkan untuk demo Soft Delete
}

func main() {
	// --- Bagian 1: Koneksi Database ---
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

	// --- Bagian 3: Auto Migrate ---
	log.Println("Running auto migration...")
	// AutoMigrate lagi setelah menambahkan DeletedAt ke User
	err = db.AutoMigrate(&Product{}, &User{})
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
	log.Println("Database auto migration finished. Tables created/updated.")

	// --- Re-Create Data for Demo ---
	log.Println("\n--- Re-Creating Data for Delete Demo ---")
	// Hati-hati! Menghapus semua data produk dan user
	db.Unscoped().Exec("DELETE FROM products") // Gunakan Unscoped untuk menghapus permanen jika soft delete aktif
	db.Unscoped().Exec("DELETE FROM users")
	db.Exec("DELETE FROM sqlite_sequence WHERE name='products'")
	db.Exec("DELETE FROM sqlite_sequence WHERE name='users'")

	productToDeleteHard := Product{Code: "HARDDEL", Price: 10}
	db.Create(&productToDeleteHard)
	log.Printf("Product created for hard delete: %+v\n", productToDeleteHard)

	userToDeleteSoft := User{Name: "User Soft Delete", Email: "softdelete@example.com"}
	db.Create(&userToDeleteSoft)
	log.Printf("User created for soft delete: %+v\n", userToDeleteSoft)

	userToDeletePermanent := User{Name: "User Perm Delete", Email: "permdelete@example.com"}
	db.Create(&userToDeletePermanent)
	log.Printf("User created for permanent delete: %+v\n", userToDeletePermanent)

	log.Println("Demo data created for delete.")

	// --- Bagian 4: Operasi CRUD - Delete ---
	log.Println("\n--- Deleting Data ---")

	// Hard Delete Product (karena Product pakai gorm.Model, tapi kita hapus by ID)
	// Atau bisa db.Delete(&Product{}, productToDeleteHard.ID)
	result := db.Delete(&productToDeleteHard)
	if result.Error != nil {
		log.Printf("Failed to hard delete product: %v\n", result.Error)
	} else {
		log.Printf("Product ID %d hard deleted. Rows affected: %d\n", productToDeleteHard.ID, result.RowsAffected)
	}

	// Coba cari produk yang sudah di hard delete (seharusnya tidak ketemu)
	var deletedProductCheck Product
	result = db.First(&deletedProductCheck, productToDeleteHard.ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Println("Product not found after hard delete (expected).")
	} else if result.Error != nil {
		log.Printf("Error finding product after hard delete: %v\n", result.Error)
	} else {
		log.Println("Product found after hard delete (UNEXPECTED!).")
	}

	log.Println("") // Newline for clarity

	// Soft Delete User (User pakai DeletedAt)
	result = db.Delete(&userToDeleteSoft) // GORM akan SET DeletedAt
	if result.Error != nil {
		log.Printf("Failed to soft delete user: %v\n", result.Error)
	} else {
		log.Printf("User ID %d soft deleted. Rows affected: %d\n", userToDeleteSoft.ID, result.RowsAffected)
		// userToDeleteSoft.DeletedAt sekarang terisi timestamp
		log.Printf("User after soft delete: %+v\n", userToDeleteSoft)
	}

	// Coba cari user yang sudah di soft delete (seharusnya tidak ketemu by default)
	var softDeletedUserCheck User
	result = db.First(&softDeletedUserCheck, userToDeleteSoft.ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Println("User not found after soft delete (expected, default query excludes soft deleted).")
	} else if result.Error != nil {
		log.Printf("Error finding user after soft delete: %v\n", result.Error)
	} else {
		log.Println("User found after soft delete (UNEXPECTED!).")
	}

	// Cari user yang sudah di soft delete MENGGUNAKAN Unscoped()
	var unscopedSoftDeletedUserCheck User
	result = db.Unscoped().First(&unscopedSoftDeletedUserCheck, userToDeleteSoft.ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Println("User not found using Unscoped (UNEXPECTED!).")
	} else if result.Error != nil {
		log.Printf("Error finding user using Unscoped: %v\n", result.Error)
	} else {
		log.Printf("User found using Unscoped: %+v\n", unscopedSoftDeletedUserCheck) // DeletedAt field should be populated
	}

	// Menghapus permanen (ketika soft delete aktif) menggunakan Unscoped()
	log.Println("\nPerforming permanent delete...")
	result = db.Unscoped().Delete(&userToDeletePermanent)
	if result.Error != nil {
		log.Printf("Failed to permanent delete user: %v\n", result.Error)
	} else {
		log.Printf("User ID %d permanently deleted. Rows affected: %d\n", userToDeletePermanent.ID, result.RowsAffected)
	}

	// Coba cari user yang sudah di permanent delete (seharusnya tidak ketemu)
	var permanentDeletedUserCheck User
	result = db.Unscoped().First(&permanentDeletedUserCheck, userToDeletePermanent.ID) // Gunakan Unscoped untuk memastikan cek
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Println("User not found after permanent delete (expected).")
	} else if result.Error != nil {
		log.Printf("Error finding user after permanent delete: %v\n", result.Error)
	} else {
		log.Println("User found after permanent delete (UNEXPECTED!).")
	}

	log.Println("Application finished.")
}
