/*
有一个特殊打字机，它由一个 圆盘 和一个 指针 组成， 圆盘上标有小写英文字母 'a' 到 'z'。只有 当指针指向某个字母时，它才能被键入。指针 初始时 指向字符 'a' 。


每一秒钟，你可以执行以下操作之一：

将指针 顺时针 或者 逆时针 移动一个字符。
键入指针 当前 指向的字符。
给你一个字符串 word ，请你返回键入 word 所表示单词的 最少 秒数 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-time-to-type-word-using-special-typewriter
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func minTimeToType(word string) int {
    last := 'a'
    result := 0
    for _,b := range word {
        step := int(b-last)
        if step < 0 {
            step *= -1
        }
        if step > 13 {
            step = 26 - step
        }
        last = b
        result += step
    }
    result += len(word)
    return result
}