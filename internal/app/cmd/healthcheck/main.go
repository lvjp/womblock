package healthcheck

import (
	"fmt"
	"net/http"

	"github.com/lvjp/womblock/internal/pkg/cmd/util"
	"github.com/lvjp/womblock/pkg/buildinfo"
)

func Run(ctx *util.Context) error {
	endpoint := fmt.Sprintf("http://%s", *ctx.Config.Server.ListenAddress)
	fmt.Fprintln(ctx.Output, "Endpoint:", endpoint)

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return fmt.Errorf("healthcheck request forging: %v", err)
	}

	req.Header.Set("User-Agent", "womblock/"+buildinfo.Get().Revision)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("healthcheck request execution: %v", err)
	}
	// Closing the response body as we don't need to read it.
	resp.Body.Close()
	fmt.Fprintln(ctx.Output, "Status:", resp.Status)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("healthcheck unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
