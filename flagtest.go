package main
import (
    "fmt"
    "flag"
)

var (
	configFile   = flag.String("config", "", "Config file for MarsR.")
	printVersion = flag.Bool("version", false, "show version")
)

var (
	version  = "0.1.2"
	codename = "MarsR"
	intro    = "A Xray backend that supports v2board"
)

func showVersion() {
	fmt.Printf("%s %s (%s) \n", codename, version, intro)
}

func main(){

    flag.Parse()
	showVersion()
	if *configFile != "" {
	
	} else {
		// Set default config path
    //add 1.1 branch
	}
	if *printVersion { //返回版本号
		return
	}

    }