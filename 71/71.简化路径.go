/*
 * @lc app=leetcode.cn id=71 lang=golang
 *
 * [71] 简化路径
 */

// @lc code=start
package main
import "strings"
func simplifyPath(path string) string {
	/*
	array := strings.Split(path,"/")
	if len(array) == 0 {
		return string("/")
	}
	now := []string{}
	for len(array) > 0 {
		if array[0] == string(".."){
			if len(now) > 0 {
				now = now[:len(now)-1]
			}
		}else if array[0] != string(".") && array[0] != string(""){
			now = append(now,array[0])
		}
		array = array[1:]
	}
	return strings.Join([]string{"/",strings.Join(now,"/")},"")
	*/

	
	array := strings.Fields(strings.Join(strings.Split(path,"/")," "))
	if len(array) == 0 {
		return string("/")
	}
	now := []string{}
	for len(array) > 0 {
		if array[0] == string("..") {
			if len(now) > 0 {
				now = now[:len(now)-1]
			}
		}else if array[0] != string(".") {
			now = append(now,array[0])
		}
		array = array[1:]
	}
	return strings.Join([]string{"/",strings.Join(now,"/")},"")
	

	/*想法不错，但不对
	back := strings.Count(path,string(".."))
	array := strings.Fields(strings.Join(strings.Split(strings.ReplaceAll(path,string("."),""),"/")," "))
	if len(array) >= back {
		array = array[:len(array)-back]
	}else{
		return string("/")
	}
	return strings.Join([]string{"/",strings.Join(array,"/")},"")
	*/
}
// @lc code=end

