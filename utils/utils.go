package utils

// utils
func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ObjectId(path, host string) string {
	return "https://" + host + "/o/" + path
}
