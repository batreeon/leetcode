/*
 * @lc app=leetcode.cn id=2043 lang=golang
 *
 * [2043] 简易银行系统
 */

// @lc code=start
type Bank struct {
	n int
	balance []int64
}


func Constructor(balance []int64) Bank {
	return Bank{
		n: len(balance),
		balance: balance,
	}
}


func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if !this.isAccountLegal(account1) || !this.isAccountLegal(account2) {
		return false
	}
	if this.balance[account1-1] < money {
		return false
	}
	this.balance[account1-1] -= money
	this.balance[account2-1] += money
	return true
}

func (this *Bank) isAccountLegal(account int) bool {
	n := this.n 
	if account >= 1 && account <= n {
		return true 
	}
	return false
}

func (this *Bank) Deposit(account int, money int64) bool {
	if !this.isAccountLegal(account) {
		return false
	}
	this.balance[account-1] += money
	return true
}


func (this *Bank) Withdraw(account int, money int64) bool {
	if !this.isAccountLegal(account) {
		return false
	}
	if this.balance[account-1] < money {
		return false
	}
	this.balance[account-1] -= money
	return true
}


/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */
// @lc code=end

