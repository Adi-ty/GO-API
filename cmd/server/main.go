package main

import (
	"fmt"

	"github.com/Adi-ty/GO-API/internal/comment"
	"github.com/Adi-ty/GO-API/internal/db"
	transportHttp "github.com/Adi-ty/GO-API/internal/transport/http"
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

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("GO REST API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
