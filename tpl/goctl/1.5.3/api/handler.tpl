package {{.PkgName}}

import (
	"net/http"
    {{if .HasRequest}}"github.com/zeromicro/go-zero/rest/httpx"
	{{end}}"smartpower/common/response"

	{{.ImportPackages}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
			response.WithParameterError(w,err)
			return
		}

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
        response.Response(w, {{if .HasResp}}resp{{else}}nil{{end}}, err)
	}
}
