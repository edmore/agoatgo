package main

import(
	"github.com/garyburd/redigo/redis"
        "fmt"
        "strings"
        "log"
	"os/exec"
        "os"
        "io"
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
	app_root := "/Users/etmoyo/Sites/camera_dashboard"
	fmt.Println(app_root)

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
			fmt.Println("Processing feed for", venue_name, "...", cam_url, login_cridentials)
			dir := app_root+"/public/feeds/"+venue_name
			os.MkdirAll(dir, 0755)

			fmt.Println(ffmpeg)
             		fmt.Println(openRTSP)

			cmd := exec.Command(ffmpeg, "-version")
			// if cmd.StdoutPipe() not called output goes to /dev/null
			// stdout, err := cmd.StdoutPipe()
			//checkError(err)
		        stderr, err := cmd.StderrPipe()
		        checkError(err)
                        //start command
			err = cmd.Start()
			checkError(err)
			//go io.Copy(os.Stdout, stdout)
			go io.Copy(os.Stderr, stderr)
			cmd.Wait()

			messages <- "done"
		}()
		fmt.Println(<-messages)
	}
}
