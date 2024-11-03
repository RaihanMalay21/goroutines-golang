package belajar_golang_goroutines

import (
	"fmt"
	"time"
	"testing"
	"strconv"
)

func TestCreateChannel(t *testing.T) {
	Channel := make(chan string)
	defer close(Channel)

	go func() {
		time.Sleep(2 * time.Second)
		Channel <- "Raihan Malay"
		fmt.Println("Eksekusi program selesai")
	}()
	
	data := <- Channel
	fmt.Println(data)
	
	time.Sleep(5 *time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 *time.Second)
	channel <- "Raihan Malay Ganteng dan kaya"
}

func TestChannelAsParameter(t *testing.T) {
	Channel := make(chan string)
	defer close(Channel)

	go GiveMeResponse(Channel)

	data := <- Channel
	fmt.Println(data)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Raihan Malay ganteng, kaya, dan masuk surga"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestChannelInAndOut(t *testing.T) {
	Channel := make(chan string)
	defer close(Channel)

	go OnlyIn(Channel)
	go OnlyOut(Channel)

	time.Sleep(3 *time.Second)
}


func TestBufferedChannel(t *testing.T) {
	Channel := make(chan string, 3)
	defer close(Channel)

	go func() {
		Channel <- "Raihan"
		Channel <- "Malay"
		Channel <- "Ganteng"
	}()

	go func() {
		fmt.Println(<-Channel)
		fmt.Println(<-Channel)
		fmt.Println(<-Channel)
	}()

	time.Sleep(3 *time.Second)
}

func TestRangeChannel(t *testing.T) {
	Channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			Channel <- "Perulangan Ke " + strconv.Itoa(i)
		}
		close(Channel)
	}()

	for data := range Channel {
		fmt.Println(data)
	}

	fmt.Println("selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select{
		case data := <-channel1:
			fmt.Println(data)
			counter++
		case data := <-channel2:
			fmt.Println(data)
			counter++
		}
		if counter == 2 {
			break
		}
	}

}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select{
		case data := <-channel1:
			fmt.Println(data)
			counter++
		case data := <-channel2:
			fmt.Println(data)
			counter++
		default:
			fmt.Println("menunggu data")
		}

		if counter == 2 {
			break
		}
	}
}