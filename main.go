package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type IndexContext struct {
	Ip string
}

func getIp(r *http.Request) string {
	ip := r.Header.Get("X-FORWARDED-FOR")
	if ip == "" {
		ip = r.RemoteAddr
	}
	return strings.Split(ip, ":")[0]
}

func index(template *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := template.Execute(w, IndexContext{Ip: getIp(r)})
		if err != nil {
			panic(err)
		}
	}
}

func simple(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, getIp(r))
}

func main() {
	t := template.Must(template.New("index.tmpl").ParseFiles("index.tmpl"))
	http.HandleFunc("/", index(t))
	http.HandleFunc("/simple", simple)

	http.ListenAndServe(":9000", nil)
}
