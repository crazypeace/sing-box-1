//go:build !with_gvisor && !with_anytls_only

package wireguard

import "github.com/sagernet/sing-tun"

func newStackDevice(options DeviceOptions) (Device, error) {
	return nil, tun.ErrGVisorNotIncluded
}

func newSystemStackDevice(options DeviceOptions) (Device, error) {
	return nil, tun.ErrGVisorNotIncluded
}
