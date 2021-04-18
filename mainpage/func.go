package mainpage

import (
	"net/http"
	"strings"
)

func GetIp(r *http.Request) string {
	ip := r.Header.Get("X-FORWARDED-FOR")
	if ip == "" {
		ip = r.RemoteAddr
	}
	return strings.Split(ip, ":")[0]
}
