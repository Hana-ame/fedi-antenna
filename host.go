package core

var hostmap = make(map[string]string)

func initHost() {
	hostmap["localhost:3000"] = "fedi.moonchan.xyz"
}

// convert an altername to it's origin.
func Host(alias string) string {
	host, exist := hostmap[alias]
	if exist {
		return host
	}
	return alias
}
