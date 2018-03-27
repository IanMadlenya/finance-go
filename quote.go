package finance

// QuoteType alias for asset class.
type QuoteType string

// MarketState alias for market state.
type MarketState string

const (
	// QuoteTypeEquity the returned quote should be an equity.
	QuoteTypeEquity QuoteType = "EQUITY"
	// QuoteTypeIndex the returned quote should be an index.
	QuoteTypeIndex QuoteType = "INDEX"
	// QuoteTypeOption the returned quote should be an option contract.
	QuoteTypeOption QuoteType = "OPTION"
	// QuoteTypeCurrency the returned quote should be a currency pair.
	QuoteTypeCurrency QuoteType = "CURRENCY"
	// QuoteTypeFuture the returned quote should be a futures contract.
	QuoteTypeFuture QuoteType = "FUTURE"
	// QuoteTypeETF the returned quote should be an etf.
	QuoteTypeETF QuoteType = "ETF"
	// QuoteTypeMutualFund the returned quote should be an mutual fund.
	QuoteTypeMutualFund QuoteType = "MUTUALFUND"

	// MarketStatePrePre pre-pre market state.
	MarketStatePrePre MarketState = "PREPRE"
	// MarketStatePostPost post-post market state.
	MarketStatePostPost MarketState = "POSTPOST"
)

// QuoteParams determines which quote to
// retrieve and aids in marshalling
// quote responses.
type QuoteParams struct {
	Params    `form:"*"`
	Symbol    string    `form:"symbols"`
	QuoteType QuoteType `form:"-"`
}

// QuoteListParams determines which quotes to fetch.
type QuoteListParams struct {
	Symbols []string
}

// APIResponse wraps a generic api response.
type APIResponse struct {
	*QuoteList
	Error *Error `json:"error"`
}

// QuoteResponse represents a quote response.
type QuoteResponse struct {
	APIResponse `json:"quoteResponse"`
}

// QuoteList is a list of quotes as returned from a quote endpoint.
type QuoteList struct {
	ListMeta
	Values []*Quote `json:"result"`
}

// Quote is the fundamental quote structure.
type Quote struct {
	Language                          string      `json:"language"`
	QuoteType                         QuoteType   `json:"quoteType"`
	QuoteSourceName                   string      `json:"quoteSourceName"`
	Currency                          string      `json:"currency"`
	EsgPopulated                      bool        `json:"esgPopulated"`
	Tradeable                         bool        `json:"tradeable"`
	ShortName                         string      `json:"shortName"`
	ExchangeDataDelayedBy             int         `json:"exchangeDataDelayedBy"`
	PriceHint                         int         `json:"priceHint"`
	PostMarketChangePercent           float64     `json:"postMarketChangePercent"`
	PostMarketTime                    int         `json:"postMarketTime"`
	Market                            string      `json:"market"`
	Exchange                          string      `json:"exchange"`
	RegularMarketPrice                float64     `json:"regularMarketPrice"`
	RegularMarketTime                 int         `json:"regularMarketTime"`
	RegularMarketChange               float64     `json:"regularMarketChange"`
	RegularMarketOpen                 float64     `json:"regularMarketOpen"`
	RegularMarketDayHigh              float64     `json:"regularMarketDayHigh"`
	RegularMarketDayLow               float64     `json:"regularMarketDayLow"`
	RegularMarketVolume               int         `json:"regularMarketVolume"`
	MarketState                       MarketState `json:"marketState"`
	PostMarketPrice                   float64     `json:"postMarketPrice"`
	PostMarketChange                  float64     `json:"postMarketChange"`
	RegularMarketChangePercent        float64     `json:"regularMarketChangePercent"`
	RegularMarketPreviousClose        float64     `json:"regularMarketPreviousClose"`
	Bid                               float64     `json:"bid"`
	Ask                               float64     `json:"ask"`
	BidSize                           int         `json:"bidSize"`
	AskSize                           int         `json:"askSize"`
	FullExchangeName                  string      `json:"fullExchangeName"`
	LongName                          string      `json:"longName"`
	FiftyTwoWeekLowChange             float64     `json:"fiftyTwoWeekLowChange"`
	FiftyTwoWeekLowChangePercent      float64     `json:"fiftyTwoWeekLowChangePercent"`
	FiftyTwoWeekHighChange            float64     `json:"fiftyTwoWeekHighChange"`
	FiftyTwoWeekHighChangePercent     float64     `json:"fiftyTwoWeekHighChangePercent"`
	FiftyTwoWeekLow                   float64     `json:"fiftyTwoWeekLow"`
	FiftyTwoWeekHigh                  float64     `json:"fiftyTwoWeekHigh"`
	DividendDate                      int         `json:"dividendDate"`
	EarningsTimestamp                 int         `json:"earningsTimestamp"`
	EarningsTimestampStart            int         `json:"earningsTimestampStart"`
	EarningsTimestampEnd              int         `json:"earningsTimestampEnd"`
	TrailingAnnualDividendRate        float64     `json:"trailingAnnualDividendRate"`
	TrailingPE                        float64     `json:"trailingPE"`
	TrailingAnnualDividendYield       float64     `json:"trailingAnnualDividendYield"`
	FinancialCurrency                 string      `json:"financialCurrency"`
	AverageDailyVolume3Month          int         `json:"averageDailyVolume3Month"`
	AverageDailyVolume10Day           int         `json:"averageDailyVolume10Day"`
	SharesOutstanding                 int         `json:"sharesOutstanding"`
	BookValue                         float64     `json:"bookValue"`
	FiftyDayAverage                   float64     `json:"fiftyDayAverage"`
	FiftyDayAverageChange             float64     `json:"fiftyDayAverageChange"`
	EpsTrailingTwelveMonths           float64     `json:"epsTrailingTwelveMonths"`
	EpsForward                        float64     `json:"epsForward"`
	FiftyDayAverageChangePercent      float64     `json:"fiftyDayAverageChangePercent"`
	TwoHundredDayAverage              float64     `json:"twoHundredDayAverage"`
	TwoHundredDayAverageChange        float64     `json:"twoHundredDayAverageChange"`
	TwoHundredDayAverageChangePercent float64     `json:"twoHundredDayAverageChangePercent"`
	MarketCap                         int64       `json:"marketCap"`
	ForwardPE                         float64     `json:"forwardPE"`
	PriceToBook                       float64     `json:"priceToBook"`
	SourceInterval                    int         `json:"sourceInterval"`
	ExchangeTimezoneName              string      `json:"exchangeTimezoneName"`
	ExchangeTimezoneShortName         string      `json:"exchangeTimezoneShortName"`
	GmtOffSetMilliseconds             int         `json:"gmtOffSetMilliseconds"`
	Symbol                            string      `json:"symbol"`
}
