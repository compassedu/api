package main

import (
	"fmt"
	"os"
	"time"

	"api.compass.education/services"
)

func main() {
	s := os.Getenv("COMPASS_SCHOOLID")
	c, u, err := services.Login(os.Getenv("COMPASS_USERNAME"), os.Getenv("COMPASS_PASSWORD"), s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Cookies: ", c)
	fmt.Println("UserId: ", u)
	t := time.Now()
	events, err := services.GetCalendarEventsByUser(c, s, u, t, t)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(events)
	staff, err := services.GetAllStaff(c, s, u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(staff)
	locations, err := services.GetAllLocations(c, s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(locations)
}
