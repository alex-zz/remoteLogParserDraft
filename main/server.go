package main

import (
	"fmt"
	"github.com/alex-zz/remoteLogParserDraft/lib/config"
	"github.com/vjeantet/jodaTime"
	"time"
	"github.com/alex-zz/remoteLogParserDraft/lib/search/adapter/ssh"
	"errors"
	"github.com/alex-zz/remoteLogParserDraft/lib/search/adapter/pool"
)

func main() {
	//testConfig()
	//testPool()
	//testDate()

	c, _ := config.Load()
	initPoolList(c)
}

func initPoolList(c *config.Config) map[string]map[string]*pool.Pool {

	var poolList map[string]map[string]*pool.Pool

	for _, project := range c.Projects {
		for _, env := range project.Environments {
			name := env.Settings.Connection
			settings := getConnectionSettings(name, c)

			switch settings.Adapter {
			case "{{adapter.ssh}}":
				factory := ssh.Factory{}
				factory.ConnectionConfig = settings

				poolConfig := pool.Config{
					Cap: env.Settings.ConnectionPoolCapacity,
					InitCap: env.Settings.ConnectionPoolInitCapacity,
					Lifetime: time.Second * 30,
					Timeout: time.Second * 10,
					Factory: &factory,
				}
				p, _ := pool.New(poolConfig)
				poolList[project.Name][env.Name] = p

			default:
				errors.New("incorrect adapter")
			}
		}
	}

	return poolList
}

func getConnectionSettings(name string, c *config.Config) *config.Connection {

	var settings *config.Connection

	for _, connection := range c.Connections {
		if connection.Name == name {
			settings = &connection
		}
	}

	return settings
}


func testDate() {
	date := jodaTime.Format("YYYY.MM.dd", time.Now())
	fmt.Println(date)

	dateTime, _ := jodaTime.Parse("dd/MMMM/yyyy:HH:mm:ss", "30/August/2015:21:44:25")
	fmt.Println(dateTime.String())
}

func testConfig() {
	config.Load()

	//fmt.Print(err)
	//fmt.Print(c)
	//c.GetConnectionsConfiguration()

	//fmt.Print(c)
}

func testPool() {
	fmt.Println("Pool test")

	/*c := config.Config{}
	c.Load()

	settings := c.ConnectionList[0].Settings

	factory := ssh.Factory{}
	factory.Settings = &settings

	p, _ := pool.New(2, 1, time.Duration(2 * time.Second), time.Duration(10 * time.Second), &factory)

	item, _ := p.Get()
	(*item.GetAdapter()).Find()
	item.Release()*/

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
