/*
 * @lc app=leetcode.cn id=1678 lang=javascript
 *
 * [1678] 设计 Goal 解析器
 */

// @lc code=start
/**
 * @param {string} command
 * @return {string}
 */
var interpret = function(command) {
    let result = new String();
    for (; command.length > 0; ) {
        if (command[0] == 'G') {
            result += 'G';
            command = command.slice(1,);
        }else if (command.slice(0,2) == "()") {
            result += "o";
            command = command.slice(2,);
        }else{
            result += "al";
            command = command.slice(4,);
        }
    }
    return result;
};
// @lc code=end

