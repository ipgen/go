// This library is the official implementation of the IPGen Spec
package ipgen

import (
	"fmt"
	"github.com/dchest/blake2b"
	"net"
)

// Generates an IP address
func IP(name, cidr string) (ip net.IP, err error) {
	ip, netwk, err := net.ParseCIDR(cidr)
	if err != nil {
		return
	}
	prefix, bits := netwk.Mask.Size()
	// We are not going to allow IP addresses with the /32 prefix for IPv4
	// addresses and those with a /128 prefix for IPv6 because they don't
	// leave any room for IP addresses to be generated.
	if prefix == bits {
		err = fmt.Errorf("%s is already a full IP address", ip)
		return
	}
	switch bits {
	case 128: // This an IPv6 address
		ip, err = ip6(name, ip, prefix)
	case 32: // This is an IPv4 address
		ip6prefix := 128 - 32 + prefix
		ip, err = ip6(name, ip.To16(), ip6prefix)
		if err != nil {
			return
		}
		ip = ip.To4()
	default:
		err = fmt.Errorf("Unsupported IP address version")
	}
	return
}

// Both IPv6 and IPv4 addresses are processed the same way.
// This is the function that does the heavy lifting.
func ip6(name string, address net.IP, prefix int) (ip net.IP, err error) {
	networkLen := prefix / 4
	var ipHash string
	// To get the `ipHash` we need to expand the IP address
	// ignoring any colons.
	for _, v := range address {
		ipHash += fmt.Sprintf("%02x", v)
	}
	networkHash := ipHash[0:networkLen]
	addressLen := 32 - networkLen
	// Blake2b hashes always have a total length that's
	// a multiple of 2.
	blakeLen := (addressLen / 2) + (addressLen % 2)
	addressHash := hash(name, uint8(blakeLen))
	ipHash = networkHash + addressHash
	// Format the IP hash to expanded IPv6
	ipAddr := fmt.Sprintf("%s:%s:%s:%s:%s:%s:%s:%s",
		ipHash[0:4],
		ipHash[4:8],
		ipHash[8:12],
		ipHash[12:16],
		ipHash[16:20],
		ipHash[20:24],
		ipHash[24:28],
		ipHash[28:32],
	)
	ip = net.ParseIP(ipAddr)
	return
}

// Generates an IPv6 subnet ID
func Subnet(name string) string {
	return hash(name, 2)
}

// Generates the blake2b hash
func hash(name string, outlen uint8) string {
	h, _ := blake2b.New(&blake2b.Config{Size: outlen})
	h.Write([]byte(name))
	return fmt.Sprintf("%x", h.Sum(nil))
}
