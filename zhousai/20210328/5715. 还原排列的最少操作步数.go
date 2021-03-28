/*
给你一个偶数 n​​​​​​ ，已知存在一个长度为 n 的排列 perm ，其中 perm[i] == i​（下标 从 0 开始 计数）。

一步操作中，你将创建一个新数组 arr ，对于每个 i ：

如果 i % 2 == 0 ，那么 arr[i] = perm[i / 2]
如果 i % 2 == 1 ，那么 arr[i] = perm[n / 2 + (i - 1) / 2]
然后将 arr​​ 赋值​​给 perm 。

要想使 perm 回到排列初始值，至少需要执行多少步操作？返回最小的 非零 操作步数。
*/
func reinitializePermutation(n int) int {
    //n为偶数时，头尾是不会变的，我们盯着一个元素，只要他还原了，其他也就还原了
    rp := func(n int) int {
        changes := make([]int,n)
        changes[0],changes[n-1] = 1,1
        for i := 1 ; i < n/2; i++ {
            changes[i] = i
        }
        for i := n/2 ; i < n ; i++ {
            changes[i] = -1 * (n-1-i)
        }
        times := 1
        index := 1
        change := changes[index]
        for change != 0 {
            index = index+changes[index]
            change += changes[index]
            times++
        }
        return times
    }
    if n%2 == 0 {
        return rp(n)
    }else{
        return rp(n+1)
    }
}