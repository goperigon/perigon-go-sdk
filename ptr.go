package perigon

// Ptr returns a pointer to v. Use it for optional request fields.
func Ptr[T any](v T) *T {
	return &v
}
