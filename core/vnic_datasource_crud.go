package core

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type VnicDatasourceCrud struct {
	D        *schema.ResourceData
	Client   client.BareMetalClient
	Resource *baremetal.Vnic
}

func (v *VnicDatasourceCrud) Get() (e error) {
	id := v.D.Get("vnic_id").(string)

	v.Resource, e = v.Client.GetVnic(id)
	return
}

func (v *VnicDatasourceCrud) SetData() {
	if v.Resource != nil {
		v.D.SetId(v.Resource.ID)
		v.D.Set("id", v.Resource.ID)
		v.D.Set("availability_domain", v.Resource.AvailabilityDomain)
		v.D.Set("compartment_id", v.Resource.CompartmentID)
		v.D.Set("display_name", v.Resource.DisplayName)
		v.D.Set("state", v.Resource.State)
		v.D.Set("private_ip_address", v.Resource.PrivateIPAddress)
		v.D.Set("public_ip_address", v.Resource.PublicIPAddress)
		v.D.Set("subnet_id", v.Resource.SubnetID)
	}
	return
}