package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	filePath    = "./01_input.txt"
	timeFormat  = "2006-01-02 15:04"
	beginShift  = "begins shift"
	fallsAsleep = "falls asleep"
	wakesUp     = "wakes up"
)

type timeRecord struct {
	timestamp time.Time
	message   string
}

type sleepSchedule struct {
	id         string
	sleepMins  []int
	totalSleep int
}

func main() {
	input := readInput()
	parsed := parseInput(input)
	sorted := sortInput(parsed)
	schedules := fillSchedules(sorted)

	guardID, min := findMostSleptMinute(schedules)

	guardIDNum, _ := strconv.Atoi(guardID)
	fmt.Println(min * guardIDNum)
}

func readInput() []string {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return strings.Split(strings.TrimRight(string(file), "\n"), "\n")
}

func parseInput(s []string) []timeRecord {
	tr := make([]timeRecord, len(s))

	for i, ss := range s {
		sa := strings.Split(ss, "] ")
		t, err := time.Parse(timeFormat, strings.TrimLeft(sa[0], "["))
		if err != nil {
			fmt.Println(err.Error())
		}

		tr[i] = timeRecord{
			message:   sa[1],
			timestamp: t,
		}
	}

	return tr
}

func sortInput(tr []timeRecord) []timeRecord {
	sort.Slice(tr, func(i, j int) bool {
		return tr[i].timestamp.Unix() < tr[j].timestamp.Unix()
	})

	return tr
}

func fillSchedules(tr []timeRecord) map[string]sleepSchedule {
	sm := make(map[string]sleepSchedule)
	currentGuard := ""
	sleepStart := 0

	for _, t := range tr {
		if ok, _ := regexp.MatchString(beginShift, t.message); ok {
			currentGuard = findGuardID(t.message)
		} else if ok, _ := regexp.MatchString(fallsAsleep, t.message); ok {
			sleepStart = t.timestamp.Minute()
		} else {
			ss, ok := sm[currentGuard]

			if ok {
				ss.totalSleep += t.timestamp.Minute() - sleepStart
				ss.sleepMins = appendMins(ss.sleepMins, sleepMins(t.timestamp.Minute(), sleepStart))

				sm[currentGuard] = ss
			} else {
				sm[currentGuard] = sleepSchedule{
					id:         currentGuard,
					sleepMins:  sleepMins(t.timestamp.Minute(), sleepStart),
					totalSleep: t.timestamp.Minute() - sleepStart,
				}
			}
		}
	}

	return sm
}

func findMostSleptMinute(sm map[string]sleepSchedule) (string, int) {
	globalMaxOccurence := 0
	globalMaxMinute := 0
	globalGuardID := ""

	for gid, schedule := range sm {
		sleepMins := schedule.sleepMins

		mm := map[int]int{}

		for _, min := range sleepMins {
			mm[min]++
		}

		currentMaxOccurence := 0
		maxMinute := 0

		for m, occ := range mm {
			if occ > currentMaxOccurence {
				currentMaxOccurence = occ
				maxMinute = m
			}
		}

		if currentMaxOccurence > globalMaxOccurence {
			globalMaxOccurence = currentMaxOccurence
			globalMaxMinute = maxMinute
			globalGuardID = gid
		}
	}

	return globalGuardID, globalMaxMinute
}

func findGuardID(s string) string {
	return strings.Split(strings.Split(s, "#")[1], " ")[0]
}

func sleepMins(end, start int) []int {
	is := []int{}

	for i := start; i < end; i++ {
		is = append(is, i)
	}

	return is
}

func appendMins(s, m []int) []int {
	for _, mm := range m {
		s = append(s, mm)
	}

	return s
}
