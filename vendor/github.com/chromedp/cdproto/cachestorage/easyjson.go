// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package cachestorage

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

func easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage(in *jlexer.Lexer, out *RequestEntriesReturns) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "cacheDataEntries":
			if in.IsNull() {
				in.Skip()
				out.CacheDataEntries = nil
			} else {
				in.Delim('[')
				if out.CacheDataEntries == nil {
					if !in.IsDelim(']') {
						out.CacheDataEntries = make([]*DataEntry, 0, 8)
					} else {
						out.CacheDataEntries = []*DataEntry{}
					}
				} else {
					out.CacheDataEntries = (out.CacheDataEntries)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *DataEntry
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(DataEntry)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.CacheDataEntries = append(out.CacheDataEntries, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "hasMore":
			out.HasMore = bool(in.Bool())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage(out *jwriter.Writer, in RequestEntriesReturns) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.CacheDataEntries) != 0 {
		const prefix string = ",\"cacheDataEntries\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v2, v3 := range in.CacheDataEntries {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.HasMore {
		const prefix string = ",\"hasMore\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.HasMore))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RequestEntriesReturns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RequestEntriesReturns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RequestEntriesReturns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RequestEntriesReturns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage1(in *jlexer.Lexer, out *RequestEntriesParams) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "cacheId":
			out.CacheID = CacheID(in.String())
		case "skipCount":
			out.SkipCount = int64(in.Int64())
		case "pageSize":
			out.PageSize = int64(in.Int64())
		case "pathFilter":
			out.PathFilter = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage1(out *jwriter.Writer, in RequestEntriesParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"cacheId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.CacheID))
	}
	{
		const prefix string = ",\"skipCount\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.SkipCount))
	}
	{
		const prefix string = ",\"pageSize\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.PageSize))
	}
	if in.PathFilter != "" {
		const prefix string = ",\"pathFilter\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.PathFilter))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RequestEntriesParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RequestEntriesParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RequestEntriesParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RequestEntriesParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage1(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage2(in *jlexer.Lexer, out *RequestCachedResponseReturns) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "response":
			if in.IsNull() {
				in.Skip()
				out.Response = nil
			} else {
				if out.Response == nil {
					out.Response = new(CachedResponse)
				}
				(*out.Response).UnmarshalEasyJSON(in)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage2(out *jwriter.Writer, in RequestCachedResponseReturns) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Response != nil {
		const prefix string = ",\"response\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Response).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RequestCachedResponseReturns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RequestCachedResponseReturns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RequestCachedResponseReturns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RequestCachedResponseReturns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage2(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage3(in *jlexer.Lexer, out *RequestCachedResponseParams) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "cacheId":
			out.CacheID = CacheID(in.String())
		case "requestURL":
			out.RequestURL = string(in.String())
		case "requestHeaders":
			if in.IsNull() {
				in.Skip()
				out.RequestHeaders = nil
			} else {
				in.Delim('[')
				if out.RequestHeaders == nil {
					if !in.IsDelim(']') {
						out.RequestHeaders = make([]*Header, 0, 8)
					} else {
						out.RequestHeaders = []*Header{}
					}
				} else {
					out.RequestHeaders = (out.RequestHeaders)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *Header
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(Header)
						}
						(*v4).UnmarshalEasyJSON(in)
					}
					out.RequestHeaders = append(out.RequestHeaders, v4)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage3(out *jwriter.Writer, in RequestCachedResponseParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"cacheId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.CacheID))
	}
	{
		const prefix string = ",\"requestURL\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RequestURL))
	}
	{
		const prefix string = ",\"requestHeaders\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.RequestHeaders == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.RequestHeaders {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil {
					out.RawString("null")
				} else {
					(*v6).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RequestCachedResponseParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RequestCachedResponseParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RequestCachedResponseParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RequestCachedResponseParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage3(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage4(in *jlexer.Lexer, out *RequestCacheNamesReturns) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "caches":
			if in.IsNull() {
				in.Skip()
				out.Caches = nil
			} else {
				in.Delim('[')
				if out.Caches == nil {
					if !in.IsDelim(']') {
						out.Caches = make([]*Cache, 0, 8)
					} else {
						out.Caches = []*Cache{}
					}
				} else {
					out.Caches = (out.Caches)[:0]
				}
				for !in.IsDelim(']') {
					var v7 *Cache
					if in.IsNull() {
						in.Skip()
						v7 = nil
					} else {
						if v7 == nil {
							v7 = new(Cache)
						}
						(*v7).UnmarshalEasyJSON(in)
					}
					out.Caches = append(out.Caches, v7)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage4(out *jwriter.Writer, in RequestCacheNamesReturns) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Caches) != 0 {
		const prefix string = ",\"caches\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v8, v9 := range in.Caches {
				if v8 > 0 {
					out.RawByte(',')
				}
				if v9 == nil {
					out.RawString("null")
				} else {
					(*v9).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RequestCacheNamesReturns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RequestCacheNamesReturns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RequestCacheNamesReturns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RequestCacheNamesReturns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage4(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage5(in *jlexer.Lexer, out *RequestCacheNamesParams) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "securityOrigin":
			out.SecurityOrigin = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage5(out *jwriter.Writer, in RequestCacheNamesParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"securityOrigin\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SecurityOrigin))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RequestCacheNamesParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RequestCacheNamesParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RequestCacheNamesParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RequestCacheNamesParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage5(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage6(in *jlexer.Lexer, out *Header) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "value":
			out.Value = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage6(out *jwriter.Writer, in Header) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"value\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Value))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Header) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Header) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Header) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Header) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage6(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage7(in *jlexer.Lexer, out *DeleteEntryParams) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "cacheId":
			out.CacheID = CacheID(in.String())
		case "request":
			out.Request = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage7(out *jwriter.Writer, in DeleteEntryParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"cacheId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.CacheID))
	}
	{
		const prefix string = ",\"request\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Request))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DeleteEntryParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DeleteEntryParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DeleteEntryParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DeleteEntryParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage7(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage8(in *jlexer.Lexer, out *DeleteCacheParams) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "cacheId":
			out.CacheID = CacheID(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage8(out *jwriter.Writer, in DeleteCacheParams) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"cacheId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.CacheID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DeleteCacheParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DeleteCacheParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DeleteCacheParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DeleteCacheParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage8(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage9(in *jlexer.Lexer, out *DataEntry) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "requestURL":
			out.RequestURL = string(in.String())
		case "requestMethod":
			out.RequestMethod = string(in.String())
		case "requestHeaders":
			if in.IsNull() {
				in.Skip()
				out.RequestHeaders = nil
			} else {
				in.Delim('[')
				if out.RequestHeaders == nil {
					if !in.IsDelim(']') {
						out.RequestHeaders = make([]*Header, 0, 8)
					} else {
						out.RequestHeaders = []*Header{}
					}
				} else {
					out.RequestHeaders = (out.RequestHeaders)[:0]
				}
				for !in.IsDelim(']') {
					var v10 *Header
					if in.IsNull() {
						in.Skip()
						v10 = nil
					} else {
						if v10 == nil {
							v10 = new(Header)
						}
						(*v10).UnmarshalEasyJSON(in)
					}
					out.RequestHeaders = append(out.RequestHeaders, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "responseTime":
			out.ResponseTime = float64(in.Float64())
		case "responseStatus":
			out.ResponseStatus = int64(in.Int64())
		case "responseStatusText":
			out.ResponseStatusText = string(in.String())
		case "responseType":
			(out.ResponseType).UnmarshalEasyJSON(in)
		case "responseHeaders":
			if in.IsNull() {
				in.Skip()
				out.ResponseHeaders = nil
			} else {
				in.Delim('[')
				if out.ResponseHeaders == nil {
					if !in.IsDelim(']') {
						out.ResponseHeaders = make([]*Header, 0, 8)
					} else {
						out.ResponseHeaders = []*Header{}
					}
				} else {
					out.ResponseHeaders = (out.ResponseHeaders)[:0]
				}
				for !in.IsDelim(']') {
					var v11 *Header
					if in.IsNull() {
						in.Skip()
						v11 = nil
					} else {
						if v11 == nil {
							v11 = new(Header)
						}
						(*v11).UnmarshalEasyJSON(in)
					}
					out.ResponseHeaders = append(out.ResponseHeaders, v11)
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage9(out *jwriter.Writer, in DataEntry) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"requestURL\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RequestURL))
	}
	{
		const prefix string = ",\"requestMethod\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RequestMethod))
	}
	{
		const prefix string = ",\"requestHeaders\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.RequestHeaders == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v12, v13 := range in.RequestHeaders {
				if v12 > 0 {
					out.RawByte(',')
				}
				if v13 == nil {
					out.RawString("null")
				} else {
					(*v13).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"responseTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.ResponseTime))
	}
	{
		const prefix string = ",\"responseStatus\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ResponseStatus))
	}
	{
		const prefix string = ",\"responseStatusText\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ResponseStatusText))
	}
	{
		const prefix string = ",\"responseType\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.ResponseType).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"responseHeaders\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.ResponseHeaders == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v14, v15 := range in.ResponseHeaders {
				if v14 > 0 {
					out.RawByte(',')
				}
				if v15 == nil {
					out.RawString("null")
				} else {
					(*v15).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DataEntry) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DataEntry) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DataEntry) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DataEntry) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage9(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage10(in *jlexer.Lexer, out *CachedResponse) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "body":
			out.Body = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage10(out *jwriter.Writer, in CachedResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"body\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Body))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CachedResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CachedResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CachedResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CachedResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage10(l, v)
}
func easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage11(in *jlexer.Lexer, out *Cache) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "cacheId":
			out.CacheID = CacheID(in.String())
		case "securityOrigin":
			out.SecurityOrigin = string(in.String())
		case "cacheName":
			out.CacheName = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage11(out *jwriter.Writer, in Cache) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"cacheId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.CacheID))
	}
	{
		const prefix string = ",\"securityOrigin\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SecurityOrigin))
	}
	{
		const prefix string = ",\"cacheName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.CacheName))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Cache) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Cache) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComChromedpCdprotoCachestorage11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Cache) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Cache) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComChromedpCdprotoCachestorage11(l, v)
}
