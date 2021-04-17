package main

import (
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

func main() {
	t := template.Must(template.New("index.tmpl").ParseFiles("index.tmpl"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := t.Execute(w, IndexContext{Ip: getIp(r)})
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":9000", nil)
}
