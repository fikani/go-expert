package database

import (
	"app-example/internal/entity"
	pkg_entity "app-example/pkg/entity"
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type ProductDBTestSuite struct {
	suite.Suite
	productDB *Product
	db        *sql.DB
}

func TestProductDBSuite(t *testing.T) {
	suite.Run(t, new(ProductDBTestSuite))
}

func (suite *ProductDBTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	err = CreateProductTables(db)
	if err != nil {
		panic(err)
	}

	suite.db = db
	suite.productDB = NewProduct(db)
}

func (suite *ProductDBTestSuite) TearDownTest() {
	suite.db.Close()
}

func (suite *ProductDBTestSuite) TestProductDB_Create_ShouldPersistProduct() {
	product, _ := entity.NewProduct("Product A", 1000)
	err := suite.productDB.Create(product)
	suite.Nil(err)

	foundProduct, err := suite.productDB.FindByID(product.ID)
	suite.Nil(err)
	suite.Equal(product.ID.String(), foundProduct.ID.String())
	suite.Equal(product.Name, foundProduct.Name)
	suite.Equal(product.Price, foundProduct.Price)
}

func (suite *ProductDBTestSuite) TestProductDB_FindAll_ShouldReturnProducts() {
	productA, _ := entity.NewProduct("Product A", 1000)
	productB, _ := entity.NewProduct("Product B", 2000)
	productC, _ := entity.NewProduct("Product C", 3000)

	suite.productDB.Create(productA)
	suite.productDB.Create(productB)
	suite.productDB.Create(productC)

	products, err := suite.productDB.FindAll(0, 10, "asc")
	suite.Nil(err)
	suite.Len(products, 3)
	suite.Equal(productA.ID.String(), products[0].ID.String())
	suite.Equal(productB.ID.String(), products[1].ID.String())
	suite.Equal(productC.ID.String(), products[2].ID.String())
}

func (suite *ProductDBTestSuite) TestProductDB_FindAll_ShouldReturnProductsInDescendingOrder() {
	productA, _ := entity.NewProduct("Product A", 1000)
	productB, _ := entity.NewProduct("Product B", 2000)
	productC, _ := entity.NewProduct("Product C", 3000)

	suite.productDB.Create(productA)
	suite.productDB.Create(productB)
	suite.productDB.Create(productC)

	products, err := suite.productDB.FindAll(0, 10, "desc")
	suite.Nil(err)
	suite.Len(products, 3)
	suite.Equal(productC.ID.String(), products[0].ID.String())
	suite.Equal(productB.ID.String(), products[1].ID.String())
	suite.Equal(productA.ID.String(), products[2].ID.String())
}

func (suite *ProductDBTestSuite) TestProductDB_FindAll_ShouldReturnProductsInAscendingOrder() {
	productA, _ := entity.NewProduct("Product A", 1000)
	productB, _ := entity.NewProduct("Product B", 2000)
	productC, _ := entity.NewProduct("Product C", 3000)

	suite.productDB.Create(productA)
	suite.productDB.Create(productB)
	suite.productDB.Create(productC)

	products, err := suite.productDB.FindAll(0, 10, "asc")
	suite.Nil(err)
	suite.Len(products, 3)
	suite.Equal(productA.ID.String(), products[0].ID.String())
	suite.Equal(productB.ID.String(), products[1].ID.String())
	suite.Equal(productC.ID.String(), products[2].ID.String())
}

func (suite *ProductDBTestSuite) TestProductDB_FindAll_ShouldReturnProductsInDescendingOrderWithLimit() {
	productA, _ := entity.NewProduct("Product A", 1000)
	productB, _ := entity.NewProduct("Product B", 2000)
	productC, _ := entity.NewProduct("Product C", 3000)

	suite.productDB.Create(productA)
	suite.productDB.Create(productB)
	suite.productDB.Create(productC)

	products, err := suite.productDB.FindAll(0, 2, "desc")
	suite.Nil(err)
	suite.Len(products, 2)
	suite.Equal(productC.ID.String(), products[0].ID.String())
	suite.Equal(productB.ID.String(), products[1].ID.String())
}

func (suite *ProductDBTestSuite) TestProductDB_FindAll_ShouldReturnProductsInDescendingOrderWithLimitAndPage() {
	productA, _ := entity.NewProduct("Product A", 1000)
	productB, _ := entity.NewProduct("Product B", 2000)
	productC, _ := entity.NewProduct("Product C", 3000)

	suite.productDB.Create(productA)
	suite.productDB.Create(productB)
	suite.productDB.Create(productC)

	products, err := suite.productDB.FindAll(1, 1, "desc")
	suite.Nil(err)
	suite.Len(products, 1)
	suite.Equal(productB.ID.String(), products[0].ID.String())
}

func (suite *ProductDBTestSuite) TestProductDB_FindByID_ShouldReturnProduct() {
	product, _ := entity.NewProduct("Product A", 1000)
	suite.productDB.Create(product)

	foundProduct, err := suite.productDB.FindByID(product.ID)
	suite.Nil(err)
	suite.Equal(product.ID.String(), foundProduct.ID.String())
	suite.Equal(product.Name, foundProduct.Name)
	suite.Equal(product.Price, foundProduct.Price)
}

func (suite *ProductDBTestSuite) TestProductDB_FindByID_ShouldReturnErrorWhenProductNotFound() {
	products, err := suite.productDB.FindByID(pkg_entity.NewID())
	suite.NotNil(err)
	suite.Empty(products)
}

func (suite *ProductDBTestSuite) TestProductDB_FindAll_WhenSortIsInvalid_ShouldReturnAsASC() {
	productA, _ := entity.NewProduct("Product A", 1000)
	productB, _ := entity.NewProduct("Product B", 2000)
	productC, _ := entity.NewProduct("Product C", 3000)

	suite.productDB.Create(productA)
	suite.productDB.Create(productB)
	suite.productDB.Create(productC)

	products, err := suite.productDB.FindAll(0, 10, "invalid")
	suite.Nil(err)
	suite.Len(products, 3)
}

func (suite *ProductDBTestSuite) TestProductDB_FindAll_WhenLimitIsZero_ShouldUseDefaultLimit() {
	productA, _ := entity.NewProduct("Product A", 1000)
	productB, _ := entity.NewProduct("Product B", 2000)
	productC, _ := entity.NewProduct("Product C", 3000)

	suite.productDB.Create(productA)
	suite.productDB.Create(productB)
	suite.productDB.Create(productC)

	products, err := suite.productDB.FindAll(0, 0, "asc")
	suite.Nil(err)
	suite.Len(products, 3)
}

func (suite *ProductDBTestSuite) TestProductDB_FindAll_WhenLimitIsGreaterThan100_ShouldUseMaxLimit() {
	for i := 0; i < 101; i++ {
		product, _ := entity.NewProduct("Product A", 1000)
		suite.productDB.Create(product)
	}
	products, err := suite.productDB.FindAll(0, 101, "asc")
	suite.Nil(err)
	suite.Len(products, 100)
}

func (suite *ProductDBTestSuite) TestProductDB_FindAll_WhenPageIsNegative_ShouldUsePageZero() {
	productA, _ := entity.NewProduct("Product A", 1000)
	productB, _ := entity.NewProduct("Product B", 2000)
	productC, _ := entity.NewProduct("Product C", 3000)

	suite.productDB.Create(productA)
	suite.productDB.Create(productB)
	suite.productDB.Create(productC)

	products, err := suite.productDB.FindAll(-1, 2, "asc")
	suite.Nil(err)
	suite.Len(products, 2)
	suite.Equal(productA.ID, products[0].ID)
	suite.Equal(productB.ID, products[1].ID)
}
