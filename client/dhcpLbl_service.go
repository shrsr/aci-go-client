package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/container"
	"github.com/ciscoecosystem/aci-go-client/models"
)

func (sm *ServiceManager) CreateDHCPRelayLabel(name string, bridge_domain string, tenant string, description string, dhcpLblattr models.DHCPRelayLabelAttributes) (*models.DHCPRelayLabel, error) {
	rn := fmt.Sprintf("dhcplbl-%s", name)
	parentDn := fmt.Sprintf("uni/tn-%s/BD-%s", tenant, bridge_domain)
	dhcpLbl := models.NewDHCPRelayLabel(rn, parentDn, description, dhcpLblattr)
	err := sm.Save(dhcpLbl)
	return dhcpLbl, err
}

func (sm *ServiceManager) ReadDHCPRelayLabel(name string, bridge_domain string, tenant string) (*models.DHCPRelayLabel, error) {
	dn := fmt.Sprintf("uni/tn-%s/BD-%s/dhcplbl-%s", tenant, bridge_domain, name)
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	dhcpLbl := models.DHCPRelayLabelFromContainer(cont)
	return dhcpLbl, nil
}

func (sm *ServiceManager) DeleteDHCPRelayLabel(name string, bridge_domain string, tenant string) error {
	dn := fmt.Sprintf("uni/tn-%s/BD-%s/dhcplbl-%s", tenant, bridge_domain, name)
	return sm.DeleteByDn(dn, models.DhcplblClassName)
}

func (sm *ServiceManager) UpdateDHCPRelayLabel(name string, bridge_domain string, tenant string, description string, dhcpLblattr models.DHCPRelayLabelAttributes) (*models.DHCPRelayLabel, error) {
	rn := fmt.Sprintf("dhcplbl-%s", name)
	parentDn := fmt.Sprintf("uni/tn-%s/BD-%s", tenant, bridge_domain)
	dhcpLbl := models.NewDHCPRelayLabel(rn, parentDn, description, dhcpLblattr)

	dhcpLbl.Status = "modified"
	err := sm.Save(dhcpLbl)
	return dhcpLbl, err

}

func (sm *ServiceManager) ListDHCPRelayLabel(bridge_domain string, tenant string) ([]*models.DHCPRelayLabel, error) {

	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/BD-%s/dhcpLbl.json", baseurlStr, tenant, bridge_domain)

	cont, err := sm.GetViaURL(dnUrl)
	list := models.DHCPRelayLabelListFromContainer(cont)

	return list, err
}

func (sm *ServiceManager) CreateRelationdhcpRsDhcpOptionPolFromDHCPRelayLabel(parentDn, tnDhcpOptionPolName string) error {
	dn := fmt.Sprintf("%s/rsdhcpOptionPol", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnDhcpOptionPolName": "%s","annotation":"orchestrator:terraform"
								
			}
		}
	}`, "dhcpRsDhcpOptionPol", dn, tnDhcpOptionPolName))

	jsonPayload, err := container.ParseJSON(containerJSON)
	if err != nil {
		return err
	}

	req, err := sm.client.MakeRestRequest("POST", fmt.Sprintf("%s.json", sm.MOURL), jsonPayload, true)
	if err != nil {
		return err
	}

	cont, _, err := sm.client.Do(req)
	if err != nil {
		return err
	}
	fmt.Printf("%+v", cont)

	return nil
}

func (sm *ServiceManager) ReadRelationdhcpRsDhcpOptionPolFromDHCPRelayLabel(parentDn string) (interface{}, error) {
	baseurlStr := "/api/node/class"
	dnUrl := fmt.Sprintf("%s/%s/%s.json", baseurlStr, parentDn, "dhcpRsDhcpOptionPol")
	cont, err := sm.GetViaURL(dnUrl)

	contList := models.ListFromContainer(cont, "dhcpRsDhcpOptionPol")

	if len(contList) > 0 {
		dat := models.G(contList[0], "tnDhcpOptionPolName")
		return dat, err
	} else {
		return nil, err
	}

}