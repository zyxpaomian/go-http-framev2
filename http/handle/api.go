package handle

import (
	"encoding/json"
	"io/ioutil"
	"go-http-frame/common"
	log "go-http-frame/common/formatlog"
	"go-http-frame/controller"
	"net/http"
)

// 用户认证
func apiUserAuth(res http.ResponseWriter, req *http.Request) {
	type Request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	reqContent, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.Errorf("[http] 请求报文解析失败")
		common.ReqBodyInvalid(res)
		return
	}

	request := &Request{}
	if err := common.ParseJsonStr(string(reqContent), request); err != nil {
		log.Errorln("[http] 解析模板JSON失败")
		common.ResMsg(res, 400, err.Error())
		return
	}

	type Response struct {
		Token string `json:"token"`
	}

	token, err := controller.UserController.GetToken(request.Username, request.Password)
	if err != nil {
		log.Errorln("[http] 用户认证失败")
		common.ResMsg(res, 400, err.Error())
		return		
	} 

	response := &Response{Token: token}
	result, err := json.Marshal(response)
	if err != nil {
		log.Errorf("[http] apiUserAuth JSON生成失败, %v", err.Error())
		common.ResMsg(res, 500, err.Error())
		return
	}
	common.ResMsg(res, 200, string(result))
}

// 获取所有用户名
func apiGetAllUser(res http.ResponseWriter, req *http.Request) {
	users, err := controller.UserController.GetAllUsers()
	if err != nil {
		log.Errorf("[http] apiGetAllUser 数据处理失败, %v", err.Error())
		common.ResMsg(res, 500, err.Error())
		return
	}

	response, err := json.Marshal(users)
	if err != nil {
		log.Errorf("[http] apiGetAllAgents JSON生成失败, %v", err.Error())
		common.ResMsg(res, 400, err.Error())
		return
	}
	common.ResMsg(res, 200, string(response))
}
