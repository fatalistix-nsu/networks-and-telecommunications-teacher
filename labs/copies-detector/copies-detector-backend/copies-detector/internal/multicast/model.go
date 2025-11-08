package multicast

type MessageType byte

const (
	JoinOrRefresh MessageType = iota
	Leave
)

const (
	typeSize       int = 1
	nameLengthSize int = 2
	maxNameLength  int = 500
)
