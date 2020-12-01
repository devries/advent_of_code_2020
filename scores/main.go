package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type StarTs interface{}

type Completion struct {
	GetStarTs string `json:"get_star_ts"`
}

type Member struct {
	LocalScore         int                        `json:"local_score"`
	CompletionDayLevel map[int]map[int]Completion `json:"completion_day_level"`
	Stars              int                        `json:"stars"`
	GlobalScore        int                        `json:"global_score"`
	Id                 string                     `json:"id"`
	Name               string                     `json:"name"`
	LastStarTs         StarTs                     `json:"last_star_ts"`
}

type Scoreboard struct {
	Members map[int]Member `json:"members"`
	Event   string         `json:"event"`
	OwnerId string         `json:"owner_id"`
}

func StarTsValue(sts StarTs) int64 {
	switch sts := sts.(type) {
	case string:
		ts, err := strconv.ParseInt(sts, 10, 64)
		check(err)
		return ts
	case int:
		return int64(sts)
	default:
		return 0
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <scorefile>\n", os.Args[0])
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	check(err)

	decoder := json.NewDecoder(f)

	var s Scoreboard

	err = decoder.Decode(&s)
	check(err)

	// Print out daily results for each member
	memberNumbers := []int{}
	for k, _ := range s.Members {
		memberNumbers = append(memberNumbers, k)
	}

	eventYear, err := strconv.ParseInt(s.Event, 10, 64)
	check(err)

	now := time.Now()

	for i := 1; i < 26; i++ {
		dayStart := time.Date(int(eventYear), 12, i, 5, 0, 0, 0, time.UTC)
		if dayStart.After(now) {
			break
		}
		for _, j := range []int{1, 2} {
			fmt.Printf("Day %2d part %d:\n", i, j)
			for _, n := range memberNumbers {
				completions, ok := s.Members[n].CompletionDayLevel[i]
				if !ok {
					continue
				}
				completion, ok := completions[j]
				if !ok {
					continue
				}

				ts, err := strconv.ParseInt(completion.GetStarTs, 10, 64)
				check(err)
				doneAt := time.Unix(ts, 0)
				dur := doneAt.Sub(dayStart)

				fmt.Printf("%25s: %s\n", s.Members[n].Name, fmtDuration(dur))
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}

	for _, n := range memberNumbers {
		ts := StarTsValue(s.Members[n].LastStarTs)
		if ts == 0 {
			fmt.Printf("%20s did not complete any stars\n", s.Members[n].Name)
		} else {
			finished := time.Unix(ts, 0)
			fmt.Printf("%20s finished %d starts on %s\n", s.Members[n].Name, s.Members[n].Stars, finished.Format("January 2, 2006 at 03:04 PM"))
		}
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute

	if h > 0 {
		return fmt.Sprintf("%4d h %2d m", h, m)
	} else {
		return fmt.Sprintf("       %2d m", m)
	}
}
