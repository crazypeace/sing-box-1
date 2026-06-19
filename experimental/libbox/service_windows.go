//go:build !with_anytls_only

package libbox

import "os"

func dup(fd int) (nfd int, err error) {
	return 0, os.ErrInvalid
}
