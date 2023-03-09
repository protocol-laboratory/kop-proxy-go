package main

type mock struct {
}

func (m *mock) Auth(username string, password string, clientId string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mock) AuthTopic(username string, password, clientId, topic, permissionType string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mock) AuthTopicGroup(username string, password, clientId, consumerGroup string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mock) AuthGroupTopic(topic, groupId string) bool {
	//TODO implement me
	panic("implement me")
}

func (m *mock) SubscriptionName(groupId string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mock) PulsarTopic(username, topic string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mock) PartitionNum(username, topic string) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mock) ListTopic(username string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mock) HasFlowQuota(username, topic string) bool {
	//TODO implement me
	panic("implement me")
}
