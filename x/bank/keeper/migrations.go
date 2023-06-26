package keeper

import (
	"cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/exported"
	v2 "github.com/cosmos/cosmos-sdk/x/bank/migrations/v2"
	v3 "github.com/cosmos/cosmos-sdk/x/bank/migrations/v3"
	v4 "github.com/cosmos/cosmos-sdk/x/bank/migrations/v4"
)

type MigrationKeeper interface {
	getStoreService() store.KVStoreService
	getCodec() codec.BinaryCodec
}

func (k BaseKeeper) getStoreService() store.KVStoreService {
	return k.storeService
}

func (k BaseKeeper) getCodec() codec.BinaryCodec {
	return k.cdc
}

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper         MigrationKeeper
	legacySubspace exported.Subspace
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper MigrationKeeper, legacySubspace exported.Subspace) Migrator {
	return Migrator{keeper: keeper, legacySubspace: legacySubspace}
}

// Migrate1to2 migrates from version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v2.MigrateStore(ctx, m.keeper.getStoreService(), m.keeper.getCodec())
}

// Migrate2to3 migrates x/bank storage from version 2 to 3.
func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	return v3.MigrateStore(ctx, m.keeper.getStoreService(), m.keeper.getCodec())
}

// Migrate3to4 migrates x/bank storage from version 3 to 4.
func (m Migrator) Migrate3to4(ctx sdk.Context) error {
	return v4.MigrateStore(ctx, m.keeper.getStoreService(), m.legacySubspace, m.keeper.getCodec())
}
