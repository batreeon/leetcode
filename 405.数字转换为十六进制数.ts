/*
 * @lc app=leetcode.cn id=405 lang=typescript
 *
 * [405] 数字转换为十六进制数
 */

// @lc code=start
function toHex(num: number): string {
    if (num === 0) {
        return "0"
    }
    let result: string = "";
    let remainder: number = 0;
    let map: Map<number, string> = new Map ([
        [10,'a'],
        [11,'b'],
        [12,'c'],
        [13,'d'],
        [14,'e'],
        [15,'f'],
    ])
    if (num<0) {
        num = num + 2**32;
    }
    while(num > 0) {
        let quotient: number = Math.floor(num/16);
        remainder = num - (quotient * 16);
        if (remainder < 10) {
            result = remainder + result;
        }else{
            result = map.get(remainder) + result;
        }
        num = quotient;
    }

    return result;
};
// @lc code=end

