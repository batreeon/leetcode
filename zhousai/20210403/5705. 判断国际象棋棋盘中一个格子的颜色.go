func squareIsWhite(coordinates string) bool {
    x,y := coordinates[0]-'0',(coordinates[1]-'a')+1
	if x&1 == 1 {
		if y&1 == 1 {
			return false
		}else{
			return true
		}
	}else{
		if y&1 == 1 {
			return true
		}else{
			return false
		}
	}
	return false
}