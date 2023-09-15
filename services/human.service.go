package services

import "hng-stage2/resource"

type HumanService interface {
	GetAllHuman()([]*resource.DbHuman, error)
	GetHumanbyID(string) (*resource.DbHuman, error)
	UpdateHuman(string, *resource.Human) (*resource.DbHuman, error)
	DeleteHuman(string) error
	CreateHuman(*resource.Human) (*resource.DbHuman, error)
}
