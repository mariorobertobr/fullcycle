package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/mariorobertobr/codebank/infrastructure/repository"
	"github.com/mariorobertobr/codebank/usecase"
	"log"
)

func main() {
	fmt.Println("hello")
	db := setupDb()
	defer db.Close()

	//cc := domain.NewCreditCard()
	//cc.Number = "1234"
	//cc.Name = "mario"
	//cc.ExpirationYear = 2021
	//cc.ExpirationMonth = 7
	//cc.CVV = 123
	//cc.Limit = 1000
	//cc.Balance = 0
	//
	//repo := repository.NewTransactionRepositoryDb(db)
	//err := repo.CreateCreditCard(*cc)
	//if err != nil {
	//	fmt.Println(err)
	//}
}

func setupUseCase(db *sql.DB) usecase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDb(db)
	usecase := usecase.NewUseCaseTransaction(transactionRepository)

	return usecase
}

func setupDb() *sql.DB {
	psqlinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"db", "5432", "postgres", "root", "codebank")

	db, err := sql.Open("postgres", psqlinfo)
	if err != nil {
		log.Fatalln("error connecting to db")
	}
	return db
}
