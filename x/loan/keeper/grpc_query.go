package keeper

import (
	"github.com/kazumal/loan/x/loan/types"
)

var _ types.QueryServer = Keeper{}
