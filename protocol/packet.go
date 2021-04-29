package protocol

import (
	"bytes"
	"encoding/binary"

	pb "github.com/castyapp/libcasty-protocol-go/proto"
	"github.com/golang/protobuf/proto"
)

const (
	ProtoMask uint32 = 0x80000000
	EMsgMask         = ^ProtoMask
)

// Represents an incoming, partially unread message.
type Packet struct {
	EMsg    pb.EMSG
	IsProto bool
	Data    []byte
}

func NewPacket(data []byte) (*Packet, error) {
	var rawEMsg uint32
	dataBuffer := bytes.NewBuffer(data)
	err := binary.Read(dataBuffer, binary.LittleEndian, &rawEMsg)
	if err != nil {
		return nil, err
	}
	emsg := newEMSG(rawEMsg)
	return &Packet{
		EMsg:    emsg,
		IsProto: emsg != pb.EMSG_INVALID,
		Data:    dataBuffer.Bytes(),
	}, nil
}

func newEMSG(e uint32) pb.EMSG {
	return pb.EMSG(e & EMsgMask)
}

func (p *Packet) ReadProtoMsg(body proto.Message) error {
	if err := proto.Unmarshal(p.Data, body); err != nil {
		return err
	}
	return nil
}
