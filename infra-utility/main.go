package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"fmt"
	"context"
	"time"
	"bufio"
	"io"
	"encoding/json"
	"errors"
	
    "github.com/tidwall/gjson" // more info https://github.com/tidwall/gjson
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
)


var dockerRegistryUserID = ""

type ErrorDetail struct {
	Message string `json:"message"`
}

type ErrorLine struct {
	Error       string      `json:"error"`
	ErrorDetail ErrorDetail `json:"errorDetail"`
}

func imageBuild(dockerClient *client.Client, path string ,tagName string, buildArgs map[string]*string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()

	tar, err := archive.TarWithOptions(path, &archive.TarOptions{})
	if err != nil {
		return err
	}

	opts := types.ImageBuildOptions{
		Dockerfile: "Dockerfile",
		Tags: []string{tagName},
   		NoCache: true,
   		Remove: true,
   		BuildArgs: buildArgs,
	}
	res, err := dockerClient.ImageBuild(ctx, tar, opts)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	err = print(res.Body)
	if err != nil {
		return err
	}

	return nil
}

func print(rd io.Reader) error {
	var lastLine string

	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {
		lastLine = scanner.Text()
		fmt.Println(scanner.Text())
	}

	errLine := &ErrorLine{}
	json.Unmarshal([]byte(lastLine), errLine)
	if errLine.Error != "" {
		return errors.New(errLine.Error)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func infrautility(utility_name string){
	// read `versions.json` file
	path := "./"+utility_name+"/versions.json"
	content, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal("Error while opening file: ", err)
	}
    
    alpine_version := gjson.Get(string(content),"utilites.alpine_version")
    
    log.Printf("%s\n",alpine_version)
    fmt.Println(alpine_version)
 
    buildArgs := map[string]*string{
			    "ALPINE_VERSION": &alpine_version.Str,
			  	}
    
    fmt.Println(buildArgs)
    
    mydir, err := os.Getwd()
    if err != nil {
        fmt.Println(err)
    }
    dockerPath := mydir+"/"+string(utility_name)
    fmt.Println(dockerPath)

    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
    
    version := "0.0.1"
	tagName := string(utility_name)+":"+version
	
	err = imageBuild(cli, dockerPath, tagName, buildArgs)
	
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func main(){

	utility_name := os.Args[1]

	if strings.Compare(string(utility_name),  "infrautility") == 0 {
 			infrautility(string(utility_name))
	}
}