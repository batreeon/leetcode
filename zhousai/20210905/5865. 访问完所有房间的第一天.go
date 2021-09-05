func firstDayBeenInAllRooms(nextVisit []int) int {
	n := len(nextVisit)
	counts := make([]int,n)
	seen := 0
	day := -1
	room := 0
	for seen != n {
		day++
		counts[room]++
		if counts[room] == 1 {
			seen++
		}
		if counts[room] & 1 == 1 {
			room = nextVisit[room]
		}else{
			room = (room + 1) % n
		}
	}
	return day
}