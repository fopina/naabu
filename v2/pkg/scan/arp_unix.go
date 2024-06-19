//go:build linux || darwin

package scan

func init() {
	ArpRequestAsync = func(ip string) {}
}
