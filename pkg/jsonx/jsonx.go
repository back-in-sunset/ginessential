package jsonx

import (
	"io"

	"github.com/goccy/go-json"
)

// alias
var (
	Marshal                 = json.Marshal
	MarshalNoEscape         = json.MarshalNoEscape
	MarshalContext          = json.MarshalContext
	MarshalWithOption       = json.MarshalWithOption
	MarshalIndent           = json.MarshalIndent
	MarshalIndentWithOption = json.MarshalIndentWithOption

	Unmarshal           = json.Unmarshal
	UnmarshalContext    = json.UnmarshalContext
	UnmarshalWithOption = json.UnmarshalWithOption
	UnmarshalNoEscape   = json.UnmarshalNoEscape

	NewDecode = json.NewDecoder
	Encoder   = json.NewEncoder
)

// EncoderOption ..
type EncoderOption func(*json.Encoder)

// EncoderWithIndentOpt ..
func EncoderWithIndentOpt(prefix string, indent string) func(*json.Encoder) {
	return func(e *json.Encoder) {
		e.SetIndent(prefix, indent)
	}
}

// EncodeW encode v to w
func EncodeW(w io.Writer, v interface{}, opts ...EncoderOption) error {
	encoder := Encoder(w)

	for _, opt := range opts {
		opt(encoder)
	}

	err := encoder.Encode(v)
	if err != nil {
		return err
	}

	return nil
}

// DecodeR decode r to v
func DecodeR(r io.Reader, v interface{}) error {
	return NewDecode(r).Decode(v)
}
