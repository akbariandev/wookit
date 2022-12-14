package coupon

import (
	"github.com/akbariandev/wookit/src"
	"github.com/akbariandev/wookit/src/base"
)

type Coupon struct {
	ID                        int           `json:"id"`
	Code                      string        `json:"code"`
	Amount                    string        `json:"amount"`
	DateCreated               string        `json:"date_created"`
	DateCreatedGmt            string        `json:"date_created_gmt"`
	DateModified              string        `json:"date_modified"`
	DateModifiedGmt           string        `json:"date_modified_gmt"`
	DiscountType              string        `json:"discount_type"`
	Description               string        `json:"description"`
	DateExpires               interface{}   `json:"date_expires"`
	DateExpiresGmt            interface{}   `json:"date_expires_gmt"`
	UsageCount                int           `json:"usage_count"`
	IndividualUse             bool          `json:"individual_use"`
	ProductIds                []interface{} `json:"product_ids"`
	ExcludedProductIds        []interface{} `json:"excluded_product_ids"`
	UsageLimit                interface{}   `json:"usage_limit"`
	UsageLimitPerUser         interface{}   `json:"usage_limit_per_user"`
	LimitUsageToXItems        interface{}   `json:"limit_usage_to_x_items"`
	FreeShipping              bool          `json:"free_shipping"`
	ProductCategories         []interface{} `json:"product_categories"`
	ExcludedProductCategories []interface{} `json:"excluded_product_categories"`
	ExcludeSaleItems          bool          `json:"exclude_sale_items"`
	MinimumAmount             string        `json:"minimum_amount"`
	MaximumAmount             string        `json:"maximum_amount"`
	EmailRestrictions         []interface{} `json:"email_restrictions"`
	UsedBy                    []interface{} `json:"used_by"`
	MetaData                  []interface{} `json:"meta_data"`
	Links                     base.Links    `json:"_links"`
}

type GetCouponsListParams struct {
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
}
