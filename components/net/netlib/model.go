package netlib

type ThroughputLogger struct {
}

type NetworkThroughput struct {
	In  uint64
	Out uint64
}

type ThroughputList []NetworkThroughput
