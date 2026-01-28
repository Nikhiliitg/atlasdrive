// package main

// import (
// 	"log"
// 	"net/http"
// 	"database/sql"

// 	httpadapter "github.com/Nikhiliitg/atlasdrive/internal/adapters/http"
// 	"github.com/Nikhiliitg/atlasdrive/internal/adapters/postgres"
// 	fileapp "github.com/Nikhiliitg/atlasdrive/internal/application/file"
// 	folderapp "github.com/Nikhiliitg/atlasdrive/internal/application/folder"
// 	_ "github.com/lib/pq"
// 	redisadapter "github.com/Nikhiliitg/atlasdrive/internal/adapters/redis"
// )

// func main() {
// 	db, err := sql.Open("postgres", "postgres://localhost/atlasdrive?sslmode=disable")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Write side
// 	folderRepo := postgres.NewFolderRepo(db)
// 	fileRepo := postgres.NewFileRepo(db)

// 	// Read side (Postgres)
// 	folderQuery := postgres.NewFolderQueryRepo(db)

// 	// Application handlers
// 	createFolderHandler := folderapp.NewCreateFolderHandler(folderRepo)
// 	redisClient := redisadapter.NewClient()
// 	folderCache := redisadapter.NewFolderCache(redisClient)

// 	cachedQuery := folderapp.NewCachedFolderQuery(
// 		folderQuery,
// 		folderCache,
// 	)

// 	listFolderHandler := folderapp.NewListFolderContentsHandler(cachedQuery)

// 	createFileHandler := fileapp.NewCreateFileHandler(fileRepo)

// 	// HTTP
// 	handler := httpadapter.NewHandler(
// 		createFolderHandler,
// 		listFolderHandler,
// 		createFileHandler,
// 	)

// 	authHandler := httpadapter.NewAuthHandler(db)
// 	router := httpadapter.NewRouter(handler, authHandler)


// 	log.Println("Server running on :8080")
// 	http.ListenAndServe(":8080", router)
// }
package main

import (
	"database/sql"
	"log"
	"net/http"

	httpadapter "github.com/Nikhiliitg/atlasdrive/internal/adapters/http"
	postgresadapter "github.com/Nikhiliitg/atlasdrive/internal/adapters/postgres"
	redisadapter "github.com/Nikhiliitg/atlasdrive/internal/adapters/redis"

	fileapp "github.com/Nikhiliitg/atlasdrive/internal/application/file"
	folderapp "github.com/Nikhiliitg/atlasdrive/internal/application/folder"

	_ "github.com/lib/pq"
)

func main() {
	// --------------------
	// Database
	// --------------------
	db, err := sql.Open("postgres", "postgres://localhost/atlasdrive?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// --------------------
	// Write repositories
	// --------------------
	folderRepo := postgresadapter.NewFolderRepo(db)
	fileRepo := postgresadapter.NewFileRepo(db)

	// --------------------
	// Read repository (Postgres)
	// --------------------
	folderQuery := postgresadapter.NewFolderQueryRepo(db)

	// --------------------
	// Redis
	// --------------------
	redisClient := redisadapter.NewClient()
	folderCache := redisadapter.NewFolderCache(redisClient)

	// --------------------
	// Application handlers
	// --------------------
	createFolderHandler := folderapp.NewCreateFolderHandler(folderRepo)

	listFolderHandler := folderapp.NewListFolderContentsHandler(
		folderQuery,   // FolderQuery
		folderCache,   // Redis cache
	)

	createFileHandler := fileapp.NewCreateFileHandler(fileRepo)

	// --------------------
	// HTTP handlers
	// --------------------
	handler := httpadapter.NewHandler(
		createFolderHandler,
		listFolderHandler,
		createFileHandler,
	)

	authHandler := httpadapter.NewAuthHandler(db)
	router := httpadapter.NewRouter(handler, authHandler)

	// --------------------
	// Start server
	// --------------------
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
