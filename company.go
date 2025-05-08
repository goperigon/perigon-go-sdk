// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package perigon

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/goperigon/perigon-go-sdk/internal/apijson"
	"github.com/goperigon/perigon-go-sdk/internal/apiquery"
	"github.com/goperigon/perigon-go-sdk/internal/requestconfig"
	"github.com/goperigon/perigon-go-sdk/option"
	"github.com/goperigon/perigon-go-sdk/packages/param"
	"github.com/goperigon/perigon-go-sdk/packages/respjson"
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
	opts = append(r.Options[:], opts...)
	path := "v1/companies/all"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Company search result
type CompanyListResponse struct {
	NumResults int64                       `json:"numResults,required"`
	Results    []CompanyListResponseResult `json:"results,required"`
	Status     int64                       `json:"status,required"`
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
	ID                string                                `json:"id,nullable"`
	Address           string                                `json:"address,nullable"`
	AltNames          []string                              `json:"altNames,nullable"`
	Ceo               string                                `json:"ceo,nullable"`
	City              string                                `json:"city,nullable"`
	Country           string                                `json:"country,nullable"`
	Description       string                                `json:"description,nullable"`
	Domains           []string                              `json:"domains,nullable"`
	Favicon           string                                `json:"favicon,nullable"`
	FullTimeEmployees int64                                 `json:"fullTimeEmployees,nullable"`
	GlobalRank        int64                                 `json:"globalRank,nullable"`
	Industry          string                                `json:"industry,nullable"`
	IsActivelyTrading bool                                  `json:"isActivelyTrading,nullable"`
	IsAdr             bool                                  `json:"isAdr,nullable"`
	IsEtf             bool                                  `json:"isEtf,nullable"`
	IsFund            bool                                  `json:"isFund,nullable"`
	Logo              string                                `json:"logo,nullable"`
	MonthlyVisits     int64                                 `json:"monthlyVisits,nullable"`
	Naics             string                                `json:"naics,nullable"`
	Name              string                                `json:"name,nullable"`
	PrimaryRecordID   string                                `json:"primaryRecordId,nullable"`
	Revenue           string                                `json:"revenue,nullable"`
	Sector            string                                `json:"sector,nullable"`
	Sic               string                                `json:"sic,nullable"`
	State             string                                `json:"state,nullable"`
	Symbols           []CompanyListResponseResultSymbol     `json:"symbols,nullable"`
	UpdatedAt         string                                `json:"updatedAt,nullable"`
	WebResources      CompanyListResponseResultWebResources `json:"webResources"`
	YearFounded       int64                                 `json:"yearFounded,nullable"`
	Zip               string                                `json:"zip,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Address           respjson.Field
		AltNames          respjson.Field
		Ceo               respjson.Field
		City              respjson.Field
		Country           respjson.Field
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
	Exchange          string `json:"exchange,nullable"`
	ExchangeShortName string `json:"exchangeShortName,nullable"`
	IpoDate           string `json:"ipoDate,nullable"`
	Symbol            string `json:"symbol,nullable"`
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
	About     string `json:"about,nullable"`
	Blog      string `json:"blog,nullable"`
	Careers   string `json:"careers,nullable"`
	Events    string `json:"events,nullable"`
	Facebook  string `json:"facebook,nullable"`
	Instagram string `json:"instagram,nullable"`
	Linkedin  string `json:"linkedin,nullable"`
	Medium    string `json:"medium,nullable"`
	Reddit    string `json:"reddit,nullable"`
	Sitemap   string `json:"sitemap,nullable"`
	Threads   string `json:"threads,nullable"`
	Tiktok    string `json:"tiktok,nullable"`
	Updates   string `json:"updates,nullable"`
	Wellfound string `json:"wellfound,nullable"`
	Wikipedia string `json:"wikipedia,nullable"`
	X         string `json:"x,nullable"`
	Youtube   string `json:"youtube,nullable"`
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
	// Filter by company industry classifications. Supports Boolean operators (AND, OR,
	// NOT), exact phrases with quotes, and wildcards (\* and ?) for flexible
	// searching.
	Industry param.Opt[string] `query:"industry,omitzero" json:"-"`
	// Filter for companies that went public on or after this date. Accepts ISO 8601
	// format (e.g., 2023-01-01T00:00:00) or yyyy-mm-dd format.
	IpoFrom param.Opt[time.Time] `query:"ipoFrom,omitzero" format:"date-time" json:"-"`
	// Filter for companies that went public on or before this date. Accepts ISO 8601
	// format (e.g., 2023-12-31T23:59:59) or yyyy-mm-dd format.
	IpoTo param.Opt[time.Time] `query:"ipoTo,omitzero" format:"date-time" json:"-"`
	// Search within company names. Supports Boolean operators (AND, OR, NOT), exact
	// phrases with quotes, and wildcards (\* and ?) for flexible searching.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Filter for companies with at least this many employees.
	NumEmployeesFrom param.Opt[int64] `query:"numEmployeesFrom,omitzero" json:"-"`
	// Filter for companies with no more than this many employees.
	NumEmployeesTo param.Opt[int64] `query:"numEmployeesTo,omitzero" json:"-"`
	// The specific page of results to retrieve in the paginated response. Starts at 0.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Primary search query for filtering companies across name, alternative names,
	// domains, and ticker symbols. Supports Boolean operators (AND, OR, NOT), exact
	// phrases with quotes, and wildcards (\* and ?) for flexible searching.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Filter by company sector classifications. Supports Boolean operators (AND, OR,
	// NOT), exact phrases with quotes, and wildcards (\* and ?) for flexible
	// searching.
	Sector param.Opt[string] `query:"sector,omitzero" json:"-"`
	// The number of companies to return per page in the paginated response.
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	// Filter by unique company identifiers. Multiple values create an OR filter.
	ID []string `query:"id,omitzero" json:"-"`
	// Filter by company headquarters country. Multiple values create an OR filter.
	Country []string `query:"country,omitzero" json:"-"`
	// Filter by company domains or websites (e.g., apple.com, microsoft.com). Multiple
	// values create an OR filter.
	Domain []string `query:"domain,omitzero" json:"-"`
	// Filter by stock exchange where companies are listed (e.g., NASDAQ, NYSE).
	// Multiple values create an OR filter.
	Exchange []string `query:"exchange,omitzero" json:"-"`
	// Filter by company stock ticker symbols (e.g., AAPL, MSFT, GOOGL). Multiple
	// values create an OR filter.
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
