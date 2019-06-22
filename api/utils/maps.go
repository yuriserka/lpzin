package utils

import "sort"

// OrdenaMap ordena os indices do map para que possa imprimir na ordem certa, é util
// para os menus
func OrdenaMap(m map[int]string) []int {
	sortedIndexes := make([]int, 0, len(m))
	for i := range m {
		sortedIndexes = append(sortedIndexes, i)
	}
	sort.Ints(sortedIndexes)

	return sortedIndexes
}
