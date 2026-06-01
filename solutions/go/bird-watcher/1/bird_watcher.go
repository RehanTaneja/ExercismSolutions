package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var BirdCount int = 0
    for i := range birdsPerDay{
        BirdCount += birdsPerDay[i]
    }
    return BirdCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var BirdCount int = 0
    var Birds_week []int = birdsPerDay[(week-1)*7:((week-1)*7)+7]
    for i := range Birds_week{
        BirdCount += Birds_week[i]
    }
    return BirdCount
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := range birdsPerDay{
        if i == 0 || i%2 == 0{
            birdsPerDay[i] += 1
        }
    }
    return birdsPerDay
}
