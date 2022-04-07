package

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
)
-example

import (
"github.com/gophercloud/gophercloud"
"github.com/gophercloud/gophercloud/openstack"
)

-example

import (
"github.com/gophercloud/gophercloud"
"github.com/gophercloud/gophercloud/openstack"
)
type OpenStackClient struct {
	IdentityEndpoint string `json:"identityEndpoint"`
	Region           string `json:"region"`
	ProjectName      string `json:"projectName"`
	ProjectDomain    string `json:"projectDomain"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	DomainName       string `json:"domainName"`
}

func (c *OpenStackClient) GetAuth() (*gophercloud.ProviderClient, error) {
	projectScope := gophercloud.AuthScope{
		ProjectName: c.ProjectName,
		DomainName:  c.ProjectDomain,
	}
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: c.IdentityEndpoint,
		Username:         c.Username,
		Password:         c.Password,
		DomainName:       c.DomainName,
		Scope:            &projectScope,
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		return nil, err
	}
	return provider, nil
}
