// Найдите и исправьте проблемы в коде ниже
// у парковки ограниченное кол-во мест для парковки - обязательное условие
// Программа должна отработать корректно и завершить работу без зависаний

package main

import (
	"fmt"
	"sync"
	"time"
)

type ParkingLot struct {
	slots int64
}

func (p *ParkingLot) Park(carID int64) {
	fmt.Printf("Машина (%d) паркуется \n", carID)
	time.Sleep(time.Second) // время стоянки
	fmt.Printf("Машина (%d) уехала с парковки\n", carID)
}

func main() {
	parking := &ParkingLot{slots: 3}

	var wg sync.WaitGroup

	carIDs := []int64{1, 2, 3, 4, 5, 6}

	for _, carID := range carIDs {
		wg.Add(1)
		go func(id int64) {
			defer wg.Done()

			parking.Park(id)
		}(carID)
	}

	wg.Wait()
}
