package main;
import (
    "fmt"
    "log"
    "net/http"
)
func main() {
    http.HandleFunc("/helloworld", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hello, World!")
    })
    fmt.Printf("Server running (port=3000), route: http://localhost:3000/helloworld\n")
    if err := http.ListenAndServe(":3000", nil); err != nil {
        log.Fatal(err)
    }
}