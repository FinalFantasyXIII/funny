package main

import (
	"dms/staff"
	"fmt"
	"github.com/emirpasic/gods/trees/binaryheap"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"sync"
	"time"
)

//假设Theater是一个舞台，并且负责节目顺序安排
type Theater  struct {
	//接受节目单的队列
	c	chan staff.Player
	//节目播放队列，提供优先级
	q	*binaryheap.Heap
	//当前表演者
	currentPlayer staff.Player
	//来把锁
	mutex sync.Mutex
}

func NewTheater() *Theater{
	t := &Theater{
		c:             make(chan staff.Player,1024),
		q:             binaryheap.NewWith(staff.PlayerComparator),
		currentPlayer: nil,
	}
	return t
}

func main() {
	j := &staff.Jack{
		Base:         staff.Base{
			Name:       "jack",
			ShowName:   "calculate numbers",
			Peroration: "jack , jack , jack",
			LV:         rand.Int(),
		},
		IntegerArray: []int{3,4,2,11,56,43,1,63,92,89,23,12},
	}

	r := &staff.Rose{
		Base:        staff.Base{
			Name:       "rose",
			ShowName:   "calculate strings",
			Peroration: "Thanks everyone!",
			LV:         rand.Int(),
		},
		StringArray: []string{"dsd","23","232df","lll","qwewe"},
	}
	p := &staff.PhoenixLegend{
		Base:        staff.Base{
			Name:       "phoenix",
			ShowName:   "sing",
			Peroration: "phoenix fly!",
			LV:         rand.Int(),
		},
		Song: "I'm sb, I'm sb",
	}
	f := &staff.FallOutBoy{
		Base:        staff.Base{
			Name:       "phoenix",
			ShowName:   "sing",
			Peroration: "we are fall out boy!",
			LV:         rand.Int(),
		},
		Song: "my songs know what you did in the dark!",
	}


	t := NewTheater()

	t.c <- j
	t.c <- r
	t.c <- p
	t.c <- f

	//派出工作线程接受表演节目单
	go func() {
		for v := range t.c{
			t.mutex.Lock()
			t.q.Push(v)
			t.mutex.Unlock()
		}
	}()
	time.Sleep(time.Second)

	for !t.q.Empty() {
		t.mutex.Lock()
		v, _ := t.q.Pop()
		t.mutex.Unlock()

		t.currentPlayer = v.(staff.Player)
		t.currentPlayer.Introduce()
		t.currentPlayer.Play()
		t.currentPlayer.Thanks()
		fmt.Println("----------------------------")
	}

}


