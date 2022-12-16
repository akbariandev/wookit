package order

import (
	"github.com/akbariandev/wookit/src"
	"github.com/akbariandev/wookit/src/base"
)

type OrderStatus string
type Weekday uint

const (
	OrderStatusAny        OrderStatus = "any"
	OrderStatusPending                = "pending"
	OrderStatusProcessing             = "processing"
	OrderStatusOnHold                 = "on-hold"
	OrderStatusCompleted              = "completed"
	OrderStatusCancelled              = "cancelled"
	OrderStatusRefunded               = "refunded"
	OrderStatusFailed                 = "failed"
	OrderStatusTrash                  = "trash"
)

type OrderBilling struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Company   string `json:"company"`
	Address1  string `json:"address_1"`
	Address2  string `json:"address_2"`
	City      string `json:"city"`
	State     string `json:"state"`
	Postcode  string `json:"postcode"`
	Country   string `json:"country"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type OrderShipping struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Company   string `json:"company"`
	Address1  string `json:"address_1"`
	Address2  string `json:"address_2"`
	City      string `json:"city"`
	State     string `json:"state"`
	Postcode  string `json:"postcode"`
	Country   string `json:"country"`
}

type OrderLineItems struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	ProductID   int           `json:"product_id"`
	VariationID int           `json:"variation_id"`
	Quantity    int           `json:"quantity"`
	TaxClass    string        `json:"tax_class"`
	Subtotal    string        `json:"subtotal"`
	SubtotalTax string        `json:"subtotal_tax"`
	Total       string        `json:"total"`
	TotalTax    string        `json:"total_tax"`
	Taxes       []base.Tax    `json:"taxes"`
	MetaData    []interface{} `json:"meta_data"`
	Sku         string        `json:"sku"`
	Price       float32       `json:"price"`
}

type OrderTaxLine struct {
	ID               int           `json:"id"`
	RateCode         string        `json:"rate_code"`
	RateID           int           `json:"rate_id"`
	Label            string        `json:"label"`
	Compound         bool          `json:"compound"`
	TaxTotal         string        `json:"tax_total"`
	ShippingTaxTotal string        `json:"shipping_tax_total"`
	MetaData         []interface{} `json:"meta_data"`
}

type OrderShippingLine struct {
	ID          int           `json:"id"`
	MethodTitle string        `json:"method_title"`
	MethodID    string        `json:"method_id"`
	Total       string        `json:"total"`
	TotalTax    string        `json:"total_tax"`
	Taxes       []interface{} `json:"taxes"`
	MetaData    []interface{} `json:"meta_data"`
}

type Order struct {
	ID                 int                 `json:"id"`
	ParentID           int                 `json:"parent_id"`
	Number             string              `json:"number"`
	OrderKey           string              `json:"order_key"`
	CreatedVia         string              `json:"created_via"`
	Version            string              `json:"version"`
	Status             string              `json:"status"`
	Currency           string              `json:"currency"`
	DateCreated        string              `json:"date_created"`
	DateCreatedGmt     string              `json:"date_created_gmt"`
	DateModified       string              `json:"date_modified"`
	DateModifiedGmt    string              `json:"date_modified_gmt"`
	DiscountTotal      string              `json:"discount_total"`
	DiscountTax        string              `json:"discount_tax"`
	ShippingTotal      string              `json:"shipping_total"`
	ShippingTax        string              `json:"shipping_tax"`
	CartTax            string              `json:"cart_tax"`
	Total              string              `json:"total"`
	TotalTax           string              `json:"total_tax"`
	PricesIncludeTax   bool                `json:"prices_include_tax"`
	CustomerID         int                 `json:"customer_id"`
	CustomerIPAddress  string              `json:"customer_ip_address"`
	CustomerUserAgent  string              `json:"customer_user_agent"`
	CustomerNote       string              `json:"customer_note"`
	Billing            OrderBilling        `json:"billing"`
	Shipping           OrderShipping       `json:"shipping"`
	PaymentMethod      string              `json:"payment_method"`
	PaymentMethodTitle string              `json:"payment_method_title"`
	TransactionID      string              `json:"transaction_id"`
	DatePaid           string              `json:"date_paid"`
	DatePaidGmt        string              `json:"date_paid_gmt"`
	DateCompleted      interface{}         `json:"date_completed"`
	DateCompletedGmt   interface{}         `json:"date_completed_gmt"`
	CartHash           string              `json:"cart_hash"`
	MetaData           []base.MetaData     `json:"meta_data"`
	LineItems          []OrderLineItems    `json:"line_items"`
	TaxLines           []OrderTaxLine      `json:"tax_lines"`
	ShippingLines      []OrderShippingLine `json:"shipping_lines"`
	FeeLines           []interface{}       `json:"fee_lines"`
	CouponLines        []interface{}       `json:"coupon_lines"`
	Refunds            []interface{}       `json:"refunds"`
	Links              base.Links          `json:"_links"`
}

type GetOrdersListParams struct {
	*src.WooConfig
	Context        base.WooContext `json:"context,omitempty"`
	Page           uint32          `json:"page,omitempty"`
	PerPage        uint8           `json:"per_page,omitempty"`
	Search         string          `json:"search,omitempty"`
	After          string          `json:"after,omitempty"`
	Before         string          `json:"before,omitempty"`
	ModifiedAfter  string          `json:"modified_after,omitempty"`
	ModifiedBefore string          `json:"modified_before,omitempty"`
	DatesAreGmt    bool            `json:"dates_are_gmt,omitempty"`
	Exclude        []string        `json:"exclude,omitempty"`
	Include        []string        `json:"include,omitempty"`
	Offset         uint32          `json:"offset,omitempty"`
	Order          base.Order      `json:"order,omitempty"`
	OrderBy        base.OrderBy    `json:"order_by,omitempty"`
	Code           string          `json:"code,omitempty"`
	Parent         []string        `json:"parent,omitempty"`
	ParentExclude  []string        `json:"parent_exclude,omitempty"`
	Status         []OrderStatus   `json:"status,omitempty"`
	Customer       int32           `json:"customer,omitempty"`
	Product        int32           `json:"product,omitempty"`
	Dp             int32           `json:"dp,omitempty"`
}
