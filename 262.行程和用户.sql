--
-- @lc app=leetcode.cn id=262 lang=mysql
--
-- [262] 行程和用户
--

-- @lc code=start
# Write your MySQL query statement below
-- SELECT T.Request_at AS Day ,
--         ROUND(
--             SUM(IF(T.Status='complete',0,1))--统计取消的订单个数
--             /
--             COUNT(T.Status)
--             ,2
--         )AS 'Cancellation Rate'
-- FROM Trips AS T 
-- WHERE T.Client_Id NOT IN (
--         SELECT Users_Id
--         FROM Users
--         WHERE Banned='Yes'
--     ) AND
--     T.Driver_Id NOT IN (
--         SELECT Users_Id
--         FROM Users
--         WHERE Banned='Yes'
--     ) AND T.Request_at BETWEEN '2013-10-01' AND '2013-10-03'
-- GROUP BY T.Request_at
SELECT T.request_at AS 'Day', 
	ROUND(
			SUM(
				IF(T.STATUS = 'completed',0,1)
			)
			/ 
			COUNT(T.STATUS),
			2
	) AS 'Cancellation Rate'
FROM trips AS T
WHERE 
T.Client_Id NOT IN (
	SELECT users_id
	FROM users
	WHERE banned = 'Yes'
)
AND
T.Driver_Id NOT IN (
	SELECT users_id
	FROM users
	WHERE banned = 'Yes'
)
AND T.request_at BETWEEN '2013-10-01' AND '2013-10-03'
GROUP BY T.request_at
-- @lc code=end

