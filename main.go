package main

import (
	"fmt"
	"vaqua/config"
	"vaqua/db"
	"vaqua/handler"
	"vaqua/repository"
	"vaqua/routes"
	"vaqua/services"

	"log"
	"vaqua/redis"
)

func main() {
	// load up variables
	config.LoadEnv()

	// connect to database
	db.InitDb()

    // connect to Redis
	if err := redis.ConnectRedis(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	fmt.Println("Connected to Redis successfully!")

	// initialise the repo
	userRepo := &repository.UserRepo{}
	transferRequestRepo := &repository.TransferRepo{}
	transactionRepo := &repository.TransactionRepo{}

	// initialise the service
	userService := &services.UserService{Repo: userRepo}
	transferService := &services.TransferServices{Repo: *transferRequestRepo}
	transactionService := &services.TransactionService{Repo: transactionRepo}

	// initialise the handler
	userHandler := &handler.UserHandler{Service: userService}
	transferRequestHandler := &handler.TransferHandler{Service: transferService}
	transactionHandler := &handler.TransactionHandler{Service: transactionService}

	// define routes
	db := db.InitDb()
	r := routes.SetupRouter(userHandler, transferRequestHandler, transactionHandler, db)

	// start the server
	fmt.Println("server is running on localhost:8080...")
	r.Run(":8080")

}


//todo 
// 1. finish gettransaction and gettransaction by id in this branch (dashboard_income)- test endpoint - git add . - git commit
// 2. checkout to a new branch from this branch (git checkout -b 'branchname')- git push
// 3. go to github create a PR from the new branch to main 

