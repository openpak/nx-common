package internalkey

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStrictAndLenient(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	t.Setenv("NX_INTERNAL_KEY", "")
	t.Setenv("OPENPAK_INTERNAL_KEY", "")
	if Strict(r) || !Lenient(r) {
		t.Fatal("no key: strict must refuse, lenient must allow")
	}
	t.Setenv("NX_INTERNAL_KEY", "k")
	if Strict(r) || Lenient(r) {
		t.Fatal("a missing header passed")
	}
	r.Header.Set(Header, "k")
	if !Strict(r) || !Lenient(r) {
		t.Fatal("the right key was refused")
	}
	ok := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(204) })
	rec := httptest.NewRecorder()
	Guard(ok).ServeHTTP(rec, httptest.NewRequest("GET", "/", nil))
	if rec.Code != 401 {
		t.Fatalf("guard without the header: %d", rec.Code)
	}
}
