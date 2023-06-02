package main

import (
	"context"
	"fmt"

	"github.com/Adi-ty/GO-API/internal/comment"
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
	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	cmtService := comment.NewService(db)

	cmtService.PostComment(
		context.Background(),
		comment.Comment{
			ID:     "9a31bf83-28dc-4b8d-bf70-7d347a24ff2e",
			Slug:   "manual-test",
			Author: "Aditya",
			Body:   "Hello World!",
		},
	)

	fmt.Println(cmtService.GetComment(
		context.Background(),
		"9a31bf83-28dc-4b8d-bf70-7d347a24ff2e",
	))

	return nil
}

func main() {
	fmt.Println("GO REST API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
