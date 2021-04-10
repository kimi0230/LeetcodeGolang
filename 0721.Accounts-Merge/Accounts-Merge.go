package accountsmerge

import (
	"LeetcodeGolang/template"
	"sort"
)

// TODO: 優化
func AccountsMerge(accounts [][]string) [][]string {
	return [][]string{}
}

// 暴力解 : Union Find
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
