package v2_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/crypto-org-chain/cronos/x/cronos/exported"
	v2 "github.com/crypto-org-chain/cronos/x/cronos/migrations/v2"
	"github.com/crypto-org-chain/cronos/x/cronos/types"
	"github.com/stretchr/testify/require"
)

type mockSubspace struct {
	ps types.Params
}

func newMockSubspace(ps types.Params) mockSubspace {
	return mockSubspace{ps: ps}
}

func (ms mockSubspace) GetParamSet(ctx sdk.Context, ps exported.ParamSet) {
	*ps.(*types.Params) = ms.ps
}

func TestMigrate(t *testing.T) {
	storeKey := sdk.NewKVStoreKey(types.ModuleName)
	ctx := testutil.DefaultContext(storeKey, sdk.NewTransientStoreKey("test"))
	store := ctx.KVStore(storeKey)
	cdc := simapp.MakeTestEncodingConfig().Codec
	legacySubspace := newMockSubspace(types.DefaultParams())
	v2.Migrate(ctx, store, legacySubspace, cdc)
	var p types.Params
	require.NoError(t, cdc.Unmarshal(store.Get(types.ParamsKey), &p))
	require.Equal(t, legacySubspace.ps, p)
}
