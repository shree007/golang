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

func infrautility(utility_name string, docker_function string){

	log.Printf("%s\n", docker_function)
	// read `versions.json` file
	path := "./"+utility_name+"/versions.json"
	content, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal("Error while opening file: ", err)
	}
    
    alpine_version := gjson.Get(string(content),"utilites.alpine_version")
    log.Printf("%s\n alpine_version: ",alpine_version)
    python3_version := gjson.Get(string(content),"utilites.python3_version")
    log.Printf("%s\n python3_version: ",python3_version)
    aws_cli_version := gjson.Get(string(content),"utilites.aws_cli_version")
    log.Printf("%s\n aws_cli_version: ",aws_cli_version)
    kubectl_version := gjson.Get(string(content),"utilites.kubectl_version")
    log.Printf("%s\n kubectl_version: ",kubectl_version)
    terraform_version := gjson.Get(string(content),"utilites.terraform_version")
    log.Printf("%s\n terraform_version: ",terraform_version)
    terragrunt_version := gjson.Get(string(content),"utilites.terragrunt_version")
    log.Printf("%s\n terragrunt_version: ",terragrunt_version)

    buildArgs := map[string]*string{
			    "ALPINE_VERSION": &alpine_version.Str,
			    "PYTHON3_VERSION": &python3_version.Str,
			    "AWSCLI_VERSION": &aws_cli_version.Str,
			    "KUBECTL_VERSION": &kubectl_version.Str,
			    "TERRAFORM_VERSION": &terraform_version.Str,
			    "TERRAGRUNT_VERSION": &terragrunt_version.Str,
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

	if strings.Compare(docker_function, "dockerBuild") == 0 {
	err = imageBuild(cli, dockerPath, tagName, buildArgs)
	
	if err != nil {
		fmt.Println(err.Error())
		return
	}
  } 
}

func main(){

	utility_name := os.Args[1]
	docker_function := os.Args[2]

	if strings.Compare(string(utility_name),  "infrautility") == 0 {
 			infrautility(string(utility_name), string(docker_function))
	}
}