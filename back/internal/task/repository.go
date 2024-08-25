package task

import "sort"

// criar método de ordenação por
// prioridade
// mais perto de vencer

func OrderCloseToExpire() {
	sort.Slice(ListTask, func(i, j int) bool {
		return ListTask[i].DueDate < ListTask[j].DueDate
	})
}

func OrderFarToExpire() {
	sort.Slice(ListTask, func(i, j int) bool {
		return ListTask[i].DueDate > ListTask[j].DueDate
	})
}

func OrderByHightDifficulty() {
	sort.Slice(ListTask, func(i, j int) bool {
		return ListTask[i].Difficulty > ListTask[j].Difficulty
	})
}
