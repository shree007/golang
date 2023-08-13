package main

import (
	"io/ioutil"
	"log"
	"os/exec"
	"os"
	"strings"
	"fmt"
	"context"
	"archive/tar"
	"bytes"
	
    "github.com/tidwall/gjson" // more info https://github.com/tidwall/gjson
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func getcurrentdir() string {
  command_to_execute := "pwd"
  cmd := exec.Command(command_to_execute)
  stdout, err := cmd.Output()
  fmt.Println(err)
  return string(stdout)
}

func buildDockerImage(dockerPath string, alpine_version string){
	tagName := "latest"
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatal(err, " :unable to init client")
	}

	buf := new(bytes.Buffer)
	tw  := tar.NewWriter(buf)
	defer tw.close()

   dockerFile := "Dockerfile"
   dockerFileReader, err := os.Open(dockerPath)
   if err != nil {
   		log.Fatal(err, " : unable to open Dockerfile")
   }
   readDockerFile, err := ioutil.ReadAll(dockerFileReader)
   if err != nil {
   		log.Fatal(err, " :unable to read dockerfile")
   }

   tarHeader := &tar.Header{
   	Name: dockerFile,
   	Size: int64(len(dockerFile)),
   }
   err = tw.WriteHeader(tarHeader)
   if err != nil {
   		log.Fatal(err, " :unable to write tar header")
   }

   _, err = tw.Write(readDockerFile)
   if err != nil{
   		log.Fatal(err, " :unable to write tar body")
   }
   dockerFileTarReader := bytes.NewReader(buf.Bytes())

   // add any build args
   buildArgs := make(map[string]*string)
   buildArgs["ALPINE_VERSION"] =  alpine_version

   imageBuildResponse, err := cli.ImageBuild(
   	ctx,
   	dockerFileTarReader.
   	types.ImageBuildOptions{
   		Context: dockerFileTarReader,
   		Dockerfile: dockerFile,
   		Tags: []string{tagName},
   		NoCache: true,
   		Remove: true,
   		BuildArgs: buildArgs,
   	})
   	if err != nil {
   		log.Fatal(err, " : unable to build docker image")
   	}
   	defer imageBuildResponse.Body.Close()
   	_, err = io.Copy(os.stdout, imageBuildResponse.Body)

   	if err != nil {
   		log.Fatal(err, " : unable to read image build response")
   	}

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
    
    dockerPath := getcurrentdir() + "/" + utility_name
    fmt.Println(dockerPath)

    buildDockerImage(dockerPath, alpine_version)

}

func main(){

	utility_name := os.Args[1]

	if strings.Compare(string(utility_name),  "infrautility") == 0 {
 			infrautility(string(utility_name))
	}
}