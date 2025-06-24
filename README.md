# Microblog

# Getting started

## Go installation

1. Install Go
```
$ brew install golang
```

2. Verify installation
```
$ go version
```

3. Create workspace
```
$ mkdir -p $HOME/go/{bin,src}
```

4. Apply changes to bash_profile
```
$ source ~/.bash_profile
```

## Clone the project

1. Clone the project
```
$ git clone https://github.com/samirgattas/microblog.git
$ cd microblog
```

2. Download missing dependencies
```
$ go mod tidy
```

## Install docker

1. Install docker
```
$ brew install docker
```

2. Install colima
```
$ brew install colima
```

## Run the project

The project use the port 8080. It is important to check if the port is free
```
lsof -i :8080
```

### Locally

Run the project
```
$ make run
```

### Docker container

1. Run colima
```
$ colima start
```

2. Build the docker image
```
$ make dockerbuild
```

3. Run docker
```
$ make dockerrun
```

# Usage

To know the API calls, consult the swagger