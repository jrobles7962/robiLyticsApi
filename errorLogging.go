package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"os"
)

func errorToLog(logData string, err error) {
	f, err := os.OpenFile("/var/log/robyLytics.error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("ERROR: Cannot write to log file")
		panic(err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println(logData, err)
}

func getNumDevelopers() int {
	redisConn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		errorToLog("Cannot connect to Redis server", err)
	}
	numDevelopers, err := redis.Int(redisConn.Do("SCARD", "data:developers"))
	if err != nil {
		errorToLog("Cannot obtain the number of developers from data:developers SET", err)
	}
	return numDevelopers

}
