package main

import (
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
	DeletedAt gorm.DeletedAt `gorm:"index"` // Soft Delete
	// Relasi One-to-Many: Satu User punya banyak Post
	Posts []Post // Slice of Post. GORM akan mencocokkan UserID di tabel posts
}

// Model baru untuk demonstrasi relasi
type Post struct {
	gorm.Model // ID, CreatedAt, UpdatedAt, DeletedAt
	Title      string
	Body       string
	UserID     uint // Kolom foreign key yang menghubungkan Post ke User. Harus ada.
	// Relasi Many-to-One: Banyak Post dimiliki oleh satu User
	// Field ini opsional, digunakan GORM untuk memuat data User saat Preload
	User User
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
	// AutoMigrate semua model, termasuk Post
	err = db.AutoMigrate(&Product{}, &User{}, &Post{})
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
	log.Println("Database auto migration finished. Tables created/updated.")

	// --- Re-Create Data for Demo (minimal untuk demo relasi) ---
	log.Println("\n--- Re-Creating Minimal Data for Relasi Demo ---")
	db.Unscoped().Exec("DELETE FROM products")
	db.Unscoped().Exec("DELETE FROM users")
	db.Unscoped().Exec("DELETE FROM posts") // Hapus data post juga
	db.Exec("DELETE FROM sqlite_sequence WHERE name='products'")
	db.Exec("DELETE FROM sqlite_sequence WHERE name='users'")
	db.Exec("DELETE FROM sqlite_sequence WHERE name='posts'")

	user1 := User{Name: "User A", Email: "user_a@example.com"}
	user2 := User{Name: "User B", Email: "user_b@example.com"}
	db.Create(&[]User{user1, user2})
	log.Printf("Created users with IDs: %d, %d\n", user1.ID, user2.ID)

	// Buat beberapa post, kaitkan dengan user
	postsToCreate := []Post{
		{Title: "Post Pertama User A", Body: "Ini isi post pertama", UserID: user1.ID},
		{Title: "Post Kedua User A", Body: "Ini isi post kedua", UserID: user1.ID},
		{Title: "Post Pertama User B", Body: "Ini isi post user B", UserID: user2.ID},
	}
	db.Create(&postsToCreate)
	log.Println("Created demo posts linked to users.")

	// --- Bagian 4: Operasi CRUD & Relasi ---
	log.Println("\n--- Demonstrating Basic Relasi (Preload) ---")

	// Ambil User pertama dan PRELOAD semua Post terkait
	var userWithPosts User
	// db.First(&userWithPosts, 1) // Jika pakai ini, user.Posts akan kosong

	result := db.Preload("Posts").First(&userWithPosts, user1.ID) // Load User1 dan Posts-nya
	if result.Error != nil {
		log.Printf("Failed to find user with posts: %v\n", result.Error)
	} else {
		log.Printf("Found User '%s' with Posts:\n", userWithPosts.Name)
		log.Printf("Number of posts loaded: %d\n", len(userWithPosts.Posts))
		for i, post := range userWithPosts.Posts {
			log.Printf("  Post %d: ID=%d, Title='%s'\n", i+1, post.ID, post.Title)
		}
	}

	// Ambil semua User dan PRELOAD semua Post terkait
	var allUsersWithPosts []User
	result = db.Preload("Posts").Find(&allUsersWithPosts)
	if result.Error != nil {
		log.Printf("Failed to find all users with posts: %v\n", result.Error)
	} else {
		log.Println("\nFound All Users with Posts:")
		for _, user := range allUsersWithPosts {
			log.Printf("- User '%s' has %d posts.\n", user.Name, len(user.Posts))
		}
	}

	// Mencari Post dan PRELOAD informasi User-nya (relasi Many-to-One)
	var postWithUser Post
	result = db.Preload("User").First(&postWithUser, postsToCreate[0].ID) // Ambil post pertama dan load User-nya
	if result.Error != nil {
		log.Printf("Failed to find post with user: %v\n", result.Error)
	} else {
		log.Printf("\nFound Post '%s' by User '%s' (loaded via Preload):\n", postWithUser.Title, postWithUser.User.Name)
	}

	log.Println("Application finished.")
}
