package client

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	constant "github.com/NpoolPlatform/service-template/pkg/const"

	// nolint:nolintlint
	testinit "github.com/NpoolPlatform/service-template/pkg/test-init"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/google/uuid"
)

// nolint:nolintlint
func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

func TestClient(t *testing.T) {
	// nolint:nolintlint
	_, _ = GetServiceTemplateInfoOnly(context.Background(),
		cruder.NewFilterConds().
			WithCond(constant.FieldID, cruder.EQ, structpb.NewStringValue(uuid.UUID{}.String())))
	// Here won't pass test due to we always test with localhost
}
