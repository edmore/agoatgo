package main

import(
	"github.com/garyburd/redigo/redis"
        "fmt"
        "strings"
        "os"
        "log"
	"os/exec"
)

func get_key(v string, key string) string {
	out := []string{"venue:",v, key}
	return strings.Join(out, "")
}

func get_path(p string) string {
	path, err := exec.LookPath("ffmpeg")
	if err != nil {
		log.Fatal("LookPath: ", err)
	}
	return path
}

func main() {
	app_root := "/Users/etmoyo/Sites/camera_dashboard"
	fmt.Println(app_root)

        c, err := redis.Dial("tcp", ":6379")
        defer c.Close()

	if err != nil {
		fmt.Println(err)
	}

	venue_list, _ := redis.Strings(c.Do("LRANGE", "venues", 0, -1))
	messages := make(chan string)

	for _, v := range venue_list {
		venue_name, _ := redis.String(c.Do("GET", get_key(v, ":venue_name")))
		cam_url, _ := redis.String(c.Do("GET", get_key(v, ":cam_url")))
		cam_user, _ := redis.String(c.Do("GET", get_key(v, ":cam_user")))
		cam_password, _ := redis.String(c.Do("GET", get_key(v, ":cam_password")))

                login_cridentials := ""
		if (cam_user != ""){
			login_cridentials = strings.Join([]string{"-u ", cam_user, " ", cam_password}, "")
		}

		ffmpeg_path := get_path("ffmpeg")

		go func(){
			fmt.Println("Processing feed for", venue_name, "...", cam_url, login_cridentials)
			cmd := strings.Join([]string{app_root, "/public/feeds/", venue_name}, "")
			os.MkdirAll(cmd, 0755)

			fmt.Println(ffmpeg_path)
                        path := get_path("ffmpeg")
			out, err := exec.Command(path, "-version", ">/dev/null").Output()
			if err != nil {
				log.Fatal(err)
			}
			os.Stdout.Write(out)
			messages <- "done"
		}()
		fmt.Println(<-messages)
	}
}
