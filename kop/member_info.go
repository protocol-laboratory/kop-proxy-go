package kop

type partitionTopicSet []string

type MemberInfo struct {
	memberId        string
	groupId         string
	groupInstanceId *string
	clientId        string
	fetchRole       bool
	partitionTopics *partitionTopicSet
}

func (mi *MemberInfo) getPartitionTopics() []string {
	if mi.partitionTopics == nil {
		return nil
	}
	return *mi.partitionTopics
}

func (mi *MemberInfo) addPartitionTopics(p string) {
	if mi.partitionTopics == nil {
		mi.partitionTopics = new(partitionTopicSet)
	}
	for _, value := range *mi.partitionTopics {
		if value == p {
			return
		}
	}
	*mi.partitionTopics = append(*mi.partitionTopics, p)
}
