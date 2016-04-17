# d-sole
Docker security hole, allow to execute a script on the host from a container.

## Concept

The idea is to start an http server on the docker gateway, to allow every container to send a get to this endpoint.
This is more or less secure thanks to the secret hash you have to guess.
If it is not that secure, please let me know!
Actually, I have an awful doubt right now, will other container see the get request? (As this is not https)

## Get Started

### Make it to your taste

Open `main.go` and replace `/path/to/my/script` and the secret hash.
Modify the `d-sole.service` file to put the user of your choice.
We recommand an unpriviledged user, but if you need to run things as root, we recommand to add this user to sudo and give him the right just to execte that particular script.

### Compile it

As you guess, we'll not install golang just to compile that :)

```
docker run --rm -v "$PWD":/usr/src/d-sole -w /usr/src/d-sole golang:1.6 go build -v
cp d-sole /opt/bin
```

### Launch it

```
cp d-sole.service /etc/systemd/system
systemctl daemon-reload
systemctl start d-sole
```

### Profit!

From inside a container:
```
curl http://172.17.0.1:8080
```

Watch your script being executed!

Be careful, keep in mind the meaning of the name of this repo!
