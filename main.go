package main

import (
	"errors"
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

	// --- Bagian 4: Operasi CRUD - Create (dibuat lagi untuk memastikan data ada saat Read) ---
	// Hapus/comment kode create dari branch sebelumnya jika ingin mulai bersih setiap run
	// atau biarkan jika ingin data terakumulasi. Untuk demo Read, sebaiknya data dibuat dulu.

	log.Println("\n--- Re-Creating Data for Read Demo ---")
	// Hapus data lama (opsional, untuk demo berulang)
	// db.Exec("DELETE FROM tasks")
	// db.Exec("DELETE FROM users") // Hati-hati dengan ID auto-increment
	// atau drop table jika ingin bersih total
	// db.Migrator().DropTable(&Task{}, &User{})
	// db.AutoMigrate(&Task{}, &User{}) // Migrate lagi setelah drop

	// Tambah beberapa produk
	productsToCreate := []Product{
		{Code: "P001", Price: 50},
		{Code: "P002", Price: 75},
		{Code: "P003", Price: 120},
	}
	db.Create(&productsToCreate)

	// Tambah beberapa user
	usersToCreate := []User{
		{Name: "Budi Santoso", Email: "budi@example.com"},
		{Name: "Siti Aminah", Email: "siti@example.com"},
		{Name: "Agus Dharma", Email: "agus@example.com"},
	}
	db.Create(&usersToCreate)
	log.Println("Demo data created.")

	// --- Bagian 4: Operasi CRUD - Read ---
	log.Println("\n--- Reading Data ---")

	// Baca satu produk berdasarkan Primary Key (ID)
	var productByID Product
	result := db.First(&productByID, 1) // Ambil produk dengan ID 1
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Println("Product with ID 1 not found.")
		} else {
			log.Printf("Error finding product by ID: %v\n", result.Error)
		}
	} else {
		log.Printf("Found Product by ID 1: %+v\n", productByID)
	}

	// Baca satu user berdasarkan kondisi Where
	var userByEmail User
	result = db.Where("email = ?", "siti@example.com").First(&userByEmail)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Println("User with email siti@example.com not found.")
		} else {
			log.Printf("Error finding user by email: %v\n", result.Error)
		}
	} else {
		log.Printf("Found User by Email: %+v\n", userByEmail)
	}

	// Baca semua user
	var allUsers []User
	result = db.Find(&allUsers)
	if result.Error != nil {
		log.Printf("Failed to find all users: %v\n", result.Error)
	} else {
		log.Printf("Found %d users:\n", len(allUsers))
		for _, user := range allUsers {
			log.Printf("- %+v\n", user)
		}
	}

	// Baca produk dengan kondisi (harga > 50)
	var expensiveProducts []Product
	result = db.Where("price > ?", 50).Find(&expensiveProducts)
	if result.Error != nil {
		log.Printf("Failed to find expensive products: %v\n", result.Error)
	} else {
		log.Printf("Found %d expensive products (price > 50):\n", len(expensiveProducts))
		for _, p := range expensiveProducts {
			log.Printf("- %+v\n", p)
		}
	}

	// Baca produk dengan kondisi IN
	var specificProducts []Product
	result = db.Where("code IN ?", []string{"P001", "P003"}).Find(&specificProducts)
	if result.Error != nil {
		log.Printf("Failed to find specific products by code: %v\n", result.Error)
	} else {
		log.Printf("Found %d specific products (codes P001, P003):\n", len(specificProducts))
		for _, p := range specificProducts {
			log.Printf("- %+v\n", p)
		}
	}

	// Baca user diurutkan berdasarkan nama, ambil 1 saja (Limit)
	var firstUserByName User
	result = db.Order("name asc").First(&firstUserByName) // First() akan otomatis Limit(1)
	if result.Error != nil {
		log.Printf("Failed to find first user by name: %v\n", result.Error)
	} else {
		log.Printf("First user by name (ordered asc, First): %+v\n", firstUserByName)
	}

	// Baca semua produk diurutkan berdasarkan harga menurun (Order)
	var productsSortedDesc []Product
	result = db.Order("price desc").Find(&productsSortedDesc)
	if result.Error != nil {
		log.Printf("Failed to find products sorted desc: %v\n", result.Error)
	} else {
		log.Printf("Found %d products sorted by price desc:\n", len(productsSortedDesc))
		for _, p := range productsSortedDesc {
			log.Printf("- %+v\n", p)
		}
	}

	// Pagination: Ambil 1 user, lewati 1 user (Limit + Offset)
	var paginatedUsers []User
	result = db.Limit(1).Offset(1).Find(&paginatedUsers)
	if result.Error != nil {
		log.Printf("Failed to find paginated users: %v\n", result.Error)
	} else {
		log.Printf("Found %d paginated users (Limit 1, Offset 1):\n", len(paginatedUsers))
		for _, u := range paginatedUsers {
			log.Printf("- %+v\n", u)
		}
	}

	log.Println("Application finished.")
}
