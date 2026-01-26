func longestCommonPrefix(strs []string) string {
    i := 0
    l := len(strs)
    for {
        common := true
        for j:=1; j<l; j++ {
            if len(strs[0]) <= i || len(strs[j]) <= i || strs[0][i] != strs[j][i] {
                common = false
                break;
            }
        }
        if !common || len(strs[0]) <= i {
            break;
        }
        i++
    }
    return strs[0][0:i]
}