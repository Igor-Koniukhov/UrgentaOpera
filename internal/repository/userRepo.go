package repository

import (
	"context"
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
	"todo/internal/configs"
	"todo/internal/models"
)

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, int, error)
	GetAllUsers() ([]models.User, error)
	GetAllFreeUsers() ([]models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id int) (models.User, error)
	CheckToken(token string) (passRes models.PasswordReset, err error)
	WriteToResetPassword(pr models.PasswordReset) error
	ResetPassword(email, p string) error
}
type UserRepo struct {
	App   *configs.AppConfig
	user  models.User
	users []models.User
	DB    *sql.DB
}

func NewUserRepo(app *configs.AppConfig, db *sql.DB) *UserRepo {
	return &UserRepo{App: app, DB: db}
}
func (u *UserRepo) CreateUser(user *models.User) (*models.User, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("INSERT INTO %s (name, email, phone, password, status_free)"+
		" VALUES(?,?,?,?, true) ", models.TableUsers)
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return nil, 0, err
	}
	res, err := u.DB.ExecContext(ctx, sqlStmt,
		user.Name,
		user.Email,
		user.Phone,
		pass)
	if err != nil {
		fmt.Println(err)
		return nil, 0, err
	}
	userId, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return nil, 0, err
	}
	return user, int(userId), nil
}
func (u *UserRepo) GetAllUsers() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	sqlStmt := fmt.Sprintf("SELECT id, name, phone, email, status_free, "+
		" created_at, updated_at FROM %s ",
		models.TableUsers)
	results, err := u.DB.QueryContext(ctx, sqlStmt)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for results.Next() {
		err = results.Scan(
			&u.user.ID,
			&u.user.Name,
			&u.user.Phone,
			&u.user.Email,
			&u.user.StatusFree,
			&u.user.CreatedAt,
			&u.user.UpdatedAt,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		u.users = append(u.users, u.user)
	}
	return u.users, nil

}
func (u *UserRepo) GetAllFreeUsers() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("SELECT id, name, phone, email, status_free,"+
		" created_at, updated_at FROM %s WHERE status_free=true",
		models.TableUsers)
	results, err := u.DB.QueryContext(ctx, sqlStmt)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for results.Next() {
		err = results.Scan(
			&u.user.ID,
			&u.user.Name,
			&u.user.Phone,
			&u.user.Email,
			&u.user.StatusFree,
			&u.user.CreatedAt,
			&u.user.UpdatedAt,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		u.users = append(u.users, u.user)
	}
	return u.users, nil
}
func (u *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("SELECT * FROM %s WHERE email = ? ", models.TableUsers)
	row := u.DB.QueryRowContext(ctx, sqlStmt, email)
	err := row.Scan(
		&u.user.ID,
		&u.user.Name,
		&u.user.Email,
		&u.user.Phone,
		&u.user.Password,
		&u.user.StatusFree,
		&u.user.CreatedAt,
		&u.user.UpdatedAt,
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &u.user, nil
}
func (u *UserRepo) GetUserByID(id int) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("SELECT id, name, phone, email, status_free,"+
		" created_at, updated_at FROM %s WHERE id = ? ",
		models.TableUsers)
	row := u.DB.QueryRowContext(ctx, sqlStmt, id)
	err := row.Scan(
		&u.user.ID,
		&u.user.Name,
		&u.user.Phone,
		&u.user.Email,
		&u.user.StatusFree,
		&u.user.CreatedAt,
		&u.user.UpdatedAt)
	if err != nil {
		fmt.Println(err)
		return models.User{}, err
	}
	return u.user, nil
}
func (u *UserRepo) CheckToken(token string) (passRes models.PasswordReset, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("SELECT id, email, token FROM %s WHERE token = ? ", models.TablePassReset)
	row := u.DB.QueryRowContext(ctx, sqlStmt, token)
	err = row.Scan(
		&passRes.ID,
		&passRes.Email,
		&passRes.Token,
	)
	if err != nil {
		fmt.Println(err)
		return models.PasswordReset{}, err
	}
	return passRes, nil
}
func (u *UserRepo) WriteToResetPassword(rp models.PasswordReset) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("INSERT INTO  %s (email, token) VALUES(?,?) ", models.TablePassReset)
	_, err := u.DB.ExecContext(ctx, sqlStmt, rp.Email, rp.Token)
	if err != nil {
		return err
	}
	return nil
}
func (u *UserRepo) ResetPassword(email, p string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	password, _ := bcrypt.GenerateFromPassword([]byte(p), 14)
	sqlStmt := fmt.Sprintf("UPDATE  %s SET password=? WHERE email=? ", models.TableUsers)
	_, err := u.DB.ExecContext(ctx, sqlStmt, password, email)
	if err != nil {
		return err
	}
	return nil
}
