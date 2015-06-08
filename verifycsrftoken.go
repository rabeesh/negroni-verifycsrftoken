package verifycsrftoken

import (
    "log"
    "net/http"
    "time"
    "os"
)

type VerifyCsrfToken struct {
    Token string
    Logger  *log.Logger
}

// NewVerifyCsrfToken returns a new instance of NewVerifyCsrfToken
func NewVerifyCsrfToken(token string) *VerifyCsrfToken {
    return &VerifyCsrfToken{
        Logger:  log.New(os.Stdout, "[negroni verify csrftoken] ", 0),
        Token: token,
    }
}

func (ct *VerifyCsrfToken) ServeHTTP(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {

    if ct.TokensMatch(req) {
        ct.AddCookie(rw)
        next(rw, req)
        return
    }

    ct.Logger.Printf("PANIC: %s", "Token mismatch")
    next(rw, req)
}

func (ct *VerifyCsrfToken) AddCookie(rw http.ResponseWriter) {
    cookie := &http.Cookie{
        Name : "XSRF-TOKEN",
        Value : ct.Token,
        Path : "/",
        Expires : time.Now().Add(240 * time.Minute),
        HttpOnly : true,
    }
    http.SetCookie(rw, cookie)
}

func (ct *VerifyCsrfToken) TokensMatch(req *http.Request) bool {
    token := req.Header.Get("X-CSRF-TOKEN")
    if ct.Token == token {
        return true
    }
    return false
}
