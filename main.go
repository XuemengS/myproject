package main

import(
        "fmt"
        "log"
        "net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hello, welcome to the unofficial site of Strawberry Switchblade!")
}

func main(){
        mux := http.NewServeMux()
        mux.HandleFunc("/", handler)
        log.Fatal(http.ListenAndServe(":3030", mux))
}

