package sprig

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHostByName(t *testing.T) {
	tpl := `{{"www.google.com" | getHostByName}}`

	resolvedIP, _ := runRaw(tpl, nil)

	ip := net.ParseIP(resolvedIP)
	assert.NotNil(t, ip)
	assert.NotEmpty(t, ip)
}

func TestIsIPv6(t *testing.T) {
	tpl := `{{"2001:db8::68" | isIPv6}}`
	isIPv6, _ := runRaw(tpl, nil)
	assert.Equal(t, "true", isIPv6)

	tpl = `{{"2001:db8::68/64" | isIPv6}}`
	isIPv6, _ = runRaw(tpl, nil)
	assert.Equal(t, "false", isIPv6)

	tpl = `{{"192.168.1.100" | isIPv6}}`
	isIPv6, _ = runRaw(tpl, nil)
	assert.Equal(t, "false", isIPv6)
}

func TestIsIPv4(t *testing.T) {
	tpl := `{{"192.168.1.100" | isIPv4}}`
	isIPv4, _ := runRaw(tpl, nil)
	assert.Equal(t, "true", isIPv4)

	tpl = `{{"192.168.1.100/32" | isIPv4}}`
	isIPv4, _ = runRaw(tpl, nil)
	assert.Equal(t, "false", isIPv4)

	tpl = `{{"::FFFF:192.0.2.0" | isIPv4}}`
	isIPv4, _ = runRaw(tpl, nil)
	assert.Equal(t, "false", isIPv4)
}
