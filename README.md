# Go - AutoCommit

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



## Generate the binary 


__For Windows on Linux machine__

    # Inside `src` folder
    GOOS=windows GOARCH=amd64 go build -o goautocommit.exe main.go



__For the same machine__

    # Inside `src` folder
    go build -o goautocommit main.go
