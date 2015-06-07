package verifycsrftoken

import (
    "testing"
    "github.com/codegangsta/negroni"
    "bytes"
    "net/http"
    "net/http/httptest"
    "strings"
)

func TestNewVerifyCsrfToken(t *testing.T) {

    response := httptest.NewRecorder()
    response.Body = new(bytes.Buffer)

    token := "xxx"
    vt := NewVerifyCsrfToken(token)
    n := negroni.New()
    n.Use(vt)

    req, err := http.NewRequest("GET", "http://localhost:3000/", nil)
    if err != nil {
        t.Error(err)
    }

    req.Header.Add("X-CSRF-TOKEN", token)

    n.ServeHTTP(response, req)
    cookieString := response.Header().Get("Set-Cookie")

    if strings.Index(cookieString, token) == -1 {
        t.Errorf("Expected response header set-cookie Match with %v ", token)
    }

}

