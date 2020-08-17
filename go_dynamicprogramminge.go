package main

/*
* 动态规划（Dynamic Programming）
* 需要在给定约束条件下优化某种指标时，动态规划很有用。
* 问题可分解为离散子问题时，可使用动态规划来解决。
* 每种动态规划解决方案都涉及网格。
* 单元格中的值通常就是你要优化的值。
* 每个单元格都是一个子问题，因此你需要考虑如何将问题分解为子
问题。
* 没有放之四海皆准的计算动态规划解决方案的公式。
*/

//计算两个字符串的最长相似字符串数
func subsequence(a, b string) int {
	cell := createMatrix(len(a)+1, len(b)+1)

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				cell[i][j] = cell[i-1][j-1] + 1
			} else {
				cell[i][j] = cell[i-1][j]

				if cell[i][j] < cell[i][j-1] {
					cell[i][j] = cell[i][j-1]
				}
			}
		}
	}
	return cell[len(a)][len(b)]
}

func createMatrix(rows, cols int) [][]int {
	cell := make([][]int, rows)
	for i := range cell {
		cell[i] = make([]int, cols)
	}
	return cell
}
//获取两个字符串的最长相似字符串
func substring(a, b string) string {
	lcs := 0
	lastSubIndex := 0
	cell := createMatrix(len(a)+1, len(b)+1)

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				cell[i][j] = cell[i-1][j-1] + 1

				if cell[i][j] > lcs {
					lcs = cell[i][j]
					lastSubIndex = i
				}
			} else {
				cell[i][j] = 0
			}
		}
	}
	return a[lastSubIndex-lcs : lastSubIndex]
}