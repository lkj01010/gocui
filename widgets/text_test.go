package widgets

import (
    "fmt"
    "testing"
    "time"
)

func TestAllText(t *testing.T) {
    var i int = 0
    var tcnt int = 0
    for {
        fmt.Printf("%c", i)
        i++
        //从零一直打印
        time.Sleep(50*time.Millisecond)
        //如果打印的太快，有时会不出结果，所以要停顿一下
        //以下几行是每隔60个换一下行，方便观察结果
        tcnt++
        if tcnt%60 == 0 {
            fmt.Println()
        }
    }
}
