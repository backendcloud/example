package openstack

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
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

func (c *OpenStackClient) ListServer() ([]servers.Server, error) {
	client, err := c.GetServerProvider()
	if err != nil {
		return nil, err
	}
	opts := servers.ListOpts{}
	pager, _ := servers.List(client, opts).AllPages()
	allPages, err := servers.ExtractServers(pager)
	if err != nil {
		return nil, err
	}
	return allPages, nil
}
