package cors

import (
	"fmt"
	"net/http"
	"time"
)


func MiddleWareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Before handler; middleware start")
		start := time.Now()
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "POST,GET,OPTIONS,PUT, DELETE")
		w.Header().Add("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding")
		handler.ServeHTTP(w, r)
		fmt.Printf("Middleware finished; %s", time.Since(start))
	})
}