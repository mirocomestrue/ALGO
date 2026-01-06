//go:build ignore
// +build ignore

package algo

var S, P string
var arr [21][21]int

func isMatchIdx(i int, j int) bool {
	if arr[i][j] != 0 {
		if arr[i][j] == 1 {
			return true
		}
		return false
	}

	ret := false

	// x*
	if len(P) > j+1 && P[j+1] == '*' && len(S) > i {
		// use || no use
		ret = ((P[j] == '.' || P[j] == S[i]) && isMatchIdx(i+1, j)) ||
			isMatchIdx(i, j+2)
	} else if len(P) > j+1 && P[j+1] == '*' {
		// no use
		ret = isMatchIdx(i, j+2)
	} else if len(P) > j && P[j] == '.' && len(S) > i {
		// .
		ret = isMatchIdx(i+1, j+1)
	} else if len(P) > j && len(S) > i && P[j] == S[i] {
		ret = isMatchIdx(i+1, j+1)
	} else if len(P) == j && len(S) == i {
		ret = true
	}

	if ret == true {
		arr[i][j] = 1
	} else {
		arr[i][j] = -1
	}
	return ret
}

func isMatch(s string, p string) bool {
	S = s
	P = p
	for i := range arr {
		for j := range arr[i] {
			arr[i][j] = 0
		}
	}

	ret := isMatchIdx(0, 0)

	return ret
}
