package config

import (
   "github.com/BurntSushi/toml"
   "io/ioutil"
)

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

func(c *Config) Load(path string) (error){

   b, err := ioutil.ReadFile(path)

   if err != nil {
      return err
   }

   if _, err := toml.Decode(string(b), c); err != nil {
      return err
   }
   return nil
}

/*var Conf Config
func init(){
   Conf.Queue.Mode = "local"
   Conf.Queue.Local.Capacity = 1000
   Conf.Storage.Mode = "local"
   Conf.Storage.Mode = "local"
}*/