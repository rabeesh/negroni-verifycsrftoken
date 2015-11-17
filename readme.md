# negroni verify csrftoken

Negroni middleware/handler for csrf token verification

## Usage


~~~ go
package main

import (
    "fmt"
    "github.com/codegangsta/negroni"
    "github.com/rabeesh/negroni-verifycsrftoken"
    "net/http"
)

// Example of session manager
type SessionManager struct {}

func (m *SessionManager) Get(r *http.Request, k string) string {
    // for token "csrftoken", return hoyyy
    return "hoyyy"
}

func main() {

  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Welcome to the home page!")
  })

  session := &SessionManager{}
  n := negroni.Classic()

  n.Use(verifycsrftoken.NewVerifyCsrfToken("csrftoken", session))
  n.UseHandler(mux)
  n.Run(":8080")
}
~~~
