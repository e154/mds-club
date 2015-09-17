package settings

import (
    "fmt"
    "io/ioutil"
    "os"
    "os/user"

    "github.com/astaxie/beego/config"
)

const (
    CONF_NAME string = "app.conf"
    APP_NAME string = "mds"
    APP_VERSION = "1.0.0"
    permMode os.FileMode = 0666
)

// Singleton
var instantiated *Settings = nil

func SettingsPtr() *Settings {
    if instantiated == nil {
        instantiated = new(Settings);
    }
    return instantiated;
}

type Settings struct {
    HomeDir string
    cfg config.ConfigContainer
    dir string
    AppVersion string
}

func (s *Settings) GetHomeDir() (string, error) {
    
    if len(s.HomeDir) != 0 {
        return s.HomeDir, nil
    }

    user, err := user.Current()
    if err != nil {
        return s.HomeDir, err
    }

    s.HomeDir = user.HomeDir

    return s.HomeDir, nil
}

func (s *Settings) Init() *Settings {

    if len(s.HomeDir) == 0 {
        s.GetHomeDir()
    }

    s.dir = fmt.Sprintf("%s/.%s/", s.HomeDir, APP_NAME)
    s.AppVersion = APP_VERSION

//    create app conf dir
    fileList, _ := ioutil.ReadDir(s.HomeDir)

    var exist bool
    for _, file := range fileList {
        if file.Name() == "."+APP_NAME {
            exist = true
            break
        }
    }

    if !exist {
        dir := fmt.Sprintf(`%s/.%s`, s.HomeDir, APP_NAME)
        os.MkdirAll(dir, os.ModePerm)
    }

    return s
}

func (s *Settings) Save() (*Settings, error) {

    if len(s.HomeDir) == 0 {
        s.GetHomeDir()
    }

    if _, err := os.Stat(s.dir + CONF_NAME); os.IsNotExist(err) {
        ioutil.WriteFile(s.dir + CONF_NAME, []byte{}, permMode)
    }

    cfg, err := config.NewConfig("ini", s.dir + CONF_NAME)
    if err != nil {
        return s, err
    }

    cfg.Set("app_version", s.AppVersion)

    if err := cfg.SaveConfigFile(s.dir + CONF_NAME); err != nil {
        fmt.Printf("err with create conf file: %s\n", s.dir + CONF_NAME)
        return s, err
    }

    return s, nil
}

func (s *Settings) Load() (*Settings, error) {

    if _, err := os.Stat(s.dir + CONF_NAME); os.IsNotExist(err) {
        return s.Save()
    }

    // read config file
    cfg, err := config.NewConfig("ini", s.dir + CONF_NAME)
    if err != nil {
        return s, err
    }

    if cfg.String("app_version") != APP_VERSION {
        s.Save()
        return s.Load()
    }

    s.AppVersion = cfg.String("app_version")

    return s, nil
}
