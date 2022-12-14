package wookit

import (
	"github.com/akbariandev/wookit/src"
	"github.com/akbariandev/wookit/src/coupon"
)

type wooKitServices struct {
	*coupon.CouponService
}

type WooKit interface {
	coupon.ICoupon
}

func NewWooKit(address, cs, ck string) WooKit {
	Config := &src.WooConfig{
		Address: address,
		CS:      cs,
		CK:      ck,
	}

	return &wooKitServices{
		CouponService: coupon.NewCouponService(Config),
	}
}
