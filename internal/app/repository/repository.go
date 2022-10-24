package repository

import (
	"github.com/cavelms/config"
)

type Repository struct {
	DB   MongoDB
	Mail Mail
	RDBS Redis
}

type mongodb struct {
	*config.DB
}

type MongoDB interface {
	Create(m interface{}) error
	FetchByID(m interface{}) error
	FetchByEmail(m interface{}) error
	FetchAll(ml interface{}, m interface{}) error
	UpdateOne(m interface{}) error
	Delete(m interface{}) error
	DeleteMany(m interface{}) error
	FetchByUserID(ml interface{}, m interface{}) error
	UpdateManyWhere(m interface{}, field, value string) error
}

func NewRepository() *Repository {
	return &Repository{
		DB:   newMongoDBRepository(),
		Mail: newEmailRepository(),
		RDBS: newRedisRepository(),
	}
}
