package base

type WooContext string
type Order string
type OrderBy string

const (
	ContextView WooContext = "view"
	ContextEdit            = "edit"

	OrderAsc  Order = "asc"
	OrderDesc       = "desc"

	OrderDate    OrderBy = "date"
	OrderId              = "id"
	OrderInclude         = "include"
	OrderTitle           = "title"
	OrderSlug            = "slug"
)

const (
	GetContinentsEndpoint = "data/continents"
	GetCurrenciesEndpoint = "data/currencies"
)

type Continent struct {
	Code      string    `json:"code,omitempty" bson:"c,omitempty"`
	Name      string    `json:"name,omitempty" bson:"n,omitempty"`
	Countries []Country `json:"countries,omitempty"  bson:"cr,omitempty"`
	Links     Links     `json:"_links,omitempty" bson:"l,omitempty"`
}

type State struct {
	Code interface{} `json:"code,omitempty" bson:"c,omitempty"`
	Name string      `json:"name,omitempty" bson:"n,omitempty"`
}

type Country struct {
	Code          string  `json:"code,omitempty" bson:"c,omitempty"`
	Name          string  `json:"name,omitempty" bson:"n,omitempty"`
	CurrencyCode  string  `json:"currency_code,omitempty" bson:"cc,omitempty"`
	CurrencyPos   string  `json:"currency_pos,omitempty" bson:"cp,omitempty"`
	DecimalSep    string  `json:"decimal_sep,omitempty" bson:"ds,omitempty"`
	DimensionUnit string  `json:"dimension_unit,omitempty" bson:"du,omitempty"`
	NumDecimals   int     `json:"num_decimals,omitempty" bson:"nd,omitempty"`
	ThousandSep   string  `json:"thousand_sep,omitempty" bson:"ts,omitempty"`
	WeightUnit    string  `json:"weight_unit,omitempty" bson:"wu,omitempty"`
	States        []State `json:"states,omitempty" bson:"s,omitempty"`
}

type Self struct {
	Href string `json:"href,omitempty" bson:"h,omitempty"`
}

type Collection struct {
	Href string `json:"href,omitempty" bson:"h,omitempty"`
}

type Links struct {
	Self       []Self       `json:"self,omitempty" bson:"s,omitempty"`
	Collection []Collection `json:"collection,omitempty" bson:"c,omitempty"`
}

type MetaData struct {
	ID    int    `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Tax struct {
	ID       int    `json:"id"`
	Total    string `json:"total"`
	Subtotal string `json:"subtotal"`
}

type Billing struct {
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

type Shipping struct {
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

type Currency struct {
	Code   string `json:"code,omitempty" bson:"c,omitempty"`
	Name   string `json:"name,omitempty" bson:"n,omitempty"`
	Symbol string `json:"symbol,omitempty" bson:"s,omitempty"`
	Links  Links  `json:"_links,omitempty" bson:"l,omitempty"`
}

/*
func (w *wookit.WooKit) GetContinents() (continents *[]Continent, err error) {
	url := makeUrl(w, GetContinentsEndpoint)
	res, err := http.NewRequest(http2.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http2.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bodyBytes, &continents)
		if err != nil {
			return nil, err
		}
		return continents, err
	}

	return nil, handleResponse(res)
}

func (w *wookit.WooKit) GetCurrencies() (currencies *[]Currency, err error) {
	url := makeUrl(w, GetCurrenciesEndpoint)
	res, err := http.NewRequest(http2.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http2.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bodyBytes, &currencies)
		if err != nil {
			return nil, err
		}
		return currencies, err
	}

	return nil, handleResponse(res)
}
*/
