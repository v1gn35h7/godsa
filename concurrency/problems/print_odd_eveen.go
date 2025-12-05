package problems

import (
	"fmt"
	"sync"
)

func PrintNumbers(n int) { fmt.Print(n) }

type POrE struct {
	OddCh  chan int
	EvenCh chan int
	ZeroCh chan int
}

func (p *POrE) PrintOdd(n int) {
	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			<-p.OddCh
			PrintNumbers(i)
			p.ZeroCh <- 2
		}
	}
}

func (p *POrE) PrintEven(n int) {
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			<-p.EvenCh
			PrintNumbers(i)
			if i == n {
				close(p.ZeroCh)
				return
			}
			p.ZeroCh <- 1
		}
	}
}

func (p *POrE) PrintZero(n int) {
	for i := 0; i < n; i++ {
		z := <-p.ZeroCh
		PrintNumbers(0)
		if z%2 == 0 {
			p.EvenCh <- 1
		} else {
			p.OddCh <- 1
		}
	}
}

func NewPOrE() *POrE {
	return &POrE{
		OddCh:  make(chan int),
		EvenCh: make(chan int),
		ZeroCh: make(chan int),
	}
}

func main_() {
	p := NewPOrE()
	n := 4

	var wg sync.WaitGroup

	go func() {
		defer wg.Done()
		wg.Add(1)
		p.PrintZero(n)
	}()

	go func() {
		defer wg.Done()
		wg.Add(1)
		p.PrintOdd(n)
	}()

	go func() {
		defer wg.Done()
		wg.Add(1)
		p.PrintEven(n)
	}()

	p.ZeroCh <- 1

	wg.Wait()
}
