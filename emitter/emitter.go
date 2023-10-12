package emitter

import (
	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/types"
)

// Emitter receives new messages and emits through outlet.
type Emitter interface {
	Work(w *io.StreamWriter) error
	Stop()
	Emit(m *types.Message)	// goroutine safe; use channel to implement queue
}
