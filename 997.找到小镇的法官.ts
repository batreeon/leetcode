/*
 * @lc app=leetcode.cn id=997 lang=typescript
 *
 * [997] 找到小镇的法官
 */

// @lc code=start
function findJudge(n: number, trust: number[][]): number {
    if (n == 1) {
        return 1;
    }

    let trustOther: number[] = new Array<number>(n+1).fill(0);
    let beTrusted: number[] = new Array<number>(n+1).fill(0);
    let beTrustedAll: number[] = new Array<number>;

    for (let t of trust) {
        trustOther[t[0]]++;
        beTrusted[t[1]]++;
        if (beTrusted[t[1]] == n-1) {
            beTrustedAll.push(t[1]);
        }
    }

    for (let betrustall of beTrustedAll) {
        if (trustOther[betrustall] == 0) {
            return betrustall;
        }
    }

    return -1;
};
// @lc code=end

