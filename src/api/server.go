package api

import (
	"fmt"
	"github.com/ejiro-edwin/Blog/src/api/router"
	"github.com/ejiro-edwin/Blog/src/config"
	"log"
	"net/http"
)

func Run()  {
	fmt.Printf("\n\tListening [::]:%d\n",config.PORT)
	listen(config.PORT)
}

func listen(port int)  {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",port), r))

}
