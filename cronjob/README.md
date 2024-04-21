
Create CronJob in Go


$ cd cronjob

$ go mod tidy

$ go run main.go





```
INFO[2023-10-31T00:38:39+05:30] Let's create new cron
*********** Start cron with one scheduled job ***********
INFO[2023-10-31T00:38:39+05:30] Start cron
INFO[2023-10-31T00:38:39+05:30] Cron Info: [{ID:1 Schedule:0x1400012a080 Next:2023-10-31 00:39:00 +0530 IST Prev:0001-01-01 00:00:00 +0000 UTC WrappedJob:0x10260eb20 Job:0x10260eb20}]
INFO[2023-10-31T00:39:00+05:30] [Job 1]Every minute Jobs Running
INFO[2023-10-31T00:40:00+05:30] [Job 1]Every minute Jobs Running
*********** Funcs may also be added to a running Cron ***********
INFO[2023-10-31T00:40:39+05:30] Add new job to a running cron
INFO[2023-10-31T00:40:39+05:30] Cron Info: [{ID:1 Schedule:0x1400012a080 Next:2023-10-31 00:41:00 +0530 IST Prev:2023-10-31 00:40:00 +0530 IST WrappedJob:0x10260eb20 Job:0x10260eb20} {ID:2 Schedule:0x1400009e040 Next:2023-10-31 00:42:00 +0530 IST Prev:0001-01-01 00:00:00 +0000 UTC WrappedJob:0x10260eb90 Job:0x10260eb90}]
INFO[2023-10-31T00:41:00+05:30] [Job 1]Every minute Jobs Running
INFO[2023-10-31T00:42:00+05:30] [Job 1]Every minute Jobs Running
INFO[2023-10-31T00:42:00+05:30] [Job 2]Every two minutes job
INFO[2023-10-31T00:43:00+05:30] [Job 1]Every minute Jobs Running
INFO[2023-10-31T00:44:00+05:30] [Job 1]Every minute Jobs Running
INFO[2023-10-31T00:44:00+05:30] [Job 2]Every two minutes job
INFO[2023-10-31T00:45:00+05:30] [Job 1]Every minute Jobs Running
*********** Remove Job2 and add new Job2 that run every 1 minute ***********
INFO[2023-10-31T00:45:39+05:30] Remove Job2 and add new Job2 with schedule run every minute
INFO[2023-10-31T00:46:00+05:30] [Job 1]Every minute Jobs Running
INFO[2023-10-31T00:46:00+05:30] [Job 2]Every one minute job
INFO[2023-10-31T00:47:00+05:30] [Job 2]Every one minute job
INFO[2023-10-31T00:47:00+05:30] [Job 1]Every minute Jobs Running
INFO[2023-10-31T00:48:00+05:30] [Job 1]Every minute Jobs Running
INFO[2023-10-31T00:48:00+05:30] [Job 2]Every one minute job
INFO[2023-10-31T00:49:00+05:30] [Job 2]Every one minute job
INFO[2023-10-31T00:49:00+05:30] [Job 1]Every minute Jobs Running
INFO[2023-10-31T00:50:00+05:30] [Job 2]Every one minute job
INFO[2023-10-31T00:50:00+05:30] [Job 1]Every minute Jobs Running
```