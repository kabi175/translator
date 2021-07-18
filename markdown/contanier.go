package markdown

import "fmt"

type contanier struct {
	start int
	end   int
	tag   int8
}
type contaniers []contanier

func (c contaniers) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c contaniers) Len() int {
	return len(c)
}

func (c contaniers) Less(i, j int) bool {
	if c[i].start < c[j].start {
		return true
	}
	if c[i].start == c[j].start && c[i].end < c[j].end {
		return true
	}
	return false
}

func (c contaniers) Overlap() contaniers {
	var (
		valid contaniers
		last  int
	)

	for cur := range c {
		last = len(valid) - 1
		if cur-1 >= 0 && last >= 0 && valid[last].end >= c[cur].start {
			fmt.Println("pop valid ", valid[last-1])
			valid = valid[:last-1]
		}
		valid = append(valid, c[cur])
	}
	return valid
}
