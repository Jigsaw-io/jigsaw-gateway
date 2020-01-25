# [JIGSAW GATEWAY](https://jigsaw-gateway.herokuapp.com)

![version](https://img.shields.io/badge/version-1.0.0-blue.svg) ![license](https://img.shields.io/badge/license-MIT-blue.svg)

This project implements the Jigsaw Blockchain-gateway.

# Getting Started

- GO Installation
https://docs.google.com/document/d/1yTEBXge5kwnq5u6AUcPC61OALerW1SFoAh6KqPme-MY/edit?usp=sharing

- Enter GOPATH Directory
```    
    cd %GOPATH%/src/github.com/Jigsaw-io
```

- Clone the repo
```    
    git clone https://github.com/Jigsaw-io/jigsaw-server.git
```

- Install packages
```    
    go get -u github.com/golang/dep/cmd/dep
    dep ensure   
```

- Build 
```
    go build
```

- Run
```
    go run main.go
```

# Git Workflow

Master branch is the main development branch. Do not commit directly to master branch.

Staging branch is for the staging environment which mimics production environment.

**TODO: figure out who, when and how to merge to staging branch**



