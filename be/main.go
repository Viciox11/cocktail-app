package main

import (
	"flag"
	"github.com/ioartigiano/ioartigiano-be/internal/ingredients"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/cors"

	"github.com/ioartigiano/ioartigiano-be/internal/entity"

	"github.com/ioartigiano/ioartigiano-be/internal/version"

	"github.com/ioartigiano/ioartigiano-be/sqldb"
	"github.com/jinzhu/gorm"

	"github.com/ioartigiano/ioartigiano-be/internal/config"
	log "github.com/sirupsen/logrus"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

var (
	number string
	build  string
)

//var file = flag.String("config", "config/local.json", "path to the config file")

func main() {
	flag.Parse()

	config.SetLogConfig()

	cfg := config.Load()
	dbCfg := cfg.Database
	db, err := sqldb.ConnectToDB(dbCfg.Dialect(), dbCfg.ConnectionInfo())
	if err != nil {
		log.Fatalf("Error connecting to DB %s", err.Error())
	}

	db.LogMode(true)
	db.AutoMigrate(
		&entity.Drink{},
		&entity.Ingredient{},
	)

	db.Model(&entity.Drink{})//.AddForeignKey("product_id", "drinks(id)", "RESTRICT", "RESTRICT")
	db.Model(&entity.Ingredient{})//.AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")

	defer db.Close()

	log.Info(cfg)
	server := &http.Server{
		Addr:    cfg.Address + ":" + cfg.Port,
		Handler: buildHandler(db, cfg),
	}

	log.Info("Listening ", server.Addr)

	err = server.ListenAndServe()
	log.Fatalln(err)
}

// buildHandler sets up the HTTP routing and builds an HTTP handler.
func buildHandler(db *gorm.DB, cfg *config.Config) http.Handler {
	entityValidator := validator.New()
	//all APIs are under "/api/v1" path prefix

	router := mux.NewRouter()
	router.Use(loggingMiddleware)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8100"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization,Content-Type"},
		AllowedMethods:   []string{"GET,POST,PUT,DELETE,PATCH,OPTIONS"},
	})

	// all APIs under noAuthGroup does not need that user is authenticated
	noAuthGroup := router.PathPrefix("/api/v1").Subrouter()

	version.RegisterHandlers(noAuthGroup, number, build)
	ingredients.RegisterHandlers(noAuthGroup, ingredients.NewService(ingredients.NewRepository(db), entityValidator))

	// handler for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	router.Handle("/docs", sh).Methods(http.MethodGet)
	router.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	handler := c.Handler(router)
	return handler
}

//Example function middleware, TODO remove
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}