# hzlistindexes

## Requirements

* Go 1.18 or better

## Build

Just running `make` will build the executable `hzlistindexes`

## Usage

```
$ ./hzlistindexes -c config.json -m employees

2022/04/28 14:59:17 INFO : trying to connect to cluster: dev
2022/04/28 14:59:17 INFO : connected to cluster: dev
2022/04/28 14:59:17 INFO : 

Members {size:1, ver:1} [
        Member 127.0.0.1:5701 - 21c4344f-e329-4327-afc0-63dfc9c6a39d
]

001: name: age, type: 0 on key: __key, attrs: [age]
```

Note that the JSON configuration is specific to the Go client. Find more at the [documentation](https://pkg.go.dev/github.com/hazelcast/hazelcast-go-client#hdr-Configuration).