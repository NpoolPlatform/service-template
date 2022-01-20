package version

import (
	npool "github.com/NpoolPlatform/message/npool"

	logger "github.com/NpoolPlatform/go-service-framework/pkg/logger" //nolint
	cv "github.com/NpoolPlatform/go-service-framework/pkg/version"    //nolint

	"golang.org/x/xerrors"
)

func Version() (*npool.VersionResponse, error) {
	info, err := cv.GetVersion()
	if err != nil {
		logger.Sugar().Errorf("get service version error: %+w", err)
		return nil, xerrors.Errorf("get service version error: %w", err)
	}
	return &npool.VersionResponse{
		Info: info,
	}, nil
}
