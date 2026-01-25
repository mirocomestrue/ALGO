var str [16]string


func rti(i int, val int, one string, five string, ten string, base int) (int, int) { // next idx, value
    // 9
    if str[i] == one && str[i+1] == ten {
        return i+2, 9*base + val
    }

    // 4
    if str[i] == one && str[i+1] == five {
        return i+2, 4*base + val
    }

    num := 0

    // 5
    if str[i] == five {
        i = i+1
        num = num+5
    }
    
    // 0, 1, 2, 3
    for {
        if str[i] != one {
            break
        }
        i = i+1
        num = num+1
    }
    return i, num*base+val

}

func romanToInt(s string) int {
    for i:=0; i<16; i++ {
        if i >= len(s) {
            str[i] = ""
        } else {
            str[i] = s[i:i+1]
        }
    }

    idx, val := 0, 0
    idx, val = rti(idx, val, "M", "NaN", "NaN", 1000)
    idx, val = rti(idx, val, "C", "D", "M", 100)
    idx, val = rti(idx, val, "X", "L", "C", 10)
    idx, val = rti(idx, val, "I", "V", "X", 1)
    return val
}