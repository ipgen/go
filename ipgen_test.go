package ipgen

import "testing"

func TestIPIsValid(t *testing.T) {
	ip, err := IP("c0a010fb-2632-40cb-a105-90297cba567a", "fd52:f6b0:3162::/48")
	if err != nil {
		t.Fatal(err)
	}
	if ip == nil {
		t.Fatal("Was expecting an IP address but none was produced")
	}
}

func TestV4In6(t *testing.T) {
	ip, err := IP("c0a010fb-2632-40cb-a105-90297cba567a", "::10.0.0.0/8")
	if err != nil {
		t.Fatal(err)
	}
	if ip.String() != "c9:31ec:e663:766c:2e4:826e:2391:6a3c" {
		t.Fatalf("Expected: %s Got: %s\n", "c9:31ec:e663:766c:2e4:826e:2391:6a3c", ip)
	}
}
