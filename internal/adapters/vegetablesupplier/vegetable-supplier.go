package vegetablesupplier

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

type VegetableSupplier struct {
	url string
}

func NewVegetableSupplier() *VegetableSupplier {
	return &VegetableSupplier{
		url: config.VegetableSupplier,
	}
}

func (vs *VegetableSupplier) VegInformation(resChan chan *dtos.Item) {
	res, err := vs.GetVegInformation()
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

func (vs *VegetableSupplier) GetVegInformation() ([]dtos.Item, error) {
	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, vs.url, nil)
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

	res := []dtos.VegetableItem{}
	finalres := []dtos.Item{}
	iteminfo := dtos.Item{}
	err = json.Unmarshal(content, &res)
	if err != nil {
		return nil, err
	}
	for _, v := range res {
		iteminfo.ID = v.ProductID
		iteminfo.Name = v.ProductName
		iteminfo.Price = v.Price
		iteminfo.Quantity = v.Quantity
		finalres = append(finalres, iteminfo)
	}
	return finalres, nil
}
