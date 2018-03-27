package quote

import (
	"fmt"
	"testing"

	finance "github.com/piquette/finance-go"
	_ "github.com/piquette/finance-go/testing"
	assert "github.com/stretchr/testify/require"
)

func TestQuoteGet(t *testing.T) {
	p := &finance.QuoteParams{
		Symbol:    "AAPL",
		QuoteType: finance.QuoteTypeEquity,
	}
	q, err := Get(p)
	fmt.Println(q)

	assert.Nil(t, err)
	assert.NotNil(t, q)
}

func TestQuoteGetNil(t *testing.T) {
	q, err := Get(nil)
	assert.Nil(t, q)
	assert.NotNil(t, err)
}
