package daos

import (
	"mbrdi/food-aggregator/dtos"

	cache "github.com/patrickmn/go-cache"
)

type IFADaos interface {
	GetItem(itemName string) (interface{}, bool)
	AddItem(v dtos.Item)
	CacheStock() map[string]cache.Item
}

type ItemInfo struct {
}

func NewFA() IFADaos {
	return &ItemInfo{}
}

func (c *ItemInfo) GetItem(itemName string) (interface{}, bool) {
	ch := New()
	return ch.Get(itemName)
}

func (c *ItemInfo) AddItem(v dtos.Item) {
	ch := New()
	ch.Add(v.Name, v, cache.NoExpiration)
}

func (c *ItemInfo) CacheStock() map[string]cache.Item {
	ch := New()
	return ch.Items()
}
