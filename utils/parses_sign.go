package utils

import "strings"

// signature
// "SHA-256=8RIlimPwETDMkWQMI59d0gm9dqhzKGtX0CsEcahxxOE=" => "SHA-256", "8RIlimPwETDMkWQMI59d0gm9dqhzKGtX0CsEcahxxOE="
func ParseDigest(d string) (algorithm, digest string) {
	arr := strings.SplitN(d, "=", 2)
	if len(arr) != 2 {
		return
	}
	return arr[0], arr[1]
}
