package verifycsrftoken

import (
	"bytes"
	"github.com/codegangsta/negroni"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type SsManger struct {
	token string
}

func (s *SsManger) Get(r *http.Request, k string) (token string) {
	return s.token
}

func TestNewVerifyCsrfToken(t *testing.T) {

	response := httptest.NewRecorder()
	response.Body = new(bytes.Buffer)

	token := "xxx"
	session := &SsManger{token: token}

	vt := NewVerifyCsrfToken("token", session)

	n := negroni.New()
	n.Use(vt)

	r, err := http.NewRequest("GET", "http://localhost:3000/", nil)
	if err != nil {
		t.Error(err)
	}

	// test for read request
	r.Method = "GET"
	r.Header.Add("X-CSRF-TOKEN", token)

	n.ServeHTTP(response, r)
	cookieString := response.Header().Get("Set-Cookie")

	if strings.Index(cookieString, token) == -1 {
		t.Errorf("Expected response header set-cookie Match with %v ", token)
	}

	// rest for status forbidden
	response = httptest.NewRecorder()
	r, err = http.NewRequest("GET", "http://localhost:3000/", nil)
	if err != nil {
		t.Error(err)
	}

	// test for read request
	r.Method = "POST"
	r.Header.Add("X-CSRF-TOKEN", "blaaa")

	n.ServeHTTP(response, r)

	if response.Code != 403 {
		t.Errorf("Expected response header set-cookie Match with %v ", token)
	}
}
