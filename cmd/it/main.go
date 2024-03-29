package main

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/protocol-laboratory/kop-proxy-go/kop"
	"os"
	"os/signal"
)

func main() {
	config := &kop.Config{}
	config.PulsarConfig.Host = "localhost"
	config.PulsarConfig.HttpPort = 8080
	config.PulsarConfig.TcpPort = 6650
	config.NetConfig.Host = "localhost"
	config.NetConfig.Port = 9092
	config.ClusterId = "it_kop"
	config.AdvertiseHost = "localhost"
	config.AdvertisePort = 9092
	config.NeedSasl = false
	config.MaxConn = int32(500)
	config.MaxConsumersPerGroup = 1
	config.GroupMinSessionTimeoutMs = 0
	config.GroupMaxSessionTimeoutMs = 60000
	config.MaxFetchRecord = 100
	config.MinFetchWaitMs = 10
	config.MaxFetchWaitMs = 200
	config.ContinuousOffset = true
	config.PulsarTenant = "public"
	config.PulsarNamespace = "default"
	config.OffsetTopic = "kafka_offset"
	config.AutoCreateOffsetTopic = true
	config.GroupCoordinatorType = kop.GroupCoordinatorTypeMemory
	config.InitialDelayedJoinMs = 3000
	config.RebalanceTickMs = 100
	e := &ItKopImpl{}
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
	for range interrupt {
		impl.Close()
		return
	}
}

var pulsarClient, _ = pulsar.NewClient(pulsar.ClientOptions{URL: "pulsar://localhost:6650"})

type ItKopImpl struct {
}

func (e ItKopImpl) Auth(username string, password string, clientId string) (bool, error) {
	return true, nil
}

func (e ItKopImpl) AuthTopic(username string, password, clientId, topic, permissionType string) (bool, error) {
	return true, nil
}

func (e ItKopImpl) AuthTopicGroup(username string, password, clientId, consumerGroup string) (bool, error) {
	return true, nil
}

func (e ItKopImpl) AuthGroupTopic(topic, groupId string) bool {
	return true
}

func (e ItKopImpl) SubscriptionName(groupId string) (string, error) {
	return groupId, nil
}

func (e ItKopImpl) PulsarTopic(username, topic string) (string, error) {
	return "persistent://public/default/" + topic, nil
}

func (e ItKopImpl) PartitionNum(username, topic string) (int, error) {
	pulsarTopic, err := e.PulsarTopic(username, topic)
	if err != nil {
		return 0, err
	}
	partitions, err := pulsarClient.TopicPartitions(pulsarTopic)
	if err != nil {
		return 0, err
	}
	return len(partitions), nil
}

func (e ItKopImpl) ListTopic(username string) ([]string, error) {
	return nil, nil
}

func (e ItKopImpl) HasFlowQuota(username, topic string) bool {
	return true
}
