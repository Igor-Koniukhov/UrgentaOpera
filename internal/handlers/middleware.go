package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"todo/internal/configs"
	"todo/services/jwtoken"
)

type MiddlewareI interface {
	CORS(next http.Handler) http.Handler
	JSON(next http.Handler) http.Handler
	AuthCheck(next http.HandlerFunc) http.HandlerFunc
}

type Middleware struct {
	App *configs.AppConfig
}

func NewMiddleware(app *configs.AppConfig) *Middleware {
	return &Middleware{App: app}
}

func (m *Middleware) CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setupResponseCORS(&w, r)
		next.ServeHTTP(w, r)
	})
}

func (m *Middleware) JSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func setupResponseCORS(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Add("Access-Control-Allow-Origin", "http://127.0.0.1:8081")
	(*w).Header().Add("Access-Control-Allow-Credentials", "true")
	(*w).Header().Add("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
	(*w).Header().Add("Access-Control-Allow-Headers", "Access-Control-Allow-Headers,"+
		" Origin, Authorization ,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method,"+
		" Access-Control-Request-Headers")
	//catch preflight request
	if r.Method == "OPTIONS" {
		(*w).WriteHeader(http.StatusOK)
		return
	}
}

func (m *Middleware) AuthCheck(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookies := r.Cookies()
		if len(cookies) >= 0 {
			for _, ck := range cookies {
				if ck.Name == "Authorization" {
					if ck.Value == "" {
						fmt.Println("Authorization header is empty!")
						return
					}
					if ck.Value != "" {
						tokenString, err := jwtoken.GetTokenFromBearerString(ck.Value)
						if err != nil {
							fmt.Println(err)
							return
						}
						creds, err := jwtoken.ValidateToken(tokenString, jwtoken.AccessSecret)
						if err != nil {
							m.App.InfoString = make(map[string]string)
							m.App.InfoString["jwt_status"] = "expired"
							err := json.NewEncoder(w).Encode(&m.App.InfoString)
							if err != nil {
								log.Fatalln(err)
							}
							return
						}
						ctx := context.WithValue(r.Context(), "user_id", creds.ID)
						next.ServeHTTP(w, r.WithContext(ctx))
					} else {
						next.ServeHTTP(w, r)
					}
				}
			}
		}
	}
}
