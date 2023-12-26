/*
* Simple Actor model
* Refenrence:
 */

package concurrency

import (
	"context"
	"fmt"
	"runtime"
	"sync"
)

type actor struct {
	Id    int
	Inbox chan int
	Limit int
	wg    *sync.WaitGroup
}

func (a *actor) Execute(ctx context.Context) {
	fmt.Println("configured wg", &a.wg)

	a.wg.Add(1)
	go func() {
		for {
			select {
			case task := <-a.Inbox:
				fmt.Println(a.Id, ": Executed task: ", task)
			case <-ctx.Done():
				a.wg.Done()
				fmt.Println("Stopping Actor: ", a.Id)
			default:
				a.Id = a.Id
			}
		}

		// fmt.Println("Stopping Actor: ", a.Id)

	}()
}

func (a *actor) Submit(task int) error {
	if len(a.Inbox) >= a.Limit {
		return fmt.Errorf("actor queue full")
	}
	a.Inbox <- task
	return nil
}

type ActorPools struct {
	actors    map[int]*actor
	maxActors int
	cancel    context.CancelFunc
	WG        *sync.WaitGroup
}

func NewActorPool() *ActorPools {
	wg := new(sync.WaitGroup)
	ap := ActorPools{
		actors:    make(map[int]*actor),
		maxActors: runtime.NumCPU(),
		WG:        wg,
	}

	for i := 0; i < ap.maxActors; i++ {
		ap.actors[i] = &actor{
			Id:    i,
			Limit: 1,
			Inbox: make(chan int, 100),
			wg:    wg,
		}
	}
	//	fmt.Println(ap.actors)
	return &ap
}

func (ap *ActorPools) Start() *sync.WaitGroup {
	ctx, cancel := context.WithCancel(context.Background())
	ap.cancel = cancel
	for _, a := range ap.actors {
		//ap.WG.Add(1)
		a.Execute(ctx)
	}

	fmt.Println("sent wg", &ap.WG)

	return ap.WG
}

func (ap *ActorPools) Submit(task int) {
	k := task % ap.maxActors

	ap.actors[k].Submit(task)
}

func (ap *ActorPools) Stop() {
	fmt.Println("Stopping the Actors pool")
	ap.cancel()
}
