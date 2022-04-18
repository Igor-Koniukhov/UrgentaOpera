package repository

import (
	"context"
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
	"todo/backend/internal/configs"
	models2 "todo/backend/internal/models"
)

type UserRepository interface {
	CreateUser(user *models2.User) (*models2.User, int, error)
	GetAllUsers() ([]models2.User, error)
	GetUserByEmail(email string) (*models2.User, error)
	GetUserByID(id int) (models2.User, error)
	CheckToken(token string) (passRes models2.PasswordReset, err error)
	WriteToResetPassword(pr models2.PasswordReset) error
	ResetPassword(email, p string) error
}
type UserRepo struct {
	App   *configs.AppConfig
	user  models2.User
	users []models2.User
	DB    *sql.DB
}

func NewUserRepo(app *configs.AppConfig, db *sql.DB) *UserRepo {
	return &UserRepo{App: app, DB: db}
}
func (u *UserRepo) CreateUser(user *models2.User) (*models2.User, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("INSERT INTO %s (name, email, phone, password)"+
		" VALUES(?,?,?,?) ", models2.TableUsers)
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
func (u *UserRepo) GetAllUsers() ([]models2.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	sqlStmt := fmt.Sprintf("SELECT id, name, phone, email, "+
		" created_at, updated_at FROM %s ",
		models2.TableUsers)
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

func (u *UserRepo) GetUserByEmail(email string) (*models2.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("SELECT id, name, email, phone, password,"+
		" created_at, updated_at FROM %s WHERE email = ? ", models2.TableUsers)
	row := u.DB.QueryRowContext(ctx, sqlStmt, email)
	err := row.Scan(
		&u.user.ID,
		&u.user.Name,
		&u.user.Email,
		&u.user.Phone,
		&u.user.Password,
		&u.user.CreatedAt,
		&u.user.UpdatedAt,
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &u.user, nil
}
func (u *UserRepo) GetUserByID(id int) (models2.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("SELECT id, name, phone, email,"+
		" created_at, updated_at FROM %s WHERE id = ? ",
		models2.TableUsers)
	row := u.DB.QueryRowContext(ctx, sqlStmt, id)
	err := row.Scan(
		&u.user.ID,
		&u.user.Name,
		&u.user.Phone,
		&u.user.Email,
		&u.user.CreatedAt,
		&u.user.UpdatedAt)
	if err != nil {
		fmt.Println(err)
		return models2.User{}, err
	}
	return u.user, nil
}
func (u *UserRepo) CheckToken(token string) (passRes models2.PasswordReset, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("SELECT id, email, token FROM %s WHERE token = ? ", models2.TablePassReset)
	row := u.DB.QueryRowContext(ctx, sqlStmt, token)
	err = row.Scan(
		&passRes.ID,
		&passRes.Email,
		&passRes.Token,
	)
	if err != nil {
		fmt.Println(err)
		return models2.PasswordReset{}, err
	}
	return passRes, nil
}
func (u *UserRepo) WriteToResetPassword(rp models2.PasswordReset) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	sqlStmt := fmt.Sprintf("INSERT INTO  %s (email, token) VALUES(?,?) ", models2.TablePassReset)
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
	sqlStmt := fmt.Sprintf("UPDATE  %s SET password=? WHERE email=? ", models2.TableUsers)
	_, err := u.DB.ExecContext(ctx, sqlStmt, password, email)
	if err != nil {
		return err
	}
	return nil
}
