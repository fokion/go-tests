package main

import ("fmt";"net/http")

func main() {
    fmt.Printf("hello, world\n")
    //we can run the app by go run src/github.com/fokion/hello/hello.go 
    //we can install it by doing go install github.com/fokion/hello
    beyondHelloStuff()
    setupServer()
    
}
func beyondHelloStuff(){
  var x int //variable declaration
  x = 3 //simple assignment
  
  y := 5 //declare and assign 
  
  sum, prod := multipleReturns(x,y)
  fmt.Printf("%v and prod %v ",sum,prod)
}
//function can have parameters and multiple return values...
//so the x , y are the arguments and the sum and prod is the signature of what
//is returned and all are type int
func multipleReturns(x, y int)(sum , prod int){
  return x + y, x * y
}
func setupServer(){
  //the func  in the HandleFunc is the handler of the request on the endpoint
  http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
    fmt.Fprintf(w,"hello from server")
  })
  fmt.Println(http.ListenAndServe("localhost:8181",nil))
}