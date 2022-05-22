// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package domain

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

func easyjsonB2ba5d8eDecodeCodexInternalPkgDomain(in *jlexer.Lexer, out *MovieSummary) {
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
		case "href":
			out.Href = string(in.String())
		case "poster":
			out.Poster = string(in.String())
		case "title":
			out.Title = string(in.String())
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
func easyjsonB2ba5d8eEncodeCodexInternalPkgDomain(out *jwriter.Writer, in MovieSummary) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"href\":"
		out.RawString(prefix[1:])
		out.String(string(in.Href))
	}
	{
		const prefix string = ",\"poster\":"
		out.RawString(prefix)
		out.String(string(in.Poster))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MovieSummary) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB2ba5d8eEncodeCodexInternalPkgDomain(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MovieSummary) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB2ba5d8eDecodeCodexInternalPkgDomain(l, v)
}
func easyjsonB2ba5d8eDecodeCodexInternalPkgDomain1(in *jlexer.Lexer, out *MovieResponse) {
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
		case "movie":
			(out.Movie).UnmarshalEasyJSON(in)
		case "related":
			if in.IsNull() {
				in.Skip()
				out.Related = nil
			} else {
				in.Delim('[')
				if out.Related == nil {
					if !in.IsDelim(']') {
						out.Related = make([]MovieSummary, 0, 1)
					} else {
						out.Related = []MovieSummary{}
					}
				} else {
					out.Related = (out.Related)[:0]
				}
				for !in.IsDelim(']') {
					var v1 MovieSummary
					(v1).UnmarshalEasyJSON(in)
					out.Related = append(out.Related, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "reviews":
			if in.IsNull() {
				in.Skip()
				out.Comments = nil
			} else {
				in.Delim('[')
				if out.Comments == nil {
					if !in.IsDelim(']') {
						out.Comments = make([]Comment, 0, 0)
					} else {
						out.Comments = []Comment{}
					}
				} else {
					out.Comments = (out.Comments)[:0]
				}
				for !in.IsDelim(']') {
					var v2 Comment
					(v2).UnmarshalEasyJSON(in)
					out.Comments = append(out.Comments, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "reviewex":
			out.ReviewExist = string(in.String())
		case "userrating":
			out.UserRating = string(in.String())
		case "collectionsInfo":
			if in.IsNull() {
				in.Skip()
				out.CollectionsInfo = nil
			} else {
				in.Delim('[')
				if out.CollectionsInfo == nil {
					if !in.IsDelim(']') {
						out.CollectionsInfo = make([]CollectionInfo, 0, 2)
					} else {
						out.CollectionsInfo = []CollectionInfo{}
					}
				} else {
					out.CollectionsInfo = (out.CollectionsInfo)[:0]
				}
				for !in.IsDelim(']') {
					var v3 CollectionInfo
					(v3).UnmarshalEasyJSON(in)
					out.CollectionsInfo = append(out.CollectionsInfo, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonB2ba5d8eEncodeCodexInternalPkgDomain1(out *jwriter.Writer, in MovieResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"movie\":"
		out.RawString(prefix[1:])
		(in.Movie).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"related\":"
		out.RawString(prefix)
		if in.Related == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v4, v5 := range in.Related {
				if v4 > 0 {
					out.RawByte(',')
				}
				(v5).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"reviews\":"
		out.RawString(prefix)
		if in.Comments == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v6, v7 := range in.Comments {
				if v6 > 0 {
					out.RawByte(',')
				}
				(v7).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"reviewex\":"
		out.RawString(prefix)
		out.String(string(in.ReviewExist))
	}
	{
		const prefix string = ",\"userrating\":"
		out.RawString(prefix)
		out.String(string(in.UserRating))
	}
	{
		const prefix string = ",\"collectionsInfo\":"
		out.RawString(prefix)
		if in.CollectionsInfo == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.CollectionsInfo {
				if v8 > 0 {
					out.RawByte(',')
				}
				(v9).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MovieResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB2ba5d8eEncodeCodexInternalPkgDomain1(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MovieResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB2ba5d8eDecodeCodexInternalPkgDomain1(l, v)
}
func easyjsonB2ba5d8eDecodeCodexInternalPkgDomain2(in *jlexer.Lexer, out *MovieBasic) {
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
		case "ID":
			out.Id = string(in.String())
		case "poster":
			out.Poster = string(in.String())
		case "title":
			out.Title = string(in.String())
		case "rating":
			out.Rating = string(in.String())
		case "info":
			out.Info = string(in.String())
		case "description":
			out.Description = string(in.String())
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
func easyjsonB2ba5d8eEncodeCodexInternalPkgDomain2(out *jwriter.Writer, in MovieBasic) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix[1:])
		out.String(string(in.Id))
	}
	{
		const prefix string = ",\"poster\":"
		out.RawString(prefix)
		out.String(string(in.Poster))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"rating\":"
		out.RawString(prefix)
		out.String(string(in.Rating))
	}
	{
		const prefix string = ",\"info\":"
		out.RawString(prefix)
		out.String(string(in.Info))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MovieBasic) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB2ba5d8eEncodeCodexInternalPkgDomain2(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MovieBasic) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB2ba5d8eDecodeCodexInternalPkgDomain2(l, v)
}
func easyjsonB2ba5d8eDecodeCodexInternalPkgDomain3(in *jlexer.Lexer, out *Movie) {
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
		case "ID":
			out.Id = string(in.String())
		case "poster":
			out.Poster = string(in.String())
		case "title":
			out.Title = string(in.String())
		case "originalTitle":
			out.TitleOriginal = string(in.String())
		case "rating":
			out.Rating = string(in.String())
		case "info":
			out.Info = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "trailerHref":
			out.Trailer = string(in.String())
		case "year":
			out.ReleaseYear = string(in.String())
		case "country":
			out.Country = string(in.String())
		case "motto":
			out.Motto = string(in.String())
		case "director":
			out.Director = string(in.String())
		case "budget":
			out.Budget = string(in.String())
		case "gross":
			out.Gross = string(in.String())
		case "duration":
			out.Duration = string(in.String())
		case "cast":
			if in.IsNull() {
				in.Skip()
				out.Actors = nil
			} else {
				in.Delim('[')
				if out.Actors == nil {
					if !in.IsDelim(']') {
						out.Actors = make([]Cast, 0, 2)
					} else {
						out.Actors = []Cast{}
					}
				} else {
					out.Actors = (out.Actors)[:0]
				}
				for !in.IsDelim(']') {
					var v10 Cast
					(v10).UnmarshalEasyJSON(in)
					out.Actors = append(out.Actors, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "genres":
			if in.IsNull() {
				in.Skip()
				out.Genres = nil
			} else {
				in.Delim('[')
				if out.Genres == nil {
					if !in.IsDelim(']') {
						out.Genres = make([]GenreInMovie, 0, 2)
					} else {
						out.Genres = []GenreInMovie{}
					}
				} else {
					out.Genres = (out.Genres)[:0]
				}
				for !in.IsDelim(']') {
					var v11 GenreInMovie
					(v11).UnmarshalEasyJSON(in)
					out.Genres = append(out.Genres, v11)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonB2ba5d8eEncodeCodexInternalPkgDomain3(out *jwriter.Writer, in Movie) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix[1:])
		out.String(string(in.Id))
	}
	{
		const prefix string = ",\"poster\":"
		out.RawString(prefix)
		out.String(string(in.Poster))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"originalTitle\":"
		out.RawString(prefix)
		out.String(string(in.TitleOriginal))
	}
	{
		const prefix string = ",\"rating\":"
		out.RawString(prefix)
		out.String(string(in.Rating))
	}
	{
		const prefix string = ",\"info\":"
		out.RawString(prefix)
		out.String(string(in.Info))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"trailerHref\":"
		out.RawString(prefix)
		out.String(string(in.Trailer))
	}
	{
		const prefix string = ",\"year\":"
		out.RawString(prefix)
		out.String(string(in.ReleaseYear))
	}
	{
		const prefix string = ",\"country\":"
		out.RawString(prefix)
		out.String(string(in.Country))
	}
	{
		const prefix string = ",\"motto\":"
		out.RawString(prefix)
		out.String(string(in.Motto))
	}
	{
		const prefix string = ",\"director\":"
		out.RawString(prefix)
		out.String(string(in.Director))
	}
	{
		const prefix string = ",\"budget\":"
		out.RawString(prefix)
		out.String(string(in.Budget))
	}
	{
		const prefix string = ",\"gross\":"
		out.RawString(prefix)
		out.String(string(in.Gross))
	}
	{
		const prefix string = ",\"duration\":"
		out.RawString(prefix)
		out.String(string(in.Duration))
	}
	{
		const prefix string = ",\"cast\":"
		out.RawString(prefix)
		if in.Actors == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v12, v13 := range in.Actors {
				if v12 > 0 {
					out.RawByte(',')
				}
				(v13).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"genres\":"
		out.RawString(prefix)
		if in.Genres == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v14, v15 := range in.Genres {
				if v14 > 0 {
					out.RawByte(',')
				}
				(v15).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Movie) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB2ba5d8eEncodeCodexInternalPkgDomain3(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Movie) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB2ba5d8eDecodeCodexInternalPkgDomain3(l, v)
}
func easyjsonB2ba5d8eDecodeCodexInternalPkgDomain4(in *jlexer.Lexer, out *CollectionInfo) {
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
		case "collection":
			out.Collection = string(in.String())
		case "hasMovie":
			out.HasMovie = bool(in.Bool())
		case "bookmarkId":
			out.BookmarkId = uint64(in.Uint64())
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
func easyjsonB2ba5d8eEncodeCodexInternalPkgDomain4(out *jwriter.Writer, in CollectionInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"collection\":"
		out.RawString(prefix[1:])
		out.String(string(in.Collection))
	}
	{
		const prefix string = ",\"hasMovie\":"
		out.RawString(prefix)
		out.Bool(bool(in.HasMovie))
	}
	{
		const prefix string = ",\"bookmarkId\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.BookmarkId))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CollectionInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB2ba5d8eEncodeCodexInternalPkgDomain4(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CollectionInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB2ba5d8eDecodeCodexInternalPkgDomain4(l, v)
}
func easyjsonB2ba5d8eDecodeCodexInternalPkgDomain5(in *jlexer.Lexer, out *Cast) {
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
		case "name":
			out.Name = string(in.String())
		case "href":
			out.Href = string(in.String())
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
func easyjsonB2ba5d8eEncodeCodexInternalPkgDomain5(out *jwriter.Writer, in Cast) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"href\":"
		out.RawString(prefix)
		out.String(string(in.Href))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Cast) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB2ba5d8eEncodeCodexInternalPkgDomain5(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Cast) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB2ba5d8eDecodeCodexInternalPkgDomain5(l, v)
}