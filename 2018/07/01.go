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

	guardID := findGuardMostAsleep(schedules)
	min := findMostSleptMinute(guardID, schedules)

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

func findGuardMostAsleep(sm map[string]sleepSchedule) string {
	currentMax := 0
	id := ""

	for gid, ss := range sm {
		if currentMax < ss.totalSleep {
			currentMax = ss.totalSleep
			id = gid
		}
	}

	return id
}

func findMostSleptMinute(gid string, sm map[string]sleepSchedule) int {
	sleepMins := sm[gid].sleepMins

	mm := map[int]int{}

	for _, min := range sleepMins {
		mm[min]++
	}

	currentMax := 0
	maxMinute := 0

	for m, occ := range mm {
		if occ > currentMax {
			currentMax = occ
			maxMinute = m
		}
	}

	return maxMinute
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
	// fmt.Println(s, m)
	for _, mm := range m {
		s = append(s, mm)
	}
	// fmt.Println(s)
	return s
}
