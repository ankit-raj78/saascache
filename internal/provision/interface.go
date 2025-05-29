package provision

import "context"

// Provisioner knows how to provision caches.
type Provisioner interface {
    ProvisionCache(ctx context.Context, spec Spec) (Result, error)
}

// Spec defines what to provision.
type Spec struct {
    Engine    string
    Pattern   string
    NodeCount int
}

// Result holds stubbed endpoint data.
type Result struct {
    Endpoint  string
    AuthToken string
}
