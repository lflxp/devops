package db

import (
	"github.com/astaxie/beego"
	bolt "github.com/coreos/bbolt"
)

var Db *bolt.DB

func init() {
	var err error
	Db, err = bolt.Open(beego.AppConfig.String("dbname"), 0600, nil)
	if err != nil {
		panic(err)
	}
	beego.Informational("初始化数据库......")

	// go func() {
	// 	// Grab the initial stats.
	// 	prev := db.Stats()

	// 	for {
	// 		// Wait for 10s.
	// 		time.Sleep(10 * time.Second)

	// 		// Grab the current stats and diff them.
	// 		stats := db.Stats()
	// 		diff := stats.Sub(&prev)

	// 		// Encode stats to JSON and print to STDERR.
	// 		json.NewEncoder(os.Stderr).Encode(diff)

	// 		// Save stats for the next loop.
	// 		prev = stats
	// 	}
	// }()
}

func Stats() {
	Db.Info()
}
