package ratelimiter_test

import (
	"log"
	"testing"

	"github.com/go-chassis/go-chassis/v2/core/config"
	"github.com/go-chassis/go-chassis/v2/core/config/model"
	"github.com/go-chassis/go-chassis/v2/core/handler"
	"github.com/go-chassis/go-chassis/v2/core/invocation"
	"github.com/go-chassis/go-chassis/v2/core/lager"
	"github.com/go-chassis/go-chassis/v2/examples/schemas/helloworld"
	"github.com/go-chassis/go-chassis/v2/middleware/ratelimiter"
	"github.com/stretchr/testify/assert"
)

func init() {
	lager.Init(&lager.Options{
		LoggerLevel: "INFO",
	})
}
func initEnv() {

	config.Init()
}

func TestProviderRateLimiterDisable(t *testing.T) {
	t.Log("testing providerratelimiter handler with qps enabled as false")
	initEnv()

	c := handler.Chain{}
	c.AddHandler(&ratelimiter.ProviderRateLimiterHandler{})

	config.GlobalDefinition = &model.GlobalCfg{}
	config.GlobalDefinition.ServiceComb.FlowControl.Provider.QPS.Enabled = false
	i := &invocation.Invocation{
		SourceMicroService: "service1",
		SchemaID:           "schema1",
		OperationID:        "SayHello",
		Args:               &helloworld.HelloRequest{Name: "peter"},
	}
	c.Next(i, func(r *invocation.Response) {
		assert.NoError(t, r.Err)
		log.Println(r.Result)
	})

}

func TestProviderRateLimiterHandler_Handle(t *testing.T) {
	t.Log("testing providerratelimiter handler with qps enabled as true")

	initEnv()
	c := handler.Chain{}
	c.AddHandler(&ratelimiter.ProviderRateLimiterHandler{})

	config.GlobalDefinition = &model.GlobalCfg{}
	config.GlobalDefinition.ServiceComb.FlowControl.Provider.QPS.Enabled = true
	i := &invocation.Invocation{
		MicroServiceName: "service1",
		SchemaID:         "schema1",
		OperationID:      "SayHello",
		Args:             &helloworld.HelloRequest{Name: "peter"},
	}
	c.Next(i, func(r *invocation.Response) {
		assert.NoError(t, r.Err)
		log.Println(r.Result)
	})
}

func TestProviderRateLimiterHandler_Handle_SourceMicroService(t *testing.T) {
	t.Log("testing providerratelimiter handler with source microservice and qps enabled as true")

	initEnv()
	c := handler.Chain{}
	c.AddHandler(&ratelimiter.ProviderRateLimiterHandler{})

	config.GlobalDefinition = &model.GlobalCfg{}
	config.GlobalDefinition.ServiceComb.FlowControl.Provider.QPS.Enabled = true
	i := &invocation.Invocation{
		SourceMicroService: "service1",
		SchemaID:           "schema1",
		OperationID:        "SayHello",
		Args:               &helloworld.HelloRequest{Name: "peter"},
	}
	c.Next(i, func(r *invocation.Response) {
		assert.NoError(t, r.Err)
		log.Println(r.Result)
	})
}
