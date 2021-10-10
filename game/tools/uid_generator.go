package tools

import (
	"container/heap"
	"errors"
)

// An IntHeap is a min-heap of ints.
type intheap []int

func (h intheap) Len() int           { return len(h) }
func (h intheap) Less(i, j int) bool { return h[i] < h[j] }
func (h intheap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intheap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *intheap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type UID_Generator struct {
	uid_chan chan int
	uid_pool *intheap
	maximum  int
}

func New_UID_Generator() *UID_Generator {
	g := &UID_Generator{}
	g.uid_chan = make(chan int)
	g.uid_pool = &intheap{}
	heap.Init(g.uid_pool)
	go g.init()
	return g
}

func (g *UID_Generator) Next() int {
	return <-g.uid_chan
}

func (g *UID_Generator) Recycle(val int) {
	g.recycle(val)
}

func (g *UID_Generator) init() {
	for {
		g.uid_chan <- g.generate()
	}
}

func (g *UID_Generator) generate() int {
	if len(*g.uid_pool) == 0 {
		g.maximum += 1
		return g.maximum - 1
	} else {
		return heap.Pop(g.uid_pool).(int)
	}
}

func (g *UID_Generator) recycle(val int) error {
	if val > g.maximum {
		return errors.New("recycled value greater than maximum")
	} else if val == g.maximum {
		g.maximum -= 1
	} else {
		heap.Push(g.uid_pool, val)
	}
	return nil
}
