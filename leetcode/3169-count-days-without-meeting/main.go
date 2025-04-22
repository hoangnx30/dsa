package main

import (
	"fmt"
	"sort"
)

func countDays(days int, meetings [][]int) int {
	prevMeetingDay := 0
	result := 0

	sort.SliceStable(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	for index, meeting := range meetings {
		start := meeting[0]
		end := meeting[1]

		if start > prevMeetingDay {
			result += start - prevMeetingDay - 1
		}
		prevMeetingDay = max(prevMeetingDay, end)

		if index == len(meetings)-1 && days != prevMeetingDay {
			result += days - prevMeetingDay
		}

		fmt.Println(index, start, prevMeetingDay, result)

	}

	return result
}

func countDaysBySubtractingMeetingDays(days int, meetings [][]int) int {
	prevMeetingDay := 0

	sort.SliceStable(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	for _, meeting := range meetings {
		start := meeting[0]
		end := meeting[1]

		start = max(prevMeetingDay+1, start)
		length := end - start + 1
		days -= max(length, 0)
		prevMeetingDay = max(prevMeetingDay, end)
	}

	return days
}

func main() {
	days := 8
	meetings := [][]int{{3, 4}, {4, 8}, {2, 5}, {3, 8}}

	fmt.Println(countDaysBySubtractingMeetingDays(days, meetings))
}
