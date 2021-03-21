package main

type AuthenticationManager struct {
	tokens map[string]int
	ttl    int
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		tokens: make(map[string]int),
		ttl:    timeToLive,
	}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.tokens[tokenId] = currentTime
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	t, ok := this.tokens[tokenId]
	if !ok {
		return
	}
	if t+this.ttl <= currentTime {
		return
	}
	this.tokens[tokenId] = currentTime
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	count := 0
	for t, v := range this.tokens {
		if v+this.ttl > currentTime {
			count++
		}
	}
	return count
}

/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */
