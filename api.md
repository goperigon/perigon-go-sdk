# All

Params Types:

- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#AllEndpointSortBy">AllEndpointSortBy</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#Article">Article</a>
- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#AllListResponse">AllListResponse</a>

Methods:

- <code title="get /v1/all">client.All.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#AllService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#AllListParams">AllListParams</a>) (<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#AllListResponse">AllListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Companies

Response Types:

- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#CompanyListResponse">CompanyListResponse</a>

Methods:

- <code title="get /v1/companies/all">client.Companies.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#CompanyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#CompanyListParams">CompanyListParams</a>) (<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#CompanyListResponse">CompanyListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Journalists

Response Types:

- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#Journalist">Journalist</a>
- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#NameCount">NameCount</a>
- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#JournalistListResponse">JournalistListResponse</a>

Methods:

- <code title="get /v1/journalists/{id}">client.Journalists.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#JournalistService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#Journalist">Journalist</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/journalists/all">client.Journalists.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#JournalistService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#JournalistListParams">JournalistListParams</a>) (<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#JournalistListResponse">JournalistListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# People

Response Types:

- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#ImageHolder">ImageHolder</a>
- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#WikidataDateHolder">WikidataDateHolder</a>
- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#WikidataLabelHolder">WikidataLabelHolder</a>
- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#PersonListResponse">PersonListResponse</a>

Methods:

- <code title="get /v1/people/all">client.People.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#PersonService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#PersonListParams">PersonListParams</a>) (<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#PersonListResponse">PersonListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Sources

Params Types:

- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#SortBy">SortBy</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#SourceTopStatHolder">SourceTopStatHolder</a>
- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#SourceListResponse">SourceListResponse</a>

Methods:

- <code title="get /v1/sources/all">client.Sources.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#SourceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#SourceListParams">SourceListParams</a>) (<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#SourceListResponse">SourceListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Stories

Response Types:

- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#NewsCluster">NewsCluster</a>
- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#StoryListResponse">StoryListResponse</a>

Methods:

- <code title="get /v1/stories/all">client.Stories.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#StoryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#StoryListParams">StoryListParams</a>) (<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#StoryListResponse">StoryListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Summarize

Response Types:

- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#SummarizeNewResponse">SummarizeNewResponse</a>

Methods:

- <code title="post /v1/summarize">client.Summarize.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#SummarizeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#SummarizeNewParams">SummarizeNewParams</a>) (<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#SummarizeNewResponse">SummarizeNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Topics

Response Types:

- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#TopicListResponse">TopicListResponse</a>

Methods:

- <code title="get /v1/topics/all">client.Topics.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#TopicService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#TopicListParams">TopicListParams</a>) (<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#TopicListResponse">TopicListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Vector

## News

Params Types:

- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#ArticleSearchFilterParam">ArticleSearchFilterParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#VectorNewsSearchResponse">VectorNewsSearchResponse</a>

Methods:

- <code title="post /v1/vector/news/all">client.Vector.News.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#VectorNewsService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#VectorNewsSearchParams">VectorNewsSearchParams</a>) (<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk">perigonsdk</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk#VectorNewsSearchResponse">VectorNewsSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
