package common

import (
	"log"
	"os"
	"strconv"
	"strings"
)

//Reads ENV variable and converts it to integer.
//Uses default value if error occurs.
func ReadEnvIntParam(val *int, def int, env string) {
	*val = def //default value
	if v := os.Getenv(env); v != "" {
		if iv, err := strconv.Atoi(v); err != nil {
			log.Printf(
				"readEnvIntParam: invalid %s value: %s: using default", env, err)
		} else {
			*val = iv
		}
	}
}

//Call log.Fatalln if err not equal to nil.
func FatalIf(err error, v ...string) {
	if err != nil {
		log.Fatalln(strings.Join(v, ":"), err)
	}
}

//Call panic if err not equal to nil.
func PanicIf(err error, v ...string) {
	if err != nil {
		panic(strings.Join(v, ":") + " " + err.Error())
	}
}

//Call log.Println if err not equal to nil.
func LogIf(err error, v ...string) {
	if err != nil {
		log.Println(strings.Join(v, ":"), err)
	}
}
