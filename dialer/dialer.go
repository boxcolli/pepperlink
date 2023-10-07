package dialer

import "github.com/boxcolli/go-transistor/types"

// Dialer takes care for grpc clients and their connections.
type Dialer interface {
	Dial(m types.Member)
	Add(m types.Member)
	Delete(m types.Member)
}
