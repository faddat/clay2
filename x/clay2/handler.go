package clay2

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/faddat/clay2/x/clay2/keeper"
	"github.com/faddat/clay2/x/clay2/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
    // this line is used by starport scaffolding
		case types.MsgCreatePost:
			return handleMsgCreatePost(ctx, k, msg) # 1
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
