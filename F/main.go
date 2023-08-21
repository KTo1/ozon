package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type timeLine struct {
	start time.Time
	end   time.Time
}

func (tl *timeLine) init(in *string) error {
	times := strings.Replace(*in, "-", " ", 1)
	timesArray := strings.Fields(times)
	timeStart, err := time.Parse("02/Jan/2006:15:04:05", "02/Jan/2006:"+timesArray[0])
	if err != nil {
		return err
	}
	tl.start = timeStart
	timeEnd, err := time.Parse("02/Jan/2006:15:04:05", "02/Jan/2006:"+timesArray[1])
	if err != nil {
		return err
	}

	tl.end = timeEnd

	if tl.start.After(tl.end) {
		return errors.New("Start after end!")
	}

	return nil
}

func (tl *timeLine) hasIntersection(other *timeLine) bool {
	//if (a2 < a1 and b2 < a1) or (a2 > b1 and b2 > b1):
	//print('пустое множество')
	if (other.start.Before(tl.start) && other.end.Before(tl.start)) || (other.start.After(tl.end) && other.end.After(tl.end)) {
		return false
	}

	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount, timeCount int
	var times string

	fmt.Fscan(in, &testCount)

	for i := 0; i < testCount; i++ {
		timesArray := make([]timeLine, 0)
		yesOrNo := true
		fmt.Fscan(in, &timeCount)
		for j := 0; j < timeCount; j++ {
			fmt.Fscan(in, &times)

			if !yesOrNo {
				in = bufio.NewReader(os.Stdin)
				break
			}

			temp := &timeLine{}
			err := temp.init(&times)
			if err != nil {
				yesOrNo = false
			}

			for _, v := range timesArray {
				if v.hasIntersection(temp) {
					yesOrNo = false
				}
			}

			timesArray = append(timesArray, *temp)
		}

		if yesOrNo {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}

	}
}
