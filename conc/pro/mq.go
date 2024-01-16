package pro

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/vicky1115050/godsa/ds"
)

type Broker interface {
	CreateTopic(string, int) (bool, error)
	Publish(string, int) (bool, error)
	Consume(string) (int, error)
	Commit(int) (bool, error)
}

type topic struct {
	name               string
	data               *ds.Queue
	offset             int64
	count              int
	partition          int
	partitions         map[int]chan int
	partitionConsumers map[int]int
	consumers          map[string]int
	mux                *sync.Mutex
}

type broker struct {
	topics map[string]*topic
	mux    *sync.Mutex
}

func NewBroker() *broker {
	return &broker{
		topics: make(map[string]*topic),
		mux:    &sync.Mutex{},
	}
}

func (b *broker) Publish(topicName string, record int) (bool, error) {
	tp, ok := b.topics[topicName]
	if !ok {
		return false, fmt.Errorf("Topic not found")
	}

	tp.data.Push(ds.NewNode(record))

	return true, nil
}

func (b *broker) CreateTopic(topicName string) (bool, error) {
	tp, ok := b.topics[topicName]
	if !ok {
		tp = &topic{
			name:               topicName,
			data:               ds.NewQueue(),
			offset:             0,
			count:              0,
			mux:                &sync.Mutex{},
			partition:          1,
			partitions:         make(map[int]chan int),
			partitionConsumers: make(map[int]int),
			consumers:          make(map[string]int),
		}
		b.mux.Lock()
		defer b.mux.Unlock()
		b.topics[topicName] = tp
		return true, nil
	} else {
		return false, nil
	}
}

func (b *broker) Consume(topicName, consumerId string, partitionId int) (int, int64, error) {

	tp, ok := b.topics[topicName]

	if !ok {
		return 0, 0, fmt.Errorf("No Topic found")
	}

	if consumer, ok := tp.consumers[consumerId]; !ok {
		return 0, 0, fmt.Errorf("Consumer not allowed")
	} else {
		if consumer != partitionId {
			tp.consumers[consumerId] = partitionId
		}
	}

	if tp.data.Len() == 0 {
		return 0, 0, nil
	}

	tp.mux.Lock()
	defer tp.mux.Unlock()
	record := tp.data.Load(int(tp.offset) + 1)

	if record == nil {
		return 0, 0, nil
	}
	// tp.data.Pop()
	atomic.AddInt64(&tp.offset, 1)

	return record.Data(), tp.offset, nil
}

func (b *broker) Commit(offset int) (bool, error) {
	return true, nil
}

func (b *broker) Subscribe(consumerId, topicName string) (chan int, int, error) {

	tp, ok := b.topics[topicName]

	if !ok {
		return nil, 0, fmt.Errorf("No Topic found")
	}

	if len(tp.partitions) >= tp.partition {
		tp.mux.Lock()
		defer tp.mux.Unlock()
		tp.consumers[consumerId] = 1
		c, _ := strconv.Atoi(consumerId)
		p := c % tp.partition
		wait := tp.partitions[p+1]

		return wait, p + 1, nil
	}

	tp.mux.Lock()
	defer tp.mux.Unlock()
	wait := make(chan int)
	partition := len(tp.partitions) + 1
	tp.consumers[consumerId] = partition
	tp.partitions[partition] = wait

	if tp.partitionConsumers[partition] == 0 {
		go func() {
			wait <- 1
		}()
	}
	tp.partitionConsumers[partition] += 1

	return wait, partition, nil
}

func (b *broker) UnSubscribe(consumerId, topicName string, partitionId int) (bool, error) {

	fmt.Println("Unsubscribing --> ", consumerId, "Form partiotion -->", partitionId)

	tp, ok := b.topics[topicName]

	if !ok {
		return false, fmt.Errorf("No Topic found")
	}

	if _, ok := tp.partitionConsumers[partitionId]; !ok {
		return false, nil
	}

	tp.mux.Lock()
	defer tp.mux.Unlock()
	delete(tp.consumers, consumerId)
	tp.partitionConsumers[partitionId] -= 1

	go func() {
		fmt.Println("Allowed 1 consumer")
		tp.partitions[partitionId] <- 1
	}()
	return true, nil
}

func (b *broker) Flush() {
	for _, v := range b.topics {
		fmt.Println(v.partitions)
		fmt.Println(v.consumers)
		fmt.Println(v.partitionConsumers)
		// for _, v := range v.partitions {
		// 	// close(v)
		// }
		v.data.Flush()
	}
}
