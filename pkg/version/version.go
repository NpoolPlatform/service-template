package version

import (
	"github.com/NpoolPlatform/go-service-app-template/message/npool"

	cv "github.com/NpoolPlatform/go-service-framework/pkg/version"

	"golang.org/x/xerrors"
)

func Version() (*npool.VersionResponse, error) {
	info, err := cv.GetVersion()
	if err != nil {
		return nil, xerrors.Errorf("get service version error: %w", err)
	}
	return &npool.VersionResponse{
		Info: info,
	}, nil
}
