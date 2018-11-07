package pokerbull

import (
	"gitlab.33.cn/chain33/plugin/dapp/pokerbull/commands"
	"gitlab.33.cn/chain33/plugin/dapp/pokerbull/executor"
	"gitlab.33.cn/chain33/plugin/dapp/pokerbull/rpc"
	"gitlab.33.cn/chain33/plugin/dapp/pokerbull/types"
	"gitlab.33.cn/chain33/pluginmgr"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.PokerBullX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.PokerBullCmd,
		RPC:      rpc.Init,
	})
}