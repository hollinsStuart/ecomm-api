package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	repo "github.com/hollinsStuart/ecomm-api/internal/adapters/postgresql/sqlc"
	"github.com/hollinsStuart/ecomm-api/internal/orders"
	"github.com/hollinsStuart/ecomm-api/internal/products"
	"github.com/jackc/pgx/v5"
)

type application struct {
	config config
	// logger
	db *pgx.Conn
}

// mount
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// middlewares
	r.Use(middleware.RequestID) // rate limiter
	r.Use(middleware.RealIP)    // rate limiter, tracing, analytics
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer) // for crashes

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Good!"))
		if err != nil {
			return
		}
	})

	productService := products.NewService(repo.New(app.db))
	productsHandler := products.NewHandler(productService)
	r.Get("/products", productsHandler.ListProducts)

	orderService := orders.NewService(repo.New(app.db), app.db)
	ordersHandler := orders.NewHandler(orderService)
	r.Post("/orders", ordersHandler.PlaceOrder)

	return r
}

// run
func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Minute,
	}
	log.Printf("Server started at %s", app.config.addr)
	return srv.ListenAndServe()
}

type config struct {
	addr string
	port string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
