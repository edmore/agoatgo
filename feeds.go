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
	app_root := "/usr/local/camera_dashboard"
	fmt.Println(app_root)

        c, err := redis.Dial("tcp", ":6379")
        defer c.Close()

	if err != nil {
		fmt.Println(err)
	}

	venue_list, _ := redis.Strings(c.Do("LRANGE", "venues", 0, -1))
	messages := make(chan string)

	for _, v := range venue_list {
		venue_name, _ := redis.String(c.Do("GET", get_venue_key(v, ":venue_name")))
		cam_url, _ := redis.String(c.Do("GET", get_venue_key(v, ":cam_url")))
		cam_user, _ := redis.String(c.Do("GET", get_venue_key(v, ":cam_user")))
		cam_password, _ := redis.String(c.Do("GET", get_venue_key(v, ":cam_password")))

                login_cridentials := ""
		if (cam_user != ""){
			login_cridentials = strings.Join([]string{"-u ", cam_password}, "")
		}
		fmt.Println(login_cridentials, "", cam_url)
		go func(){
			fmt.Println("Running feed for", venue_name, "...")
			messages <- "done"
		}()
		fmt.Println(<-messages)
	}
}
