/*
 * @lc app=leetcode.cn id=981 lang=golang
 *
 * [981] 基于时间的键值存储
 */

// @lc code=start
package main
// import "sort"
type Value struct {
	v string
	time int
}

type TimeMap struct {
	m map[string][]Value
}


/** Initialize your data structure here. */
func Constructor() TimeMap {
	return TimeMap{map[string][]Value{}}
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
	/*
	length := len(this.m[key])
	index := sort.Search(length,func(i int) bool {return this.m[key][i].time >= timestamp})
	if index == length {
		this.m[key] = append(this.m[key],Value{value,timestamp})
	}else{
		this.m[key] = append(this.m[key],Value{})
		copy(this.m[key][index+1:],this.m[key][index:length-1])
		this.m[key][index] = Value{value,timestamp}
	}
	*/
	// 啊，我去，没认真读题，set操作时间戳都是递增的
	this.m[key] = append(this.m[key],Value{value,timestamp})
}


func (this *TimeMap) Get(key string, timestamp int) string {
	// sort.Slice(this.m[key],func(i,j int) bool {return this.m[key][i].time < this.m[key][j].time})
	for i := len(this.m[key])-1 ; i >= 0 ; i-- {
		if this.m[key][i].time <= timestamp {
			return this.m[key][i].v
		}
	}
	return ""
}


/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
// @lc code=end

