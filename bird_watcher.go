package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    var s int
    
	for i:=0; i<len(birdsPerDay); i++ {
        s+=birdsPerDay[i]
    }
return s
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	w := week-1
    var s int
    for i := w*7; i<week*7; i++ {
        s += birdsPerDay[i]
    
    }
return s
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i:= 0; i<len(birdsPerDay); i+=2 {
        birdsPerDay[i] = birdsPerDay[i]+1
    }
return birdsPerDay
}
