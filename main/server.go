package main

import (
	"fmt"
	"github.com/alex-zz/remoteLogParserDraft/lib/search/adapter/ssh"
	"github.com/alex-zz/remoteLogParserDraft/lib/search/adapter/pool"
	"time"
	"github.com/alex-zz/remoteLogParserDraft/lib/config"
	"github.com/vjeantet/jodaTime"
)

func main() {

	//testConfig()
	//testPool()
	//testTimeout()
	testDate()

	//c := ssh.Connection{}
	//connection, _ := ssh_connection.NewSshConnection("192.168.42.42", 22, "vagrant", "vagrant")
	//result, _  := connection.RunCommand("ls -la")
	//fmt.Println(result)
}

func testDate() {
	date := jodaTime.Format("YYYY.MM.dd", time.Now())
	fmt.Println(date)

	dateTime, _ := jodaTime.Parse("dd/MMMM/yyyy:HH:mm:ss", "30/August/2015:21:44:25")
	fmt.Println(dateTime.String())
}

func testConfig() {
	c := config.Config{}
	c.Load()

	settings := c.ConnectionList[0].Settings

	fmt.Print(settings)
}

func testTimeout() {

	c := make(chan int, 10)

	c <- 1
	c <- 2

	fmt.Println(len(c))

	ticker := time.NewTicker(2 * time.Second)

	fmt.Println("Pool test")

	for range ticker.C {
		fmt.Println("Tick")

		var s []int

		for len(c) > 0 {
			item := <- c
			fmt.Println("Iter")
			s = append(s, item)
			fmt.Println(len(c))
		}

		/*for item := range c {
			fmt.Println("Iter")
			s = append(s, item)
			fmt.Println(len(c))
		}*/

		for _, item := range s {
			c <- item
		}
	}

	fmt.Println("Pool test 2")
}

func testPool() {
	fmt.Println("Pool test")

	c := config.Config{}
	c.Load()

	settings := c.ConnectionList[0].Settings

	factory := ssh.Factory{}
	factory.Settings = &settings

	p, _ := pool.New(2, 1, time.Duration(2 * time.Second), time.Duration(10 * time.Second), &factory)

	item, _ := p.Get()
	(*item.GetAdapter()).Find()
	item.Release()

	/*go func() {
		item, _ := p.Get()
		(*item.GetAdapter()).Find()
		item.Release()
	}()

	go func() {
		item, _ := p.Get()
		(*item.GetAdapter()).Find()
		item.Release()
	}()

	go func() {
		item, _ := p.Get()
		(*item.GetAdapter()).Find()
		item.Release()
	}()*/

	time.Sleep(20)
}