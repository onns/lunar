package lunar

import (
	"fmt"
	"testing"
	"time"
)

/*
@Time : 2021/11/28 14:51
@Author : onns
@File : /lunar_test.go
*/

func TestParse(t *testing.T) {
	tt := time.Date(1996,8,4,0,0,0,0,time.Now().Location())
	lunar := Parse(tt)
	fmt.Println(lunar)
	fmt.Println(lunar.AddDate(1,0,0))
	fmt.Println(lunar.AddDate(1,0,0).ToSolar().String())
	fmt.Println(Parse(time.Date(1996,8,4,0,0,0,0,time.Now().Location())))
	fmt.Println(Parse(time.Date(1996,2,19,0,0,0,0,time.Now().Location())))
	fmt.Println(Parse(time.Date(1997,2,7,0,0,0,0,time.Now().Location())))
}
