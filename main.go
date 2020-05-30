package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jhonmike/transaction/account"
	"github.com/jhonmike/transaction/config"
	"github.com/jhonmike/transaction/model"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	cfg := config.MustReadFromEnv()

	db := model.NewDatabase(cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPass, cfg.DbBase)

	r := mux.NewRouter()
	account.MakeAccountHandlers(r, db)

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%s", cfg.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
