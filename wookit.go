package wookit

import (
	"github.com/akbariandev/wookit/src"
	"github.com/akbariandev/wookit/src/coupon"
	"github.com/akbariandev/wookit/src/order"
)

type wooKitServices struct {
	*coupon.CouponService
	*order.OrderService
}

type WooKit interface {
	coupon.ICoupon
	order.IOrder
}

func NewWooKit(address, cs, ck string) WooKit {
	Config := &src.WooConfig{
		Address: address,
		CS:      cs,
		CK:      ck,
	}

	return &wooKitServices{
		CouponService: coupon.NewCouponService(Config),
		OrderService:  order.NewOrderService(Config),
	}
}
