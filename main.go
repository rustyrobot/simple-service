//  Copyright 2016 Mirantis, Inc.
//
//  Licensed under the Apache License, Version 2.0 (the "License"); you may
//  not use this file except in compliance with the License. You may obtain
//  a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//  WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See then
//  License for the specific language governing permissions and limitations
//  under the License.
	
package main

import (
	"flag"
	"fmt"
	"net/http"
	"log"
	"time"
	"math/big"
)

var bind = flag.String("bind", ":80", "Bind to endpoint")
var counter big.Int
const counterPeriod = 1


func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Current time: %s!", time.Now())
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "Counter: %s", counter.String())
	fmt.Fprintf(w, "\n")
}

func runCounter() {
	for {
		counter.Add(&counter, big.NewInt(1))
		time.Sleep(time.Duration(counterPeriod) * time.Second)
	}
}


func main() {
	flag.Parse()
	log.Printf("Starting service on port %s", *bind)
	http.HandleFunc("/", handler)
	go runCounter()

	if err := http.ListenAndServe(*bind, nil); err != nil {
		log.Fatal(err)
	}

	log.Println("Stopping service...")
}
