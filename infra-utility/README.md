https://fission.io/
https://www.openfaas.com/

$ go mod tidy
$ go run main.go infrautility


$ logs

{"stream":"Step 1/4 : FROM node:12"}
{"stream":"\n"}
{"stream":" ---\u003e 212cfb481ff8\n"}
{"stream":"Step 2/4 : WORKDIR /src"}
{"stream":"\n"}
{"stream":" ---\u003e Running in 3cf90b49a9d9\n"}
{"stream":"Removing intermediate container 3cf90b49a9d9\n"}
{"stream":" ---\u003e d8de4a6c85d1\n"}
{"stream":"Step 3/4 : COPY . ."}
{"stream":"\n"}
{"stream":" ---\u003e cc963fecdb99\n"}
{"stream":"Step 4/4 : CMD [ \"node\", \"app.js\" ]"}
{"stream":"\n"}
{"stream":" ---\u003e Running in dc9a9bb6b99e\n"}
{"stream":"Removing intermediate container dc9a9bb6b99e\n"}
{"stream":" ---\u003e ba7eebd4501d\n"}
{"aux":{"ID":"sha256:ba7eebd4501d8eb3bc9e8e785e5636a4a04ee73b1c504c75df755f05e286a19b"}}
{"stream":"Successfully built ba7eebd4501d\n"}
{"stream":"Successfully tagged latest:latest\n"}
