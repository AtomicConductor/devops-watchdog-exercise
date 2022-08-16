package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	HEALTHY   = "healthy"
	RUNNING   = "running"
	STARTING  = "starting"
	UNHEALTHY = "unhealthy"
)

var (
	containers      []container
	startTime       = time.Now()
	plmStatus       = STARTING
	rvlStatus       = STARTING
	tenSecTimer     = time.NewTimer(time.Duration(10) * time.Second)
	fifteenSecTimer = time.NewTimer(time.Duration(15) * time.Second)
)

type container struct {
	Id      string   `json:"id"`
	Names   []string `json:"names"`
	Image   string   `json:"image"`
	ImageID string   `json:"imageID"`
	State   string   `json:"state"`
	Status  string   `json:"status"`
}

func outputContainers() []container {
	return []container{
		container{
			Id:      "bb4746c2124e75b9891c451e40bb3da37f94e32090edb95a256a2e73bf4306dc",
			Names:   []string{"/conductor_hlm"},
			Image:   "gcr.io/conductor/hlm:20220808134517",
			ImageID: "sha256:b2d8853ef92f0916dd1be0c1aab8dcfafc979471d56d1fd8aa5a663e18a23f87",
			State:   RUNNING,
			Status:  fmt.Sprintf("Up %s (%s)", time.Since(startTime).Round(time.Second), STARTING),
		},
		container{
			Id:      "32d6abc33c399cd1cfcabe07f4922bc3ac1858725d44be2cbb405184ccf6d095",
			Names:   []string{"/conductor_plm"},
			Image:   "gcr.io/conductor/plm:20220808134517",
			ImageID: "sha256:62ac29fa0629270f13de19a419de8561f84e6a32d16c4f025eef708ef58ec490",
			State:   RUNNING,
			Status:  fmt.Sprintf("Up %s (%s)", time.Since(startTime).Round(time.Second), plmStatus),
		},
		container{
			Id:      "917948e476dad84e1f465f1f0b18ba4ebaf347b6b07c9dcf8aa257098cb943c2",
			Names:   []string{"/conductor_rvl"},
			Image:   "gcr.io/conductor/rvl:20220808134517",
			ImageID: "sha256:4d52848b3ba75052bee0089a8b0f40d0db7f3c0f31ecfa78e11559e50d90f181",
			State:   RUNNING,
			Status:  fmt.Sprintf("Up %s (%s)", time.Since(startTime).Round(time.Second), rvlStatus),
		},
	}
}

func handleContainers(w http.ResponseWriter, r *http.Request) {
	containers = outputContainers()
	response, _ := json.Marshal(containers)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func updateContainerStatus() {
	for {
		select {
		case <-tenSecTimer.C:
			plmStatus = HEALTHY
		case <-fifteenSecTimer.C:
			rvlStatus = UNHEALTHY
			return
		}
	}
}

func main() {
	go updateContainerStatus()
	mux := http.NewServeMux()
	mux.HandleFunc("/containers", handleContainers)
	http.ListenAndServe(":8081", mux)
}
