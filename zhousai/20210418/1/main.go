func checkIfPangram(sentence string) bool {
    cnt := map[rune]bool{}
    for _,s := range sentence {
        cnt[s] = true
        if len(cnt) == 26 {
            return true
        }
    }
    return false
}