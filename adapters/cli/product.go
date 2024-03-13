package cli

import (
	"fmt"

	"github.com/eduardo6722/go-hexagonal/application"
)

func Run(service application.IProductService, action string, productId string, productName string, price float64) (string, error) {
	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s", product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		service.Enable(product)
		result = fmt.Sprintf("Product with ID %s has been enabled", product.GetID())
	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		service.Disable(product)
		result = fmt.Sprintf("Product with ID %s has been disabled", product.GetID())
	case "get":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s - Name %s - Price %f - Status %s", product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	default:
		return result, fmt.Errorf("invalid action")
	}

	return result, nil
}
