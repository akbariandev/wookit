package productAttribute

import (
	"encoding/json"
	"errors"
	httpHandler "github.com/akbariandev/wookit/internal/http"
	"github.com/akbariandev/wookit/pkg/woocommerce"
	"io"
	"net/http"
	"strconv"
)

const (
	getProductAttributeListEndPoint = "products/attributes"
)

type ProductAttributeService struct {
	*woocommerce.WooConfig
}

type IProductAttribute interface {
	GetProductAttributeList(*GetProductAttributeListParams) ([]*ProductAttribute, int32, int32, error)
}

func NewProductAttributeService(config *woocommerce.WooConfig) *ProductAttributeService {
	return &ProductAttributeService{
		config,
	}
}

func (c *ProductAttributeService) GetProductAttributeList(params *GetProductAttributeListParams) (productAttributes []*ProductAttribute, totalOrders, totalPages int32, err error) {
	if params == nil {
		params = &GetProductAttributeListParams{}
	}
	params.WooConfig = c.WooConfig
	request, err := httpHandler.NewGetRequest(c.WooConfig.Address, getProductAttributeListEndPoint, params, nil)
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
		err = json.NewDecoder(res.Body).Decode(&productAttributes)
		if err != nil {
			return nil, 0, 0, err
		}

		tp, _ := strconv.Atoi(res.Header.Get("X-WP-TotalPages"))
		t, _ := strconv.Atoi(res.Header.Get("X-WP-Total"))
		return productAttributes, int32(tp), int32(t), err
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, 0, 0, err
	}

	return nil, 0, 0, errors.New(string(bodyBytes))
}
