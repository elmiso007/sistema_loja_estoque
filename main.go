package main

import (
	"net/http"

	"elmiso007/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
