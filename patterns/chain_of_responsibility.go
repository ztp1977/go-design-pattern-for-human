package patterns

import "github.com/k0kubun/pp"

type (
	ChainOfResponsibility struct{}

	iAccount interface {
		getNext() iAccount
		setNext(a iAccount)
		canPay(price int64) bool
		pay(price int64)
	}

	// 这里可以写共同的函数
	payer struct {
		balance     int64
		nextAccount iAccount
	}

	// 通过override模拟继承
	bank struct {
		*payer
	}

	paypal struct {
		*payer
	}

	bitcoin struct {
		*payer
	}
)

// TODO *payer, payer的区别
func (m *payer) pay(price int64) {
	if m.canPay(price) {
		m.balance -= price
		pp.Printf("pay -%s\n", price)
		return
	}
	nextAccount := m.getNext()
	if nextAccount == nil {
		pp.Println("None of accounts have enough balance")
		return
	}
	pp.Println("can not pay balance")
	nextAccount.pay(price)
}

func (m *payer) getNext() iAccount {
	return m.nextAccount
}

func (m *payer) setNext(a iAccount) {
	m.nextAccount = a
}

func (m *payer) setBalance(balance int64) {
	m.balance = balance
}

func (m *payer) canPay(price int64) bool {
	return m.balance >= price
}

func newBank(balance int64) *bank {
	return &bank{payer: &payer{balance: balance}}
}

func newPaypal(balance int64) *paypal {
	return &paypal{payer: &payer{balance: balance}}
}

func newBitcoin(balance int64) *bitcoin {
	return &bitcoin{payer: &payer{balance: balance}}
}

func (ChainOfResponsibility) Do() {

	desc.SetDesc("ChainOfResponsibility", "如果处理不了，交给下一个函数去解决", "setNext, getNext链接方法", "next嵌套方法的实现", "网站路由的函数是这个实现，每个函数留下一个onNext", "")
	desc.print()

	bank := newBank(100)
	paypal := newPaypal(150)
	bitcoin := newBitcoin(1000)

	// 主要是setNext这个方法， 把不同方法连起来
	bank.setNext(paypal)
	paypal.setNext(bitcoin)

	//bank.setNext(paypal)
	//paypal.setNext(bitcoin)

	bank.pay(256)

	pp.Println(bank)
}
