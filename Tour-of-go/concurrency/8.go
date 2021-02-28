// Implemented by Hunachi

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	internalWalk(t, ch)
	close(ch)
}

func internalWalk(t *tree.Tree, ch chan int){
	if (t == nil) {
		return
	}
	internalWalk(t.Left, ch)
	ch <- t.Value
	internalWalk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	go Walk(t1, ch1)
	ch2 := make(chan int)
	go Walk(t2, ch2)
	for {
		c1, ok1 := <-ch1
		c2, ok2 := <-ch2
		if c1 != c2 {
			return false
		}
		if !ok1 && !ok2 {
			return true
		} else if !ok1 || !ok2 {
			return false
		}
	}
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(2), ch)
	Neko:
	for {
		select {
			case c, ok := <-ch:
				if !ok {
					break Neko
				}
				fmt.Println(c)
		}
	}
	flag := Same(tree.New(2),tree.New(1))
	if flag {
		fmt.Println("Same trees")
	} else {
		fmt.Println("Not Same trees")
	}
}