package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"todo/backend/internal/configs"
	models2 "todo/backend/internal/models"
	"todo/backend/internal/repository"
	"todo/backend/services/jwtoken"
)

type UserI interface {
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	Registration(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	LogOut(w http.ResponseWriter, r *http.Request)
	Forgot(w http.ResponseWriter, r *http.Request)
	Reset(w http.ResponseWriter, r *http.Request)
}
type User struct {
	App  *configs.AppConfig
	repo repository.Repository
}

func NewUser(app *configs.AppConfig, repo repository.Repository) *User {
	return &User{App: app, repo: repo}
}
func (u *User) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	users, err := u.repo.GetAllUsers()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.NewEncoder(w).Encode(&users)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (u *User) Registration(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	message := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	if data["password"] != data["password_confirm"] {
		message["message"] = "Password do not match!"
		err = json.NewEncoder(w).Encode(message)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	if _, ok := u.CheckUserExist(data["email"]); ok {
		message["registration_status"] = " UserI with that email exist!"
		fmt.Println(ok, " UserI with that email exist!")
		return
	}
	user := models2.User{
		Name:     data["name"],
		Email:    data["email"],
		Phone:    data["phoneNumber"],
		Password: data["password"],
	}
	_, id, err := u.repo.CreateUser(&user)
	if err != nil {
		return
	}
	message["registration_status"] = " Success! Congratulations you registered! Now Login!"
	message["id"] = id
	err = json.NewEncoder(w).Encode(&message)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (u *User) Login(w http.ResponseWriter, r *http.Request) {
	var logRequest models2.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&logRequest)
	if err != nil {
		return
	}

	message := make(map[string]interface{})
	user, ok := u.LoginValidation(logRequest)
	if !ok {
		message["status"] = "false"
		message["message"] = "UserI doesn't exist!"
		message["user_name"] = ""
		message["jwt_status"] = "seted"
		err = json.NewEncoder(w).Encode(&message)
		if err != nil {
			fmt.Println(err)
			return
		}
		w.WriteHeader(http.StatusForbidden)
		return
	}
	token, err := jwtoken.TokenGenerator(w, user.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	setToken := &http.Cookie{
		Name:     "Authorization",
		Value:    "Bearer " + token.AccessToken,
		HttpOnly: true,
		SameSite: 0,
	}
	http.SetCookie(w, setToken)
	message["user_id"] = user.ID
	message["status"] = "true"
	message["user_name"] = user.Name
	message["email"] = user.Email
	message["message"] = "You logging!"
	message["access_token"] = token.AccessToken
	message["refresh_token"] = token.RefreshToken

	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.NewEncoder(w).Encode(&message)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (u *User) LogOut(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" logout")
	cookies := r.Cookies()
	if len(cookies) >= 0 {
		for _, ck := range cookies {
			if ck.Name == "Authorization" {
				ck.MaxAge = -1
				http.SetCookie(w, ck)
			}
		}
	}
}
func (u *User) Forgot(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Println(err)
		return
	}
	token := RandStringRunes(12)
	passwordReset := models2.PasswordReset{
		Email: data["email"],
		Token: token,
	}
	err := u.repo.WriteToResetPassword(passwordReset)
	if err != nil {
		fmt.Println(err)
		return
	}
	from := "admin@example.com"
	to := []string{
		data["email"],
	}
	url := "http://localhost:8080/reset/" + token
	message := []byte("Click <a href=\"" + url + "\">here</a>to reset your password!")
	err = smtp.SendMail("0.0.0.0:1025", nil, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	messageMap := make(map[string]string)
	messageMap["message"] = "success"
	err = json.NewEncoder(w).Encode(&messageMap)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (u *User) Reset(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	messageMap := make(map[string]string)
	if data["password"] != data["password_confirm"] {
		w.WriteHeader(http.StatusBadRequest)
		messageMap["message"] = "Password do not match!"
		err = json.NewEncoder(w).Encode(messageMap)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	passwordReset, err := u.repo.CheckToken(data["token"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		messageMap["message"] = "Invalid token!"
	}
	err = u.repo.ResetPassword(passwordReset.Email, data["password"])
	if err != nil {
		fmt.Println(err)
		return
	}
	messageMap["message"] = "success"
	err = json.NewEncoder(w).Encode(messageMap)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (u *User) LoginValidation(logReq models2.LoginRequest) (user *models2.User, ok bool) {

	user, ok = u.CheckUserExist(logReq.Email)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(logReq.Password)); err != nil {
		fmt.Println(err)
		return &models2.User{}, ok
	}
	fmt.Println(ok, " it ok")
	return user, ok
}
func (u *User) CheckUserExist(email string) (*models2.User, bool) {
	user, err := u.repo.GetUserByEmail(email)
	fmt.Println(user, " info about user")
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return &models2.User{}, false
	}
	return user, true
}
func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
