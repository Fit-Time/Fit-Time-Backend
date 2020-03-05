package controllers

import (
	"net/http"
	"testing"
)

func TestHello(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test",
			args: args{
				w: nil,
				req: &http.Request{
					Method:           "",
					URL:              nil,
					Proto:            "",
					ProtoMajor:       0,
					ProtoMinor:       0,
					Header:           nil,
					Body:             nil,
					GetBody:          nil,
					ContentLength:    0,
					TransferEncoding: nil,
					Close:            false,
					Host:             "",
					Form:             nil,
					PostForm:         nil,
					MultipartForm:    nil,
					Trailer:          nil,
					RemoteAddr:       "",
					RequestURI:       "",
					TLS:              nil,
					Cancel:           nil,
					Response:         nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Hello(tt.args.w, tt.args.req)
		})
	}
}
