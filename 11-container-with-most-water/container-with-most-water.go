func maxArea(height []int) int {
    // naive approach -> 10 ^ 5 ^ 2 = 10 ^ 10 .. 안되겠지?
    // (a, b) 와 (b, c) 를 알때 (a, c) 구하기? -> merge 하는 방법이 당장 생각은 안난다.
    // height 범위가 있으니 얘를 써보자...
    
    // k 에 대해 젤 큰거, 젤 작은거 를 저장하기?
    var left [10001]int

    // 0,1->0
    // 2,3,4,5,6,7,8 -> 1
    // ...
    // 10000 -> k 

    var right [10001]int

    // 0,1,2,3,4,5,6,7 -> 마지막
    // 8 -> 마지막-2

    // 이런식으로 투포인터 접근? O(n) 에 가능

    var i, h, maxH int
    // left[h] 채우기
    for {
        if i >= len(height) {
            break;
        }
        maxH = max(maxH, height[i])
        if height[i] >= h {
            left[h] = i
            h++
            continue;
        }
        i++
    }
    // right[h] 채우기
    i, h = len(height)-1, 0
    for {
        if i < 0 {
            break;
        }
        if height[i] >= h {
            right[h] = i
            h++
            continue;
        }
        i--
    }

    ret := 0
    // k * (젤큰거 - 젤작은거) 의 max 구하기. O(h) 에 가능
    for i=0; i<=maxH; i++ {
        area := (right[i] - left[i]) * i
        ret = max(ret, area)
    }

    return ret
}