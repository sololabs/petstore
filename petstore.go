// Copyright 2015 go-swagger maintainers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/sololabs/petstore/api"
)

func main() {
	port := flag.Uint("p", 8080, "listener port")
	flag.Parse()

	petstoreAPI, err := api.NewPetstore()
	if err != nil {
		log.Fatalln(err)
	}
	logAllRequests := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %#v", r)
		petstoreAPI.ServeHTTP(w, r)
	})
	log.Printf("Serving petstore api on http://127.0.0.1:%v/swagger-ui/", *port)
	_ = http.ListenAndServe(fmt.Sprintf(":%v", *port), logAllRequests)
}
