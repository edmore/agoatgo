package main

import(
	"fmt"
        "github.com/garyburd/redigo/redis"
        "strings"
)

func get_venue_key(v string, attribute string) string {
	out := []string{"venue:",v, attribute}
	return strings.Join(out, "")
}

func main() {
	//app_root := "/usr/local/camera_dashboard"

        c, err := redis.Dial("tcp", ":6379")
        defer c.Close()

	if err != nil {
		fmt.Println(err)
	}

	venue_list, err := redis.Strings(c.Do("LRANGE", "venues", 0, -1))

	if err != nil {
		fmt.Println(err)
	}

	messages := make(chan string)
	for _, v := range venue_list {
		venue_name := get_venue_key(v, ":venue_name")

		go func(){
			fmt.Println("Running feed for", venue_name)
			messages <- "done"
		}()
		fmt.Println(<-messages)
	}
}
