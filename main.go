package main

import (
	"fmt"
	"vaqua/config"
	"vaqua/db"
	"vaqua/handler"
	"vaqua/repository"
	"vaqua/routes"
	"vaqua/services"
)

func main() {
	// load up variables
	config.LoadEnv()

	// connect to database
	db.InitDb()

	// initialise the repo
	userRepo := &repository.UserRepo{}
	transferRequestRepo := &repository.TransferRequestRepo{}
	transactionRepo := &repository.TransactionRepo{}

	// initialise the service
	userService := &services.UserService{Repo: userRepo}
	transferRequestService := &services.TransferRequestService{Repo: transferRequestRepo}
	transactionService := &services.TransactionService{Repo: transactionRepo}

	// initialise the handler
	userHandler := &handler.UserHandler{Service: userService}
	transferRequestHandler := &handler.TransferRequestHandler{Service: transferRequestService}
	transactionHandler := &handler.TransactionHandler{Service: transactionService}

	// define routes
	db := db.InitDb()
	r := routes.SetupRouter(userHandler, transferRequestHandler, transactionHandler, db)

	// start the server
	fmt.Println("server is running on localhost:8080...")
	r.Run(":8080")

}
