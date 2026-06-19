// range — over slice, array, map, string, channel
//
// for i, v := range slice
// for k, v := range map
// for i, r := range string  // runes
// for v := range slice      // index only if one var
package gobasics

func MapKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func SliceWithIndex(nums []int) [][2]int {
	out := make([][2]int, len(nums))
	for i, v := range nums {
		out[i] = [2]int{i, v}
	}
	return out
}
