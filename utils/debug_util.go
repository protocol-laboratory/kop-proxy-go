package utils

func DebugTopicMatch(kMap, pMap map[string]bool, kTopic, pTopic string) bool {
	return kMap[kTopic] && pMap[pTopic]
}

func DebugKafkaTopicMatch(m map[string]bool, topic string) bool {
	return m[topic]
}

func DebugPulsarPartitionTopicMatch(m map[string]bool, partitionTopic string) bool {
	return m[partitionTopic]
}
