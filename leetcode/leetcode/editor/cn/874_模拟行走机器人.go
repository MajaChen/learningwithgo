package leetcode

import "math"

//leetcode submit region begin(Prohibit modification and deletion)

type coordinate struct {
	x int
	y int
}

func robotSim(commands []int, obstacles [][]int) int {

	maximalDistance, px, py := math.MinInt32, 0, 0

	obstacleMapping := make(map[coordinate]bool)
	for _, obstacle := range obstacles {
		obstacleMapping[coordinate{obstacle[0], obstacle[1]}] = true
	}

	funcMapping := make(map[int]func(int))
	funcMapping[0] = func(step int) {
		var i, j int
		for i, j = px+1, py; i <= px+step && !obstacleMapping[coordinate{i, j}]; i++ {
		}
		px = i - 1
	}
	funcMapping[90] = func(step int) {
		var i, j int
		for i, j = px, py+1; j <= py+step && !obstacleMapping[coordinate{i, j}]; j++ {
		}
		py = j - 1
	}
	funcMapping[180] = func(step int) {
		var i, j int
		for i, j = px-1, py; i >= px-step && !obstacleMapping[coordinate{i, j}]; i-- {
		}
		px = i + 1
	}
	funcMapping[270] = func(step int) {
		var i, j int
		for i, j = px, py-1; j >= py-step && !obstacleMapping[coordinate{i, j}]; j-- {
		}
		py = j + 1
	}

	angleMapping := map[int]int{-1: -90, -2: 90}

	angle := 90
	for _, command := range commands {
		if command < 0 {
			// 角度变换
			angle += angleMapping[command] + 360
			angle %= 360
			continue
		}
		funcMapping[angle](command)
		if dist := int(math.Pow(float64(px), 2) + math.Pow(float64(py), 2)); dist > maximalDistance {
			maximalDistance = dist
		}
	}
	return maximalDistance
}

//leetcode submit region end(Prohibit modification and deletion)
