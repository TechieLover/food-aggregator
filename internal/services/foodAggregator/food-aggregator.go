package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"mbrdi/food-aggregator/dtos"
	"mbrdi/food-aggregator/internal/adapters/fruitsupplier"
	"mbrdi/food-aggregator/internal/adapters/grainsupplier"
	"mbrdi/food-aggregator/internal/adapters/vegetablesupplier"
	"mbrdi/food-aggregator/internal/daos"
	"strconv"
	"strings"
)

type IItemsInfo interface {
	GetItem(name string) (*dtos.Item, error)
	GetItemByNameAndQuantity(name string, quantity int) (*dtos.Item, error)
	GetItemByNameQuantityAndPrice(name string, quantity int, price float64) (*dtos.Item, error)
	ShowSummary() ([]dtos.Item, error)
	GetItemInFastMode(name string) (*dtos.Item, error)
}

type ItemsInfo struct {
	cac daos.IFADaos
	fs  fruitsupplier.IFruitSupplier
	gs  grainsupplier.IGrainSupplier
	vs  vegetablesupplier.IVegSupplier
}

func New() IItemsInfo {
	return &ItemsInfo{
		cac: daos.NewFA(),
		fs:  fruitsupplier.New(),
		gs:  grainsupplier.New(),
		vs:  vegetablesupplier.New(),
	}
}

func (c *ItemsInfo) GetItem(name string) (*dtos.Item, error) {
	finalresp := []dtos.Item{}
	resp, err := c.fs.GetFruitsInformation()
	if err != nil {
		fmt.Println("Error received from fruit seller", err)
	}
	finalresp = append(finalresp, resp...)
	resp, err = c.gs.GetGrainInformation()
	if err != nil {
		fmt.Println("Error received from grain seller", err)
	}
	finalresp = append(finalresp, resp...)
	resp, err = c.vs.GetVegInformation()
	if err != nil {
		fmt.Println("Error received from vegtable seller", err)
	}
	finalresp = append(finalresp, resp...)
	for _, v := range finalresp {
		if strings.ToLower(v.Name) == strings.ToLower(name) {
			return &v, nil
		}
	}
	return nil, errors.New("Not_found")
}

func (c *ItemsInfo) GetItemByNameAndQuantity(name string, quantity int) (*dtos.Item, error) {
	finalresp := []dtos.Item{}
	resp, err := c.fs.GetFruitsInformation()
	if err != nil {
		fmt.Println("Error received from fruit seller", err)
	}
	finalresp = append(finalresp, resp...)
	resp, err = c.gs.GetGrainInformation()
	if err != nil {
		fmt.Println("Error received from grain seller", err)
	}
	finalresp = append(finalresp, resp...)
	resp, err = c.vs.GetVegInformation()
	if err != nil {
		fmt.Println("Error received from vegtable seller", err)
	}
	finalresp = append(finalresp, resp...)
	for _, v := range finalresp {
		if strings.ToLower(v.Name) == strings.ToLower(name) && v.Quantity >= quantity {
			return &v, nil
		}
	}
	return nil, errors.New("Not_found")
}

func (c *ItemsInfo) GetItemByNameQuantityAndPrice(name string, quantity int, price float64) (*dtos.Item, error) {
	val, found := c.cac.GetItem(name)
	if found {
		byteData, _ := json.Marshal(val)
		item := &dtos.Item{}
		json.Unmarshal(byteData, item)
		if strings.ToLower(item.Name) == strings.ToLower(name) && item.Quantity >= quantity && GetPrice(item.Price) <= price {
			return item, nil
		}
		return nil, errors.New("Not_found")
	}
	finalresp := []dtos.Item{}
	resp, err := c.fs.GetFruitsInformation()
	if err != nil {
		fmt.Println("Error received from fruit seller", err)
	}
	finalresp = append(finalresp, resp...)
	resp, err = c.gs.GetGrainInformation()
	if err != nil {
		fmt.Println("Error received from grain seller", err)
	}
	finalresp = append(finalresp, resp...)
	resp, err = c.vs.GetVegInformation()
	if err != nil {
		fmt.Println("Error received from vegtable seller", err)
	}
	finalresp = append(finalresp, resp...)
	for _, v := range finalresp {
		c.cac.AddItem(v)
	}
	val, found = c.cac.GetItem(name)
	if found {
		fmt.Println(val)
		byteData, _ := json.Marshal(val)
		item := &dtos.Item{}
		json.Unmarshal(byteData, item)
		if strings.ToLower(item.Name) == strings.ToLower(name) && item.Quantity >= quantity && GetPrice(item.Price) <= price {
			return item, nil
		}
		return nil, errors.New("Not_found")
	}
	return nil, errors.New("Not_found")
}

func (c *ItemsInfo) ShowSummary() ([]dtos.Item, error) {
	finalresp := []dtos.Item{}
	m := c.cac.CacheStock()
	fmt.Println(m)
	for _, val := range m {
		byteData, _ := json.Marshal(val.Object)
		item := &dtos.Item{}
		json.Unmarshal(byteData, item)
		finalresp = append(finalresp, *item)
	}
	if len(finalresp) == 0 {
		return nil, errors.New("Not_found")
	}
	return finalresp, nil
}

func GetPrice(val string) float64 {
	t := strings.Replace(val, "$", "", -1)
	prc, _ := strconv.ParseFloat(t, 64)
	return prc
}

func (c *ItemsInfo) GetItemInFastMode(name string) (*dtos.Item, error) {
	itemschan := make(chan *dtos.Item)
	itemschan1 := make(chan *dtos.Item)
	itemschan2 := make(chan *dtos.Item)
	go c.fs.FruitInformation(itemschan)
	go c.gs.GrainInformation(itemschan1)
	go c.vs.VegInformation(itemschan2)
	for val := range itemschan {
		if strings.ToLower(val.Name) == strings.ToLower(name) {
			return val, nil
		}
	}

	for val := range itemschan1 {
		if strings.ToLower(val.Name) == strings.ToLower(name) {
			return val, nil
		}
	}

	for val := range itemschan2 {
		if strings.ToLower(val.Name) == strings.ToLower(name) {
			return val, nil
		}
	}
	return nil, errors.New("Not found")
}
