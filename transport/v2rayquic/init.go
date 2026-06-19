//go:build with_quic && !with_anytls_only

package v2rayquic

import "github.com/sagernet/sing-box/transport/v2ray"

func init() {
	v2ray.RegisterQUICConstructor(NewServer, NewClient)
}
