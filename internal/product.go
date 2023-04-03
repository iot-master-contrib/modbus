package internal

import (
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/lib"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"modbus/model"
	"xorm.io/xorm"
)

var Products lib.Map[model.Product]

func LoadProducts() error {
	var products []*model.Product
	err := db.Engine.Find(&products)
	if err != nil {
		if err == xorm.ErrNotExist {
			return nil
		}
		return err
	}
	for _, m := range products {
		err := LoadProduct(m)
		if err != nil {
			log.Error(err)
		}
	}
	return nil
}

func LoadProduct(m *model.Product) error {
	Products.Store(m.Id, m)
	return nil
}

func GetProduct(id string) *model.Product {
	return Products.Load(id)
}
