package main

import (
	"github.com/protocol-laboratory/kop-proxy-go/kop"
	"os"
	"os/signal"
)

func main() {
	config := &kop.Config{
		PulsarConfig: kop.PulsarConfig{
			Host:     "localhost",
			HttpPort: 8080,
			TcpPort:  6650,
		},
		ClusterId:                "test_kop",
		AdvertiseHost:            "localhost",
		AdvertisePort:            9092,
		NeedSasl:                 false,
		MaxConn:                  int32(500),
		MaxConsumersPerGroup:     1,
		GroupMaxSessionTimeoutMs: 60000,
		MaxFetchRecord:           100,
		MinFetchWaitMs:           10,
		MaxFetchWaitMs:           5000,
		ContinuousOffset:         true,
		PulsarTenant:             "public",
		PulsarNamespace:          "default",
		OffsetTopic:              "kafka_offset",
		GroupCoordinatorType:     kop.Standalone,
		InitialDelayedJoinMs:     3000,
		RebalanceTickMs:          100,
	}
	e := &mock{}
	impl, err := kop.NewKop(e, config)
	if err != nil {
		panic(err)
	}
	err = impl.Run()
	if err != nil {
		panic(err)
	}
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	for {
		<-interrupt
	}
}
