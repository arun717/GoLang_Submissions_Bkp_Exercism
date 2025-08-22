package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var total, i int
    for i = 0; i < len(birdsPerDay); i++ {
        total += birdsPerDay[i]
    }
    return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    high := 7*week
    var i, total int
    for i = high-7; i < high; i++ {
        total += birdsPerDay[i]
    }
    return total
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	var i int
    for i = 0; i < len(birdsPerDay); i++ {
        if i%2 == 0 {
            birdsPerDay[i] += 1
        }
    }
    return birdsPerDay
}
