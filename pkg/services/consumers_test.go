package services_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "github.com/redhatinsights/app-common-go/pkg/api/v1"
	"github.com/redhatinsights/edge-api/pkg/services"
)

var _ = Describe("ConsumerService basic functions", func() {
	Describe("creation of a kafka consumer service", func() {
		Context("returns correct instance", func() {
			p := 9092
			topic := "platform.playbook-dispatcher.runs"
			config := &v1.KafkaConfig{Brokers: []v1.BrokerConfig{{Hostname: "localhost", Port: &p}}}
			s := services.NewKafkaConsumerService(config, topic)
			It("not to be nil", func() {
				Expect(s).ToNot(BeNil())
			})
		})
	})
	Describe("creation of a kafka consumer service", func() {
		Context("returns correct instance", func() {
			p := 9092
			topic := "platform.inventory.events"
			config := &v1.KafkaConfig{Brokers: []v1.BrokerConfig{{Hostname: "localhost", Port: &p}}}
			s := services.NewKafkaConsumerService(config, topic)
			It("not to be nil", func() {
				Expect(s).ToNot(BeNil())
			})
		})
	})
	Describe("Fail to create a service", func() {
		Context("returns nill", func() {
			p := 9092
			topic := "test1"
			config := &v1.KafkaConfig{Brokers: []v1.BrokerConfig{{Hostname: "localhost", Port: &p}}}
			s := services.NewKafkaConsumerService(config, topic)
			It("to be nil", func() {
				Expect(s).To(BeNil())
			})
		})
	})
	Describe("creation of a kafka consumer service", func() {
		Context("returns correct instance", func() {
			p := 9092
			topic := "platform.inventory.events"
			config := &v1.KafkaConfig{Brokers: []v1.BrokerConfig{{Hostname: "localhost", Port: &p}}}
			s := services.NewKafkaConsumerService(config, topic)
			It("not to be nil", func() {
				Expect(s).ToNot(BeNil())
			})
			s.Start()
		})
	})
})
