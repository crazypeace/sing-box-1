//go:build !windows && !with_anytls_only

package libbox

import "syscall"

func dup(fd int) (nfd int, err error) {
	return syscall.Dup(fd)
}
