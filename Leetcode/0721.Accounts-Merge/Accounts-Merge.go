package accountsmerge

import (
	"LeetcodeGolang/template"
	"sort"
)

// TODO: 優化
func AccountsMerge(accounts [][]string) (result [][]string) {
	// result := [][]string{}
	emailToIndex := map[string]int{}
	emailToName := map[string]string{}
	indexToEmails := map[int][]string{}

	for _, account := range accounts {
		name := account[0]
		for _, email := range account[1:] {
			if _, has := emailToIndex[email]; !has {
				// 紀錄 email, 對寫入到 對應的index 跟 對應的user name
				emailToIndex[email] = len(emailToIndex)
				emailToName[email] = name
			}
		}
	}
	parent := make([]int, len(emailToIndex))
	// 初始化
	for i := range parent {
		parent[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			// 遞迴一直找下去
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(from, to int) {
		parent[find(from)] = find(to)
	}

	for _, account := range accounts {
		firstIndex := emailToIndex[account[1]]
		for _, email := range account[2:] {
			union(emailToIndex[email], firstIndex)
		}
	}

	for email, index := range emailToIndex {
		index = find(index)
		indexToEmails[index] = append(indexToEmails[index], email)
	}
	for _, emails := range indexToEmails {
		for i := 0; i < len(indexToEmails); i++ {
			// fmt.Println(emailToName[indexToEmails[i][0]])
		}
		sort.Strings(emails)
		account := append([]string{emailToName[emails[0]]}, emails...)
		result = append(result, account)
	}

	return
}

func AccountsMerge2(accounts [][]string) (r [][]string) {
	uf := template.UnionFind{}
	uf.Init(len(accounts))
	// emailToID 將所有的 email 都拆開，拆開與 id(數組下標) 對應
	// idToName 將 id(數組下標) 與 name 對應
	// idToEmails 將 id(數組下標) 與整理好去重以後的 email 組對應
	emailToID, idToName, idToEmails, res := make(map[string]int), make(map[int]string), make(map[int][]string), [][]string{}
	for id, acc := range accounts {
		idToName[id] = acc[0]
		for i := 1; i < len(acc); i++ {
			pid, ok := emailToID[acc[i]]
			if ok {
				uf.Union(id, pid)
			}
			emailToID[acc[i]] = id
		}
	}
	for email, id := range emailToID {
		pid := uf.Find(id)
		idToEmails[pid] = append(idToEmails[pid], email)
	}
	for id, emails := range idToEmails {
		name := idToName[id]
		sort.Strings(emails)
		res = append(res, append([]string{name}, emails...))
	}
	return res
}

// 暴力解 : Union Find, 結果比較快
func AccountsMergeBurst(accounts [][]string) [][]string {
	if len(accounts) == 0 {
		return [][]string{}
	}
	uf, result, visited := template.UnionFind{}, [][]string{}, map[int]bool{}
	uf.Init(len(accounts))
	for i := 0; i < len(accounts); i++ {
		for j := i + 1; j < len(accounts); j++ {
			if accounts[i][0] == accounts[j][0] {
				// A, B帳號. 名字一樣
				tmpAmail, tmpBmail, flag := accounts[i][1:], accounts[j][1:], false
				for j := 0; j < len(tmpAmail); j++ {
					for k := 0; k < len(tmpBmail); k++ {
						if tmpAmail[j] == tmpBmail[k] {
							// A,B 帳號, 找到相同的email. 認定為同一人
							flag = true
							break
						}
					}
					if flag {
						break
					}
				}
				if flag {
					uf.Union(i, j)
				}
			}
		}
	}
	for i := 0; i < len(accounts); i++ {
		if visited[i] {
			continue
		}
		emails, account, tmpMap := accounts[i][1:], []string{accounts[i][0]}, map[string]string{}
		for j := i + 1; j < len(accounts); j++ {
			if uf.Find(j) == uf.Find(i) {
				visited[j] = true
				for _, v := range accounts[j][1:] {
					tmpMap[v] = v
				}
			}
		}
		// A 帳號 Email
		for _, v := range emails {
			tmpMap[v] = v
		}
		emails = []string{}
		for key := range tmpMap {
			emails = append(emails, key)
		}
		sort.Strings(emails)
		account = append(account, emails...)
		result = append(result, account)
	}
	return result
}
