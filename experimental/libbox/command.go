//go:build !with_anytls_only

package libbox

const (
	CommandLog int32 = iota
	CommandStatus
	CommandGroup
	CommandClashMode
	CommandConnections
)
