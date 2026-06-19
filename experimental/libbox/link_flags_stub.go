//go:build !unix && !with_anytls_only

package libbox

import (
	"net"
)

func linkFlags(rawFlags uint32) net.Flags {
	panic("stub!")
}
