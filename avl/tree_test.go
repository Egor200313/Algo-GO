package avl

import (
	"strconv"
	"testing"
)


func TestInsert(t *testing.T) {
	tree:= Tree{}
	tree.Insert(1)

	if tree.Root == nil {
		t.Errorf("Unexpected nil")
	}

	if tree.Root.value != 1 {
		t.Errorf("WA")
	}

	tree.Insert(2)

	if tree.Root == nil {
		t.Errorf("Unexpected nil")
	}

	if tree.Root.value != 1 {
		t.Errorf("WA")
	}

	if tree.Root.right.value != 2 {
		t.Errorf("WA")
	}

	tree.Insert(3)

	if tree.Root.value != 2 {
		t.Errorf("WA")
	}

	if tree.Root.right.value != 3 {
		t.Errorf("WA")
	}

	if tree.Root.left.value != 1 {
		t.Errorf("WA")
	}
}

func TestSearch(t *testing.T) {
	tree:= Tree{}
	tree.Insert(1)
	tree.Insert(2)
	
	if !tree.Search(1) {
		t.Errorf("WA")
	}

	if !tree.Search(2) {
		t.Errorf("WA")
	}

	if tree.Search(3) {
		t.Errorf("WA")
	}

	tree.Insert(3)
	if !tree.Search(3) {
		t.Errorf("WA")
	}

}

func TestStress(t *testing.T){
	tree:= Tree{}
	for i:= 0; i < 1000000; i++ {
		tree.Insert(i)
	}

	for i:= 0; i < 1000000; i++ {
		if !tree.Search(i) {
			t.Errorf("WA on " + strconv.Itoa(i))
		}
	}
}

