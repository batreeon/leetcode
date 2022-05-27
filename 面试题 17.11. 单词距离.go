import "sort"
func findClosest(words []string, word1 string, word2 string) int {
    indexs := map[string][]int{}
    for i, word := range words {
        if _, ok := indexs[word]; !ok {
            indexs[word] = []int{}
        }
        indexs[word] = append(indexs[word], i)
    }
    result := len(words)
    for _, index := range indexs[word1] {
        i := sort.Search(len(indexs[word2]), func(i int) bool {return indexs[word2][i] > index})
        if i == len(indexs[word2]) {
            nearIndex := indexs[word2][len(indexs[word2])-1]
            result = min(result, index - nearIndex)
        }else if i == 0 {
            nearIndex := indexs[word2][0]
            result = min(result, nearIndex - index)
        }else{
            result = min(result, min(index - indexs[word2][i-1], indexs[word2][i] - index))
        }
    }
    return result
}
func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}