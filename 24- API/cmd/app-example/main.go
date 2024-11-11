package main

import (
	"app-example/configs"
	_ "app-example/docs"
	"app-example/internal/infra/database"
	"app-example/internal/infra/webserver/handlers"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/mattn/go-sqlite3"
)

// @title App Example API
// @version 1.0
// @description This is a sample server for an app example.
// @termsOfService http://example.com/terms/
// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8082
// @BasePath /
// @securityDefinitions.api_key ApiKeyAuth
// @in header
// @name Authorization
// @tokenUrl /users/token
func main() {
	config := configs.NewConfig()
	db, err := sql.Open("sqlite3", "./your-database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	database.CreateUserTables(db)
	database.CreateProductTables(db)

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)
	userDb := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDb)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.WithValue("jwt", config.TokenAuth()))
	router.Use(middleware.WithValue("jwtExpiresIn", config.JWTExpiresIn()))

	router.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth()))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.ListAllProducts)
		r.Get("/{id}", productHandler.GetProductByID)
		r.Put("/{id}", productHandler.UpdateProductByID)
		r.Delete("/{id}", productHandler.DeleteProductByID)
	})

	router.Post("/users", userHandler.CreateUser)
	router.Post("/users/token", userHandler.GetJwtToken)

	router.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8082/docs/doc.json")))
	http.ListenAndServe(":8082", router)
}
