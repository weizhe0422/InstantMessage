package Service

import (
	"../Model"
	"../Util"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"math/rand"
	"time"
)

type UserService struct {
}

func (s *UserService) Register(mobile, plainPassWd, nickName, avatar, sex string) (user Model.User, err error) {

	//確認用戶是否已存在
	tmp := Model.User{}
	_, err = DBEngine.Where("mobile=?", mobile).Get(&tmp)
	if err != nil {
		log.Println(err.Error())
		return tmp, errors.New("無法取得帳戶是否存在，請重新嘗試註冊!")
	}

	//帳戶已存在
	if tmp.Id > 0 {
		return tmp, errors.New("該手機號碼已經註冊過，請重新確認!")
	}

	//開始註冊
	//	將資訊填入
	tmp.Mobile = mobile
	tmp.Nickname = nickName
	tmp.Avatar = avatar
	tmp.Sex = sex
	//	將輸入輸入的密碼加密後再存回去
	tmp.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	tmp.Passwd = Util.MakePasswd(plainPassWd, tmp.Salt)
	tmp.Createat = time.Now()
	tmp.Token = fmt.Sprintf("%8d", time.Now().Unix())
	//寫入資料庫
	_, err = DBEngine.InsertOne(&tmp)

	return tmp, err
}

func (s *UserService) LogIn (mobile, plainPassWd string) (user Model.User, err error){
	//確認帳戶是否存在
	tmp := Model.User{}
	_, err = DBEngine.Where("mobile=?", mobile).Get(&tmp)
	if err != nil {
		log.Println(err.Error())
		return tmp, errors.New("無法取得帳戶是否存在，請重新嘗試登入!")
	}

	//不存在: 則直接返回
	if tmp.Id <= 0{
		return tmp, errors.New("該帳戶不存在，請進行註冊!")
	}

	//存在: 驗證密碼是否正確
	//不正確: 則返回錯誤訊息
	log.Println(plainPassWd+"/"+tmp.Passwd+"/"+tmp.Salt)
	log.Println(Util.MakePasswd(plainPassWd, tmp.Salt))
	if !Util.ValidatePasswd(plainPassWd, tmp.Salt, tmp.Passwd){
		return Model.User{}, errors.New("密碼輸入錯誤，請重新輸入!")
	}

	//正確: 更新token回資料庫，並返回帳戶資訊
	token := Util.Md5Encode(fmt.Sprintf("%8d",time.Now().Unix()))
	tmp.Token = token
	_, err = DBEngine.Id(tmp.Id).Cols("token").Update(&tmp)
	if err != nil {
		return tmp, errors.New("無法更新帳戶狀態，請重新登入")
	}

	return tmp, err
}