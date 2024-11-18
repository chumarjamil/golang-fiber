package main

import (
	"fmt"
	db "go_lang_rest_api/config"
)

func main() {
	fmt.Println("Creating structure")

	db.Connect()
}
