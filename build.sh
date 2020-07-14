#!/usr/bin/env bash
export Port=3000
export FruitSupplierUrl="https://my-json-server.typicode.com/2020-abhilash/mock-api/fruits"
export VegetableSupplier="https://my-json-server.typicode.com/2020-abhilash/mock-api/vegetables"
export GrainSupplier="https://my-json-server.typicode.com/2020-abhilash/mock-api/grains"

go install
if [ $? != 0 ]; then
  echo "## Build Failed ##"
  exit
fi

echo "Restoring all vendor versions ..."
godep restore
echo "Done."


echo "Doing some cleaning ..."
go clean
echo "Done."

echo "Running goimport ..."
goimports -w=true .
echo "Done."

echo "Running go vet ..."
go vet ./internal/...
if [ $? != 0 ]; then
  exit
fi
echo "Done."

echo "Running go generate ..."
go generate ./internal/...
echo "Done."

echo "Running go format ..."
gofmt -w .
echo "Done."

echo "Running go build ..."
go build -race
if [ $? != 0 ]; then
  echo "## Build Failed ##"
  exit
fi
echo "Done."

if [ $? == 0 ]; then
    echo "Done."
	echo "## Starting service ##"
    ./food-aggregator
fi
