package woocommerce

type WooConfig struct {
	Address string `json:"-"`
	CS      string `json:"consumer_secret"`
	CK      string `json:"consumer_key"`
}
