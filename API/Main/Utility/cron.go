package Utility

import "github.com/robfig/cron/v3"

var Entries = make(map[string]cron.EntryID)
var CronI *cron.Cron = cron.New()
