package main

import (
	. "demo/openstack"
	"fmt"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
)

func main() {
	openStackClient := &OpenStackClient{
		IdentityEndpoint: "http://openstack-keystone-vip:35357/v3",
		Region:           "regionone",
		ProjectName:      "admin",
		ProjectDomain:    "Default",
		Username:         "admin",
		Password:         "Admin_PWD_864867351qsc2wdv3efb4rgn",
		DomainName:       "Default",
	}

	// create server
	createServerDetail, err := openStackClient.CreateServer(servers.CreateOpts{
		Name:      "hanwei",
		FlavorRef: "1",
		ImageRef:  "cbd95b01-f714-49d6-96a1-431d7f00d93d",
		Networks: []servers.Network{
			servers.Network{UUID: "a1512f64-3dee-4c56-9bff-d39a1b1f4ddd"}},
	})
	if err != nil {
		panic(err)
	}
	//fmt.Println("server details: %s", createServerDetail)

	// get server by serverId
	serverDetail, err := openStackClient.GetServer(createServerDetail.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println("server details: %s", serverDetail)

}
