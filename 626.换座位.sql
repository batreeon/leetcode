--
-- @lc app=leetcode.cn id=626 lang=mysql
--
-- [626] 换座位
--

-- @lc code=start
# Write your MySQL query statement below
SELECT 
    (CASE 
        WHEN id%2=1 AND counts != id THEN id+1
        WHEN id%2=1 AND counts = id THEN id
        ELSE id-1
        END
    )AS id,student
FROM seat,
    (
        SELECT COUNT(*) AS counts 
        FROM seat
    ) AS seatcount
ORDER BY id
-- @lc code=end

