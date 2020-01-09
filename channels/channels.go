package channels

import (
	"fmt"
	"sync"
	"time"
)

/*
Channels: Used to sync data between goroutines

*/

var wg = sync.WaitGroup{}

//PlayChannels1 ...
func PlayChannels1() {
	ch := make(chan int)
	wg.Add(2)
	go func() {
		i := <-ch //Recieve data from channel
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		i := 42
		ch <- i // Send the number 42 to the channel
		i = 27
		wg.Done()
	}()
	wg.Wait()
}

//PlayChannels2 ...
func PlayChannels2() {
	ch := make(chan int)
	//Generating 5 senders and 5 recievers (trouble when not symmetrical, need buffers!)
	for j := 0; j < 5; j++ {
		wg.Add(2)
		go func() {
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}()
		go func() {
			ch <- 42 //This line of code pauses this routine on this line until is space in this channel (only one message per time as channel is unbuffered)
			wg.Done()
		}()
	}
	wg.Wait()
}

//PlayChannels3 ...
func PlayChannels3() {
	/*
		Here we have goroutines which bot recieve and send
		data on channel. We do not want that usually.
	*/
	ch := make(chan int)
	wg.Add(2)
	go func() {
		i := <-ch
		fmt.Println(i)
		ch <- 27
		wg.Done()
	}()
	go func() {
		ch <- 42
		fmt.Println(<-ch)
		wg.Done()
	}()

	wg.Wait()
}

//PlayChannels4 ...
func PlayChannels4() {
	ch := make(chan int)
	wg.Add(2)
	//Want to make this under a recieve only channel
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	//Want to make this under a send only channel
	go func(ch chan<- int) {
		ch <- 42
		wg.Done()
	}(ch)
	wg.Wait()
}

//PlayChannels5 buffered channels to avoid deadlock
func PlayChannels5() {
	ch := make(chan int, 50) //Channel with buffer of 50 integerers
	wg.Add(2)
	go func() {
		i := <-ch
		fmt.Println(i)
		i = <-ch
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		ch <- 42
		ch <- 27
		wg.Done()
	}()
	wg.Wait()

	/*
		Note: This is NOT the intended usecase for buffered channels.
		The intended usecase is when sender and reciever operates
		at different frequencies. Say i.e. that a sensor sends 100 samples
		once every hour for the datapoints from that hour and then the reciever
		must process it as fast as it is possible to do (Burst transmission)
	*/

	//How to actually solve the probnlem above before the note

	wg.Add(2)
	go func(ch <-chan int) {
		for {
			if i, ok := <-ch; ok { //ok=true if channel open
				fmt.Println(i)
			} else {
				break
			}
		}

		/*
			Note:
			We could use the for range construct as this breaks
			of automatically when detects channel closed.
		*/
	}(ch)
	go func(ch chan<- int) {
		ch <- 42
		//close(ch) //Causes panic because we afterward tries to pass value to closed channel
		/*
			Note:
			You CANNOT reopen a closed channel
			and you cant even detect it except from the
			application panicing. This is a limitation of the GOlang.
			Will need a recover function if there is possibility of
			this happening.
		*/
		ch <- 27
		close(ch)
		/*
			Note:
			What we are doing here is sending two values and then
			telling the channel we are done working with it
		*/
		wg.Done()
	}(ch)
}

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)

//PlayChannels6 buffered channels to avoid deadlock
func PlayChannels6() {
	go logger()
	//The defer below is to softly shut down the logger
	defer func() {
		close(logCh)
	}()
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}

	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 * time.Millisecond)
}

func logger() {
	for entry := range logCh {
		fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
	}
}

//Now, soft shutdown of logger with select statement
var logCh2 = make(chan logEntry, 50)
var doneCh = make(chan struct{}) //Struct{} requires zero memory allocation (this is unique to struct{})
//doneCh can't send any message through, except that a message was sent or recieved
//This is a SIGNAL ONLY channel.
//Common and good practice, saves memory!

//PlayChannels7 logger with selec statement shutdown
func PlayChannels7() {
	go loggerv2()
	logCh2 <- logEntry{time.Now(), logInfo, "App is starting"}

	logCh2 <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 * time.Millisecond)
	doneCh <- struct{}{} //Take note of this notation, struct{} is type and {} is empty initialization
}

func loggerv2() {
	for {
		select {
		case entry := <-logCh2:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)

		case <-doneCh:
			break
		default:
			//Avoid blocking when no message comes in.
		}
	}
}

/*
Restricting dataflow:
	Send only: chan <- int
	Recieve only: <-chan int

Buffered channels
	- Channels block sender side till reciever is available
	- Block reciever side till message is available
	- Can decouple sender and reciever with buffered channels
		- make(chan int, 50)
	- Use buffered channels when sender and reciever have assymetric loading

For..range loops with channels
	- Use to monitor channel and process messages as they arrive
	- Loop closes when channel is closed.

Select statements
	- Allows goroutine to monitor several channels at once
		- Blocks if all channels block
		- If multiple channels recieve value simulataneously, behaviour is undefined
			- Ordering does not resolve this!

*/
