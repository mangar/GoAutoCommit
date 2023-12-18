package main

import (
    "encoding/json"
    "io/ioutil"
    "fmt"
    "github.com/sirupsen/logrus"
    "gopkg.in/natefinch/lumberjack.v2"
    "os"    
    "os/exec"
)


type Config struct {
    Log   string `json:"log"`
}

type Repo struct {
    Name      string `json:"name"`
    Directory string `json:"directory"`
}

type Configuration struct {
    Config Config `json:"config"`
    Repos  []Repo `json:"repos"`
}


func main() {

    // 
    if len(os.Args) < 2 {
        fmt.Println("Inform the config file. Ex.: goautocommit ./config.json")
        return
    }

    var configFileParam string = os.Args[1]

    // 
    configuration := loadConfiguration(configFileParam)
    log := configLog(configuration)

    log.Info("----- Starting GoAutoCommit")


    // 
    for _, repo := range configuration.Repos {
        fmt.Printf("[%s] Repo: %s\n", repo.Name, repo.Directory)
        checkAndCommit(repo.Directory, log)
    }

    log.Info("----- Finished GoAutoCommit")
}



func loadConfiguration(configFileParam string) Configuration{    
    configFile, err := ioutil.ReadFile(configFileParam)
    if err != nil {
        panic(err)
    }

    var configuration Configuration
    json.Unmarshal(configFile, &configuration)
    return configuration
}

func configLog(configuration Configuration) *logrus.Logger {
    log := logrus.New()
    // log.SetFormatter(&logrus.JSONFormatter{})
    log.SetOutput(&lumberjack.Logger{
        Filename:   configuration.Config.Log,
        MaxSize:    10, // megabytes
        MaxBackups: 3,
        MaxAge:     28, //days
        Compress:   true, // compressão de arquivos antigos
    })

    return log
}



func checkAndCommit(directory string, log *logrus.Logger) {

    log.Info("Repository: " + directory)

    //
    if _, err := os.Stat(directory + "/.git"); os.IsNotExist(err) {
        return
    }

    //
    cmd := exec.Command("git", "status", "--porcelain")
    cmd.Dir = directory
    output, err := cmd.Output()
    if err != nil {
        return
    }

    if len(output) > 0 {
        //
        cmd = exec.Command("git", "add", "-A")
        cmd.Dir = directory
        cmd.Run()

        //
        cmd = exec.Command("git", "commit", "-m", "Auto-Commit")
        cmd.Dir = directory
        cmd.Run()
    }
}