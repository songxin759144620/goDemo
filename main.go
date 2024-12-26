package main
import (
	"fmt"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello,world!")
}

func main() {
	http.HandleFunc("/login",login)
	fmt.Println("Starting server at :8080")
	if err := http.ListenAndServe(":8080", nil);err != nil {
		fmt.Println("Error starting server:",err)
	}
}