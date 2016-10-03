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
