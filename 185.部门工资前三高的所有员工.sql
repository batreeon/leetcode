--
-- @lc app=leetcode.cn id=185 lang=mysql
--
-- [185] 部门工资前三高的所有员工
--

-- @lc code=start
# Write your MySQL query statement below

-- SELECT d.Name AS Department,
--         e1.Name AS Employee,
--         e1.Salary as Salary
-- FROM  Employee e1,Department d
-- WHERE d.Id = e1.DepartmentId AND 
--         3 > (
--             SELECT COUNT(DISTINCT e2.Salary)
--             FROM Employee e2
--             WHERE e2.Salary > e1.Salary AND 
--                     e1.DepartmentId = e2.DepartmentId 
--                     /*这一句很重要，同部门的才需要比较。很奇怪，这条注释用--写会报错*/
--         )

SELECT d.Name AS Department,
        e1.Name AS Employee,
        e1.Salary as Salary
FROM  Employee e1 JOIN Department d ON d.Id = e1.DepartmentId
WHERE 3 > (
            SELECT COUNT(DISTINCT e2.Salary)
            FROM Employee e2
            WHERE e2.Salary > e1.Salary AND 
                    e1.DepartmentId = e2.DepartmentId 
                    /*这一句很重要，同部门的才需要比较。很奇怪，这条注释用--写会报错*/
        )

-- @lc code=end

