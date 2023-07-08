package race

import (
	"fmt"
	"math/rand"
	"strings"
)

type Race struct {
	TotalHorses   int
	RaceCount     int
	GoalToFind    int
	RaceSize      int
	HorseTimes    map[int]float32
	winningHorses []int
}

func New() *Race {
	var race Race
	// fmt.Print("Total horses: ")
	// fmt.Scan(&race.Total)
	// fmt.Print("Goal: ")
	// fmt.Scan(&race.Goal)
	// fmt.Print("Race size: ")
	// fmt.Scan(&race.RaceSize)
	race.TotalHorses = 25
	race.GoalToFind = 3
	race.RaceSize = 5
	race.GenerateTimes()
	return &race
}

func (r *Race) GenerateTimes() {
	r.HorseTimes = make(map[int]float32)
	for i := 0; i < r.TotalHorses; i++ {
		r.HorseTimes[i+1] = rand.Float32()
	}
	r.getWinningHorses()
}

func (r *Race) getWinningHorses() {
	var min float32 = 1
	var index int = 0
	for len(r.winningHorses) < r.TotalHorses {
		for k, v := range r.HorseTimes {
			if v < min {
				min = v
				index = k
			}
		}
		r.HorseTimes[index] = 10
		min = 1
		r.winningHorses = append(r.winningHorses, index)
	}
}

func (r *Race) ShowTimes() {
	for k, v := range r.HorseTimes {
		fmt.Printf("%d: %f\n", k, v)
	}
}

func (r *Race) RunRace() {
	horses := getNHorses(r.RaceSize)
	r.GetWinners(horses)
	r.RaceCount++
}

func getNHorses(n int) []int {
	format := ""
	for i := 0; i < n; i++ {
		format += "%d "
	}
	format = strings.TrimSpace(format)
	fmt.Printf("\nChoose the %d horses: ", n)
	horses := make([]int, n)
	for i := 0; i < n; i++ {
		var horse int
		fmt.Scan(&horse)
		horses[i] = horse
	}
	return horses
}

func (r *Race) PrintWinners() {
	fmt.Println(r.winningHorses)
}

func (r *Race) GetWinners(horses []int) {
	winners := []int{}
	for _, w := range r.winningHorses {
		for _, v := range horses {
			if v == w {
				winners = append(winners, v)
			}
		}
	}
	fmt.Println(winners)
}

func (r *Race) TestWinners() {
	horses := getNHorses(r.GoalToFind)
	for i := 0; i < r.GoalToFind; i++ {
		if horses[i] != r.winningHorses[i] {
			fmt.Println("Those are not the winners!")
			return
		}
	}
	fmt.Printf("It took you %d races, good job!\n", r.RaceCount)
}
