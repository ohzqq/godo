package godo

import "github.com/spf13/cast"

type Response map[string]any

func NewResponse(status int, body any) Response {
	return Response{
		"body":       body,
		"statusCode": status,
	}
}

func (r *Response) SetHeader(k, v string) {
	h := r.Headers()
	h[k] = v
	r["headers"] = h
}

func (r *Response) Headers() map[string]string {
	if _, ok := r["headers"]; !ok {
		return map[string]string{}
	}
	return cast.ToStringMapString(r["headers"])
}

func (r *Response) ContentType(v string) {
	r.SetHeader("Content-Type", v)
}

func (r *Response) JSON() *Response {
	r.ContentType("application/json")
	return r
}

func (r *Response) XML() *Response {
	r.ContentType("application/xml")
	return r
}

type Request map[string]any

func (r Request) Get(key string) string {
	if v, ok := r[key]; ok {
		return cast.ToString(v)
	}
	return ""
}

func (r Request) Params() map[string]any {
	p := make(map[string]any)
	for k, v := range r {
		switch k {
		case "http":
		case "__ow_headers", "__ow_method", "__ow_path":
		default:
			p[k] = v
		}
	}
	return p
}

func (r Request) Path() string {
	h := r.getHTTP()
	if p, ok := h["path"]; ok {
		return cast.ToString(p)
	}
	return ""
}

func (r Request) Method() string {
	h := r.getHTTP()
	if m, ok := h["method"]; ok {
		return cast.ToString(m)
	}
	return ""
}

func (r Request) getHTTP() map[string]any {
	if h, ok := r["http"]; ok {
		return cast.ToStringMap(h)
	}
	return map[string]any{}
}

func (r Request) Headers() map[string]string {
	h := r.getHTTP()
	if headers, ok := h["headers"]; ok {
		return cast.ToStringMapString(headers)
	}
	return map[string]string{}
}
