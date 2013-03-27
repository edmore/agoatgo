package main

import(
	"github.com/garyburd/redigo/redis"
        "fmt"
        "strings"
        "log"
	"os/exec"
        "os"
)

func getKey(v string, key string) string {
	out := []string{"venue:",v, key}
	return strings.Join(out, "")
}

func getPath(p string) string {
	path, err := exec.LookPath(p)
	checkError(err)
	return path
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}

func main() {
	app_root := "/Users/edmoremoyo/Sites/camera_dashboard_v2"
        c, err := redis.Dial("tcp", ":6379")
        defer c.Close()
	checkError(err)

	venue_list, _ := redis.Strings(c.Do("LRANGE", "venues", 0, -1))
	messages := make(chan string)

	for _, v := range venue_list {
		venue_name, _ := redis.String(c.Do("GET", getKey(v, ":venue_name")))
		cam_url, _ := redis.String(c.Do("GET", getKey(v, ":cam_url")))
		cam_user, _ := redis.String(c.Do("GET", getKey(v, ":cam_user")))
		cam_password, _ := redis.String(c.Do("GET", getKey(v, ":cam_password")))

		ffmpeg := getPath("ffmpeg")
		openRTSP := getPath("openRTSP")

                login_cridentials := ""
		if (cam_user != ""){
			login_cridentials = "-u "+cam_user+" "+cam_password
		}

		go func(){
			fmt.Println("Processing feed for", venue_name, "...")
			dir := app_root+"/public/feeds/"+venue_name
			os.MkdirAll(dir, 0755)

			feed_cmd := openRTSP+` `+login_cridentials+ ` -F `+venue_name+` -d 10 -b 300000 `+cam_url+` \
                                            && `+ffmpeg+` -i `+venue_name+`video-H264-1 -r 1 -s 320x180 -ss 5 -vframes 1 -f image2 `+app_root+`/public/feeds/`+venue_name+`/`+venue_name+`.jpeg\
                                            && rm -f `+venue_name+`video-H264-1`

			cmd := exec.Command("bash", "-c", feed_cmd)
                        //start command
			err = cmd.Start()
			checkError(err)
                        fmt.Println("Do some trivial stuff here ...")
			cmd.Wait()

			messages <- venue_name + " feed processed"
		}()
		fmt.Println(<-messages)
	}
}
