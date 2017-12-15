package main

import (
	"github.com/tokwii/crawl/scheduler"
	"github.com/tokwii/crawl/config"
	"fmt"
	"time"
	"encoding/xml"
	"io/ioutil"
	"os"
)

const CONFIG_FILE  = "config/settings.toml"

func main()  {

	err := config.Conf.LoadConfig(CONFIG_FILE)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)

	}

	startTime := time.Now()
	s := scheduler.InitSchedule(config.Conf.Scheduler.WorkerPool,
					config.Conf.Scheduler.SeedUrls)
	s.Schedule()
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")

	xmlBytes, err := xml.MarshalIndent(s.CStorage.CreateSiteMap() ,"  ", "    ")

	if err != nil {
		fmt.Errorf("error: %v\n", err)
	}

	xmlstring := []byte(xml.Header + string(xmlBytes))
	//os.Stdout.Write(xmlstring)
	ioutil.WriteFile("sitemap.xml", xmlstring, 0666)
}

