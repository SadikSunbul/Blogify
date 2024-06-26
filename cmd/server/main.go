package main

import (
	"github.com/SadikSunbul/Blogify/internal/handlers"
	"github.com/SadikSunbul/Blogify/internal/infrastructure/database"
	"github.com/SadikSunbul/Blogify/internal/infrastructure/repositories"
	"github.com/SadikSunbul/Blogify/internal/services"
	"gopkg.in/yaml.v2"
	"log"
	"net/http"
	"net/url"
	"os"
)

// Config yapılandırma dosyası
type Config struct {
	Database struct { // veritabanı bilgileri
		Connection string `yaml:"connection"` // veritabanı bağlantısı
	} `yaml:"database"` // veritabanı bilgileri
}

func main() {

	configFile, err := os.ReadFile("../db/config.yaml") // config dosyasını okur
	if err != nil {
		log.Fatalf("config.yaml dosyası okunamadı: %v", err) // hata döndürür
	}

	var config Config                         // yapılandırma dosyası
	err = yaml.Unmarshal(configFile, &config) // yapılandırma dosyasını okur
	if err != nil {                           // hata varsa
		log.Fatalf("Yapılandırma dosyası ayrıştırılamadı: %v", err)
	}

	conn, _ := url.Parse(config.Database.Connection)       // veritabanı bağlantısı
	conn.RawQuery = "sslmode=verify-ca;sslrootcert=ca.pem" // ssl bilgileri
	db := database.InitDB(conn.String())                   // veritabanını başlatır

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)
	postRepo := repositories.NewPostRepository(db)
	commentRepo := repositories.NewCommentRepository(db)
	likeRepo := repositories.NewLikeRepository(db)

	// Initialize services
	userService := services.NewUserService(userRepo)
	postService := services.NewPostService(postRepo)
	commentService := services.NewCommentService(commentRepo)
	likeService := services.NewLikeService(likeRepo)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)
	postHandler := handlers.NewPostHandler(postService)
	commentHandler := handlers.NewCommentHandler(commentService)
	likeHandler := handlers.NewLikeHandler(likeService)

	// Register routes
	http.HandleFunc("/users", userHandler.HandleUsers)
	http.HandleFunc("/posts", postHandler.HandlePosts)
	http.HandleFunc("/comments", commentHandler.HandleComments)
	http.HandleFunc("/likes", likeHandler.HandleLikes)

	// Start the server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
