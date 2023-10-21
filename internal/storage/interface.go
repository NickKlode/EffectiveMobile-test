package storage

import "emobletest/internal/storage/model"

type StorageInterface interface {
	CreateUser(u model.User) (int, error)
	DeleteUser(id int) error
	UpdateUser(id int, ui model.UpdateInput) error
	GetUser(gi model.GetInput) ([]model.User, error)
}
