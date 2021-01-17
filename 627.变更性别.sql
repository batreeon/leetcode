--
-- @lc app=leetcode.cn id=627 lang=mysql
--
-- [627] 变更性别
--

-- @lc code=start
# Write your MySQL query statement below
UPDATE salary
SET sex = (
    CASE
    WHEN sex='m' THEN 'f'
    ELSE 'm'
    END
)
-- @lc code=end

