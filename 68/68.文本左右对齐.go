/*
 * @lc app=leetcode.cn id=68 lang=golang
 *
 * [68] 文本左右对齐
 */

// @lc code=start
package main

func fullJustify(words []string, maxWidth int) []string {
	result := []string{}
	for i := 0 ; i < len(words) ; {
		lenwords := len(words[i])
		j := i+1
		for ; j < len(words) ; j++ {
			if lenwords + 1 + len(words[j]) <= maxWidth {
				lenwords += 1 + len(words[j])
			}else{
				break
			}
		}
		// j-i是单词数，那么j-i-1就是间隙数
		gap := j-i-1
		// 因为前面在计算lenwords包含了一个空格，所以计算space数时要加上gap
		space := maxWidth - lenwords + gap
		// 这一行的数据
		row := make([]byte,maxWidth)
		// 在当前行的写入位置
		idx := 0

		if j == len(words) {
		// 是最后一行，因为题目说了最后一行左对齐，不用保持空格平衡
			for k := i ; k < j ; k++ {
				copy(row[idx:],[]byte(words[k]))
				idx += len(words[k])
				// 注意如果到末尾了，就不用在单词末尾加空格了，否则在单词末尾加一个空格
				if idx < maxWidth {
					row[idx] = ' '
					idx++
				}
			}
			// 如果单词填充完了，还有剩余多的空格，则全部添加到末尾
			for ; idx < maxWidth ; idx++ {
				row[idx] = ' '
			}
		}else{
			// 不是最后一行
			for k := i ; k < j ; k++ {
				copy(row[idx:],[]byte(words[k]))
				idx += len(words[k])
				if gap > 0 {
					// 为了保持空格平衡，并且左边的尽可能多
					// nextSpace计算该空隙应该放置的空格数
					nextSpace := space/gap
					// 如果总空格数能整除gap，那么每一个间隙就放一样多的空格，否则就在靠左侧的间隙多放一个
					// 只多放一个，为了保持平衡
					if space % gap != 0 {
						nextSpace++
					}
					// 从左到右放置，没放一个，space就减去nextSpace
					space -= nextSpace
					gap--
					// 放置空格
					for ; nextSpace > 0 ; nextSpace-- {
						row[idx] = ' '
						idx++
					}
				}else if space > 0 {
					// 只有一个单词，gap为0，但是空格数不为0
					for ; space > 0 ; space-- {
						row[idx] = ' '
						idx++
					}
				}
			}
		}
		result = append(result,string(row))
		i = j
	}
	return result
}
// @lc code=end

