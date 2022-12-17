package product

import (
	"encoding/json"
	"errors"
	httpHandler "github.com/akbariandev/wookit/internal/http"
	"github.com/akbariandev/wookit/src"
	"io"
	"net/http"
	"strconv"
)

const (
	getProductListEndPoint = "products"
)

type ProductService struct {
	*src.WooConfig
}

type IProduct interface {
	GetProductsList(params *GetProductListParams) (products []*Product, totalOrders, totalPages int32, err error)
}

func NewProductService(config *src.WooConfig) *ProductService {
	return &ProductService{
		config,
	}
}

func (c *ProductService) GetProductsList(params *GetProductListParams) (products []*Product, totalOrders, totalPages int32, err error) {
	if params == nil {
		params = &GetProductListParams{}
	}
	params.WooConfig = c.WooConfig
	request, err := httpHandler.NewGetRequest(c.WooConfig.Address, getProductListEndPoint, params, nil)
	if err != nil {
		return nil, 0, 0, err
	}

	res, err := request.Call()
	if err != nil {
		return nil, 0, 0, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		//var json = jsonIter.ConfigCompatibleWithStandardLibrary
		err = json.NewDecoder(res.Body).Decode(&products)
		if err != nil {
			return nil, 0, 0, err
		}

		tp, _ := strconv.Atoi(res.Header.Get("X-WP-TotalPages"))
		t, _ := strconv.Atoi(res.Header.Get("X-WP-Total"))
		return products, int32(tp), int32(t), err
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, 0, 0, err
	}

	return nil, 0, 0, errors.New(string(bodyBytes))
}
