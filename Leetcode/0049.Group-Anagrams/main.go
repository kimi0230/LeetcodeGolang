package groupanagrams

import "sort"

// 使用計數: 時間複雜 O(nk), 空間複雜 O(nk)
// n 是單詞的個數 , k是單字的最大長度
// O(n)遍歷單詞, 遍歷字符O(nk)
func GroupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string, len(strs))

	for _, str := range strs {
		// 建立每一個string的特徵`count`, 藉由統計每個字母出現的次數
		count := [26]int{}
		for _, b := range str {
			count[b-'a']++
		}
		// 將同樣的次數的string放置在 map中
		m[count] = append(m[count], str)
	}

	// 把map中的string放入到結果的 `[][]string` array 中
	ans := [][]string{}
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}

// 排序: 時間複雜 O(nklog(k)), 空間複雜 O(klog(k))
// n 是單詞的個數 , k是單字的最大長度
// 遍歷單詞 O(n)
// 排序字符 O(klog(k))
func GroupAnagramsBySort(strs []string) [][]string {
	m := make(map[string][]string, len(strs))

	for _, str := range strs {
		// 建立每一個string的特徵, 藉由統計排序
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)

		// 將同樣的特徵的string放置在 map中
		m[sortedStr] = append(m[sortedStr], str)
	}

	// 把map中的string放入到結果的 `[][]string` array 中
	ans := [][]string{}
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}
