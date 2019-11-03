package Controller

import (
	"../Model"
	"../Service"
	"../Util"
	"fmt"
	"math/rand"
	"net/http"
)

var usrService Service.UserService

func RegisterFunc() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		request.ParseForm()
		mobile := request.PostForm.Get("mobile")
		plainPassWd := request.PostForm.Get("passwd")
		nickName := fmt.Sprintf("USER%8d", rand.Int31())
		avatar := ""
		sex := Model.USER_MAN

		user, err := usrService.Register(mobile, plainPassWd, nickName, avatar, sex)
		if err != nil {
			Util.ParseFailResult(writer, "註冊失敗: "+err.Error())
		} else {
			Util.ParseOKResult(writer, user)
		}
	}
}

func LoginFunc() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		mobile := request.PostForm.Get("mobile")
		passwd := request.PostForm.Get("passwd")

		user, err := usrService.LogIn(mobile, passwd)
		if err != nil {
			Util.ParseFailResult(writer, "登入失敗: "+err.Error())
		}else{
			Util.ParseOKResult(writer, user)
		}
	}
}
