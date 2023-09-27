package boltsink_test

import (
	"testing"

	pb "github.com/boxcolli/pepperlink/idl/gen/hello/v1"
	"google.golang.org/protobuf/encoding/protojson"
	// "github.com/stretchr/testify/assert"
)

func TestBoltSink(t *testing.T) {
	msg := pb.Hello{
		Name: "Alice",
	}
	t.Log(protojson.Marshal(&msg)) // just for log
}