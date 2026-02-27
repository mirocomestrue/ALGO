func letterCombinations(digits string) []string {
    m := map[int][]string{
        2: {"a", "b", "c"},
        3: {"d", "e", "f"},
        4: {"g", "h", "i"},
        5: {"j", "k", "l"},
        6: {"m", "n", "o"},
        7: {"p", "q", "r", "s"},
        8: {"t", "u", "v"},
        9: {"w", "x", "y", "z"},
    }

    ret := []string{""}

    
    for i:=0; i<len(digits); i++ {
        n := int(digits[i] - '0')
        strs := m[n]

        var newRet []string
        
        for _, r := range ret {
            for _, str := range strs {
                newRet = append(newRet, r+str)
            }
        }

        ret = newRet
    }
    return ret
}
