package response

import (
	"net/http"
	"smartpower/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

type Body struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err == nil {
		body.Code = xerr.OK
		body.Msg = "OK"
		body.Data = resp
	} else {
		logx.Debug(err)
		errcode := xerr.SERVER_COMMON_ERROR
		errmsg := err.Error()
		causeErr := errors.Cause(err)                // err类型
		if e, ok := causeErr.(*xerr.CodeError); ok { //自定义错误类型
			//自定义CodeError
			errcode = e.GetErrCode()
			errmsg = e.GetErrMsg()
		} else {
			if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
				grpcCode := uint32(gstatus.Code())
				if xerr.IsCodeErr(grpcCode) { //区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
					errcode = grpcCode
					errmsg = gstatus.Message()
				}
			}
		}
		body.Code = errcode
		body.Msg = errmsg

	}

	httpx.OkJson(w, body)

}

func WithParameterError(w http.ResponseWriter, err error) {
	var body Body
	body.Code = xerr.PARAM_ERROR
	body.Msg = err.Error()
	httpx.OkJson(w, body)
}
