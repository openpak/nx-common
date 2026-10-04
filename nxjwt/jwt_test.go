package nxjwt

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"
)

func TestSignParseRoundTrip(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 2048)
	tok := Sign(map[string]any{"alg": "RS256", "typ": "JWT"}, map[string]any{"sub": "abc", "n": 1}, key)
	c, err := Parse(tok, &key.PublicKey)
	if err != nil || Str(c, "sub") != "abc" {
		t.Fatalf("parse = %v, %v", c, err)
	}
	other, _ := rsa.GenerateKey(rand.Reader, 2048)
	if _, err := Parse(tok, &other.PublicKey); err == nil {
		t.Fatal("a token verified against a key that did not sign it")
	}
	if _, err := Parse(tok, nil); err != nil {
		t.Fatalf("decode without a key: %v", err)
	}
	if len(UUID4()) != 36 || len(RandHex(8)) != 16 {
		t.Fatal("id helpers changed length")
	}
}
