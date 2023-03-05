package customer

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
	getCustomerListEndPoint = "customers"
)

type CustomerService struct {
	*woocommerce.WooConfig
}

type ICustomer interface {
	GetCustomersList(params *GetCustomerListParams) (customers []*Customer, totalCustomers, totalPages int32, err error)
}

func NewCustomerService(config *woocommerce.WooConfig) *CustomerService {
	return &CustomerService{
		config,
	}
}

func (c *CustomerService) GetCustomersList(params *GetCustomerListParams) (customers []*Customer, totalCustomers, totalPages int32, err error) {
	if params == nil {
		params = &GetCustomerListParams{}
	}
	params.WooConfig = c.WooConfig
	request, err := httpHandler.NewGetRequest(c.WooConfig.Address, getCustomerListEndPoint, params, nil)
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
		err = json.NewDecoder(res.Body).Decode(&customers)
		if err != nil {
			return nil, 0, 0, err
		}

		tp, _ := strconv.Atoi(res.Header.Get("X-WP-TotalPages"))
		t, _ := strconv.Atoi(res.Header.Get("X-WP-Total"))
		return customers, int32(tp), int32(t), err
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, 0, 0, err
	}

	return nil, 0, 0, errors.New(string(bodyBytes))
}
