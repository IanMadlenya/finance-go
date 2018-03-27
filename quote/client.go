package quote

import (
	finance "github.com/piquette/finance-go"
	"github.com/piquette/finance-go/form"
)

// Client is used to invoke /quote APIs.
type Client struct {
	B finance.Backend
}

// Get returns a quote that matches the parameters specified.
func Get(params *finance.QuoteParams) (*finance.Quote, error) {
	return getC().Get(params)
}

// Get returns a quote that matches the parameters specified.
func (c Client) Get(params *finance.QuoteParams) (*finance.Quote, error) {

	if params == nil {
		return nil, &finance.Error{
			Code:        finance.ErrorCodeArguments,
			Description: finance.ErrorDescriptionSymbols,
		}
	}

	var body *form.Values
	var commonParams *finance.Params

	commonParams = &params.Params
	body = &form.Values{}
	form.AppendTo(body, params)

	resp := &finance.QuoteResponse{}
	err := c.B.Call("GET", "/v7/finance/quote", body, commonParams, resp)
	if err != nil {
		return nil, err
	}

	return resp.Values[0], resp.Error
}

// List returns several quotes.
func List(params *finance.QuoteListParams) *Iter {
	return getC().List(params)
}

// List returns several quotes.
func (c Client) List(params *finance.QuoteListParams) *Iter {
	return &Iter{}
}

// Iter is an iterator for a list of quotes.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*finance.Iter
}

// Quote returns the most recent Quote
// visited by a call to Next.
func (i *Iter) Quote() *finance.Quote {
	return i.Current().(*finance.Quote)
}

func getC() Client {
	return Client{finance.GetBackend(finance.YahooBackend)}
}
