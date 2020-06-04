package account

import fmt "fmt"
import time "time"

type Account struct { 
money int
writable int
}

func NewAccount(i int,flag int) *Account {
 return & Account {i , flag } 
 }

func (v *Account) SendAndRevoke(c Type__st_Account) {
 v . unfreeze ( ) 
 c.Send(v, v) 
 time . Sleep ( 500 * time . Microsecond ) 
 v . freeze ( ) 
 fmt . Println ( "Revoked" ) 
 }

func (v *Account) ReceiveAndWithdraw(c Type__st_Account) {
 B := c.Receive(v) 
 fmt . Println ( "value of B.money before withdrawing:" , B . query ( ) ) 
 B . withdraw ( 42 ) 
 fmt . Println ( "value of B.money after withdrawing 42:" , B . query ( ) ) 
 time . Sleep ( 1000 * time . Microsecond ) 
 B . withdraw ( 5 ) 
 fmt . Println ( "value of B.money after revocation:" , B . query ( ) ) 
 }

func (v *Account) query() int {
 return v . money 
 }

func (v *Account) withdraw(x int) {
 if v . writable > 0 && v . money >= x {
 v . money -= x 
 } 
 }

func (v *Account) unfreeze() {
 v . writable = 1 
 }

func (v *Account) freeze() {
 v . writable = 0 
 }


//import "fmt"

type type__st_Account struct {
	rs      int
	channel (chan *Account)
	users   []interface{}
}

type Type__st_Account interface {
	Receive(interface{}) *Account
	Send(*Account, interface{})
	Join(interface{}, interface{})
}

func (c *type__st_Account) Receive(ref interface{}) *Account {
	valid := false
	//fmt.Printf("[recv] ref= %p \n", ref)
	for _, user := range c.users {
		if user == ref {
			valid = true
		}
	}
	if c.rs <= 1 && valid { //receive from a send only capchan
		ret, _ := <-c.channel
		return ret
	} else {
		panic("Cannot receive: not a user of the channel")
	}
}

func (c *type__st_Account) Send(i *Account, ref interface{}) {
	valid := false
	//fmt.Printf("[send] ref= %p \n", ref)
	for _, user := range c.users {
		if user == ref {
			valid = true
		}
	}
	if c.rs >= 1 && valid {
		c.channel <- i
	} else {
		panic("Cannot send: not a user of the channel")
	}
}

//join
func (c *type__st_Account) Join(newuser interface{}, olduser interface{}) {
	flag := false
	for _, user := range c.users {
		if user == olduser {
			c.users = append(c.users, newuser)
			//fmt.Printf("[join] newuser= %p \n", newuser)
			flag = true
			break
		}
	}
	if !flag {
		panic("Cannot join: not a user of the channel")
	}
}


func New__st_Account(rs int, users []interface{}) Type__st_Account {
	return &type__st_Account{rs, make(chan *Account), users}
}