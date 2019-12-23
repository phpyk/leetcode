package list

import (
	"fmt"
	"testing"
)

func main() {
	myList := NewList()
	fmt.Printf("myList: %+v \n", myList)
	fmt.Printf("myList.root address: %p \n", &myList.root)
	fmt.Printf("myList.root.next address: %p \n", &myList.root.next)
	fmt.Println("-----Insert e1,e2------")
	e1 := Element{nil, 100}
	myList.Insert(&e1, &myList.root)
	e2 := Element{nil, 200}
	myList.Insert(&e2, &e1)

	fmt.Println("-----Insert 5 elements------")
	//此处p作为哨兵节点，作为下一次插入时的位置
	p := &e2
	for i := 0; i < 3; i++ {
		e := Element{nil, i}
		myList.Insert(&e, p)
		p = &e
	}
	fmt.Println("-----Insert e3------")
	e3 := Element{nil, 300}
	myList.Insert(&e3, p)

	fmt.Println("-----Print all Elements------")
	myList.Print()

	fmt.Println("-----Delete e2------")
	myList.Delete(&e2)
	myList.Print()

	fmt.Println("-----Reverse list------")
	myList.Reverse()
	myList.Print()

	fmt.Println("-----Trans to circle list------")
	myList.TransToCircle()
	//myList.Print()

	fmt.Println("-----Check if is circle------")
	fmt.Printf("list is Circle ? %v\n", myList.IsCircleList())


}

func TestSortedLinkedListMerge(t *testing.T) {
	l1 := NewList()
	for i:= 9;i>0;i-=2 {
		e := Element{nil,i}
		l1.Insert(&e,&l1.root)
	}
	fmt.Println("--------l1---------")
	l1.Print()
	l2 := NewList()
	for i:=10;i>0;i-=2 {
		e := Element{nil,i}
		l2.Insert(&e,&l2.root)
	}
	fmt.Println("--------l2---------")
	l2.Print()
	l3 := NewList()
	l3 = SortedLinkedListMerge(l1,l2)
	fmt.Println("--------l3---------")
	l3.Print()
}