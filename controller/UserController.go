package controller

import (
	"MusicLibrary/common"
	"MusicLibrary/model"
	"MusicLibrary/untils"
	"crypto/sha512"
	"encoding/hex"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Check if Mail exists
func isMailExist(db *gorm.DB, Mail string) bool {
	var user model.User
	db.Where("Mail = ?", Mail).First(&user)
	if user.ID != 0 {
		return true
	}

	return false
}

// Find User
func findUser(DB *gorm.DB, Mail string) model.User {
	var user model.User
	DB.Where("Mail = ?", Mail).First(&user)
	return user
}

func Register(ctx *gin.Context) {

	var requestUser = model.User{}
	ctx.Bind(&requestUser)
	name := requestUser.Name
	mail := requestUser.Mail
	password := requestUser.Password
	if !(len(mail) > 0 && len(password) > 0 && len(name) > 0) {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "请输入全部信息")
		return
	}
	// Verify
	if !untils.VerifyMail(mail) {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "邮箱格式不正确")
		return
	}
	if len(password) < 6 {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	if len(name) == 0 {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "昵称不能为空")
		return
	}
	hashedPassword := sha512.Sum512([]byte(password))
	Uid := sha512.Sum512([]byte(mail))
	newUser := model.User{
		Name:     name,
		Mail:     mail,
		Password: hex.EncodeToString(hashedPassword[:]),
		Uid:      hex.EncodeToString(Uid[:]),
	}
	DB := common.GetDB()
	if isMailExist(DB, mail) {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "邮箱已存在，请登录")
		return
	}
	DB.Create(&newUser)
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "内部错误"})
		log.Printf("token generate error : %v", err)
		return
	}
	common.Success(ctx, gin.H{"token": token}, "注册成功")
}
func Login(ctx *gin.Context) {
	var requestUser = model.User{}
	ctx.Bind(&requestUser)
	name := requestUser.Name
	mail := requestUser.Mail
	password := requestUser.Password
	// Verify
	if !untils.VerifyMail(mail) {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "邮箱格式不正确")
		return
	}
	if len(password) < 6 {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	if len(name) == 0 {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "昵称不能为空")
		return
	}
	hashedPassword := sha512.Sum512([]byte(password))
	DB := common.GetDB()
	loginUser := findUser(DB, mail)
	if loginUser.Uid == "" {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在，请注册")
		return
	}
	if loginUser.Password != hex.EncodeToString(hashedPassword[:]) {
		common.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码错误")
		return
	}
}
