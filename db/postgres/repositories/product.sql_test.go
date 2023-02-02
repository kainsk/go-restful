package repositories

import (
	"context"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestCreateProduct(t *testing.T) {
	createNewProduct(t)
}

func TestGetProduct(t *testing.T) {
	prod := createNewProduct(t)
	getProd, err := testRepo.GetProduct(context.Background(), testDB, prod.ID)

	require.NoError(t, err)
	require.NotEmpty(t, getProd)

	require.Equal(t, prod.ID, getProd.ID)
	require.Equal(t, prod.Name, getProd.Name)
}

func TestListProduct(t *testing.T) {
	createNewProduct(t)
	arg := ListProductsParams{
		Limit:  5,
		Offset: 0,
	}

	prods, err := testRepo.ListProducts(context.Background(), testDB, arg)
	require.NoError(t, err)
	require.NotNil(t, prods)
	require.GreaterOrEqual(t, len(prods), 1)
}

func TestUpdateProduct(t *testing.T) {
	prod := createNewProduct(t)
	arg := UpdateProductParams{
		NewProductName: "new product",
		ProductID:      prod.ID,
	}

	newProd, err := testRepo.UpdateProduct(context.Background(), testDB, arg)
	require.NoError(t, err)
	require.NotEmpty(t, newProd)
	require.Equal(t, prod.ID, newProd.ID)
	require.Equal(t, "new product", newProd.Name)
}

func TestDeleteProduct(t *testing.T) {
	prod := createNewProduct(t)
	err := testRepo.DeleteProduct(context.Background(), testDB, prod.ID)
	require.NoError(t, err)
}

func createNewProduct(t *testing.T) Product {
	prod, err := testRepo.CreateProduct(context.Background(), testDB, "test product")

	require.NoError(t, err)
	require.NotEmpty(t, prod)

	require.Equal(t, "test product", prod.Name)

	return prod
}
