package main

func main() {

}

// https://leetcode.cn/problems/my-calendar-i/

type pair struct {
	start int
	end   int
}

type MyCalendar []pair

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (c *MyCalendar) Book(start int, end int) bool {
	for _, p := range *c {
		if p.start < end && start < p.end {
			return false
		}
	}
	*c = append(*c, pair{
		start: start,
		end:   end,
	})
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
