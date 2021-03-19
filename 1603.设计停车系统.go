/*
 * @lc app=leetcode.cn id=1603 lang=golang
 *
 * [1603] 设计停车系统
 */

// @lc code=start
type ParkingSystem struct {
	parking *[3]int
}


func Constructor(big int, medium int, small int) ParkingSystem {
	parking := [3]int{}
	parking[0] = big
	parking[1] = medium
	parking[2] = small
	ps := ParkingSystem{&parking}
	return ps
}


func (this *ParkingSystem) AddCar(carType int) bool {
	if (this.parking)[carType-1] <= 0 {
		return false
	}
	(this.parking)[carType-1]--
	return true
}


/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
// @lc code=end

