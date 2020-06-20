package ratelimiter

import (
	"sync"
	"time"
)

// CreditsBalanceRateLimiter is a reconfigurable rate limiter based on leaky bucket algorithm, formulated in terms of a
// credits balance that is replenished every time CheckCredit() method is called (tick) by the amount proportional
// to the time elapsed since the last tick, up to max of creditsPerSecond. A call to CheckCredit() takes a cost
// of an item we want to pay with the balance. If the balance exceeds the cost of the item, the item is "purchased"
// and the balance reduced, indicated by returned value of true. Otherwise the balance is unchanged and return false.
//
// This can be used to limit a rate of messages emitted by a service by instantiating the Rate Limiter with the
// max number of messages a service is allowed to emit per second, and calling CheckCredit(1.0) for each message
// to determine if the message is within the rate limit.
//
// It can also be used to limit the rate of traffic in bytes, by setting creditsPerSecond to desired throughput
// as bytes/second, and calling CheckCredit() with the actual message size.
type CreditsBalanceRateLimiter struct {
	mu sync.Mutex

	creditsPerSecond float64
	balance float64
	maxBalance float64
	lastTick time.Time

	timeNow func() time.Time
}

// NewCreditsBalanceRateLimiter creates a new CreditsBalanceRateLimiter
func NewCreditsBalanceRateLimiter(creditsPerSecond, maxBalance float64) *CreditsBalanceRateLimiter {
	return &CreditsBalanceRateLimiter{
		creditsPerSecond: creditsPerSecond,
		balance:          maxBalance,
		maxBalance:       maxBalance,
		lastTick:         time.Now(),
		timeNow:          time.Now,
	}
}

// CheckCredit tries to reduce the current balance by itemCost provided that the current balance
// is not lest than itemCost.
func (rl *CreditsBalanceRateLimiter) CheckCredit(itemCost float64) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	// if we have enough credits to pay for current item, then reduce balance and allow
	if rl.balance > itemCost {
		rl.balance -= itemCost
		return true
	}
	// otherwise check if balance can be increased due to time elapsed, and try again
	rl.updateBalance()
	if rl.balance >= itemCost {
		rl.balance -= itemCost
		return true
	}
	return false
}

// updateBalance recalculates current balance based on time elapsed. Must be called while holding a lock.
func (rl *CreditsBalanceRateLimiter) updateBalance() {
	// calculate how much time passed since the last tick, and update current tick
	currentTime := rl.timeNow()
	elapsedTime := currentTime.Sub(rl.lastTick)
	rl.lastTick = currentTime
	// calculate how much credit have we accumulated since the last tick
	rl.balance += elapsedTime.Seconds() * rl.creditsPerSecond
	if rl.balance > rl.maxBalance {
		rl.balance = rl.maxBalance
	}
}

// Update changes the main parameters of the rate limiter in-place, while retaining
// the current accumulated balance (pro-rated to the new maxBalance value). Using this method
// instead of creating a new rate limiter helps to avoid thundering herd when sampling
// strategies are updated.
func (rl *CreditsBalanceRateLimiter) Update(creditsPersecond, maxBalance float64) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	rl.updateBalance() // get up to date balance
	rl.balance = rl.balance * maxBalance / rl.maxBalance
	rl.creditsPerSecond = creditsPersecond
	rl.maxBalance = maxBalance
}