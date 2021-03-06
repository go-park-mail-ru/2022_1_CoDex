// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package gendelivery

import (
	domain "codex/internal/pkg/domain"
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson89020eaDecodeCodexInternalPkgGenresDeliveryRest(in *jlexer.Lexer, out *Genres) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Genres, 0, 2)
			} else {
				*out = Genres{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 domain.Genre
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson89020eaEncodeCodexInternalPkgGenresDeliveryRest(out *jwriter.Writer, in Genres) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Genres) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89020eaEncodeCodexInternalPkgGenresDeliveryRest(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Genres) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89020eaDecodeCodexInternalPkgGenresDeliveryRest(l, v)
}
