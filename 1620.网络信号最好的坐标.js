/*
 * @lc app=leetcode.cn id=1620 lang=javascript
 *
 * [1620] 网络信号最好的坐标
 */

// @lc code=start
/**
 * @param {number[][]} towers
 * @param {number} radius
 * @return {number[]}
 */
var bestCoordinate = function(towers, radius) {
    let maxX = towers.sort((a,b) => b[0]-a[0])[0][0];
    let maxY = towers.sort((a,b) => b[1]-a[1])[0][1];

    let result = [0,0];
    let maxIntensity = 0;

    for (let x = 0; x <= maxX; x++) {
        for (let y = 0; y <= maxY; y++) {
            let intensity = 0;
            for (let tower of towers) {
                let xi = tower[0], yi = tower[1], qi = tower[2];
                let distance2 = ((x - xi)**2 + (y - yi)**2);
                if (distance2 <= radius**2) {
                    intensity += Math.floor(qi / (1 + distance2**0.5));
                }
            }

            if (intensity > maxIntensity) {
                result[0] = x;
                result[1] = y;
                maxIntensity = intensity;
            }
        }
    }
    return result;
};
// @lc code=end

