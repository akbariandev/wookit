package product

import (
	"github.com/akbariandev/wookit/pkg/woocommerce"
	"github.com/akbariandev/wookit/pkg/woocommerce/base"
)

// BackOrder Options: no, notify and yes. Default is no.
type BackOrder string

// ProductType Options: simple, grouped, external and variable. Default is simple.
type ProductType string

// CatalogVisibility Options: visible, catalog, search and hidden. Default is visible.
type CatalogVisibility string

// TaxStatus Options: taxable, shipping and none. Default is taxable
type TaxStatus string

// StockStatus Options: instock, outofstock, onbackorder. Default is instock.
type StockStatus string

// ProductStatus Options: any, draft, pending, private and publish. Default is any.
type ProductStatus string

const (
	BackOrderNo     BackOrder = "no"
	BackOrderNotify           = "notify"
	BackOrderYes              = "yes"

	ProductTypeSimple   ProductType = "simple"
	ProductTypeGrouped              = "grouped"
	ProductTypeExternal             = "external"
	ProductTypeVariable             = "variable"

	CatalogVisibilityVisible CatalogVisibility = "visible"
	CatalogVisibilityCatalog                   = "catalog"
	CatalogVisibilitySearch                    = "search"
	CatalogVisibilityHidden                    = "hidden"

	TaxStatusTaxable  TaxStatus = "taxable"
	TaxStatusShipping           = "shipping"
	TaxStatusNone               = "none"

	StockStatusInStock     StockStatus = "instock"
	StockStatusOutOfStock              = "outofstock"
	StockStatusOnBackOrder             = "onbackorder"

	ProductStatusAny     ProductStatus = "any"
	ProductStatusDraft                 = "draft"
	ProductStatusPending               = "pending"
	ProductStatusPrivate               = "private"
	ProductStatusPublish               = "publish"
)

type ProductDimension struct {
	Length string `json:"length"`
	Width  string `json:"width"`
	Height string `json:"height"`
}

type ProductDownload struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	File string `json:"file"`
}

type ProductCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type ProductTag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type ProductImage struct {
	ID              int    `json:"id"`
	DateCreated     string `json:"date_created"`
	DateCreatedGmt  string `json:"date_created_gmt"`
	DateModified    string `json:"date_modified"`
	DateModifiedGmt string `json:"date_modified_gmt"`
	Src             string `json:"pkg"`
	Name            string `json:"name"`
	Alt             string `json:"alt"`
}

type ProductAttribute struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Position  int      `json:"position"`
	Visible   bool     `json:"visible"`
	Variation bool     `json:"variation"`
	Options   []string `json:"options"`
}

type ProductDefaultAttribute struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Option string `json:"option"`
}

type Product struct {
	ID                int                       `json:"id"`
	Name              string                    `json:"name"`
	Slug              string                    `json:"slug"`
	Permalink         string                    `json:"permalink"`
	DateCreated       string                    `json:"date_created"`
	DateCreatedGmt    string                    `json:"date_created_gmt"`
	DateModified      string                    `json:"date_modified"`
	DateModifiedGmt   string                    `json:"date_modified_gmt"`
	Type              ProductType               `json:"type"`
	Status            string                    `json:"status"`
	Featured          bool                      `json:"featured"`
	CatalogVisibility CatalogVisibility         `json:"catalog_visibility"`
	Description       string                    `json:"description"`
	ShortDescription  string                    `json:"short_description"`
	Sku               string                    `json:"sku"`
	Price             string                    `json:"price"`
	RegularPrice      string                    `json:"regular_price"`
	SalePrice         string                    `json:"sale_price"`
	DateOnSaleFrom    interface{}               `json:"date_on_sale_from"`
	DateOnSaleFromGmt interface{}               `json:"date_on_sale_from_gmt"`
	DateOnSaleTo      interface{}               `json:"date_on_sale_to"`
	DateOnSaleToGmt   interface{}               `json:"date_on_sale_to_gmt"`
	PriceHTML         string                    `json:"price_html"`
	OnSale            bool                      `json:"on_sale"`
	Purchasable       bool                      `json:"purchasable"`
	TotalSales        interface{}               `json:"total_sales,omitempty"`
	Virtual           bool                      `json:"virtual"`
	Downloadable      bool                      `json:"downloadable"`
	Downloads         []ProductDownload         `json:"downloads"`
	DownloadLimit     int                       `json:"download_limit"`
	DownloadExpiry    int                       `json:"download_expiry"`
	ExternalURL       string                    `json:"external_url"`
	ButtonText        string                    `json:"button_text"`
	TaxStatus         TaxStatus                 `json:"tax_status"`
	TaxClass          string                    `json:"tax_class"`
	ManageStock       bool                      `json:"manage_stock"`
	StockQuantity     interface{}               `json:"stock_quantity"`
	StockStatus       StockStatus               `json:"stock_status"`
	Backorders        BackOrder                 `json:"backorders"`
	BackordersAllowed bool                      `json:"backorders_allowed"`
	BackOrdered       bool                      `json:"backordered"`
	SoldIndividually  bool                      `json:"sold_individually"`
	Weight            string                    `json:"weight"`
	Dimensions        ProductDimension          `json:"dimensions"`
	ShippingRequired  bool                      `json:"shipping_required"`
	ShippingTaxable   bool                      `json:"shipping_taxable"`
	ShippingClass     string                    `json:"shipping_class"`
	ShippingClassID   int                       `json:"shipping_class_id"`
	ReviewsAllowed    bool                      `json:"reviews_allowed"`
	AverageRating     string                    `json:"average_rating"`
	RatingCount       int                       `json:"rating_count"`
	RelatedIds        []int                     `json:"related_ids"`
	UpsellIds         []int                     `json:"upsell_ids"`
	CrossSellIds      []int                     `json:"cross_sell_ids"`
	ParentID          int                       `json:"parent_id"`
	PurchaseNote      string                    `json:"purchase_note"`
	Categories        []ProductCategory         `json:"categories"`
	Tags              []ProductTag              `json:"tags"`
	Images            []ProductImage            `json:"images"`
	Attributes        []ProductAttribute        `json:"attributes"`
	DefaultAttributes []ProductDefaultAttribute `json:"default_attributes"`
	Variations        []int                     `json:"variations"`
	GroupedProducts   []int                     `json:"grouped_products"`
	MenuOrder         int                       `json:"menu_order"`
	MetaData          []base.MetaData           `json:"meta_data"`
	Links             base.Links                `json:"_links"`
}

type GetProductListParams struct {
	*woocommerce.WooConfig
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
	Parent         []string        `json:"parent,omitempty"`
	ParentExclude  []string        `json:"parent_exclude,omitempty"`
	Slug           string          `json:"slug,omitempty"`
	Status         ProductStatus   `json:"status,omitempty"`
	Type           ProductType     `json:"type,omitempty"`
	Sku            string          `json:"sku,omitempty"`
	Featured       bool            `json:"featured,omitempty"`
	Category       string          `json:"category,omitempty"`
	Tag            string          `json:"tag,omitempty"`
	ShippingClass  string          `json:"shipping_class,omitempty"`
	Attribute      string          `json:"attribute,omitempty"`
	AttributeTerm  string          `json:"attribute_term,omitempty"`
	TaxClass       string          `json:"tax_class,omitempty"`
	OnSale         bool            `json:"on_sale,omitempty"`
	MinPrice       string          `json:"min_price,omitempty"`
	MaxPrice       string          `json:"max_price,omitempty"`
	StockStatus    StockStatus     `json:"stock_status,omitempty"`
}
