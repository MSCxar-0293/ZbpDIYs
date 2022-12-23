package card21

import (
	"math/rand"
	"strconv"
)

func rdnum(max int) int {
	a := rand.Intn(max)
	return a
}

func puid(player string) int64 {
	playeruid, _ := strconv.ParseInt(player, 10, 64)
	return playeruid
}

func judger21A(cards []string) int {
	var s int
	for k := 0; k <= len(cards)-1; k = k + 1 {
		value := cards[k][6:]
		switch {
		case value == "A":
			value = "1"
		case value == "J" || value == "Q" || value == "K":
			value = "10"
		}
		valuei, _ := strconv.Atoi(value)
		s = s + valuei
	}
	score := s - 21
	return score
}

func judger21B(cards []string) int {
	var s int
	for k := 0; k <= len(cards)-1; k = k + 1 {
		value := cards[k][6:]
		switch {
		case value == "A":
			value = "11"
		case value == "J" || value == "Q" || value == "K":
			value = "10"
		}
		valuei, _ := strconv.Atoi(value)
		s = s + valuei
	}
	score := s - 21
	return score
}

func judgerfin(result1, result2 int) int {
	if result1 < 0 && result2 > 0 {
		return result1
	} else if result1 > 0 && result2 < 0 {
		return result2
	} else {
		if result1 < result2 {
			return result2
		} else {
			return result1
		}
	}
}

func judger17A(cards []string) int {
	var s int
	for k := 0; k <= len(cards)-1; k = k + 1 {
		value := cards[k][6:]
		switch {
		case value == "A":
			value = "1"
		case value == "J" || value == "Q" || value == "K":
			value = "10"
		}
		valuei, _ := strconv.Atoi(value)
		s = s + valuei
	}
	point := s - 17
	return point
}

func judger17B(cards []string) int {
	var s int
	for k := 0; k <= len(cards)-1; k = k + 1 {
		value := cards[k][6:]
		switch {
		case value == "A":
			value = "10"
		case value == "J" || value == "Q" || value == "K":
			value = "10"
		}
		valuei, _ := strconv.Atoi(value)
		s = s + valuei
	}
	point := s - 17
	return point
}

// Contains 数组是否包含某元素
func Contains(slice []string, s string) int {
	for index, value := range slice {
		if value == s {
			return index
		}
	}
	return -1
}

func or(s1, s2 bool) bool {
	if s1 || s2 {
		return true
	} else {
		return false
	}
}

func winner(m map[string]int) string {
	max := -10000
	winnerr := ""
	for num := range m {
		if m[num] > max {
			max = m[num]
			winnerr = num
		}
	}
	return winnerr
}
