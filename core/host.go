package core

import "github.com/Hana-ame/fedi-antenna/core/utils"

// convert an altername to it's origin.
func Host(alias string) string {
	host, exist := utils.AliasMap[alias]
	if exist {
		return host
	}
	return alias
}
