package handle

import (
	"go-http-frame/http"
	//go_http "net/http"
)

func InitHandle(r *http.WWWMux) {
	// api相关的接口
	initAPIMapping(r)
}

func initAPIMapping(r *http.WWWMux) {
	// 用户认证
	r.RegistURLMapping("/v1/api/user/userauth", "POST", false, apiUserAuth)
	// 获取所有用户
	r.RegistURLMapping("/v1/api/user/getalluser", "GET", true, apiGetAllUser)
	
}

