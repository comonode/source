package mint

import (
	"github.com/Source-Protocol-Cosmos/source/v2/x/mint/keeper"
	"github.com/Source-Protocol-Cosmos/source/v2/x/mint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis new mint genesis
func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, ak types.AccountKeeper, data *types.GenesisState) {
	keeper.SetMinter(ctx, data.Minter)
	keeper.SetParams(ctx, data.Params)
	ak.GetModuleAccount(ctx, types.ModuleName)
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) *types.GenesisState {
	minter := keeper.GetMinter(ctx)
	params := keeper.GetParams(ctx)
	return types.NewGenesisState(minter, params)
}
