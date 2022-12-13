package coupon

import (
	"github.com/akbariandev/wookit/src"
	"github.com/akbariandev/wookit/src/base"
)

type WooContext string
type Order string
type OrderBy string

const (
	view WooContext = "view"
	edit            = "edit"

	asc  Order = "asc"
	desc       = "desc"

	date    OrderBy = "date"
	id              = "id"
	include         = "include"
	title           = "title"
	slug            = "slug"
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
	context         WooContext `json:"context,omitempty"`
	page            uint32     `json:"page,omitempty"`
	per_page        uint8      `json:"per_Page,omitempty"`
	search          string     `json:"search,omitempty"`
	after           string     `json:"after,omitempty"`
	before          string     `json:"before,omitempty"`
	modified_after  string     `json:"modified_After,omitempty"`
	modified_before string     `json:"modified_Before,omitempty"`
	dates_are_gmt   bool       `json:"dates_Are_Gmt,omitempty"`
	exclude         []string   `json:"exclude,omitempty"`
	include         []string   `json:"include,omitempty"`
	offset          uint32     `json:"offset,omitempty"`
	order           Order      `json:"order,omitempty"`
	orderby         OrderBy    `json:"orderby,omitempty"`
	code            string     `json:"code,omitempty"`
}
