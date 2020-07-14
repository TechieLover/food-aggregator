package fruitsupplier

import "mbrdi/food-aggregator/dtos"

type IFruitSupplier interface {
	GetFruitsInformation() ([]dtos.Item, error)
	FruitInformation(resChan chan *dtos.Item)
}

func New() IFruitSupplier {
	return NewFruitSupplier()
}
