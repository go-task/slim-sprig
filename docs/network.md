# Network Functions

Sprig network manipulation functions.

## getHostByName

The `getHostByName` receives a domain name and returns the ip address.

```
getHostByName "www.google.com" would return the corresponding ip address of www.google.com
```

## isIPv4

The `isIPv4` returns `true` if a string is an IPv4 address.

```
isIPv4 "192.168.1.100" returns `true`
isIPv4 "fe80::1ff:fe23:4567:890a" returns `false`
```

## isIPv6

The `isIPv6` returns `true` if a string is an IPv6 address.

```
isIPv4 "192.168.1.100" returns `false`
isIPv4 "fe80::1ff:fe23:4567:890a" returns `true`
```
