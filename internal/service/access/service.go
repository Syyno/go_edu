package access

import "context"

type AccessService interface {
	Check(ctx context.Context, endpointUri string) error
}
