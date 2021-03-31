/*
 * @lc app=leetcode.cn id=621 lang=golang
 *
 * [621] 任务调度器
 */

// @lc code=start
package main
import "sort"

type sortedTask struct {
	task byte
	time int	//记录该任务下一次可以执行的时间点，初始为1
	n    *int	//记录该任务剩余的任务数，这里用指针，是因为要修改这个值
}

// 按每个任务的剩余任务数降序排列
type sortedTasks []sortedTask

func (a sortedTasks) Len() int           { return len(a) }
func (a sortedTasks) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortedTasks) Less(i, j int) bool { return *(a[i].n) > *(a[j].n) }

func leastInterval(tasks []byte, n int) int {
	m := map[byte]sortedTask{}
	for _,task := range tasks {
		if _,ok := m[task] ; !ok {
			taskNum := 1
			m[task] = sortedTask{task:task,time:1,n:&taskNum}
		}else{
			*(m[task].n)++
		}
	}

	sortedtasks := sortedTasks{}
	for _,v := range m {
		sortedtasks = append(sortedtasks,v)
	}

	sort.Sort(sortedtasks)
	allTime := 1
	l := len(sortedtasks)
	numTasks := l

	// 当剩余的任务总数为0时，退出
	for ; numTasks != 0 ; allTime++ {
		// 优先选取剩余任务数多的任务
		for i := 0 ; i < numTasks ; i++ {
			// if *(sortedtasks[i].n) == 0 {
			// 	continue
			// }
			if sortedtasks[i].time <= allTime {
				sortedtasks[i].time = sortedtasks[i].time+n+1
				*(sortedtasks[i].n)--
				if *(sortedtasks[i].n) == 0 {
					// 不需要下面这一句，我们每进行了一个任务之后就还要排序的，任务数为0自然排到最后去了
					// 通过numTasks调整遍历范围，自然就达到了删除0任务数的任务
					// copy(sortedtasks[i:],sortedtasks[i+1:])
					numTasks--
				}
				sort.Sort(sortedtasks)
				// 企图删掉任务数为0的任务，但提升很小
				sortedtasks = sortedtasks[:numTasks]
				break
			}
		}
	}
	return allTime-1
}

// @lc code=end

