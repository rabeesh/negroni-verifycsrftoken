package verifycsrftoken

import (
    "log"
    "net/http"
    "time"
    "os"
)

type SessionGet interface {
    Get(r *http.Request, key string) string
}

type VerifyCsrfToken struct {
    TokenKey string
    Session SessionGet
    Logger  *log.Logger
}


// A new instance of NewVerifyCsrfToken
func NewVerifyCsrfToken(key string, store SessionGet) *VerifyCsrfToken {
    return &VerifyCsrfToken {
        TokenKey: key,
        Logger:  log.New(os.Stdout, "[negroni verify csrftoken] ", 0),
        Session: store,
    }
}

func (ct *VerifyCsrfToken) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
    stoken := ct.Session.Get(r, ct.TokenKey)
    if ct.isReadRequest(r) || ct.TokensMatch(r, stoken) {
        ct.AddCookie(rw, stoken)
        next(rw, r)
        return
    }

    // throw new token mismatch exception
    ct.Logger.Printf("PANIC: %s", "Token mismatch")
    rw.WriteHeader(http.StatusForbidden);
}

// Check the HTTP request method is only for read
func (ct *VerifyCsrfToken) isReadRequest(r *http.Request) bool {
    verbs := []string{"HEAD", "GET", "OPTIONS"}
    for _, m := range(verbs) {
        if r.Method == m {
            return true
        }
    }
    return false
}

func (ct *VerifyCsrfToken) AddCookie(rw http.ResponseWriter, stoken string) {
    cookie := http.Cookie{
        Name : "XSRF-TOKEN",
        Value : stoken,
        Expires : time.Now().Add(240 * time.Minute),
    }
    http.SetCookie(rw, &cookie)
}

func (ct *VerifyCsrfToken) TokensMatch(r *http.Request, stoken string) bool {
    token := r.Header.Get("X-CSRF-TOKEN")
    if stoken == token {
        return true
    }
    return false
}
