package fapi_client

type TickerResponse struct {
	Chart struct {
		Result []struct {
			Meta struct {
				Currency             string  `json:"currency"`
				Symbol               string  `json:"symbol"`
				ExchangeName         string  `json:"exchangeName"`
				InstrumentType       string  `json:"instrumentType"`
				FirstTradeDate       int     `json:"firstTradeDate"`
				RegularMarketTime    int     `json:"regularMarketTime"`
				Gmtoffset            int     `json:"gmtoffset"`
				Timezone             string  `json:"timezone"`
				ExchangeTimezoneName string  `json:"exchangeTimezoneName"`
				RegularMarketPrice   float64 `json:"regularMarketPrice"`
				ChartPreviousClose   float64 `json:"chartPreviousClose"`
				PriceHint            int     `json:"priceHint"`
				CurrentTradingPeriod struct {
					Pre struct {
						Timezone  string `json:"timezone"`
						Start     int    `json:"start"`
						End       int    `json:"end"`
						Gmtoffset int    `json:"gmtoffset"`
					} `json:"pre"`
					Regular struct {
						Timezone  string `json:"timezone"`
						Start     int    `json:"start"`
						End       int    `json:"end"`
						Gmtoffset int    `json:"gmtoffset"`
					} `json:"regular"`
					Post struct {
						Timezone  string `json:"timezone"`
						Start     int    `json:"start"`
						End       int    `json:"end"`
						Gmtoffset int    `json:"gmtoffset"`
					} `json:"post"`
				} `json:"currentTradingPeriod"`
				DataGranularity string   `json:"dataGranularity"`
				Range           string   `json:"range"`
				ValidRanges     []string `json:"validRanges"`
			} `json:"meta"`
			Timestamp   []int `json:"timestamp"`
			Comparisons []struct {
				Symbol             string    `json:"symbol"`
				High               []float64 `json:"high"`
				Low                []float64 `json:"low"`
				ChartPreviousClose float64   `json:"chartPreviousClose"`
				Close              []float64 `json:"close"`
				Open               []float64 `json:"open"`
			} `json:"comparisons"`
			Indicators struct {
				Quote []struct {
					Low    []float64 `json:"low"`
					High   []float64 `json:"high"`
					Open   []float64 `json:"open"`
					Close  []float64 `json:"close"`
					Volume []int     `json:"volume"`
				} `json:"quote"`
				Adjclose []struct {
					Adjclose []float64 `json:"adjclose"`
				} `json:"adjclose"`
			} `json:"indicators"`
		} `json:"result"`
		Error interface{} `json:"error"`
	} `json:"chart"`
}

