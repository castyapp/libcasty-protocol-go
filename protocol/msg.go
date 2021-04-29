package protocol

import (
	"bytes"
	"encoding/binary"
	"net"

	pb "github.com/castyapp/libcasty-protocol-go/proto"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/golang/protobuf/proto"
)

func NewMsgProtobuf(eMsg pb.EMSG, body proto.Message) (buffer *bytes.Buffer, err error) {
	buffer = new(bytes.Buffer)

	// Writing emsg
	if err := binary.Write(buffer, binary.LittleEndian, pb.EMSG(uint32(eMsg))); err != nil {
		return nil, err
	}

	if body != nil {

		// Wrinting body if not nil
		body, err := proto.Marshal(body)
		if err != nil {
			return nil, err
		}

		if _, err := buffer.Write(body); err != nil {
			return nil, err
		}

	}

	return buffer, nil
}

func BrodcastMsgProtobuf(conn net.Conn, eMsg pb.EMSG, body proto.Message) (err error) {
	buf, err := NewMsgProtobuf(eMsg, body)
	if err != nil {
		return
	}
	if err = wsutil.WriteServerMessage(conn, ws.OpBinary, buf.Bytes()); err != nil {
		return
	}
	return
}
