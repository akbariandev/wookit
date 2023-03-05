package order

import (
	"errors"
	httpHandler "github.com/akbariandev/wookit/internal/http"
	"github.com/akbariandev/wookit/pkg/woocommerce"
	jsonIter "github.com/json-iterator/go"
	"io"
	"net/http"
	"strconv"
)

const (
	getOrderListEndPoint = "orders"
)

type OrderService struct {
	*woocommerce.WooConfig
}

type IOrder interface {
	GetOrdersList(params *GetOrdersListParams) (orders []*Order, totalOrders, totalPages int32, err error)
}

func NewOrderService(config *woocommerce.WooConfig) *OrderService {
	return &OrderService{
		config,
	}
}

func (c *OrderService) GetOrdersList(params *GetOrdersListParams) (orders []*Order, totalOrders, totalPages int32, err error) {
	if params == nil {
		params = &GetOrdersListParams{}
	}
	params.WooConfig = c.WooConfig
	request, err := httpHandler.NewGetRequest(c.WooConfig.Address, getOrderListEndPoint, params, nil)
	if err != nil {
		return nil, 0, 0, err
	}

	res, err := request.Call()
	if err != nil {
		return nil, 0, 0, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		var json = jsonIter.ConfigCompatibleWithStandardLibrary
		err = json.NewDecoder(res.Body).Decode(&orders)
		if err != nil {
			return nil, 0, 0, err
		}

		tp, _ := strconv.Atoi(res.Header.Get("X-WP-TotalPages"))
		t, _ := strconv.Atoi(res.Header.Get("X-WP-Total"))
		return orders, int32(tp), int32(t), err
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, 0, 0, err
	}

	return nil, 0, 0, errors.New(string(bodyBytes))
}
