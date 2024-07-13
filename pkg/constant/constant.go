package constant

type ContextKey string

const (
	RequestIdHeader ContextKey = "requestId"
	TraceIdHeader   ContextKey = "traceId"
	UserIdHeader    ContextKey = "userId"
	SpanIdHeader    ContextKey = "spanId"
)

const (
	AuthorizationHeader = "Authorization"
	BearerPrefix        = "Bearer"
)
