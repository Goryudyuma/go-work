package main

import (
	auth "github.com/abbot/go-http-auth"
	"net/http"

	"golang.org/x/net/http2"
)

func secret(user, realm string) string {
	if user == "SS" {
		return "9fb34736a82f452e274ab7da86576925"
	}
	return ""
}

func main() {
	var srv http.Server
	srv.Addr = ":12345"
	authenticator := auth.NewDigestAuthenticator("063.jp", secret)
	http.HandleFunc("/", authenticator.Wrap(func(res http.ResponseWriter, req *auth.AuthenticatedRequest) {
		http.FileServer(http.Dir("/var/www/html/SS/")).ServeHTTP(res, &req.Request)
	}))
	http2.ConfigureServer(&srv, &http2.Server{})
	srv.ListenAndServeTLS("../key/server.crt", "../key/server.key")
}
