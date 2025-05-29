package provision

import "context"

// LocalProvisioner simulates cache provisioning.
type LocalProvisioner struct{}

func (p *LocalProvisioner) ProvisionCache(ctx context.Context, spec Spec) (Result, error) {
    return Result{
        Endpoint:  "localhost:6379",
        AuthToken: "",
    }, nil
}
