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

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Welcome to the home page!")
  })

  n := negroni.Classic()
  n.Use(verifycsrftoken.NewVerifyCsrfToken("xxx"))
  n.UseHandler(mux)
  n.Run(":3000")
}
~~~
