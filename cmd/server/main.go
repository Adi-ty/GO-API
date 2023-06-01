package main

import "fmt"

// Run - is going to be responsible
// for the instantiation and startup of our
// go application
func Run() error {
	fmt.Println("Starting up our application")
	return nil
}

func main() {
	fmt.Println("GO REST API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
