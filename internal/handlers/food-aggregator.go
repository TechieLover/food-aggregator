package handlers

import (
	services "mbrdi/food-aggregator/internal/services/foodAggregator"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func setFoodAggregatorRoutes(router *httprouter.Router) {
	router.GET("/v1/buy-item/:item_name", ItemDetailsByName)
	router.GET("/v1/buy-item-qty/:item_name", ItemDetailsByNameAndQuant)
	router.GET("/v1/buy-item-qty-price/:item_name", ItemDetailsByNameQuantAndPrice)
	router.GET("/v1/show-summary", ShowSummary)
	router.GET("/v1/fast-buy-item/:item_name", FastBuy)
}

func ItemDetailsByName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	itemName := ps.ByName("item_name")
	if strings.TrimSpace(itemName) == "" {
		writeJSONMessage("Please Provide a item name", ERR_MSG, http.StatusBadRequest, w)
		return
	}
	ser := services.New()
	resp, err := ser.GetItem(itemName)
	if err != nil {
		writeJSONMessage("NOT_FOUND", ERR_MSG, http.StatusBadRequest, w)
		return
	}
	writeJSONStruct(resp, http.StatusOK, w)
	return
}

func ItemDetailsByNameAndQuant(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	itemName := ps.ByName("item_name")
	if strings.TrimSpace(itemName) == "" {
		writeJSONMessage("Please Provide a item name", ERR_MSG, http.StatusBadRequest, w)
		return
	}
	qty := r.FormValue("qty")
	if strings.TrimSpace(qty) == "" {
		writeJSONMessage("Please provide a quantity", ERR_MSG, http.StatusBadRequest, w)
		return
	}
	qtyInt, err := strconv.Atoi(qty)
	if err != nil {
		writeJSONMessage("Please provide a proper quantity", ERR_MSG, http.StatusBadRequest, w)
		return
	}
	ser := services.New()
	resp, err := ser.GetItemByNameAndQuantity(itemName, qtyInt)
	if err != nil {
		writeJSONMessage("NOT_FOUND", ERR_MSG, http.StatusBadRequest, w)
		return
	}
	writeJSONStruct(resp, http.StatusOK, w)
	return
}

func ItemDetailsByNameQuantAndPrice(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	itemName := ps.ByName("item_name")
	if strings.TrimSpace(itemName) == "" {
		writeJSONMessage("Please Provide a item name", ERR_MSG, http.StatusBadRequest, w)
		return
	}
	qty := r.FormValue("qty")
	if strings.TrimSpace(qty) == "" {
		writeJSONMessage("Please provide a quantity", ERR_MSG, http.StatusBadRequest, w)
		return
	}
	qtyInt, err := strconv.Atoi(qty)
	if err != nil {
		writeJSONMessage("Please provide a proper quantity", ERR_MSG, http.StatusBadRequest, w)
		return
	}

	price := r.FormValue("price")
	if strings.TrimSpace(price) == "" {
		writeJSONMessage("Please provide a price", ERR_MSG, http.StatusBadRequest, w)
		return
	}
	prc, err := strconv.ParseFloat(price, 64)
	if err != nil {
		writeJSONMessage("Please provide a proper quantity", ERR_MSG, http.StatusBadRequest, w)
		return
	}
	ser := services.New()
	resp, err := ser.GetItemByNameQuantityAndPrice(itemName, qtyInt, prc)
	if err != nil {
		writeJSONMessage("NOT_FOUND", ERR_MSG, http.StatusBadRequest, w)
		return
	}
	writeJSONStruct(resp, http.StatusOK, w)
	return
}

func ShowSummary(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ser := services.New()
	resp, err := ser.ShowSummary()
	if err != nil {
		writeJSONMessage("NOT_FOUND", ERR_MSG, http.StatusBadRequest, w)
		return
	}
	writeJSONStruct(resp, http.StatusOK, w)
	return
}

func FastBuy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	itemName := ps.ByName("item_name")
	if strings.TrimSpace(itemName) == "" {
		writeJSONMessage("Please Provide a item name", ERR_MSG, http.StatusBadRequest, w)
		return
	}
	ser := services.New()
	resp, err := ser.GetItemInFastMode(itemName)
	if err != nil {
		writeJSONMessage("NOT_FOUND", ERR_MSG, http.StatusBadRequest, w)
		return
	}
	writeJSONStruct(resp, http.StatusOK, w)
	return
}
