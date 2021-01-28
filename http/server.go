package http

import (
	"github.com/codeedu/esquenta-imersao-go/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type WebServer struct {
	Products *model.Products
}

func NewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {
	e := echo.New()
	e.GET("/products", w.getAll)
	e.POST("/products", w.createProduct)
	e.Logger.Fatal(e.Start(":8080"))
}

func (w WebServer) getAll(c echo.Context) error {
	return c.JSON(http.StatusOK, w.Products)
}

func (w WebServer) createProduct(c echo.Context) error {
	product := model.NewProduct()

	if err := c.Bind(product); err != nil {
		return err
	}
	w.Products.Add(product)
	return c.JSON(http.StatusCreated, product)
}
//
//curl -X POST \
//-H 'Content-Type: application/json' \
//-d '{"name":"Carrinho"}' \
//localhost:8080/products