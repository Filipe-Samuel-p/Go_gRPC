package repository

import (
	"API_Unary/src/pb/products"
	"fmt"
	"os"

	"google.golang.org/protobuf/proto"
)

type ProductRepository struct {
}

const fileName string = "./products.txt"

func (pr *ProductRepository) loadData() (products.ProductList, error) {
	productList := products.ProductList{}

	data, err := os.ReadFile(fileName)
	if err != nil {
		return productList, fmt.Errorf("Error on read file. error: %v\n", err)
	}

	err = proto.Unmarshal(data, &productList)
	if err != nil {
		return productList, fmt.Errorf("Error on unmarshal. error: %v\n", err)
	}

	return productList, nil
}

func (pr *ProductRepository) saveData(productList products.ProductList) error {

	data, err := proto.Marshal(&productList)
	if err != nil {
		return fmt.Errorf("Error on marshal. Error: %v\n", err)
	}

	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		return fmt.Errorf("Error on write file. error: %v\n", err)
	}

	return nil
}

func (pr *ProductRepository) Create(product products.Product) (products.Product, error) {
	productList, err := pr.loadData()
	if err != nil {
		return product, err
	}

	product.Id = int32(len(productList.Products) + 1)
	productList.Products = append(productList.Products, &product)
	err = pr.saveData(productList)

	return product, nil
}

func (pr *ProductRepository) FindAll() (products.ProductList, error) {
	return pr.loadData()
}
