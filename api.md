# All

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#AllEndpointSortBy">AllEndpointSortBy</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#Article">Article</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#AllListResponse">AllListResponse</a>

Methods:

- <code title="get /v1/all">client.All.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#AllService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#AllListParams">AllListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#AllListResponse">AllListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Companies

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#CompanyListResponse">CompanyListResponse</a>

Methods:

- <code title="get /v1/companies/all">client.Companies.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#CompanyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#CompanyListParams">CompanyListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#CompanyListResponse">CompanyListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Journalists

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#Journalist">Journalist</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#NameCount">NameCount</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#JournalistListResponse">JournalistListResponse</a>

Methods:

- <code title="get /v1/journalists/{id}">client.Journalists.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#JournalistService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#Journalist">Journalist</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/journalists/all">client.Journalists.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#JournalistService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#JournalistListParams">JournalistListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#JournalistListResponse">JournalistListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# People

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#ImageHolder">ImageHolder</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#WikidataDateHolder">WikidataDateHolder</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#WikidataLabelHolder">WikidataLabelHolder</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#PersonListResponse">PersonListResponse</a>

Methods:

- <code title="get /v1/people/all">client.People.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#PersonService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#PersonListParams">PersonListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#PersonListResponse">PersonListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Sources

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#SortBy">SortBy</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#SourceTopStatHolder">SourceTopStatHolder</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#SourceListResponse">SourceListResponse</a>

Methods:

- <code title="get /v1/sources/all">client.Sources.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#SourceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#SourceListParams">SourceListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#SourceListResponse">SourceListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Stories

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#NewsCluster">NewsCluster</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#StoryListResponse">StoryListResponse</a>

Methods:

- <code title="get /v1/stories/all">client.Stories.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#StoryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#StoryListParams">StoryListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#StoryListResponse">StoryListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Summarize

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#SummarizeNewResponse">SummarizeNewResponse</a>

Methods:

- <code title="post /v1/summarize">client.Summarize.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#SummarizeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#SummarizeNewParams">SummarizeNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#SummarizeNewResponse">SummarizeNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Topics

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#TopicListResponse">TopicListResponse</a>

Methods:

- <code title="get /v1/topics/all">client.Topics.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#TopicService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#TopicListParams">TopicListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#TopicListResponse">TopicListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Vector

## News

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#ArticleSearchFilterParam">ArticleSearchFilterParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#VectorNewsSearchResponse">VectorNewsSearchResponse</a>

Methods:

- <code title="post /v1/vector/news/all">client.Vector.News.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#VectorNewsService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#VectorNewsSearchParams">VectorNewsSearchParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/perigon-sdk-go#VectorNewsSearchResponse">VectorNewsSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
