// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Sample audit-pubsub is a Cloud Run service which handles Cloud Audit Log messages with Pub/Sub data.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// HelloAuditLogEventHandler receives and processes a Pub/Sub message via a CloudEvent.
func HelloAuditLogEventHandler(w http.ResponseWriter, r *http.Request) {
	s := fmt.Sprintf("Created Pub/Sub topic: %s", string(r.Header.Get("Ce-Subject")))
	log.Printf(s)

	// Send empty reply response.
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(""))
}

func main() {
	http.HandleFunc("/", HelloAuditLogEventHandler)
	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// Start HTTP server.
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
