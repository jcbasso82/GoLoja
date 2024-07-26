package main

import (
	"fmt"
	"loja/routes"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("failed to load environment variables: %v", err)
	}
	routes.CarregaRotas()
	port := os.Getenv("PORT")
	fmt.Println("Porta:", port)
	http.ListenAndServe(":"+port, nil)
	//http.ListenAndServeTLS(":443", "/etc/letsencrypt/live/loja.juliobasso.com.br/fullchain.pem", "/etc/letsencrypt/live/loja.juliobasso.com.br/privkey.pem", nil)
}
