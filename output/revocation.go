package main

import time "time"
import account "github.com/lzy0505/gocap/output/account"
import capchan "github.com/lzy0505/gocap/output/capchan"

func main() {
 B := account . NewAccount ( 100 , 0 ) 
 C := account . NewAccount ( 0 , 0 ) 
 c := account.New__st_Account(1, [](interface{}){capchan.TopLevel}) 
 c.Join(B, capchan.TopLevel) 
 c.Join(C, capchan.TopLevel) 
 go B . SendAndRevoke ( c ) 
 go C . ReceiveAndWithdraw ( c ) 
 time . Sleep ( 3000 * time . Microsecond ) 
 }

