package main

import (
	"sort"
)

type player struct {
	id    int
	score int
}

type Leaderboard struct {
	players []*player
	exist   map[int]bool
}

func Constructor() Leaderboard {
	return Leaderboard{
		players: []*player{},
		exist:   make(map[int]bool),
	}
}

func (this *Leaderboard) AddScore(playerId int, score int) {
	if !this.exist[playerId] {
		newPlayer := &player{
			id:    playerId,
			score: score,
		}
		this.exist[playerId] = true
		this.players = append(this.players, newPlayer)
		return
	}

	for _, p := range this.players {
		if p.id == playerId {
			p.score += score
		}
	}
}

func (this *Leaderboard) Top(K int) int {
	sort.Slice(this.players, func(i, j int) bool { return this.players[i].score > this.players[j].score })

	var top int
	for i := 0; i < K; i++ {
		top += this.players[i].score
	}
	return top
}

func (this *Leaderboard) Reset(playerId int) {
	if !this.exist[playerId] {
		return
	}

	delete(this.exist, playerId)
	for i := 0; i < len(this.players); i++ {
		if this.players[i].id == playerId {
			this.players = append(this.players[:i], this.players[i+1:]...)
			break
		}
	}
}

/**
 * Your Leaderboard object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddScore(playerId,score);
 * param_2 := obj.Top(K);
 * obj.Reset(playerId);
 */
