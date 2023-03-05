package productAttribute

import (
	"github.com/akbariandev/wookit/pkg/woocommerce"
)

// ProductAttributeOrder Options: menu_order, name, name_num and id. Default is menu_order.
type ProductAttributeOrder string

const (
	ProductAttributeOrderMenuOrder ProductAttributeOrder = "menu_order"
	ProductAttributeOrderName                            = "name"
	ProductAttributeOrderNameNum                         = "name_num"
	ProductAttributeOrderId                              = "id"
)

type ProductAttribute struct {
	ID          int                   `json:"id"`
	Name        string                `json:"name"`
	Slug        string                `json:"slug"`
	Type        string                `json:"type"`
	OrderBy     ProductAttributeOrder `json:"order_by"`
	HasArchives bool                  `json:"has_archives"`
}

type GetProductAttributeListParams struct {
	*woocommerce.WooConfig
}
