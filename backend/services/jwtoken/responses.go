package jwtoken

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"strconv"
	"todo/backend/internal/configs"
	models2 "todo/backend/internal/models"
	"todo/backend/internal/repository"
)

type JwTokenI interface {
	TokenResponder(w http.ResponseWriter, logReq *models2.LoginRequest) (*models2.LoginResponse, int, error)
}

var (
	RefreshSecret = os.Getenv("REFRESH_SECRET")
	AccessSecret  = os.Getenv("ACCESS_SECRET")
)

type JwToken struct {
	App  *configs.AppConfig
	repo repository.Repository
}

func NewJwToken(app *configs.AppConfig, repo repository.Repository) *JwToken {
	return &JwToken{App: app, repo: repo}
}

func (j *JwToken) TokenResponder(w http.ResponseWriter, logReq *models2.LoginRequest) (*models2.LoginResponse, int, error) {
	user, err := j.repo.GetUserByEmail(logReq.Email)
	if err != nil {
		fmt.Println(err)
		return nil, 0, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(logReq.Password)); err != nil {
		fmt.Println(err)
		return nil, 0, err
	}
	resp, err := TokenGenerator(w, user.ID)
	if err != nil {
		if err != nil {
			log.Println(err)
			return nil, 0, err
		}
	}
	return resp, user.ID, nil
}

func TokenGenerator(w http.ResponseWriter, id int) (*models2.LoginResponse, error) {
	RefreshLifetimeMinutes, err := strconv.Atoi(os.Getenv("REFRESH_LIFE_TIME_MINUTES"))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	AccessLifetimeMinutes, err := strconv.Atoi(os.Getenv("ACCESS_LIFE_TIME_MINUTES"))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	accessString, err := GenerateToken(id, AccessLifetimeMinutes, AccessSecret)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	refreshString, err := GenerateToken(id, RefreshLifetimeMinutes, RefreshSecret)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return nil, err
	}
	resp := &models2.LoginResponse{
		AccessToken:  accessString,
		RefreshToken: refreshString,
	}
	return resp, nil
}
