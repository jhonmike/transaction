package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jhonmike/transaction/account"
	"github.com/jhonmike/transaction/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	cfg := config.MustReadFromEnv()

	dbCfg := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbBase, cfg.DbPass)
	db, err := gorm.Open("postgres", dbCfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&account.Account{}, &account.OperationType{}, &account.Transaction{})

	r := mux.NewRouter()
	account.MakeAccountHandlers(r)

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%s", cfg.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
