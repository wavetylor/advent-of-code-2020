package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func main() {
	start, _ := strconv.Atoi(os.Args[1])
	end, _ := strconv.Atoi(os.Args[2])

	body := `package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("../in.txt")

	lines := strings.Split(string(dat), "\n")

	for _, l := range lines {
		fmt.Println(l)
	}
}
`

	for i := start; i < end+1; i++ {

		dayDir := fmt.Sprintf("../%d", i)

		makeDirs(dayDir)
		createFiles(body, dayDir)
	}
}

func makeDirs(dayDir string) {

	os.Mkdir(dayDir, 0755)
	os.Mkdir(dayDir+"/challenge1", 0755)
	os.Mkdir(dayDir+"/challenge2", 0755)
}

func createFiles(body string, dayDir string) {
	b := []byte(body)
	ioutil.WriteFile(dayDir+"/challenge1/main.go", b, 0755)
	ioutil.WriteFile(dayDir+"/challenge2/main.go", b, 0755)

	url := fmt.Sprintf("https://adventofcode.com/2020/day/%s/input", string(dayDir[3:]))
	req, _ := http.NewRequest("GET", url, nil)

	//TODO can we make this dynamic?
	session, _ := os.LookupEnv("ADVENT_SESSION")
	req.AddCookie(&http.Cookie{Name: "session", Value: session})

	client := &http.Client{}
	resp, _ := client.Do(req)

	defer resp.Body.Close()

	input, _ := ioutil.ReadAll(resp.Body)

	ioutil.WriteFile(dayDir+"/in.txt", input, 0755)

}
