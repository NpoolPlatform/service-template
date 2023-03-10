package version

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/version"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

func Version() (*basetypes.VersionResponse, error) {
	info, err := version.GetVersion()
	if err != nil {
		return nil, err
	}
	return &basetypes.VersionResponse{
		Info: info,
	}, nil
}
