package version

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/go-service-framework/pkg/version"
)

func TestVersion(t *testing.T) {
	cli := resty.New()
	resp, err := cli.R().
		Get("http://localhost:32759/version")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		ver, err := version.GetVersion()
		assert.NotNil(t, err)
		assert.Equal(t, ver, resp.Body())
	}
}
