package nxapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClientSendsTheKeyAndMapsStatuses(t *testing.T) {
	t.Setenv("NX_INTERNAL_KEY", "secret")
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Internal-Key") != "secret" {
			w.WriteHeader(401)
			return
		}
		switch r.URL.Path {
		case "/internal/identity/users/u1":
			json.NewEncoder(w).Encode(User{ID: "u1", NaID: "na1", PID: 7})
		case "/internal/penne/notifications":
			var n Notification
			json.NewDecoder(r.Body).Decode(&n)
			json.NewEncoder(w).Encode(Submission{Accepted: true, Outcome: "accepted", ID: n.ID})
		default:
			w.WriteHeader(404)
		}
	}))
	defer srv.Close()
	c := New(srv.URL + "/")
	u, err := c.User(context.Background(), "u1")
	if err != nil || u.NaID != "na1" || u.PID != 7 {
		t.Fatalf("user = %+v, %v", u, err)
	}
	if _, err := c.User(context.Background(), "nobody"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("missing user: %v", err)
	}
	s, err := c.Notify(context.Background(), Notification{ID: "ev1", UserID: "u1", Kind: "friend_deleted"})
	if err != nil || !s.Accepted || s.ID != "ev1" {
		t.Fatalf("notify = %+v, %v", s, err)
	}
	c.Key = "wrong"
	if _, err := c.User(context.Background(), "u1"); err == nil {
		t.Fatal("a wrong key was not an error")
	}
}
