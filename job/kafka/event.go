/*
Copyright 2018 BlackRock, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package kafka

import (
	"time"

	"github.com/Shopify/sarama"
	"github.com/blackrock/axis/job"
)

type event struct {
	job.AbstractEvent
	kafka    *kafka
	kafkaMsg *sarama.ConsumerMessage
}

func (e *event) GetID() string {
	return e.kafka.AbstractSignal.GetID()
}

func (e *event) GetSource() string {
	return e.kafkaMsg.Topic
}

func (e *event) GetBody() []byte {
	return e.kafkaMsg.Value
}

func (e *event) GetTimestamp() time.Time {
	return e.kafkaMsg.Timestamp
}

func (e *event) GetSignal() job.Signal {
	return e.kafka
}
