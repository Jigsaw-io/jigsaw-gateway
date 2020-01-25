# [JIGSAW GATEWAY](https://jigsaw-gateway.herokuapp.com)

![version](https://img.shields.io/badge/version-1.0.0-blue.svg) ![license](https://img.shields.io/badge/license-MIT-blue.svg)

This project implements the Jigsaw Blockchain-gateway.

# Getting Started

- GO Installation
https://docs.google.com/document/d/1yTEBXge5kwnq5u6AUcPC61OALerW1SFoAh6KqPme-MY/edit?usp=sharing

- Setup Workspace
    - Enter GOPATH: `cd %GOPATH%`.
    - Enter src directory : `cd src`.
    - Create a "github.com" directory if absent : `mkdir github.com`.
    - Enter github.com directory : `cd github.com`.
    - Create a "zeemzo" directory if absent : `mkdir zeemzo`.
    - Enter zeemzo directory : `cd zeemzo`.

- Clone the repo
```    
    git clone https://github.com/zeemzo/jigsaw-gateway.git
```

- Install packages
```    
    go get -u github.com/golang/dep/cmd/dep
    cd jigsaw-gateway
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



