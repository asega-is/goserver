package main
import (
        "log"
        "fmt"
        "net/http"
        "os"
)
func Log(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
	handler.ServeHTTP(w, r)
    })
}

func handler(w http.ResponseWriter, r *http.Request) {
        h, _ := os.Hostname()
        fmt.Fprintf(w, "Hi there, I'm served from %s!", h)
}
func main() {

        http.HandleFunc("/", handler)
        http.ListenAndServe(":8484", Log(http.DefaultServeMux))
}
