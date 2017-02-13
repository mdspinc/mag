package main

import "log"

func FatalIf(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
