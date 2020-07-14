package vegetablesupplier

import "mbrdi/food-aggregator/dtos"

type IVegSupplier interface {
	GetVegInformation() ([]dtos.Item, error)
	VegInformation(resChan chan *dtos.Item)
}

func New() IVegSupplier {
	return NewVegetableSupplier()
}
