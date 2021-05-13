package congress

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
)

/// 0001号分叉
/// 用于修复质押模块错误的产出了收益
func fork_0001_20210513(state *state.StateDB) error {

	var targetAddress := []common.Address{
		common.HexToAddress("0x0000000000000000000000000000000000000000"),
	}

	state.SubBalance()


	return nil
}