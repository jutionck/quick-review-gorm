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

// Tamabahn model baru
type Product struct {
	gorm.Model
	Code  string `gorm:"unique;size:10"`
	Price uint
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
	err = db.AutoMigrate(&Task{}, &User{}, &Product{})
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
	log.Println("Database auto migration finished. Tables created/updated.")
	// --- Re-Create Data for Demo ---
	log.Println("\n--- Re-Creating Data for Update Demo ---")
	db.Exec("DELETE FROM products")                              // Hati-hati! Menghapus semua data produk.
	db.Exec("DELETE FROM users")                                 // Hati-hati! Menghapus semua data user.
	db.Exec("DELETE FROM sqlite_sequence WHERE name='products'") // Reset auto-increment SQLite
	db.Exec("DELETE FROM sqlite_sequence WHERE name='users'")    // Reset auto-increment SQLite

	productsToCreate := []Product{
		{Code: "P001", Price: 50},
		{Code: "P002", Price: 75},
	}
	db.Create(&productsToCreate)

	usersToCreate := []User{
		{Name: "Budi Santoso", Email: "budi@example.com"},
		{Name: "Siti Aminah", Email: "siti@example.com"},
	}
	db.Create(&usersToCreate)
	log.Println("Demo data created for update.")

	// --- Bagian 4: Operasi CRUD - Update ---
	log.Println("\n--- Updating Data ---")

	// Update menggunakan Save: Ambil record, ubah field, lalu Save
	var productToUpdate Product
	result := db.First(&productToUpdate, 1) // Asumsikan produk ID 1 ada
	if result.Error != nil {
		log.Printf("Product with ID 1 not found for update (Save): %v\n", result.Error)
	} else {
		log.Printf("Product before update (Save): %+v\n", productToUpdate)
		productToUpdate.Price = 150        // Ubah harga
		result = db.Save(&productToUpdate) // Simpan perubahan
		if result.Error != nil {
			log.Printf("Failed to update product using Save: %v\n", result.Error)
		} else {
			log.Printf("Product updated successfully using Save. New Price: %d\n", productToUpdate.Price)
		}
	}

	// Update menggunakan Updates (map): Ambil record, lalu Updates dengan map
	var userToUpdate User
	result = db.Where("email = ?", "siti@example.com").First(&userToUpdate) // Asumsikan user ini ada
	if result.Error != nil {
		log.Printf("User with email siti@example.com not found for update (Updates): %v\n", result.Error)
	} else {
		log.Printf("User before update (Updates): %+v\n", userToUpdate)
		result = db.Model(&userToUpdate).Updates(map[string]interface{}{"Name": "Siti Aminah Updated", "Email": "siti.updated@example.com"}) // Update Nama dan Email
		if result.Error != nil {
			log.Printf("Failed to update user using Updates (map): %v\n", result.Error)
		} else {
			log.Printf("User updated successfully using Updates (map). New Name: %s, New Email: %s\n", userToUpdate.Name, userToUpdate.Email)
		}
	}

	// Update menggunakan Updates (struct): Ambil record, lalu Updates dengan struct
	var productToUpdate2 Product
	result = db.First(&productToUpdate2, 2) // Asumsikan produk ID 2 ada
	if result.Error != nil {
		log.Printf("Product with ID 2 not found for update (Updates struct): %v\n", result.Error)
	} else {
		log.Printf("Product before update (Updates struct): %+v\n", productToUpdate2)
		result = db.Model(&productToUpdate2).Updates(Product{Price: 90}) // Hanya update Price (field lain di struct Product{Price: 90} adalah zero-value)
		if result.Error != nil {
			log.Printf("Failed to update product using Updates (struct): %v\n", result.Error)
		} else {
			log.Printf("Product updated successfully using Updates (struct). New Price: %d\n", productToUpdate2.Price)
		}
	}

	log.Println("Application finished.")
}
