/*
 * @lc app=leetcode.cn id=2374 lang=typescript
 *
 * [2374] 边积分最高的节点
 */

// @lc code=start
function edgeScore(edges: number[]): number {
    let map: Map<number,number> = new Map<number,number>();
    let maxNums: number[] = new Array<number>();
    let maxScore: number = 0;
    for(let outNum = 0; outNum < edges.length; outNum++) {
        let inNum: number = edges[outNum];
        let score:number = map.get(inNum) || 0;
        score += outNum
        if (score > maxScore) {
            maxScore = score;
            maxNums.length = 0;
            maxNums.push(inNum);
        }else if(score === maxScore) {
            maxNums.push(inNum);
        }
        map.set(inNum, score);
    }

    let result: number = edges.length;
    for (let inNum of maxNums) {
        if (inNum < result) {
            result = inNum;
        }
    }
    return result;
};
// @lc code=end

