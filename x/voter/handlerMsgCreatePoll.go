package voter

import (
	"github.com/alterkim/voter/x/voter/keeper"
	"github.com/alterkim/voter/x/voter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func handleMsgCreatePoll(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreatePoll) (*sdk.Result, error) {
	var poll = types.Poll{
		Creator: msg.Creator,
		ID:      msg.ID,
		Title:   msg.Title,
		Options: msg.Options,
	}
	k.CreatePoll(ctx, poll)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
