
# Go-Redis example

## For The General Project
Implementation of key-value storage app. In this project, only strings are used as data types.

```
├─ bin           //The folder where the binary files was created
├─ cmd           //The code that started it all
├─ config.yml    //Config file for backend
├─ go.mod        //3rd party libraries
├─ go.sum        //Sums and versions of 3rd party libraries
├─ makefile      //MakeFile for build,test and version control 
└─ pkg
   ├─ api                    //Api Layer for project
   ├─ model                  //Models for every type of object
   ├─ repository             //DB Layer
   │  ├─ key
   ├─ server                 //Server Layer for all aplication.
   ├─ service                //Service Layer
   │  ├─ key
   └─ version                //Version control&save for git

```

## ⚡️ Quick start

First of all, [download](https://golang.org/dl/) and install **Go**. :)

## Pre-Req
>Docker

## For build

```bash
make build
```
## For Test

```bash
make test
```
## For Dockerize and run

```bash
make dockerize
```