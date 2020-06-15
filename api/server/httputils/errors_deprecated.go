package httputils // import "github.com/helmutkemper/moby/api/server/httputils"
import "github.com/helmutkemper/moby/errdefs"

// GetHTTPErrorStatusCode retrieves status code from error message.
//
// Deprecated: use errdefs.GetHTTPErrorStatusCode
func GetHTTPErrorStatusCode(err error) int {
	return errdefs.GetHTTPErrorStatusCode(err)
}
