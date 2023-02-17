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
	require.Equal(t, prod.Price, getProd.Price)
}

func TestListProduct(t *testing.T) {
	prod := createNewProduct(t)
	arg := ListProductsParams{
		Limit:  5,
		Offset: 0,
		UserID: prod.UserID,
	}

	prods, err := testRepo.ListProducts(context.Background(), testDB, arg)
	require.NoError(t, err)
	require.NotNil(t, prods)
	require.GreaterOrEqual(t, len(prods), 1)
}

func TestUpdateProduct(t *testing.T) {
	prod := createNewProduct(t)
	arg := UpdateProductParams{
		Name:  "new product",
		ID:    prod.ID,
		Price: 555,
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
	user := createNewUser(t)
	arg := CreateProductParams{
		Name:   "test product",
		Price:  100,
		UserID: user.ID,
	}

	prod, err := testRepo.CreateProduct(context.Background(), testDB, arg)

	require.NoError(t, err)
	require.NotEmpty(t, prod)

	require.Equal(t, "test product", prod.Name)
	require.Equal(t, int64(100), prod.Price)
	require.Equal(t, user.ID, prod.UserID)

	return prod
}
