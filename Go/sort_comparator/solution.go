package sortcomparator

import (
	"fmt"
	"os"
	"sort"
)

type Player struct {
	Name  string
	Score int
}

func Sorting() []Player {
	number := 0
	res := make([]Player, 0)
	fmt.Fscan(os.Stdout, &number)
	for i := 0; i < number; i++ {
		player := Player{}
		fmt.Fscan(os.Stdout, &player.Name)
		fmt.Fscan(os.Stdout, &player.Score)
		res = append(res, player)
	}
	sort.Slice(res, func(i int, j int) bool {
		if res[i].Score > res[j].Score {
			return true
		} else if res[i].Score == res[j].Score {
			return res[i].Name < res[j].Name
		}
		return false
	})
	return res
}
