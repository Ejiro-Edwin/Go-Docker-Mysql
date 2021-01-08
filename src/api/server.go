package api

import (
	"fmt"
	"github.com/ejiro-edwin/Blog/src/api/router"
	"log"
	"net/http"
)

func Run()  {
	fmt.Println("\n\tListening [::]:3000\n")
	r := router.New()
	log.Fatal(http.ListenAndServe(":3000", r))
}
