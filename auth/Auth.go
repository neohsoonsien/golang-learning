package auth

import (
    "fmt"
    "net/http"

    auth "github.com/abbot/go-http-auth"
)

func Secret(user, realm string) string {
    if user == "john" {
        // password is "hello"
        return "$1$dlPL2MqE$oQmn16q49SqdmhenQuNgs1"
    }
    return realm
}

func handle(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
    fmt.Fprintf(w, "<html><body><h1>Hello, %s!</h1></body></html>", r.Username)
}

func Auth() {
    authenticator := auth.NewBasicAuthenticator("localhost", Secret)
    http.HandleFunc("/", authenticator.Wrap(handle))
    http.ListenAndServe(":8080", nil)
}