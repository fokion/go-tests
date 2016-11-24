package main

import ("fmt";"net/http";"html/template")
// Serve html templates . Look at golang.org/pkg/html/template
func main() {
  //the Must will absorb errors and halt execution if it cannot parse the template
    templates := template.Must(template.ParseFiles("templates/index.html"))
    //the func  in the HandleFunc is the handler of the request on the endpoint
    http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
      if err:= templates.ExecuteTemplate(w,"index.html",nil); err != nil{
        http.Error(w,err.Error(),http.StatusInternalServerError)
      }
    })
    fmt.Println(http.ListenAndServe("localhost:8181",nil))
}