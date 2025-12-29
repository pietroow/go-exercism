package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    var count int
    for _,num := range birdsPerDay {
        count += num
    }
    return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    // ( week - 1 ) so we can jump 7 in 7 based on idx
    startIndex := 7 * (week - 1)
    endIndex := startIndex + 7
    sliced := birdsPerDay[startIndex:endIndex]
    return TotalBirdCount(sliced)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for idx, _ := range birdsPerDay {
        if idx % 2 == 0 {
            birdsPerDay[idx] += 1
        }
    }
    return birdsPerDay
}
