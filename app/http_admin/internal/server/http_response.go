package server

import (
	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/kratos/v2/transport/http/status"

	grpcstatus "google.golang.org/grpc/status"
	stdhttp "net/http"
)

// httpResponse 响应结构体
type httpResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// EncoderResponse  请求响应封装
func EncoderResponse() http.EncodeResponseFunc {
	return func(w http.ResponseWriter, r *http.Request, v interface{}) error {
		if v == nil {
			return nil
		}
		resp := &httpResponse{
			Code:    stdhttp.StatusOK,
			Message: "",
			Data:    v,
		}
		codec := encoding.GetCodec("json")
		data, err := codec.Marshal(resp)
		if err != nil {
			return err
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(data)
		if err != nil {
			return err
		}
		return nil
	}
}

// EncoderError 错误响应封装
func EncoderError() http.EncodeErrorFunc {
	return func(w http.ResponseWriter, r *http.Request, err error) {
		if err == nil {
			return
		}
		se := &httpResponse{}
		gs, ok := grpcstatus.FromError(err)
		if !ok {
			se = &httpResponse{Code: stdhttp.StatusInternalServerError}
		}
		se = &httpResponse{
			Code:    status.FromGRPCCode(gs.Code()),
			Message: gs.Message(),
			Data:    nil,
		}
		codec, _ := http.CodecForRequest(r, "Accept")
		body, err := codec.Marshal(se)
		if err != nil {
			w.WriteHeader(stdhttp.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/"+codec.Name())
		w.WriteHeader(se.Code)
		_, _ = w.Write(body)
	}
}
