package main

func main() {

}

//https://leetcode-cn.com/problems/first-bad-version/
func firstBadVersion(n int) int {
	l, r := 0, n
	for l < r {
		mid := (r-l)/2 + l
		if isBadVersion(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 */
func isBadVersion(version int) bool {

}
