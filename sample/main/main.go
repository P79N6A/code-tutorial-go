/*
	https://gobyexample.com/
	
	bsample.Juniordata():  slice, array, map, range
		range: use for travel
		at  junior-build-in-data-struct.go
	bsample.Function: function call, multi return value function, Variadic Functions
		at function.go
	bsample.Seniorfunc: anonymous functions,Closures,Recursion
		at senior-function.go
	bsample.PointAndStruct(): pointer and struct.
			pointer like c/c++, but pointer can use dot like reference more
		at pointer-struct.go
	==========OO:object-oriented==========
	bsample.StructMethod():  method in struct, OO:Package/Encapsulation
		at struct-method.go
	bsample.Interface(): interface, OO:Inheritance and polymorphism
			interface use with struct, need implementation at struct
		at interface.go
	bsample.ErrorFunc(): error, like java~
		at error.go
	=========concurrent goroutines and channel(pipes) ====================
	bsample.RoutineFun(): goroutine is a lightweight thread of execution
		at goroutines.go
	bsample.ChannelSample(): channel like pipe, connect goroutines
		at channel.go
	bsample.ChannelBuf(): send multi msg to channel
	bsample.ChannelSyn():  use channel to synchronize execution across goroutines
	bsample.ChannelDir(): using channels as function parameters
	
	======= at channel-operation.go=======
	bsample.SelectChannel()
	bsample.TimeoutChannel()
	bsample.NonBlockChannel()
	bsample.CloseChannel()
	bsample.RangeOverChannel()
	
	=======at timer-ticker.go========
	bsample.Timers()   after target time send msg
	bsample.Tickers()   in fix interval send msg
	
	=====at statemanngement-atomic_counter-mutex.go====
	//bsample.WorkerPools()
	//bsample.AtomicCounter()
	//bsample.MutexState()
	=====at statemanagement-stateful-goroutines.go=====
	//bsample.StatefulGoroutines()
	
	===at defer-final.go====
	bsample.Defer()  like final
	
	
*/
package main

import (
	"fmt"
	"sample/basicsample"
//	"packagesample"
)

func main() {
	fmt.Println("main")
	bsample.Juniordata()
	// bsample.Function()
	//bsample.Seniorfunc()
	//bsample.PointAndStruct()
	//bsample.StructMethod()
	//bsample.Interface()
	// bsample.ErrorFunc()
	//bsample.RoutineFun()
	/*
	bsample.ChannelSample()
	bsample.ChannelBuf()
	bsample.ChannelSyn()
	bsample.ChannelDir()
	*/
	bsample.SelectChannel()
	/*
	bsample.SelectChannel()
	bsample.TimeoutChannel()
	bsample.NonBlockChannel()
	bsample.CloseChannel()
	bsample.RangeOverChannel()
	*/
	//bsample.Timers()
	//bsample.Tickers()
	
	//bsample.WorkerPools()
	//bsample.AtomicCounter()
	//bsample.MutexState()
	//bsample.StatefulGoroutines()
	
	
	//bsample.SortByFunc()
	//bsample.Sorting()
	
	//bsample.Defer()
	
	//bsample.CollectionFunc()
	
	//psample.StringFunc()
	//psample.StringFormat()
	
	//psample.RegEx()
	
	//psample.Json()
	//psample.TimeUse()
	//psample.Epoch()
	//psample.TimeFormat()
	//psample.NumParse()
	//psample.UrlParse()
	
//	psample.Sha1()
//	psample.Base64()
	
}

