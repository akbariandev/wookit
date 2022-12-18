package wookit

import (
	"github.com/akbariandev/wookit/src"
	"github.com/akbariandev/wookit/src/coupon"
	"github.com/akbariandev/wookit/src/customer"
	"github.com/akbariandev/wookit/src/order"
	"github.com/akbariandev/wookit/src/product"
	productAttribute "github.com/akbariandev/wookit/src/product_attribute"
)

type wooKitServices struct {
	*coupon.CouponService
	*order.OrderService
	*customer.CustomerService
	*product.ProductService
	*productAttribute.ProductAttributeService
}

type WooKit interface {
	coupon.ICoupon
	order.IOrder
	customer.ICustomer
	product.IProduct
	productAttribute.IProductAttribute
}

func NewWooKit(address, cs, ck string) WooKit {
	Config := &src.WooConfig{
		Address: address,
		CS:      cs,
		CK:      ck,
	}

	return &wooKitServices{
		CouponService:           coupon.NewCouponService(Config),
		OrderService:            order.NewOrderService(Config),
		CustomerService:         customer.NewCustomerService(Config),
		ProductService:          product.NewProductService(Config),
		ProductAttributeService: productAttribute.NewProductAttributeService(Config),
	}
}
