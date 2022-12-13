package coupon

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
	getCouponListEndPoint = "coupons"
)

type CouponService struct {
	*src.WooConfig
}

type ICoupon interface {
	GetCouponsList(params GetCouponsListParams) (coupons []*Coupon, totalProducts, totalPages int32, err error)
}

func NewCouponService(config *src.WooConfig) *CouponService {
	return &CouponService{
		config,
	}
}

func (c *CouponService) GetCouponsList(params GetCouponsListParams) (coupons []*Coupon, totalCoupons, totalPages int32, err error) {
	params.WooConfig = c.WooConfig
	request, err := httpHandler.NewGetRequest(c.WooConfig.Address, getCouponListEndPoint, params, nil)
	if err != nil {
		return nil, 0, 0, err
	}

	res, err := request.Call()
	if err != nil {
		return nil, 0, 0, err
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, 0, 0, err
	}

	if res.StatusCode == http.StatusOK {
		err = json.Unmarshal(bodyBytes, &coupons)
		if err != nil {
			return nil, 0, 0, err
		}
		tp, _ := strconv.Atoi(res.Header.Get("X-WP-TotalPages"))
		t, _ := strconv.Atoi(res.Header.Get("X-WP-Total"))
		return coupons, int32(tp), int32(t), err
	}

	return nil, 0, 0, errors.New(string(bodyBytes))
}
