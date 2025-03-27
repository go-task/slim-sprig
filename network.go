package sprig

import (
	"math/rand"
	"net"
	"strings"
)

func getHostByName(name string) string {
	addrs, _ := net.LookupHost(name)
	// TODO: add error handing when release v3 comes out
	return addrs[rand.Intn(len(addrs))]
}

func isIPv6(ip string) bool {
	netIP := net.ParseIP(ip)
	return netIP != nil && netIP.To4() == nil
}

func isIPv4(input string) bool {
	netIP := net.ParseIP(input)
	return netIP != nil && netIP.To4() != nil && !strings.Contains(input, ":")
}
