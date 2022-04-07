package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/flavors"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
)

func (c *OpenStackClient) GetServerProvider() (*gophercloud.ServiceClient, error) {
	provider, err := c.GetAuth()
	if err != nil {
		return nil, err
	}
	return openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: c.Region,
	})
}

func (c *OpenStackClient) CreateServer(opts servers.CreateOpts) (*servers.Server, error) {
	client, err := c.GetServerProvider()
	if err != nil {
		return nil, err
	}
	return servers.Create(client, opts).Extract()
}

func (c *OpenStackClient) GetServer(serverId string) (*servers.Server, error) {
	client, err := c.GetServerProvider()
	if err != nil {
		return nil, err
	}
	return servers.Get(client, serverId).Extract()
}

func (c *OpenStackClient) ListFlavors() ([]flavors.Flavor, error) {
	var fls []flavors.Flavor
	provider, err := c.GetAuth()
	if err != nil {
		return fls, err
	}
	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: c.Region,
	})
	if err != nil {
		return fls, err
	}
	pager, err := flavors.ListDetail(client, flavors.ListOpts{}).AllPages()
	if err != nil {
		return fls, err
	}
	allPages, err := flavors.ExtractFlavors(pager)
	if err != nil {
		return fls, err
	}
	//fmt.Println(allPages)
	return allPages, nil
}

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

	//openStackClient.ListFlavors()
	//provider, err := openstack.AuthenticatedClient(opts)
	//if err != nil {
	//	panic(err)
	//}
	//cc, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
	//	Region: os.Getenv("regionone"),
	//})
	//if err != nil {
	//	panic(err)
	//}

	//server, err := servers.Create(cc, servers.CreateOpts{
	//	Name:      "My new server!",
	//	FlavorRef: "flavor_id",
	//	ImageRef:  "image_id",
	//}).Extract()
	//if err != nil {
	//	fmt.Println("Unable to create server: %s", err)
	//}
	//fmt.Println("Server ID: %s", server.ID)

	//serverListOpts := servers.ListOpts{Name: "server_1"}
	//pager := servers.List(cc, serverListOpts)
	//listErr := pager.EachPage(func(page pagination.Page) (bool, error) {
	//	serverList, _ := servers.ExtractServers(page)
	//
	//	for _, s := range serverList {
	//		// "s" will be a servers.Server
	//		fmt.Println("Server ID: %s", s.ID)
	//	}
	//	return false, nil
	//})
	//if listErr != nil {
	//	fmt.Println("Unable to list server: %s", listErr)
	//}
}
