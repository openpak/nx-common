package nxhttp

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSplitListenerKeepsConsoleTrafficAwayFromInternalRoutes(t *testing.T) {
	console := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(201) })
	internal := http.NewServeMux()
	internal.HandleFunc("/internal/x", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(202) })
	h := SplitListener(console, internal)
	for _, c := range []struct {
		host, path string
		want       int
	}{
		{"hac.lp1.savanna.srv.nintendo.net", "/internal/x", 201}, // a console host never reaches /internal
		{"accounts.nintendo.com", "//double//slash", 201},        // and is never path-cleaned (no 301)
		{"openpak-nx-shop:20070", "/internal/x", 202},
		{"openpak-nx-shop:20070", "/top", 404},
	} {
		r := httptest.NewRequest("GET", "http://x"+c.path, nil)
		r.Host = c.host
		w := httptest.NewRecorder()
		h.ServeHTTP(w, r)
		if w.Code != c.want {
			t.Errorf("%s%s = %d, want %d", c.host, c.path, w.Code, c.want)
		}
	}
}