type QuoteResponse struct {
	QuoteResponse struct {
		Error  interface{} `json:"error"`
		Result []struct {
			Ask                               float64 `json:"ask"`
			AskSize                           int     `json:"askSize"`
			AverageDailyVolume10Day           int     `json:"averageDailyVolume10Day"`
			AverageDailyVolume3Month          int     `json:"averageDailyVolume3Month"`
			Bid                               float64 `json:"bid"`
			BidSize                           int     `json:"bidSize"`
			BookValue                         float64 `json:"bookValue"`
			Currency                          string  `json:"currency"`
			DisplayName                       string  `json:"displayName"`
			DividendDate                      int     `json:"dividendDate"`
			EarningsTimestamp                 int     `json:"earningsTimestamp"`
			EarningsTimestampEnd              int     `json:"earningsTimestampEnd"`
			EarningsTimestampStart            int     `json:"earningsTimestampStart"`
			EpsCurrentYear                    float64 `json:"epsCurrentYear"`
			EpsForward                        float64 `json:"epsForward"`
			EpsTrailingTwelveMonths           float64 `json:"epsTrailingTwelveMonths"`
			EsgPopulated                      bool    `json:"esgPopulated"`
			Exchange                          string  `json:"exchange"`
			ExchangeDataDelayedBy             int     `json:"exchangeDataDelayedBy"`
			ExchangeTimezoneName              string  `json:"exchangeTimezoneName"`
			ExchangeTimezoneShortName         string  `json:"exchangeTimezoneShortName"`
			FiftyDayAverage                   float64 `json:"fiftyDayAverage"`
			FiftyDayAverageChange             float64 `json:"fiftyDayAverageChange"`
			FiftyDayAverageChangePercent      float64 `json:"fiftyDayAverageChangePercent"`
			FiftyTwoWeekHigh                  float64 `json:"fiftyTwoWeekHigh"`
			FiftyTwoWeekHighChange            float64 `json:"fiftyTwoWeekHighChange"`
			FiftyTwoWeekHighChangePercent     float64 `json:"fiftyTwoWeekHighChangePercent"`
			FiftyTwoWeekLow                   float64 `json:"fiftyTwoWeekLow"`
			FiftyTwoWeekLowChange             float64 `json:"fiftyTwoWeekLowChange"`
			FiftyTwoWeekLowChangePercent      float64 `json:"fiftyTwoWeekLowChangePercent"`
			FiftyTwoWeekRange                 string  `json:"fiftyTwoWeekRange"`
			FinancialCurrency                 string  `json:"financialCurrency"`
			FirstTradeDateMilliseconds        int64   `json:"firstTradeDateMilliseconds"`
			ForwardPE                         float64 `json:"forwardPE"`
			FullExchangeName                  string  `json:"fullExchangeName"`
			GmtOffSetMilliseconds             int     `json:"gmtOffSetMilliseconds"`
			Language                          string  `json:"language"`
			LongName                          string  `json:"longName"`
			Market                            string  `json:"market"`
			MarketCap                         int64   `json:"marketCap"`
			MarketState                       string  `json:"marketState"`
			MessageBoardID                    string  `json:"messageBoardId"`
			PostMarketChange                  float64 `json:"postMarketChange"`
			PostMarketChangePercent           float64 `json:"postMarketChangePercent"`
			PostMarketPrice                   float64 `json:"postMarketPrice"`
			PostMarketTime                    int     `json:"postMarketTime"`
			PriceEpsCurrentYear               float64 `json:"priceEpsCurrentYear"`
			PriceHint                         int     `json:"priceHint"`
			PriceToBook                       float64 `json:"priceToBook"`
			QuoteSourceName                   string  `json:"quoteSourceName"`
			QuoteType                         string  `json:"quoteType"`
			Region                            string  `json:"region"`
			RegularMarketChange               float64 `json:"regularMarketChange"`
			RegularMarketChangePercent        float64 `json:"regularMarketChangePercent"`
			RegularMarketDayHigh              float64 `json:"regularMarketDayHigh"`
			RegularMarketDayLow               float64 `json:"regularMarketDayLow"`
			RegularMarketDayRange             string  `json:"regularMarketDayRange"`
			RegularMarketOpen                 float64 `json:"regularMarketOpen"`
			RegularMarketPreviousClose        float64 `json:"regularMarketPreviousClose"`
			RegularMarketPrice                float64 `json:"regularMarketPrice"`
			RegularMarketTime                 int     `json:"regularMarketTime"`
			RegularMarketVolume               int     `json:"regularMarketVolume"`
			SharesOutstanding                 int64   `json:"sharesOutstanding"`
			ShortName                         string  `json:"shortName"`
			SourceInterval                    int     `json:"sourceInterval"`
			Symbol                            string  `json:"symbol"`
			Tradeable                         bool    `json:"tradeable"`
			TrailingAnnualDividendRate        float64 `json:"trailingAnnualDividendRate"`
			TrailingAnnualDividendYield       float64 `json:"trailingAnnualDividendYield"`
			TrailingPE                        float64 `json:"trailingPE"`
			Triggerable                       bool    `json:"triggerable"`
			TwoHundredDayAverage              float64 `json:"twoHundredDayAverage"`
			TwoHundredDayAverageChange        float64 `json:"twoHundredDayAverageChange"`
			TwoHundredDayAverageChangePercent float64 `json:"twoHundredDayAverageChangePercent"`
		} `json:"result"`
	} `json:"quoteResponse"`
}
