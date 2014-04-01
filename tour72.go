package main

import "code.google.com/p/go-tour/tree"
import "fmt"
// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
	if t.Left!=nil{
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right!=nil{
		Walk(t.Right, ch)
	}
}
// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool{
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func(){
		Walk(t1, ch1)
		close(ch1)
	}()
	go func(){
		Walk(t2, ch2)
		close(ch2)
	}()

	for {

		v1, ok := <-ch1
		v2, ok2 := <-ch2
		if !ok && !ok2{
			return true;
		}
		if(v1!=v2){
			return false;
		}

	}
	return false

}

func main() {
	fmt.Println(Same(tree.New(1),tree.New(1)))
}
