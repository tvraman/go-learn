package sortmap

import "sort"

//SortedMap implements Interface sort

type sortedMap struct {
	m map[string]int
	s []string
}

// Implement Interface sort methods
// Methods: len, swap,  less

func (sm *sortedMap) Len() int {
	return len(sm.m)
}

func (sm *sortedMap) Less(i, j int) bool {
	return sm.m[sm.s[i]] < sm.m[sm.s[j]]
}

func (sm *sortedMap) Swap(i, j int) {
	sm.s[i], sm.s[j] = sm.s[j], sm.s[i]
}

// Exported:

func SortedKeys(m map[string]int) []string {
	// we could also use new below
	sm := &sortedMap{}
	sm.m = m
	sm.s = make([]string, len(m))
	i := 0
	for key, _ := range m {
		sm.s[i] = key
		i++
	}
	sort.Sort(sm)
	return sm.s
}
