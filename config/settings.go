package config

import (
   "github.com/BurntSushi/toml"
   "io/ioutil"
   "fmt"
)
const LOCAL_QUEUE = "local"

type storageModeConf struct{
   ServerPort string
}

type queueModeConf struct{
   Capacity int
   ServerPort string
}

type storageConf struct{
   Mode string
   Local storageModeConf
   Remote storageModeConf
}

type queueConf struct{
   Mode string
   Local queueModeConf
   Remote queueModeConf
}

type schedulerConf struct{
   WorkerPool int
   SeedUrls []string
   CrawlExtDomains bool
}

type Config struct {
   Storage      storageConf
   Queue        queueConf
   Scheduler    schedulerConf
}

var Conf Config

func (c *Config) LoadConfig(path string) (error){

   b, err := ioutil.ReadFile(path)

   if err != nil {
      return err
   }

   if _, err := toml.Decode(string(b), c); err != nil {
      return err
   }

   if c.Queue.Mode == LOCAL_QUEUE && c.Queue.Local.Capacity <= 0 {
      return fmt.Errorf("Local Queue Capacity Empty")
   }

   if len(c.Scheduler.SeedUrls) == 0 {
      return fmt.Errorf("No Seed Urls")
   }

   if c.Scheduler.WorkerPool < 1 {
      return fmt.Errorf("Invalid Number of Worker")
   }

   return nil
}
