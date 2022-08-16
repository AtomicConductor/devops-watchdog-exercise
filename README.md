## Overview
In this exercise, you will write a golang application that continually queries an http endpoint, parses the objects in the json response, and outputs informational text about them after a timeout.


## HTTP Server
The included source code (server.go) will run a local http server on your machine. It has a "/containers" endpoint that will return a json response. The response contains a list of container objects. The statuses of the containers will change over time.
An example response from this http server:

```
[
  {
    "id": "068d9117685312461f4c3d7eb038b57665ecad966b4e7e5de66e23bbcbe156e2",
    "names": [
      "/conductor_cgls"
    ],
    "image": "gcr.io/conductor/hlm:20220808134517",
    "imageID": "sha256:dd692ed0e3000806102989bb261e02b3724f2e15b09cb8c5976e289460a60aa7",
    "state": "running",
    "status": "Up 7m29s (healthy)"
  },
  {
    "id": "bb4746c2124e75b9891c451e40bb3da37f94e32090edb95a256a2e73bf4306dc",
    "names": [
      "/conductor_hlm"
    ],
    "image": "gcr.io/conductor/hlm:20220808134517",
    "imageID": "sha256:b2d8853ef92f0916dd1be0c1aab8dcfafc979471d56d1fd8aa5a663e18a23f87",
    "state": "running",
    "status": "Up 7m29s (healthy)"
  },
  {
    "id": "32d6abc33c399cd1cfcabe07f4922bc3ac1858725d44be2cbb405184ccf6d095",
    "names": [
      "/conductor_plm"
    ],
    "image": "gcr.io/conductor/plm:20220808134517",
    "imageID": "sha256:62ac29fa0629270f13de19a419de8561f84e6a32d16c4f025eef708ef58ec490",
    "state": "running",
    "status": "Up 7m29s (healthy)"
  },
  {
    "id": "917948e476dad84e1f465f1f0b18ba4ebaf347b6b07c9dcf8aa257098cb943c2",
    "names": [
      "/conductor_rvl"
    ],
    "image": "gcr.io/conductor/rvl:20220808134517",
    "imageID": "sha256:4d52848b3ba75052bee0089a8b0f40d0db7f3c0f31ecfa78e11559e50d90f181",
    "state": "running",
    "status": "Up 7m29s (healthy)"
  }
]
```

The `status` field values look like this: `Up {time} ({health_status})`
where `{health_status}` is one of `starting`, `healthy`, or `unhealthy`.


## Exercise
Write a golang application that outputs the names of any containers that don't pass our quality check. The quality check consists of querying the endpoint every 5 seconds until all the containers are running, ready, and healthy or the timeout (20 seconds) is reached, whichever comes first. We expect to see four healthy containers: `conductor_cgls`, `conductor_hlm`, `conductor_plm`, and `conductor_rvl`.

If all containers pass, the program should output
```
All containers pass!
```

If all containers do not pass, the program should output the names of the failing containers like so:
```
Failing Containers:
  Not Healthy: [/conductor_rvl]
  Not Ready: [/conductor_hlm]
  Not Running: [/conductor_cgls]
```


## Quality Check
`running` means the container is in the list of containers returned by the server.

`not running` means the container is not in the list of containers returned by the server.

<br>

`ready` means the container is running and it has a status of either "healthy" or "unhealthy".

`not ready` means the container is running and it has a status that is not "healthy" or "unhealthy". For example, while the container is starting up, the status will be "starting".

<br>

`healthy` means the container is running, ready, and has a status of "healthy".

`not healthy` means the container is running, ready, and has a status of "unhealthy".
