package service

import (
	"chkdin/data"
	"chkdin/models"
)

type ServiceUsers struct {
	DataLayer data.DataUserInfo
}

func (t ServiceUsers) GetUsersList() ([]models.User, error) {
	val, err := t.DataLayer.GetUserList()
	if err != nil {
		return nil, err
	}
	return val, nil

}

func (p ServiceUsers) GetUserDetails(userId int) (models.User, error) {
	val, err := p.DataLayer.GetUserDetails(userId)
	if err != nil {
		return models.User{}, err
	}
	return val, nil
}

func (p ServiceUsers) DeleteUser(userId int) (string, error) {
	res, err := p.DataLayer.DeleteUser(userId)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (p ServiceUsers) CreateUser(user models.User) (models.User, error) {
	res, err := p.DataLayer.CreateUser(user)
	if err != nil {
		return models.User{}, err
	}
	return res, nil
}

func (p ServiceUsers) UpdateUser(user models.User, userId int) (models.User, error) {
	res, err := p.DataLayer.UpdateUser(user, userId)
	if err != nil {
		return models.User{}, err
	}
	return res, nil
}
