package grainsupplier

import "mbrdi/food-aggregator/dtos"

type IGrainSupplier interface {
	GetGrainInformation() ([]dtos.Item, error)
	GrainInformation(resChan chan *dtos.Item)
}

func New() IGrainSupplier {
	return NewGrainSupplier()
}
