package controllers

import (
  "net/http"
  "fmt"
  "io/ioutil"
  // "log"
  "os"
  "encoding/json"

  _"github.com/khafido/simple-app-go/api/responses"
)


type Response struct {
  ID        uint32   `json:"id"`
  Nama      string   `json:"nama"`
  Username  string   `json:"username"`
  Password  string   `json:"password"`
  Foto      string   `json:"foto"`
}

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
  // responses.JSON(w, http.StatusOK, "<h1>Welcome To This Awesome API</h1>")
  fmt.Println("Starting the application...")
  response, err := http.Get("http://localhost:8080/users")

  if err != nil {
    fmt.Print(err.Error())
    os.Exit(1)
  }

  responseData, _ := ioutil.ReadAll(response.Body)
  fmt.Println(string(responseData))
  var users []Response
  json.Unmarshal(responseData, &users)
  fmt.Println(users)
}

func (server *Server) EditUser(w http.ResponseWriter, r *http.Request) {
  id, _ := r.URL.Query()["id"]
  fmt.Println("Starting the application..."+string(id[0]))
  response, err := http.Get("http://localhost:8080/users/"+string(id[0]))

  if err != nil {
    fmt.Print(err.Error())
    os.Exit(1)
  }

  responseData, _ := ioutil.ReadAll(response.Body)
  fmt.Println(string(responseData))
  var users Response
  json.Unmarshal(responseData, &users)
  fmt.Println(users)
}
