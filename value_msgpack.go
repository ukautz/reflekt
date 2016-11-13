package reflekt

import (
	"gopkg.in/vmihailenco/msgpack.v2"
)

func (this *Value) MarshalMsgpack() ([]byte, error) {
	return msgpack.Marshal(this.v)
}

func (this *Value) UnmarshalMsgpack(raw []byte) error {
	return msgpack.Unmarshal(raw, &this.v)
}