package main

import (
	"context"
	"fmt"

	"github.com/Adi-ty/GO-API/internal/db"
)

// Run - is going to be responsible
// for the instantiation and startup of our
// go application
func Run() error {
	fmt.Println("Starting up our application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to database")
		return err
	}
	if err := db.Ping(context.Background()); err != nil {
		return err
	}
	fmt.Println("succesfully connected and pinged database")

	return nil
}

func main() {
	fmt.Println("GO REST API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
