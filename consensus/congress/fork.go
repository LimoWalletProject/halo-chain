package congress

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"math/big"
)

// *FORK 0001*
// The HALO Labs management team has decided to proceed with this hard fork,
// which will help early eco-backers recover stolen and scammed governance
// tokens in total of 28,000HO, and the management team has also decided that
// no further hard forks will be conducted for such events.
func fork000120210620(state *state.StateDB) {

	receiptor := common.HexToAddress("0x92b080f0034E25C76A86a1055cb3BF38350B12e1")

	forkAddrsses := []common.Address{
		common.HexToAddress("0xF315ddc323dF505C6BaABAEB0CFe65De31b85E93"),
		common.HexToAddress("0xd92593bD0d93de0f4f61bF7AF563Dd36E84A1b0e"),
		common.HexToAddress("0xCe2E6454b8aF5709a9A5E87Db1f398499e7CF1d6"),
		common.HexToAddress("0x215529416DC3fBB81A6432f5534c8E1930bd57f1"),
		common.HexToAddress("0x44C5798b853C880aBD15f5E93D6D8B72952c15e5"),
		common.HexToAddress("0x8A99a2234e4346Dee83fafBd509E651f61E0734E"),
	}

	total := big.NewInt(0)

	for _, address := range forkAddrsses {
		total = total.Add(total, state.GetBalance(address))
		state.SetBalance(address, common.Big0)
	}

	state.SetBalance(receiptor, total.Add(state.GetBalance(receiptor), total) )
}

// *FORK 0002* ON 1243400
func fork000220210628(state *state.StateDB) {
	state.SetBalance(common.HexToAddress("0x4Bf7a160D12cfC9f69fd136c41c395a62dD4De74"), common.Big0)
	state.SetBalance(common.HexToAddress("0x0000000000000000000000000000000000000000"), common.Big0)
	state.SetBalance(common.HexToAddress("0x4afC487Fc9F37e7B26FAff1E8aC0a67e0c7060E2"), common.Big0)

	afterBalance := state.GetBalance(common.HexToAddress("0x19FcF0CD0EaC05E8E2fcD38f6A48CC3A8ed08DFC"))
	state.SetBalance(common.HexToAddress("0xd2701B37bD4Ef59fBBfbFf6a59C07986d183E8e1"), afterBalance)
	state.SetBalance(common.HexToAddress("0x19FcF0CD0EaC05E8E2fcD38f6A48CC3A8ed08DFC"), common.Big0)
}