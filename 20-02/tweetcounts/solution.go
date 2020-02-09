package main

import "fmt"

var delta map[string]int = map[string]int{
	"minute": 60,
	"hour":   60 * 60,
	"day":    60 * 60 * 24,
}

type tweet struct {
	id   int
	name string
	time int
}

type TweetCounts struct {
	tweets  []*tweet
	counter int
}

func Constructor() TweetCounts {
	return TweetCounts{
		tweets:  []*tweet{},
		counter: 0,
	}
}

func (this *TweetCounts) RecordTweet(tweetName string, time int) {
	this.counter++
	t := &tweet{
		id:   this.counter,
		name: tweetName,
		time: time,
	}
	this.tweets = append(this.tweets, t)
}

func (this *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {
	del := delta[freq]
	segments := (endTime + 1 - startTime) / del
	if (endTime+1-startTime)%del != 0 {
		segments++
	}
	fmt.Println("segments", segments)
	fqs := make([]int, segments)

	for i := 0; i < len(this.tweets); i++ {
		if this.tweets[i].name != tweetName {
			continue
		}
		if this.tweets[i].time < startTime || this.tweets[i].time > endTime {
			continue
		}
		seg := (this.tweets[i].time - startTime) % del
		fqs[seg]++
	}
	return fqs
}

/**
 * Your TweetCounts object will be instantiated and called as such:
 * obj := Constructor();
 * obj.RecordTweet(tweetName,time);
 * param_2 := obj.GetTweetCountsPerFrequency(freq,tweetName,startTime,endTime);
 */
