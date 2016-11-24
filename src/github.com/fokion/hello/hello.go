package main

import ("fmt";"net/http")

func main() {
    fmt.Printf("hello, world\n")
    //we can run the app by go run src/github.com/fokion/hello/hello.go 
    //we can install it by doing go install github.com/fokion/hello
    //
    //the func  in the HandleFunc is the handler of the request on the endpoint
    http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
      fmt.Fprintf(w,"hello from server")
    })
    fmt.Println(http.ListenAndServe("localhost:8181",nil))
}