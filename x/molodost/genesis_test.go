package molodost_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "molodost/testutil/keeper"
	"molodost/testutil/nullify"
	"molodost/x/molodost"
	"molodost/x/molodost/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MolodostKeeper(t)
	molodost.InitGenesis(ctx, *k, genesisState)
	got := molodost.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
