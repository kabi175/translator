package markdown

import (
	"sort"
	"testing"
)

func TestContanierSort(t *testing.T) {
	var c contaniers
	c = append(c, contanier{7, 8, 1})
	c = append(c, contanier{1, 2, 3})
	c = append(c, contanier{1, 2, 1})
	c = append(c, contanier{3, 4, 1})
	sort.Sort(c)
	got := sort.IsSorted(c)
	if got != true {
		t.Errorf("Contanier not sorted\n%v\n", c)
	}
}

func TestContanierOverlap(t *testing.T) {
	var c contaniers = contaniers{
		{0, 1, 1},
		{2, 3, 1},
		{2, 5, 1},
		{4, 10, 1},
		{5, 6, 1},
		{11, 12, 1},
	}
	sort.Sort(c)
	got := sort.IsSorted(c)
	if got != true {
		t.Errorf("Contanier not sorted\n%v\n", c)
	}
	newC := c.Overlap()
	t.Error(newC)
}
