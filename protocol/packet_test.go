package protocol

import (
	"testing"

	"github.com/castyapp/libcasty-protocol-go/proto"
	pb "github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func TestPacket(t *testing.T) {
	t.Parallel()
	packet, err := NewPacket([]byte(""))
	assert.EqualError(t, err, "EOF")
	assert.Nil(t, packet)
}

func TestPingPongBytesToPacket(t *testing.T) {

	// Ping bytes packet
	pingBytesPacket := []byte{1, 0, 0, 0}

	packet, err := NewPacket(pingBytesPacket)
	assert.NoError(t, err)
	assert.NotNil(t, packet)
	assert.Equal(t, packet.EMsg, proto.EMSG_PING)

	// Pong bytes packet
	pongBytesPacket := []byte{2, 0, 0, 0}

	packet, err = NewPacket(pongBytesPacket)
	assert.NoError(t, err)
	assert.NotNil(t, packet)
	assert.Equal(t, packet.EMsg, proto.EMSG_PONG)
}

func TestLogonBytesToPacket(t *testing.T) {

	// Ping bytes packet
	logonBytesPacket := []byte{
		3, 0, 0, 0, 26, 18, 115, 117,
		112, 101, 114, 45, 115, 101, 99,
		117, 114, 101, 45, 116, 111, 107, 101, 110,
	}

	packet, err := NewPacket(logonBytesPacket)
	assert.NoError(t, err)
	assert.NotNil(t, packet)
	assert.Equal(t, packet.EMsg, proto.EMSG_LOGON)

	logonPacket := new(proto.LogOnEvent)

	err = packet.ReadProtoMsg(logonPacket)
	assert.NotNil(t, packet)

	assert.Equal(t, logonPacket.Token, []byte("super-secure-token"))
}

func PacketEncodingDecodingTest(t *testing.T, emsg proto.EMSG, message pb.Message) *Packet {

	buf, err := NewMsgProtobuf(emsg, message)
	assert.NoError(t, err)

	packet, err := NewPacket(buf.Bytes())
	assert.NoError(t, err)

	assert.True(t, packet.IsProto)
	assert.Equal(t, packet.EMsg, emsg)

	return packet
}

type EventTest struct {
	name       string
	event      pb.Message
	emsg       proto.EMSG
	readerTest func(e EventTest, p *Packet)
}

func (e EventTest) Run(t *testing.T) {
	t.Parallel()
	packet := PacketEncodingDecodingTest(t, e.emsg, e.event)
	e.readerTest(e, packet)
}

func TestEvents(t *testing.T) {
	tests := []EventTest{
		{
			name: "LogOnEvent",
			event: &proto.LogOnEvent{
				Username: []byte("random-username"),
				Token:    []byte("random-token"),
				Password: []byte("random-password"),
			},
			emsg: proto.EMSG_LOGON,
			readerTest: func(e EventTest, p *Packet) {
				logonEvent := new(proto.LogOnEvent)
				err := p.ReadProtoMsg(logonEvent)
				assert.NoError(t, err)
				payload := e.event.(*proto.LogOnEvent)
				assert.Equal(t, logonEvent.Username, payload.Username)
				assert.Equal(t, logonEvent.Password, payload.Password)
				assert.Equal(t, logonEvent.Token, payload.Token)
			},
		},
		{
			name: "TheaterLogOnEvent",
			event: &proto.TheaterLogOnEvent{
				Room:  []byte("a-room"),
				Token: []byte("random-token"),
			},
			emsg: proto.EMSG_LOGON,
			readerTest: func(e EventTest, p *Packet) {
				logonEvent := new(proto.TheaterLogOnEvent)
				err := p.ReadProtoMsg(logonEvent)
				assert.NoError(t, err)
				payload := e.event.(*proto.TheaterLogOnEvent)
				assert.Equal(t, logonEvent.Room, payload.Room)
				assert.Equal(t, logonEvent.Token, payload.Token)
			},
		},
		{
			name: "ChatMsgEvent",
			event: &proto.ChatMsgEvent{
				Message: []byte("random-message"),
				Reciever: &proto.User{
					Id: "random-user-id-234o2734827304",
				},
			},
			emsg: proto.EMSG_NEW_CHAT_MESSAGE,
			readerTest: func(e EventTest, p *Packet) {
				logonEvent := new(proto.ChatMsgEvent)
				err := p.ReadProtoMsg(logonEvent)
				assert.NoError(t, err)
				payload := e.event.(*proto.ChatMsgEvent)
				assert.Equal(t, logonEvent.Message, payload.Message)
				assert.Equal(t, logonEvent.Reciever.Id, payload.Reciever.Id)
			},
		},
		{
			name: "TheaterPlay",
			event: &proto.TheaterVideoPlayer{
				TheaterId:   "random-theater-id-238472308hs",
				CurrentTime: 23.2342342,
				State:       proto.TheaterVideoPlayer_PLAYING,
				UserId:      "random-user-id-238wdfisjd",
			},
			emsg: proto.EMSG_PLAYING,
			readerTest: func(e EventTest, p *Packet) {
				logonEvent := new(proto.TheaterVideoPlayer)
				err := p.ReadProtoMsg(logonEvent)
				assert.NoError(t, err)
				payload := e.event.(*proto.TheaterVideoPlayer)
				assert.Equal(t, logonEvent.TheaterId, payload.TheaterId)
				assert.Equal(t, logonEvent.UserId, payload.UserId)
				assert.Equal(t, logonEvent.State, payload.State)
				assert.Equal(t, logonEvent.CurrentTime, payload.CurrentTime)
			},
		},
		{
			name: "TheaterPause",
			event: &proto.TheaterVideoPlayer{
				TheaterId:   "random-theater-id-238472308hs",
				CurrentTime: 23.2342342,
				State:       proto.TheaterVideoPlayer_PAUSED,
				UserId:      "random-user-id-238wdfisjd",
			},
			emsg: proto.EMSG_THEATER_PAUSE,
			readerTest: func(e EventTest, p *Packet) {

				logonEvent := new(proto.TheaterVideoPlayer)
				err := p.ReadProtoMsg(logonEvent)
				assert.NoError(t, err)

				payload := e.event.(*proto.TheaterVideoPlayer)

				assert.Equal(t, logonEvent.TheaterId, payload.TheaterId)
				assert.Equal(t, logonEvent.UserId, payload.UserId)
				assert.Equal(t, logonEvent.State, payload.State)
				assert.Equal(t, logonEvent.CurrentTime, payload.CurrentTime)

			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, test.Run)
	}
}
