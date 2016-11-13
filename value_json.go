package reflekt

import (
	"encoding/json"
)

func (this *Value) MarshalJSON() ([]byte, error) {
	return json.Marshal(this.v)
}

func (this *Value) UnmarshalJSON(raw []byte) error {
	return json.Unmarshal(raw, &this.v)
}