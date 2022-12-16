package customer

import (
	"github.com/akbariandev/wookit/src"
	"github.com/akbariandev/wookit/src/base"
)

// CustomerRole Limit result set to resources with a specific role. Options: all, administrator, editor, author, contributor, subscriber, customer and shop_manager. Default is customer
type CustomerRole string

const (
	CustomerRoleAll           CustomerRole = "all"
	CustomerRoleAdministrator              = "administrator"
	CustomerRoleEditor                     = "editor"
	CustomerRoleAuthor                     = "author"
	CustomerRoleContributor                = "contributor"
	CustomerRoleSubscriber                 = "subscriber"
	CustomerRoleCustomer                   = "customer"
	CustomerRoleShopManager                = "shop_manager"
)

type Customer struct {
	ID               int           `json:"id"`
	DateCreated      string        `json:"date_created"`
	DateCreatedGmt   string        `json:"date_created_gmt"`
	DateModified     string        `json:"date_modified"`
	DateModifiedGmt  string        `json:"date_modified_gmt"`
	Email            string        `json:"email"`
	FirstName        string        `json:"first_name"`
	LastName         string        `json:"last_name"`
	Role             string        `json:"role"`
	Username         string        `json:"username"`
	Billing          base.Billing  `json:"billing"`
	Shipping         base.Shipping `json:"shipping"`
	IsPayingCustomer bool          `json:"is_paying_customer"`
	AvatarURL        string        `json:"avatar_url"`
	MetaData         []interface{} `json:"meta_data"`
	Links            base.Links    `json:"_links"`
}

type GetCustomerListParams struct {
	*src.WooConfig
	Context base.WooContext `json:"context,omitempty"`
	Page    uint32          `json:"page,omitempty"`
	PerPage uint8           `json:"per_page,omitempty"`
	Search  string          `json:"search,omitempty"`
	Exclude []string        `json:"exclude,omitempty"`
	Include []string        `json:"include,omitempty"`
	Offset  uint32          `json:"offset,omitempty"`
	Order   base.Order      `json:"order,omitempty"`
	OrderBy base.OrderBy    `json:"order_by,omitempty"`
	Email   string          `json:"email,omitempty"`
	Role    CustomerRole    `json:"role,omitempty"`
}
