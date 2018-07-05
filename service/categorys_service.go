package service

import (
	"github.com/CanyonSYSU/Server/db"
	"github.com/CanyonSYSU/Server/entity"
	"github.com/CanyonSYSU/Server/loghelper"
	//"fmt"
	//simplejson "github.com/bitly/go-simplejson"
)

func ListAllCategorys() []string {
	categorys := db.FindAllCategorys()
	var str []string
	for _, value := range categorys {
		str = append(str, value.Categorys)
	}
	return str
}

func CategoryRegister(name string) (bool, error) {
	var v entity.Categorys
	v.Categorys = name
	if err := db.CreateCategorys(&v); err != nil {
		loghelper.Error.Println(err)
		return false, err
	} else {
		return true, err
	}
}
