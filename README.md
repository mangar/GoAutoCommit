# Go - AutoCommit

​<img src="./_extras/logo.png" width="120px">

__GoAutoCommit__ created in the Go language, specializes in making commits to selected Git repositories from time to time. Like a digital ninja, it infiltrates your repositories, drops a stylish commit, and disappears before you can say "gopher."



## Configuration

Check the `config.json` file configuration

Template:

    {
        "config":{
            "log":"/mnt/c/Users/my/loglogfile.log"
        },
        "repos":[
            {
                "name": "Temp",
                "directory": "/mnt/c/Users/my/repos/temp"
            }
        ]    
    }

- config
  - log: Directory for the log file
- repos
  - name: Name for organization
  - directory: Local repository



# Building

Go Version: `go version go1.21.5 linux/amd64`


## Generate the binary 


__For Windows on Linux machine__

    # Inside `src` folder
    GOOS=windows GOARCH=amd64 go build -o goautocommit.exe main.go



__For the same machine__

    # Inside `src` folder
    go build -o goautocommit main.go
