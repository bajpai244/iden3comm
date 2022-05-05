package packers

import (
	"encoding/json"
	core "github.com/iden3/go-iden3-core"
	"github.com/iden3/iden3comm"
	"github.com/pkg/errors"
)

// MediaTypePlainMessage is media type for plain message
const MediaTypePlainMessage iden3comm.MediaType = "application/iden3comm-plain-json"

// PlainMessagePacker is simple packer that doesn't use encryption / encoding
type PlainMessagePacker struct {
}

// Pack returns packed message to transport envelope
func (p *PlainMessagePacker) Pack(payload []byte, _ *core.ID) ([]byte, error) {

	var msgMap map[string]interface{}
	err := json.Unmarshal(payload, &msgMap)
	if err != nil {
		return nil, err
	}
	msgMap["typ"] = MediaTypePlainMessage
	return json.Marshal(msgMap)
}

// Unpack returns unpacked message from transport envelope
func (p *PlainMessagePacker) Unpack(envelope []byte) (iden3comm.Iden3Message, error) {

	var msg iden3comm.BasicMessage
	err := json.Unmarshal(envelope, &msg)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &msg, err
}

// MediaType for iden3comm
func (p *PlainMessagePacker) MediaType() iden3comm.MediaType {
	return MediaTypePlainMessage
}
