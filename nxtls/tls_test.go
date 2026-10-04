package nxtls

import (
	"crypto/tls"
	"testing"
)

func TestFixedAndOnDemandCertificates(t *testing.T) {
	conf := ServerConfig("[test]", []string{"a.example", "b.example"}, t.TempDir(), false)
	fixed, err := conf.GetCertificate(&tls.ClientHelloInfo{ServerName: "b.example"})
	if err != nil {
		t.Fatal(err)
	}
	empty, _ := conf.GetCertificate(&tls.ClientHelloInfo{})
	if fixed != empty {
		t.Fatal("no SNI must get the fixed certificate")
	}
	probe, _ := conf.GetCertificate(&tls.ClientHelloInfo{ServerName: "c.example"})
	again, _ := conf.GetCertificate(&tls.ClientHelloInfo{ServerName: "c.example"})
	if probe == fixed || probe != again {
		t.Fatal("an unknown name must get one on-demand certificate of its own")
	}
	if conf.ClientAuth != tls.NoClientCert || conf.MinVersion != tls.VersionTLS10 {
		t.Fatal("listener policy changed")
	}
}
