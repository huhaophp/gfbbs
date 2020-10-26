package common

func SortMap(s map[int]int) []int {
	var rank []int
	for  key, _ := range s {
		rank = append(rank, key)
	}
	for i := 0; i < len(rank); i++ {
		for j := i + 1; j < len(rank); j++ {
			if s[rank[i]] < s[rank[j]] {
				rank[i], rank[j] = rank[j], rank[i]
			}
		}
	}

	return rank
}