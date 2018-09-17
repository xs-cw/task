package task

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestA(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	lll := 10000
	ccc := make(chan bool, lll)
	for i := 0; i != lll; i++ {
		v := rand.Int()%1000 + 1
		d := time.Millisecond * time.Duration(v)

		Add(time.Now().Add(d), func() {
			fmt.Println("hello", d, Default.Len())
			ccc <- true
		})
	}
	Print()
	Join()
}

func TestB(t *testing.T) {
	fmt.Println("begin", time.Now())
	AddPeriodic(PeriodicIntervalCount(time.Now(), time.Second/10, 50), func() {
		fmt.Println("fork", time.Now(), Default.Len())
	})
	Print()
	Join()
}


func TestC(t *testing.T) {
	fmt.Println("begin", time.Now())
	n := AddPeriodic(PeriodicIntervalCount(time.Now(), time.Second, 5), func() {
		fmt.Println("fork", time.Now(), Default.Len())
	})
	Add(time.Now().Add(time.Second*2), func() {
		Cancel(n)
		fmt.Println("close", time.Now())
	})
	Print()
	Join()
	fmt.Println("begin", time.Now())
	n = AddPeriodic(PeriodicIntervalCount(time.Now(), time.Second, 5), func() {
		fmt.Println("fork", time.Now(), Default.Len())
	})
	Add(time.Now().Add(time.Second*2), func() {
		Cancel(n)
		fmt.Println("close", time.Now(), Default.Len())
	})
	Print()
	Join()
}

func TestD(t *testing.T) {
	fmt.Println("begin", time.Now())
	AddPeriodic(PeriodicTiming(time.Now().Add(time.Second), time.Now()), func() {
		fmt.Println("fork", time.Now())
	})
	Print()
	Join()
}
