package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/geekAshish/DriveDesk/driver"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	carService "github.com/geekAshish/DriveDesk/service/car"
	carStore "github.com/geekAshish/DriveDesk/store/car"

	engineService "github.com/geekAshish/DriveDesk/service/engine"
	engineStore "github.com/geekAshish/DriveDesk/store/engine"

	carHandler "github.com/geekAshish/DriveDesk/handler/car"
	engineHandler "github.com/geekAshish/DriveDesk/handler/engine"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	driver.InitDB()
	defer driver.CloseDB()

	db := driver.GetDB()
	carStore := carStore.New(db)
	carService := carService.NewCarService(carStore)

	engineStore := engineStore.New(db)
	engineService := engineService.NewEngineService(engineStore)

	carHandler := carHandler.NewCarHandler(carService)
	engineHandler := engineHandler.NewEngineHandler(engineService)

	router := mux.NewRouter()

	// Middleware
	// router.Use(middleware.AuthMiddleware)

	// schemaFile := "store/schema.sql"
	// if err := executeSchemaFile(db, schemaFile); err != nil {
	// 	log.Fatal("error while executing the schema file: ", err)
	// }

	router.HandleFunc("/cars/{id}", carHandler.GetCarById).Methods("GET")
	router.HandleFunc("/cars", carHandler.GetCarByBrand).Methods("GET")
	router.HandleFunc("/cars", carHandler.CreateCar).Methods("POST")
	router.HandleFunc("/cars/{id}", carHandler.UpdateCar).Methods("PUT")
	router.HandleFunc("/cars/{id}", carHandler.DeleteCar).Methods("DELETE")

	router.HandleFunc("/engine/{id}", engineHandler.GetEngineById).Methods("GET")
	router.HandleFunc("/engine", engineHandler.CreateEngine).Methods("POST")
	router.HandleFunc("/engine/{id}", engineHandler.UpdateEngine).Methods("PUT")
	router.HandleFunc("/engine/{id}", engineHandler.DeleteEngine).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("INVALID PORT NUMBER")
	}

	addr := fmt.Sprintf(":%s", port)
	log.Printf("Server listning on : %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}


func executeSchemaFile(db *sql.DB, filename string) error {
	sqlFile, err := os.ReadFile(filename);
	if err != nil {
		return err
	}

	_, err = db.Exec(string(sqlFile))
	if err != nil {
		return err
	}

	return nil
}