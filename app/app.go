package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/SaravananPitchaimuthu/Fruits/Fruits/domain"
	"github.com/SaravananPitchaimuthu/Fruits/Fruits/logger"
	"github.com/SaravananPitchaimuthu/Fruits/Fruits/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func SanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environmental variables are not defined")
	}

	if os.Getenv("host") == "" || os.Getenv("host_schema") == "" {
		log.Fatal("Environmental varibles are not defined for mysql host")
	}
}

func Start() {
	router := mux.NewRouter()
	logger.Info("Starting the Application ....")
	SanityCheck()
	//wiring
	client := getDbClient()
	fruitRepositoryDb := domain.NewFruitRepositoryDb(client)
	accountRepositoryDb := domain.NewAccountRepositoryDb(client)
	// accountRepositoryDb := domain.NewAccountRepositoryDb(client)
	//ch := CustomHandlers{service.NewFruitService(domain.NewFruitRepositoryStub())}
	ch := CustomHandlers{service.NewFruitService(fruitRepositoryDb)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDb)}

	router.HandleFunc("/fruits", ch.getAllFruits).Methods(http.MethodGet)
	router.HandleFunc("/fruits/{fruit_id:[0-9]+}", ch.getFruitById).Methods(http.MethodGet)
	router.HandleFunc("/fruits/{fruit_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	router.HandleFunc("/fruits/{fruit_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).Methods(http.MethodPost)
	localhost := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	http.ListenAndServe(fmt.Sprintf("%s:%s", localhost, port), router)
}

func getDbClient() *sqlx.DB {
	username := os.Getenv("username")
	password := os.Getenv("password")
	host := os.Getenv("host")
	port := os.Getenv("port")
	host_schema := os.Getenv("host_schema")

	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, host_schema)
	client, err := sqlx.Open("mysql", datasource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
