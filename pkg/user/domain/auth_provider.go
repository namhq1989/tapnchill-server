package domain

const (
	AuthProviderExtension = "extension"
	AuthProviderGoogle    = "google"
)

type AuthProvider struct {
	Provider string
	ID       string
	Name     string
	Email    string
}
