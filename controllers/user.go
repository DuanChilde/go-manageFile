package controllers

import (
	"encoding/json"
	"manageFile/models"
	"net/http"
	"time"

	"github.com/astaxie/beego"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	uid := models.AddUser(user)
	u.Data["json"] = res.Success("", map[string]interface{}{"uid": uid})
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users := models.GetAllUsers()
	userList := map[string]interface{}{
		"userList": users,
	}
	u.Data["json"] = res.Success("", userList)
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	int64	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	//rs := ValidateTokenMiddleware(u.Ctx.ResponseWriter, u.Ctx.Request)
	//beego.Debug(rs)
	//beego.Debug(u.Ctx.Request.Header.Get("token"))
	uid, _ := u.GetInt64(":uid")
	if uid != 0 {
		user, err := models.GetUserById(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			userInfo := map[string]interface{}{
				"user": user,
			}
			u.Data["json"] = res.Success("", userInfo)
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	int64	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	uid, _ := u.GetInt64(":uid")
	if uid != 0 {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		_, err := models.UpdateUser(uid, &user)
		if err != nil {
			u.Data["json"] = res.Fail(err.Error())
		} else {
			u.Data["json"] = res.Success("update success", make(map[string]interface{}))
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	int64	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid, _ := u.GetInt64(":uid")
	models.DeleteUser(uid)
	u.Data["json"] = res.Success("delete success", make(map[string]interface{}))
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	//isRemenbered := u.GetString("isRemenbered")
	if userId, err := models.Login(username, password); userId != 0 {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := make(jwt.MapClaims)
		claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
		claims["iat"] = time.Now().Unix()
		token.Claims = claims
		tokenString, _ := token.SignedString([]byte("duanwei"))

		userInfo := map[string]interface{}{
			"userId":   userId,
			"username": username,
			"token":    tokenString,
		}
		u.Data["json"] = res.Success("login success", userInfo)
	} else {
		u.Data["json"] = res.Fail(err.Error())
	}
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = res.Success("logout success", make(map[string]interface{}))
	u.ServeJSON()
}

func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request) bool {
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte("duanwei"), nil
		})
	if err == nil {
		if token.Valid {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
	return true
}
