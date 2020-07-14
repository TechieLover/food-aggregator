package grainsupplier

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mbrdi/food-aggregator/dtos"
	"mbrdi/food-aggregator/internal/config"
	"net/http"
)

var (
	ErrNot200 = errors.New("Status not OK")
)

type GrainSupplier struct {
	url string
}

func NewGrainSupplier() *GrainSupplier {
	return &GrainSupplier{
		url: config.GrainSupplier,
	}
}

func (gs *GrainSupplier) GrainInformation(resChan chan *dtos.Item) {
	res, err := gs.GetGrainInformation()
	if err != nil {
		fmt.Println("Error in getting a data")
		return
	}
	go func(resChan chan *dtos.Item) {
		defer close(resChan)
		for i := 0; i < len(res); i++ {
			resChan <- &res[i]
		}
	}(resChan)
}

func (gs *GrainSupplier) GetGrainInformation() ([]dtos.Item, error) {
	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, gs.url, nil)
	if err != nil {
		return nil, err
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, ErrNot200
	}

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	res := []dtos.GrainItem{}
	finalres := []dtos.Item{}
	iteminfo := dtos.Item{}
	err = json.Unmarshal(content, &res)
	if err != nil {
		return nil, err
	}
	for _, v := range res {
		iteminfo.ID = v.ItemID
		iteminfo.Name = v.ItemName
		iteminfo.Price = v.Price
		iteminfo.Quantity = v.Quantity
		finalres = append(finalres, iteminfo)
	}
	return finalres, nil
}
