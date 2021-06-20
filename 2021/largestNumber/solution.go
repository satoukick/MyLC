package main

import (
	"fmt"
	"strconv"
	"strings"
)

type moment struct {
	h int
	m int
}

func parseTime(input string) *moment {
	mome := &moment{}
	tm := strings.Split(input, ":")
	mome.h, _ = strconv.Atoi(tm[0])
	mome.m, _ = strconv.Atoi(tm[1])
	return mome
}

func numberOfRounds(startTime string, finishTime string) int {
	startMoment := parseTime(startTime)
	finishMoment := parseTime(finishTime)
	if startMoment.h > finishMoment.h || (startMoment.h == finishMoment.h && startMoment.m > finishMoment.m) {
		finishMoment.h += 24
	}

	if startMoment.m%15 != 0 {
		startMoment.m += (15 - startMoment.m%15)
		if startMoment.m == 60 {
			startMoment.h++
			startMoment.m = 0
		}
	}

	if finishMoment.m%15 != 0 {
		finishMoment.m -= finishMoment.m % 15
	}

	fmt.Println(*startMoment, *finishMoment)

	minutes := (finishMoment.h-startMoment.h)*60 + (finishMoment.m - startMoment.m)
	return minutes / 15
}
