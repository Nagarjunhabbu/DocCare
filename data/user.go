package data

import (
	"chkdin/models"
	"database/sql"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type DataUserInfo struct {
	Sql *sql.DB
}

//func to get users list present in DB
func (t DataUserInfo) GetUserList() (resp []models.User, err error) {
	query := "select id,name,email,place from users"
	rows, err := t.Sql.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var t models.User
		rows.Scan(&t.Id, &t.Name, &t.Email, &t.Place)
		resp = append(resp, t)
	}
	return
}

//func to get specified user data by passing userId
func (d DataUserInfo) GetUserDetails(userId int) (models.User, error) {
	query := "select id,name,place,email from users where id=?"
	row := d.Sql.QueryRow(query, userId)
	var t models.User
	err := row.Scan(&t.Id, &t.Name, &t.Place, &t.Email)
	if err != nil {
		return models.User{}, err
	}
	return t, nil
}

//func to delete user from DB
func (d DataUserInfo) DeleteUser(userId int) (string, error) {
	_, err := d.GetUserDetails(userId)
	if err != nil {
		return "", err
	}
	query, _ := d.Sql.Prepare("delete from users where id=?")
	_, err = query.Exec(userId)
	if err != nil {
		return "", err
	}
	return "User Deleted Successfully!", nil
}

//func to create new user in DB
func (d DataUserInfo) CreateUser(user models.User) (models.User, error) {

	query, _ := d.Sql.Prepare("insert into users(name,place,email) values (?, ?, ?)")
	resp, err := query.Exec(user.Name, user.Place, user.Email)
	if err != nil {
		return models.User{}, err
	}
	id, _ := resp.LastInsertId()
	return d.GetUserDetails(int(id))
}

//func to update particular information of specified user
func (d DataUserInfo) UpdateUser(user models.User, userId int) (models.User, error) {

	if user.Place == "" {
		query, _ := d.Sql.Prepare("update users set email=? where id=?")
		_, err := query.Exec(user.Email, userId)
		if err != nil {
			return models.User{}, err
		}
	} else if user.Email == "" {
		query, _ := d.Sql.Prepare("update users set place=? where id=?")
		_, err := query.Exec(user.Place, userId)
		if err != nil {
			return models.User{}, err
		}
	} else if user.Email != "" && user.Place != "" {
		query, _ := d.Sql.Prepare("update users set place=?,email =? where id=?")
		_, err := query.Exec(user.Place, user.Email, userId)
		if err != nil {
			return models.User{}, err
		}
	} else {
		return models.User{}, errors.New("invalid operation")
	}
	return d.GetUserDetails(userId)
}

// func to getUser details by passing user Name
func (d DataUserInfo) GetUserByName(userName string) (models.User, error) {
	query := "select id from users where name=?"
	row := d.Sql.QueryRow(query, userName)
	var t models.User
	err := row.Scan(&t.Id)
	if err != nil {
		return models.User{}, err
	}
	return t, nil
}

//func to register new user
func (d DataUserInfo) SignUpUser(user models.Login) (models.User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	password := string(hash)

	query, _ := d.Sql.Prepare("insert into users(name,password) values (?, ?)")
	resp, err := query.Exec(user.Name, password)
	if err != nil {
		return models.User{}, err
	}
	id, _ := resp.LastInsertId()
	return d.GetUserDetails(int(id))
}
