package voter_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/vrde/voter/x/voter"
	simapp "github.com/vrde/voter/x/voter/app"
	"github.com/vrde/voter/x/voter/keeper"
	"github.com/vrde/voter/x/voter/types"
)

var (
	address  = simapp.CreateTestAddrs(1)[0]
	address2 = simapp.CreateTestAddrs(2)[1]
	address3 = simapp.CreateTestAddrs(3)[2]
)

func createTestApp(isCheckTx bool) (*simapp.SimApp, sdk.Context) {
	app := simapp.Setup(isCheckTx)
	ctx := app.BaseApp.NewContext(isCheckTx, abci.Header{})

	return app, ctx
}

func TestInvalidMsg(t *testing.T) {
	app, ctx := createTestApp(false)
	h := voter.NewHandler(app.VoterKeeper)
	_, err := h(ctx, sdk.NewTestMsg())
	require.Error(t, err)
}

func TestMsgCreatePoll(t *testing.T) {
	var pollList []types.Poll
	app, ctx := createTestApp(false)
	k := app.VoterKeeper
	h := voter.NewHandler(k)
	m := types.NewMsgCreatePoll(address, "A title", "options")
	_, err := h(ctx, m)
	require.Nil(t, err)
	res, _ := keeper.ListPoll(ctx, k)
	types.ModuleCdc.MustUnmarshalJSON(res, &pollList)
	fmt.Println("antani")
	fmt.Printf("%+v\n", pollList)
}
