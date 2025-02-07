/*
 * @lc app=leetcode.cn id=59 lang=typescript
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
// 螺旋数组，额，我是模拟人脑挨个螺旋填充的
function generateMatrix(n: number): number[][] {
    let res: number[][] = new Array(n).fill(0).map(() => new Array(n).fill(0));

    const oris: number[][] = [
        [0,1],
        [1,0],
        [0,-1],
        [-1,0]
    ];

    let x: number = 1;
    let i: number = 0;
    let j: number = 0;
    res[i][j] = x;
    x++;
    let ori: number = 0;
    while(x <= n*n) {
        ori %= 4;
        i += oris[ori][0];
        j += oris[ori][1];
        if (i < 0 || i >= n || j < 0 || j >= n || res[i][j] !== 0) {
            i -= oris[ori][0];
            j -= oris[ori][1];

            ori++;
            ori %= 4;
            i += oris[ori][0];
            j += oris[ori][1];
        }
        res[i][j] = x;
        x++;
    }
    return res;
};
// @lc code=end

