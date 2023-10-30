package main



import (

	"fmt"
	"time"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
}

func main(){
	log.Info("Let's create new cron")
	cron_client := cron.New()
    cron_client.AddFunc("*/1 * * * *", func() { log.Info("[Job 1]Every minute Jobs Running \n") })
    
    
    fmt.Println("*********** Start cron with one scheduled job ***********")
	log.Info("Start cron")
	cron_client.Start()
	loggingCronEntries(cron_client.Entries())
	time.Sleep(2 * time.Minute)

    
    fmt.Println("*********** Funcs may also be added to a running Cron ***********")
	log.Info("Add new job to a running cron")
	entryID2, _ := cron_client.AddFunc("*/2 * * * *", func() { log.Info("[Job 2]Every two minutes job\n") })
	loggingCronEntries(cron_client.Entries())
	time.Sleep(5 * time.Minute)

	
	fmt.Println("*********** Remove Job2 and add new Job2 that run every 1 minute ***********")
	log.Info("Remove Job2 and add new Job2 with schedule run every minute")
	cron_client.Remove(entryID2)
	cron_client.AddFunc("*/1 * * * *", func() { log.Info("[Job 2]Every one minute job\n") })
	time.Sleep(5 * time.Minute)
    
}

func loggingCronEntries(cronEntries []cron.Entry){
		log.Infof("Cron Info: %+v\n", cronEntries)
}