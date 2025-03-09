#
# @lc app=leetcode.cn id=1604 lang=python3
#
# [1604] 警告一小时内使用相同员工卡大于等于三次的人
#

# @lc code=start
class Solution:
    def alertNames(self, keyName: List[str], keyTime: List[str]) -> List[str]:
        if len(keyName) < 3:
            return []
        
        staff: dict[str : list[str]] = {}
        for i in range(len(keyName)):
            name: str = keyName[i]
            time: str = keyTime[i]
            times: list[str] = staff.get(name, [])
            times.append(time)
            staff[name] = times
        
        result: list[str] = [
            name 
            for name, times in staff.items() 
            if len(times) >= 3 and self.alert(times)
            ]
        result.sort()

        return result

    
    # t1 is later than t2
    def duration(self, t1: str, t2: str) -> int:
        h1: int = int(t1[:2])
        m1: int = int(t1[3:])
        h2: int = int(t2[:2])
        m2: int = int(t2[3:])

        if h1 == h2:
            return m1 - m2
        return (h1 - h2 - 1) * 60 + (m1 + (60 - m2))

    def alert(self, times: list[str]) -> bool:
        times.sort()
        for i in range(len(times)-2):
            if self.duration(times[i+2], times[i]) <= 60:
                return True
        return False
# @lc code=end

