package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abeja-project/graphql-test/internal/abeja"
	"github.com/abeja-project/graphql-test/internal/database"
)

const addr = ":8000"

func main() {
	db := database.New()
	a := abeja.New(db)
	fmt.Printf("Listen on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, a))
}
