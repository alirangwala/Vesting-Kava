package keeper

import (
	"github.com/alirangwala/mychain/x/mychain/types"
)

var _ types.QueryServer = Keeper{}
