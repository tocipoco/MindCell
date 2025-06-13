package app

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// OrderBeginBlockers returns the order of begin blockers
func OrderBeginBlockers() []string {
	return []string{
		"modelregistry",
		"shardallocator",
		"inferencegateway",
		"billing",
		"reward",
		"slashing",
		"token",
	}
}

// OrderEndBlockers returns the order of end blockers
func OrderEndBlockers() []string {
	return []string{
		"token",
		"slashing",
		"reward",
		"billing",
		"inferencegateway",
		"shardallocator",
		"modelregistry",
	}
}

// OrderInitGenesis returns the order of module genesis initialization
func OrderInitGenesis() []string {
	return []string{
		"token",
		"modelregistry",
		"shardallocator",
		"inferencegateway",
		"billing",
		"reward",
		"slashing",
	}
}

// ModuleAccountPermissions returns module account permissions
func ModuleAccountPermissions() map[string][]string {
	return map[string][]string{
		"billing":  {"minter", "burner"},
		"reward":   {"minter"},
		"slashing": {"burner"},
	}
}

// BlockedAddresses returns all module account addresses that are blocked from receiving funds
func BlockedAddresses() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range ModuleAccountPermissions() {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}
	return modAccAddrs
}

// GetMaccPerms returns a mapping of the application's module account permissions.
func GetMaccPerms() map[string][]string {
	return ModuleAccountPermissions()
}
