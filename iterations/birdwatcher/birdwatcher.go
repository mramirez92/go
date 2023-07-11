package birdwatcher

// sum total of bird count array
func TotalBirdCount(birdsPerDay []int) int{
    total := 0
    for i:=0; i < len(birdsPerDay); i++{
        total += birdsPerDay[i]
    }
    return total
}

// BirdsPerWeek returns total of bird count by week
func BirdsPerWeek(birdsPerDay []int, week int) int{
    end := week * 7
    start := end - 7
    week_slice := birdsPerDay[start:end]
    return TotalBirdCount(week_slice)
}

func FixBirdCounting(birdsPerDay []int) []int{
    for i:= 0; i <len(birdsPerDay); i++{
        if i % 2 ==0 {
            birdsPerDay[i] += 1
        }
    }
    return birdsPerDay
}
