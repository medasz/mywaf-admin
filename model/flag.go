package model

import (
	uuid "github.com/satori/go.uuid"
)

type Flag struct {
	Uuid string `xorm:"notnull"`
	Name string `xorm:"notnull"`
}

func initFlag() error {
	_, err := Engine.InsertOne(Flag{
		Uuid: uuid.NewV4().String(),
		Name: "config",
	})
	return err
}

func UpdataFlag(name string) error {
	_, err := Engine.Cols("uuid").Update(&Flag{Name: name, Uuid: uuid.NewV4().String()})
	return err
}
func GetFlag(name string) (string, error) {
	flag := Flag{
		Name: name,
	}
	_, err := Engine.Get(&flag)
	return flag.Uuid, err
}
