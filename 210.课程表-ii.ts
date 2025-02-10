/*
 * @lc app=leetcode.cn id=210 lang=typescript
 *
 * [210] 课程表 II
 */

// @lc code=start
// 拓扑排序
function findOrder(numCourses: number, prerequisites: number[][]): number[] {
    // 记录已排序的结果
    let order: number[] = new Array<number>();
    // 记录入度
    let inNum: number[] = new Array<number>(numCourses).fill(0);
    // 记录后继
    let next: number[][] = new Array(numCourses).fill(0).map(() => new Array());

    prerequisites.forEach((prerequisite) => {
        inNum[prerequisite[0]]++;
        next[prerequisite[1]].push(prerequisite[0]);
    })

    inNum.forEach((value, index) => {
        if (value === 0) {
            order.push(index);
        }
    })

    let deli: number = 0;
    while(deli < order.length) {
        let del = order[deli];
        next[del].forEach((value) => {
            inNum[value]--;
            if (inNum[value] === 0) {
                order.push(value);
            }
        })
        deli++;
    }
    if (order.length != numCourses) {
        return new Array<number>();
    }
    return order;
};
// @lc code=end

