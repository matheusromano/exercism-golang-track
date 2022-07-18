package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var totalBirdCount int
	for _, bird := range birdsPerDay {
		totalBirdCount += bird
	}
	return totalBirdCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var totalBirdCountWeed int
	if week == 1 {
		for i := 0; i < 7; i++ {
			totalBirdCountWeed += birdsPerDay[i]
		}
	}
	if week == 2 {
		for i := 7; i < 15; i++ {
			totalBirdCountWeed += birdsPerDay[i]
		}
	}
	return totalBirdCountWeed
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := range birdsPerDay {
		if i%2 == 0 {
			birdsPerDay[i]++
		}
	}
	return birdsPerDay
}
