// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package perigon

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/goperigon/perigon-go-sdk/v2/internal/apijson"
	"github.com/goperigon/perigon-go-sdk/v2/internal/apiquery"
	"github.com/goperigon/perigon-go-sdk/v2/internal/requestconfig"
	"github.com/goperigon/perigon-go-sdk/v2/option"
	"github.com/goperigon/perigon-go-sdk/v2/packages/param"
	"github.com/goperigon/perigon-go-sdk/v2/packages/respjson"
)

// CompanyService contains methods and other services that help with interacting
// with the perigon API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCompanyService] method instead.
type CompanyService struct {
	Options []option.RequestOption
}

// NewCompanyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCompanyService(opts ...option.RequestOption) (r CompanyService) {
	r = CompanyService{}
	r.Options = opts
	return
}

// Browse or search for companies Perigon tracks using name, domain, ticker symbol,
// industry, and more. Supports Boolean search logic and filtering by metadata such
// as country, exchange, employee count, and IPO date.
func (r *CompanyService) List(ctx context.Context, query CompanyListParams, opts ...option.RequestOption) (res *CompanyListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/companies/all"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Company search result
type CompanyListResponse struct {
	NumResults int64                       `json:"numResults" api:"required"`
	Results    []CompanyListResponseResult `json:"results" api:"required"`
	Status     int64                       `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumResults  respjson.Field
		Results     respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyListResponse) RawJSON() string { return r.JSON.raw }
func (r *CompanyListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyListResponseResult struct {
	ID                string                                `json:"id" api:"nullable"`
	Address           string                                `json:"address" api:"nullable"`
	AltNames          []string                              `json:"altNames" api:"nullable"`
	Ceo               string                                `json:"ceo" api:"nullable"`
	Cik               string                                `json:"cik" api:"nullable"`
	City              string                                `json:"city" api:"nullable"`
	Country           string                                `json:"country" api:"nullable"`
	Cusip             string                                `json:"cusip" api:"nullable"`
	Description       string                                `json:"description" api:"nullable"`
	Domains           []string                              `json:"domains" api:"nullable"`
	Favicon           string                                `json:"favicon" api:"nullable"`
	FullTimeEmployees int64                                 `json:"fullTimeEmployees" api:"nullable"`
	GlobalRank        int64                                 `json:"globalRank" api:"nullable"`
	Industry          string                                `json:"industry" api:"nullable"`
	IsActivelyTrading bool                                  `json:"isActivelyTrading" api:"nullable"`
	IsAdr             bool                                  `json:"isAdr" api:"nullable"`
	IsEtf             bool                                  `json:"isEtf" api:"nullable"`
	IsFund            bool                                  `json:"isFund" api:"nullable"`
	Isin              string                                `json:"isin" api:"nullable"`
	Logo              string                                `json:"logo" api:"nullable"`
	MonthlyVisits     int64                                 `json:"monthlyVisits" api:"nullable"`
	Naics             string                                `json:"naics" api:"nullable"`
	Name              string                                `json:"name" api:"nullable"`
	PrimaryRecordID   string                                `json:"primaryRecordId" api:"nullable"`
	Revenue           string                                `json:"revenue" api:"nullable"`
	Sector            string                                `json:"sector" api:"nullable"`
	Sic               string                                `json:"sic" api:"nullable"`
	State             string                                `json:"state" api:"nullable"`
	Symbols           []CompanyListResponseResultSymbol     `json:"symbols" api:"nullable"`
	UpdatedAt         string                                `json:"updatedAt" api:"nullable"`
	WebResources      CompanyListResponseResultWebResources `json:"webResources" api:"nullable"`
	YearFounded       int64                                 `json:"yearFounded" api:"nullable"`
	Zip               string                                `json:"zip" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Address           respjson.Field
		AltNames          respjson.Field
		Ceo               respjson.Field
		Cik               respjson.Field
		City              respjson.Field
		Country           respjson.Field
		Cusip             respjson.Field
		Description       respjson.Field
		Domains           respjson.Field
		Favicon           respjson.Field
		FullTimeEmployees respjson.Field
		GlobalRank        respjson.Field
		Industry          respjson.Field
		IsActivelyTrading respjson.Field
		IsAdr             respjson.Field
		IsEtf             respjson.Field
		IsFund            respjson.Field
		Isin              respjson.Field
		Logo              respjson.Field
		MonthlyVisits     respjson.Field
		Naics             respjson.Field
		Name              respjson.Field
		PrimaryRecordID   respjson.Field
		Revenue           respjson.Field
		Sector            respjson.Field
		Sic               respjson.Field
		State             respjson.Field
		Symbols           respjson.Field
		UpdatedAt         respjson.Field
		WebResources      respjson.Field
		YearFounded       respjson.Field
		Zip               respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyListResponseResult) RawJSON() string { return r.JSON.raw }
func (r *CompanyListResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyListResponseResultSymbol struct {
	Exchange          string `json:"exchange" api:"nullable"`
	ExchangeShortName string `json:"exchangeShortName" api:"nullable"`
	IpoDate           string `json:"ipoDate" api:"nullable"`
	Symbol            string `json:"symbol" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Exchange          respjson.Field
		ExchangeShortName respjson.Field
		IpoDate           respjson.Field
		Symbol            respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyListResponseResultSymbol) RawJSON() string { return r.JSON.raw }
func (r *CompanyListResponseResultSymbol) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyListResponseResultWebResources struct {
	About     string `json:"about" api:"nullable"`
	Blog      string `json:"blog" api:"nullable"`
	Careers   string `json:"careers" api:"nullable"`
	Events    string `json:"events" api:"nullable"`
	Facebook  string `json:"facebook" api:"nullable"`
	Instagram string `json:"instagram" api:"nullable"`
	Linkedin  string `json:"linkedin" api:"nullable"`
	Medium    string `json:"medium" api:"nullable"`
	Reddit    string `json:"reddit" api:"nullable"`
	Sitemap   string `json:"sitemap" api:"nullable"`
	Threads   string `json:"threads" api:"nullable"`
	Tiktok    string `json:"tiktok" api:"nullable"`
	Updates   string `json:"updates" api:"nullable"`
	Wellfound string `json:"wellfound" api:"nullable"`
	Wikipedia string `json:"wikipedia" api:"nullable"`
	X         string `json:"x" api:"nullable"`
	Youtube   string `json:"youtube" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		About       respjson.Field
		Blog        respjson.Field
		Careers     respjson.Field
		Events      respjson.Field
		Facebook    respjson.Field
		Instagram   respjson.Field
		Linkedin    respjson.Field
		Medium      respjson.Field
		Reddit      respjson.Field
		Sitemap     respjson.Field
		Threads     respjson.Field
		Tiktok      respjson.Field
		Updates     respjson.Field
		Wellfound   respjson.Field
		Wikipedia   respjson.Field
		X           respjson.Field
		Youtube     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyListResponseResultWebResources) RawJSON() string { return r.JSON.raw }
func (r *CompanyListResponseResultWebResources) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyListParams struct {
	// String. Filter by company industry classifications. Supports Boolean operators
	// (AND, OR, NOT), exact phrases with quotes, and wildcards (\* and ?) for flexible
	// searching.
	Industry param.Opt[string] `query:"industry,omitzero" json:"-"`
	// Date. Filter for companies that went public on or after this date. Accepts ISO
	// 8601 format (e.g., 2023-01-01T00:00:00) or yyyy-mm-dd format.
	IpoFrom param.Opt[time.Time] `query:"ipoFrom,omitzero" format:"date-time" json:"-"`
	// Date. Filter for companies that went public on or before this date. Accepts ISO
	// 8601 format (e.g., 2023-12-31T23:59:59) or yyyy-mm-dd format.
	IpoTo param.Opt[time.Time] `query:"ipoTo,omitzero" format:"date-time" json:"-"`
	// String. Search within company names. Supports Boolean operators (AND, OR, NOT),
	// exact phrases with quotes, and wildcards (\* and ?) for flexible searching.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Integer. Filter for companies with at least this many employees.
	NumEmployeesFrom param.Opt[int64] `query:"numEmployeesFrom,omitzero" json:"-"`
	// Integer. Filter for companies with no more than this many employees.
	NumEmployeesTo param.Opt[int64] `query:"numEmployeesTo,omitzero" json:"-"`
	// Integer. The specific page of results to retrieve in the paginated response.
	// Starts at 0.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// String. Primary search query for filtering companies across name, alternative
	// names, domains, and ticker symbols. Supports Boolean operators (AND, OR, NOT),
	// exact phrases with quotes, and wildcards (\* and ?) for flexible searching.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// String. Filter by company sector classifications. Supports Boolean operators
	// (AND, OR, NOT), exact phrases with quotes, and wildcards (\* and ?) for flexible
	// searching.
	Sector param.Opt[string] `query:"sector,omitzero" json:"-"`
	// Integer. The number of companies to return per page in the paginated response.
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	// String Array. Filter by unique company identifiers. Multiple values create an OR
	// filter.
	ID []string `query:"id,omitzero" json:"-"`
	// String Array. Filter by company headquarters country. Multiple values create an
	// OR filter.
	Country []string `query:"country,omitzero" json:"-"`
	// String Array. Filter by company domains or websites (e.g., apple.com,
	// microsoft.com). Multiple values create an OR filter.
	Domain []string `query:"domain,omitzero" json:"-"`
	// String Array. Filter by stock exchange where companies are listed (e.g., NASDAQ,
	// NYSE). Multiple values create an OR filter.
	Exchange []string `query:"exchange,omitzero" json:"-"`
	// String Array. Filter by company stock ticker symbols (e.g., AAPL, MSFT, GOOGL).
	// Multiple values create an OR filter.
	Symbol []string `query:"symbol,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CompanyListParams]'s query parameters as `url.Values`.
func (r CompanyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
