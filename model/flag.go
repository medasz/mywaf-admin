package model

import (
	uuid "github.com/satori/go.uuid"
	"log"
)

type Flag struct {
	Uuid string `xorm:"notnull"`
	Name string `xorm:"notnull"`
}

func initFlag() error {
	_, err := Engine.Insert(Flag{
		Uuid: uuid.NewV4().String(),
		Name: "config",
	}, Flag{
		Uuid: uuid.NewV4().String(),
		Name: "rule",
	})
	return err
}

func UpdataFlag(name string) error {
	log.Println(name)
	_, err := Engine.Cols("uuid").Update(&Flag{Uuid: uuid.NewV4().String()},&Flag{Name: name})
	return err
}
func GetFlag(name string) (string, error) {
	flag := Flag{
		Name: name,
	}
	_, err := Engine.Get(&flag)
	return flag.Uuid, err
}
