# Shared Response Types

- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2/shared">shared</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2/shared#CategoryHolder">CategoryHolder</a>
- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2/shared">shared</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2/shared#Coordinate">Coordinate</a>
- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2/shared">shared</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2/shared#LocationHolder">LocationHolder</a>
- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2/shared">shared</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2/shared#SourceLocation">SourceLocation</a>

# All

Params Types:

- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#AllEndpointSortBy">AllEndpointSortBy</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#Article">Article</a>
- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#AllListResponse">AllListResponse</a>

Methods:

- <code title="get /v1/articles/all">client.All.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#AllService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#AllListParams">AllListParams</a>) (\*<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#AllListResponse">AllListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Companies

Response Types:

- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#CompanyListResponse">CompanyListResponse</a>

Methods:

- <code title="get /v1/companies/all">client.Companies.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#CompanyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#CompanyListParams">CompanyListParams</a>) (\*<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#CompanyListResponse">CompanyListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Journalists

Response Types:

- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#Journalist">Journalist</a>
- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#NameCount">NameCount</a>
- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#JournalistListResponse">JournalistListResponse</a>

Methods:

- <code title="get /v1/journalists/{id}">client.Journalists.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#JournalistService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#Journalist">Journalist</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/journalists/all">client.Journalists.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#JournalistService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#JournalistListParams">JournalistListParams</a>) (\*<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#JournalistListResponse">JournalistListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# People

Response Types:

- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#ImageHolder">ImageHolder</a>
- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#WikidataDateHolder">WikidataDateHolder</a>
- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#WikidataLabelHolder">WikidataLabelHolder</a>
- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#PersonListResponse">PersonListResponse</a>

Methods:

- <code title="get /v1/people/all">client.People.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#PersonService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#PersonListParams">PersonListParams</a>) (\*<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#PersonListResponse">PersonListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Sources

Params Types:

- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#SortBy">SortBy</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#SourceTopStatHolder">SourceTopStatHolder</a>
- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#SourceListResponse">SourceListResponse</a>

Methods:

- <code title="get /v1/sources/all">client.Sources.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#SourceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#SourceListParams">SourceListParams</a>) (\*<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#SourceListResponse">SourceListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Stories

Response Types:

- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#NewsCluster">NewsCluster</a>
- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#RecordStatHolder">RecordStatHolder</a>
- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#StoryListResponse">StoryListResponse</a>

Methods:

- <code title="get /v1/stories/all">client.Stories.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#StoryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#StoryListParams">StoryListParams</a>) (\*<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#StoryListResponse">StoryListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Summarize

Response Types:

- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#SummarizeNewResponse">SummarizeNewResponse</a>

Methods:

- <code title="post /v1/summarize">client.Summarize.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#SummarizeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#SummarizeNewParams">SummarizeNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#SummarizeNewResponse">SummarizeNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Topics

Response Types:

- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#TopicListResponse">TopicListResponse</a>

Methods:

- <code title="get /v1/topics/all">client.Topics.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#TopicService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#TopicListParams">TopicListParams</a>) (\*<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#TopicListResponse">TopicListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Vector

## News

Params Types:

- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#ArticleSearchFilterParam">ArticleSearchFilterParam</a>
- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#CoordinateFilterParam">CoordinateFilterParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#VectorNewsSearchResponse">VectorNewsSearchResponse</a>

Methods:

- <code title="post /v1/vector/news/all">client.Vector.News.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#VectorNewsService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#VectorNewsSearchParams">VectorNewsSearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#VectorNewsSearchResponse">VectorNewsSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Wikipedia

Params Types:

- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#WikipediaSearchFilterParam">WikipediaSearchFilterParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#WikipediaSearchResponse">WikipediaSearchResponse</a>
- <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#WikipediaVectorSearchResponse">WikipediaVectorSearchResponse</a>

Methods:

- <code title="get /v1/wikipedia/all">client.Wikipedia.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#WikipediaService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#WikipediaSearchParams">WikipediaSearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#WikipediaSearchResponse">WikipediaSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/vector/wikipedia/all">client.Wikipedia.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#WikipediaService.VectorSearch">VectorSearch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#WikipediaVectorSearchParams">WikipediaVectorSearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2">perigon</a>.<a href="https://pkg.go.dev/github.com/goperigon/perigon-go-sdk/v2#WikipediaVectorSearchResponse">WikipediaVectorSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
