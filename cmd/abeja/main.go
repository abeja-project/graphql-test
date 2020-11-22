package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abeja-project/graphql-test/internal/abeja"
)

const addr = ":8000"

func main() {
	a := abeja.New()
	fmt.Printf("Listen on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, a))
}
