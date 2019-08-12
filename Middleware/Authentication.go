package Middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"strings"
)

var Authentication = func(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		notAuth := []string{"/login"}
		requestPath := r.URL.Path

		for _, value := range notAuth {
			if value == requestPath {
				h.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		splitted := strings.Split(tokenHeader, " ")

		if len(splitted) != 2 {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		e := godotenv.Load()

		if e != nil {
			fmt.Print(e)
		}

		tokenString := splitted[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, e error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected sigin method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("CLAVE_SECRETA")), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims["Usuario"])

			if claims["pass"] == os.Getenv("CLAVE_TOKENS") {
				h.ServeHTTP(w, r)
				return
			}
		} else {
			fmt.Println("hola")
			fmt.Println(err)
		}
	})
}