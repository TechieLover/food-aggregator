package fruitsupplier

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

type FruitSupplier struct {
	url string
}

func NewFruitSupplier() *FruitSupplier {
	return &FruitSupplier{
		url: config.FruitSupplierUrl,
	}
}

func (fs *FruitSupplier) FruitInformation(resChan chan *dtos.Item) {
	res, err := fs.GetFruitsInformation()
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

func (fs *FruitSupplier) GetFruitsInformation() ([]dtos.Item, error) {
	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, fs.url, nil)
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

	res := []dtos.Item{}
	err = json.Unmarshal(content, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
