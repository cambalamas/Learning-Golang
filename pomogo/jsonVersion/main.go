package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ChimeraCoder/gojson"
	"github.com/Sirupsen/logrus"
)

// // Setup store the config options
// type Setup struct {
// 	TWork time.Duration `json:"twork"`
// 	TRest time.Duration `json:"trest"`
// }

// main is the entry point ;)
func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}

	j, err := gojson.ParseJson(f)
	j2, _ := j.(map[string]interface{})
	if err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}

	if err = f.Close(); err != nil {
		logrus.Warn(err)
	}

	for {
		fmt.Println("Work time!")
		time.Sleep(j2["twork"].(time.Duration) * time.Second)
		fmt.Println("Rest time :D")
		time.Sleep(j2["trest"].(time.Duration) * time.Second)
	}
}

//getSetup read a json file and dump it on a Setup struct
// func getSetup(file string) (s Setup) {
// 	setupFile, err := ioutil.ReadFile(file)
// 	if err != nil {
// 		logrus.Fatalf("Fail on load file \"%s\". (%v)", file, err)
// 		os.Exit(1)
// 	}
// 	err = json.Unmarshal(setupFile, &s)
// 	if err != nil {
// 		logrus.Fatalf("Can't dump json into struct. (%v)", err)
// 	}
// 	return
// }
