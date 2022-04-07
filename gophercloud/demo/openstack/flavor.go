package openstack

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/flavors"
)

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
