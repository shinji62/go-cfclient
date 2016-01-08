package cfclient

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestListUserProvidedServices(t *testing.T) {
	Convey("List User Provided Services", t, func() {
		setup(MockRoute{"GET", "/v2/user_provided_service_instances", listUserProvidedServicesPayload})
		defer teardown()
		c := &Config{
			ApiAddress:   server.URL,
			LoginAddress: fakeUAAServer.URL,
			Token:        "foobar",
		}
		client := NewClient(c)
		services := client.ListUserProvidedServices()
		So(len(services), ShouldEqual, 2)
		So(services[0].Guid, ShouldEqual, "d1e9e4e3-879f-4a58-a0b0-8a048296202b")
		So(services[0].Name, ShouldEqual, "oracle-11-docker")
		So(services[0].Credentials["host"], ShouldEqual, "192.168.99.10")
	})
}
clfm