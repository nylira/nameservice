package app

import (
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/tendermint/tendermint/libs/log"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	dbm "github.com/tendermint/tendermint/libs/db"
)

const (
	appName = "nameservice"
)

type nameserviceApp struct {
	*bam.BaseApp
}

func NewnameserviceApp(logger log.Logger, db dbm.DB) *nameserviceApp {
	// first define the top level codec that will be shared between modules
	cdc := MakeCodec()

	// BaseApp handles interactions with Tendermint through ABCI
	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc))

	var app = &nameserviceApp {
		BaseApp: bApp,
		cdc: cdc
	}

	return app
}
