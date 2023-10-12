package channel

import (
	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/types"
)

type channelStreamReader struct {
	ch <-chan *types.Message
}

// Read implements io.StreamReader.
func (r *channelStreamReader) Read() (*types.Message, error) {
	m, ok := <-(r.ch)
	if !ok {
		return nil, io.ErrClosed
	}
	return m, nil
}

func NewChannelStreamReader(ch <-chan *types.Message) io.StreamReader {
	return &channelStreamReader{
		ch: ch,
	}
}
