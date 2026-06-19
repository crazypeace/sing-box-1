//go:build !(darwin || linux) && !with_anytls_only

package libbox

import "os"

func getTunnelName(fd int32) (string, error) {
	return "", os.ErrInvalid
}
