package main

import (
	"github.com/VilimRoller/Basic-CRUD/utils"
)

func main() {
	redisClient := utils.GetDefaultRedisClient()

	utils.RegisterEndpoints(redisClient)
}
