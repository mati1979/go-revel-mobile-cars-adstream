# Introduction

This application demonstrates usage of revel framework with websockets and real time streaming using mobile.de API.

# Running in a container
The easiest way is to run this using docker and instructions available here:
```
https://registry.hub.docker.com/u/matiwinnetou/go-revel-mobile-cars-adstream/
```

# Running outside container (localhost)

# Go installation
To install golang we will use brew command tool for osx, if you are not aware of it, familarize yourself with it here (http://brew.sh/)

```
brew install go
mkdir gocode
cd gocode
export GOROOT="/usr/local/Cellar/go/1.3.3/libexec" #replace with your go version
export GOPATH=`pwd` #gocode, directory previously created, will be bound to GOPATH variable
export PATH=$GOPATH/bin:$GOPATH #bin folder of your compiled executables will end up in system path
```

# Revel
This application uses Revel web framework (http://revel.github.io/)

To install revel simply type:
```
go get github.com/revel/cmd/revel
```

this will install revel command line tool to your $GOPATH/bin

You are verify that revel is correctly installed by running a demo chat app:
revel run github.com/revel/revel/samples/chat

https://registry.hub.docker.com/u/matiwinnetou/go-revel-mobile-cars-adstream/
Once google go and revel are installed you need to run go get:
```
go get github.com/matiwinnetou/go-revel-mobile-cars-adstream
```

Once you do this and assuming you are in $GOPATH folder open a file app.conf for editing which is in conf subdirectory:
```
vim src/github.com/matiwinnetou/go-revel-mobile-cars-adstream/conf/app.conf
```
and you need to provide your mobile.de AdStream credentials:
```
adstream.user=
adstream.pass=
```
Please enter them for both dev and prod modes

Once credentils are in place simply run the command:
```
revel run github.com/matiwinnetou/go-revel-mobile-cars-adstream
```
this will start the web application on port 9000 in dev modes

You can start the application in prod mode by passing a prod variable at the end of the command:
```
revel run github.com/matiwinnetou/go-revel-mobile-cars-adstream prod
```

# TODO
1. Occasionally after some time a stream hangs, detect a timeout and restart a remote stream to mobile.de
2. Introduce ads a second variable and expose it to web gui
3. Introduce a infobox, when user clicks on an ad he can see a picture of a car and details about it, e.g. price, place from where the car is
4. Provide an ability to filter stream based on country, price, city, etc, make, model
