func form(digit int, one string, five string, ten string) string {
    if digit == 4 {
        return one + five
    }
    if digit == 9 {
        return one + ten
    }
    ret := ""
    if digit >= 5 {
        ret = five
        digit = digit-5
    }
    for i:=0; i<digit; i++ {
        ret = ret + one
    }
    return ret
}

func intToRoman(num int) string {
    a := num / 1000
    b := (num % 1000) / 100
    c := (num % 100) / 10
    d := num % 10

    return form(a, "M", "", "") + form(b, "C", "D", "M") + form(c, "X", "L", "C") + form(d, "I", "V", "X")    
}