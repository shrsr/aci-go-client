package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
)

func (sm *ServiceManager) CreateL4_L7RedirectHealthGroup(name string, tenant string, description string, nameAlias string, vnsRedirectHealthGroupAttr models.L4_L7RedirectHealthGroupAttributes) (*models.L4_L7RedirectHealthGroup, error) {
	rn := fmt.Sprintf(models.RnvnsRedirectHealthGroup, name)
	parentDn := fmt.Sprintf(models.ParentDnvnsRedirectHealthGroup, tenant)
	vnsRedirectHealthGroup := models.NewL4_L7RedirectHealthGroup(rn, parentDn, description, nameAlias, vnsRedirectHealthGroupAttr)
	err := sm.Save(vnsRedirectHealthGroup)
	return vnsRedirectHealthGroup, err
}

func (sm *ServiceManager) ReadL4_L7RedirectHealthGroup(name string, tenant string) (*models.L4_L7RedirectHealthGroup, error) {
	dn := fmt.Sprintf(models.DnvnsRedirectHealthGroup, tenant, name)

	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	vnsRedirectHealthGroup := models.L4_L7RedirectHealthGroupFromContainer(cont)
	return vnsRedirectHealthGroup, nil
}

func (sm *ServiceManager) DeleteL4_L7RedirectHealthGroup(name string, tenant string) error {
	dn := fmt.Sprintf(models.DnvnsRedirectHealthGroup, tenant, name)
	return sm.DeleteByDn(dn, models.VnsredirecthealthgroupClassName)
}

func (sm *ServiceManager) UpdateL4_L7RedirectHealthGroup(name string, tenant string, description string, nameAlias string, vnsRedirectHealthGroupAttr models.L4_L7RedirectHealthGroupAttributes) (*models.L4_L7RedirectHealthGroup, error) {
	rn := fmt.Sprintf(models.RnvnsRedirectHealthGroup, name)
	parentDn := fmt.Sprintf(models.ParentDnvnsRedirectHealthGroup, tenant)
	vnsRedirectHealthGroup := models.NewL4_L7RedirectHealthGroup(rn, parentDn, description, nameAlias, vnsRedirectHealthGroupAttr)
	vnsRedirectHealthGroup.Status = "modified"
	err := sm.Save(vnsRedirectHealthGroup)
	return vnsRedirectHealthGroup, err
}

func (sm *ServiceManager) ListL4_L7RedirectHealthGroup(tenant string) ([]*models.L4_L7RedirectHealthGroup, error) {
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/svcCont/vnsRedirectHealthGroup.json", models.BaseurlStr, tenant)
	cont, err := sm.GetViaURL(dnUrl)
	list := models.L4_L7RedirectHealthGroupListFromContainer(cont)
	return list, err
}
