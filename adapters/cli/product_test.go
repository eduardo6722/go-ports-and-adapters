package cli_test

import (
	"fmt"
	"testing"

	"github.com/eduardo6722/go-hexagonal/adapters/cli"
	"github.com/eduardo6722/go-hexagonal/application"
	mock_application "github.com/eduardo6722/go-hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func mockProduct(ctrl *gomock.Controller, id string, name string, price float64, status string) *mock_application.MockIProduct {
	productMock := mock_application.NewMockIProduct(ctrl)

	productMock.EXPECT().GetID().Return(id).AnyTimes()
	productMock.EXPECT().GetName().Return(name).AnyTimes()
	productMock.EXPECT().GetPrice().Return(price).AnyTimes()
	productMock.EXPECT().GetStatus().Return(status).AnyTimes()

	return productMock
}

func mockService(ctrl *gomock.Controller, productMock *mock_application.MockIProduct, id string, name string, price float64) *mock_application.MockIProductService {
	service := mock_application.NewMockIProductService(ctrl)
	service.EXPECT().Create(name, price).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(id).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	return service
}

func productFactory() *application.Product {
	product := application.NewProduct()
	product.Name = "Product Test"
	product.Price = 20.0
	return product
}

func TestRunCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := productFactory()

	productMock := mockProduct(ctrl, product.ID, product.Name, product.Price, product.Status)
	service := mockService(ctrl, productMock, product.ID, product.Name, product.Price)

	expectedResult := fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s", product.ID, product.Name, product.Price, product.Status)
	result, err := cli.Run(service, "create", "", product.Name, product.Price)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)
}

func TestRunEnable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := productFactory()

	productMock := mockProduct(ctrl, product.ID, product.Name, product.Price, product.Status)
	service := mockService(ctrl, productMock, product.ID, product.Name, product.Price)

	expectedResult := fmt.Sprintf("Product with ID %s has been enabled", product.ID)
	result, err := cli.Run(service, "enable", product.ID, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)
}

func TestRunDisable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := productFactory()

	productMock := mockProduct(ctrl, product.ID, product.Name, product.Price, product.Status)
	service := mockService(ctrl, productMock, product.ID, product.Name, product.Price)

	expectedResult := fmt.Sprintf("Product with ID %s has been disabled", product.ID)
	result, err := cli.Run(service, "disable", product.ID, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)
}

func TestRunGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := productFactory()

	productMock := mockProduct(ctrl, product.ID, product.Name, product.Price, product.Status)
	service := mockService(ctrl, productMock, product.ID, product.Name, product.Price)

	expectedResult := fmt.Sprintf("Product ID %s - Name %s - Price %f - Status %s", product.ID, product.Name, product.Price, product.Status)
	result, err := cli.Run(service, "get", product.ID, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)
}

func TestRunInvalidAction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mock_application.NewMockIProductService(ctrl)

	result, err := cli.Run(service, "invalid", "", "", 0)
	require.Error(t, err)
	require.Equal(t, "", result)
}
