// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package comdelivery

import (
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

func easyjson89020eaDecodeCodexInternalPkgCommentDeliveryRest(in *jlexer.Lexer, out *commentsResp) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "review":
			(out.Comment).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson89020eaEncodeCodexInternalPkgCommentDeliveryRest(out *jwriter.Writer, in commentsResp) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"review\":"
		out.RawString(prefix[1:])
		(in.Comment).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v commentsResp) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89020eaEncodeCodexInternalPkgCommentDeliveryRest(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *commentsResp) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89020eaDecodeCodexInternalPkgCommentDeliveryRest(l, v)
}
func easyjson89020eaDecodeCodexInternalPkgCommentDeliveryRest1(in *jlexer.Lexer, out *commentReq) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "movieId":
			out.MovieId = string(in.String())
		case "userId":
			out.UserId = string(in.String())
		case "reviewText":
			out.Content = string(in.String())
		case "reviewType":
			out.Type = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson89020eaEncodeCodexInternalPkgCommentDeliveryRest1(out *jwriter.Writer, in commentReq) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"movieId\":"
		out.RawString(prefix[1:])
		out.String(string(in.MovieId))
	}
	{
		const prefix string = ",\"userId\":"
		out.RawString(prefix)
		out.String(string(in.UserId))
	}
	{
		const prefix string = ",\"reviewText\":"
		out.RawString(prefix)
		out.String(string(in.Content))
	}
	{
		const prefix string = ",\"reviewType\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v commentReq) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89020eaEncodeCodexInternalPkgCommentDeliveryRest1(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *commentReq) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89020eaDecodeCodexInternalPkgCommentDeliveryRest1(l, v)
}
