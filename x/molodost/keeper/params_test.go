package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "molodost/testutil/keeper"
	"molodost/x/molodost/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MolodostKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
