//go:build !solution

package hotelbusiness

type Guest struct {
	CheckInDate  int
	CheckOutDate int
}

type Load struct {
	StartDate  int
	GuestCount int
}

func ComputeLoad(guests []Guest) []Load {
	var dateMin, dateMax int
	var outpLoads []Load
	loads := make(map[int]int)
	if len(guests) > 0 {
		dateMin = guests[0].CheckInDate
		dateMax = guests[0].CheckOutDate
	} else {
		return []Load{}
	}
	for i := 0; i < len(guests); i++ {
		if guests[i].CheckInDate < dateMin {
			dateMin = guests[i].CheckInDate
		}
		if guests[i].CheckOutDate > dateMax {
			dateMax = guests[i].CheckOutDate
		}
		for date := guests[i].CheckInDate; date < guests[i].CheckOutDate; date++ {
			loads[date]++
		}
	}
	loads[dateMax] = 0
	outpLoads = append(outpLoads, Load{dateMin, loads[dateMin]})
	if len(loads) > 1 {
		for date := dateMin + 1; date <= dateMax; date++ {
			if loads[date] == loads[date-1] {
				continue
			} else {
				outpLoads = append(outpLoads, Load{date, loads[date]})
			}
		}
	}
	return outpLoads
}
