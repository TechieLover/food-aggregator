package config

import (
	"os"
)

var (
	Port              string
	FruitSupplierUrl  string
	GrainSupplier     string
	VegetableSupplier string
)

func init() {
	Port = os.Getenv("Port")
	FruitSupplierUrl = os.Getenv("FruitSupplierUrl")
	GrainSupplier = os.Getenv("GrainSupplier")
	VegetableSupplier = os.Getenv("VegetableSupplier")
}
