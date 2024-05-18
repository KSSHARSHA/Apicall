package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/chi/v5"
	
)


func main() {
router := chi.NewRouter()
router.Use(middleware.Logger)

router.Get("/hello", basicHandler)



server := &http.Server{
	Addr: ":3000",
	Handler: router,
}
err := server.ListenAndServe()

if err != nil{
	fmt.Println("Couldn't connect to server", err)

}

}

func basicHandler (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello! World"))
}