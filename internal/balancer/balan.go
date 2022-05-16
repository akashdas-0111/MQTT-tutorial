package balancer

import (
	"github.com/segmentio/kafka-go"
)

type Custom struct {
	Partition int
}

func (lb *Custom) Balance(msg kafka.Message, partitions ...int) int {
	partition := lb.Partition
	lb.Partition += 1
	if lb.Partition > 4 {
		lb.Partition = 0
	}
	return partition
}
