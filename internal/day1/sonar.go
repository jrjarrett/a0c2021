package day1

type Sonar struct{}

func (s *Sonar) DepthIncrease(sonarSweep []int) int {
	// degenerate case
	if len(sonarSweep) == 0 {
		return 0
	}

	initialDepth := sonarSweep[0]
	var depthIncreases int = 0
	for _, depth := range sonarSweep {
		if depth > initialDepth {
			depthIncreases += 1
		}
		initialDepth = depth
	}
	return depthIncreases
}

func (s *Sonar) CreateSlidingWindows(sonarSweep []int) []int {
	var slidingWindows []int
	size := 3
	var j int
	for i := 0; i < len(sonarSweep); i += 1 {
		j = i + size
		if j > len(sonarSweep) {
			return slidingWindows
		}
		// do what do you want to with the sub-slice, here just printing the sub-slices
		//fmt.Println(sonarSweep[i:j])
		slidingWindows = append(slidingWindows, addArray(sonarSweep[i:j]...))
	}
	return slidingWindows

}
func addArray(numbs ...int) int {
	var result int
	for _, numb := range numbs {
		result += numb
	}
	return result
}
