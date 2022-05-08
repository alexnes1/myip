package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/alexnes1/myip/mainpage"
)

func simple(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, mainpage.GetIp(r))
}

func main() {
	indexPage := mainpage.New()
	http.HandleFunc("/", indexPage.GetHandler())
	http.HandleFunc("/txt", simple)

	http.ListenAndServe(os.Getenv("MYIP_ADDR"), nil)
}
