package odoo

// ChooseDeliveryPackage represents choose.delivery.package model.
type ChooseDeliveryPackage struct {
	CompanyId             *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate            *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid             *Many2One `xmlrpc:"create_uid,omitempty"`
	DeliveryPackageTypeId *Many2One `xmlrpc:"delivery_package_type_id,omitempty"`
	DisplayName           *String   `xmlrpc:"display_name,omitempty"`
	Id                    *Int      `xmlrpc:"id,omitempty"`
	MoveLineIds           *Relation `xmlrpc:"move_line_ids,omitempty"`
	ShippingWeight        *Float    `xmlrpc:"shipping_weight,omitempty"`
	WeightUomName         *String   `xmlrpc:"weight_uom_name,omitempty"`
	WriteDate             *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid              *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ChooseDeliveryPackages represents array of choose.delivery.package model.
type ChooseDeliveryPackages []ChooseDeliveryPackage

// ChooseDeliveryPackageModel is the odoo model name.
const ChooseDeliveryPackageModel = "choose.delivery.package"

// Many2One convert ChooseDeliveryPackage to *Many2One.
func (cdp *ChooseDeliveryPackage) Many2One() *Many2One {
	return NewMany2One(cdp.Id.Get(), "")
}

// CreateChooseDeliveryPackage creates a new choose.delivery.package model and returns its id.
func (c *Client) CreateChooseDeliveryPackage(cdp *ChooseDeliveryPackage) (int64, error) {
	ids, err := c.CreateChooseDeliveryPackages([]*ChooseDeliveryPackage{cdp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateChooseDeliveryPackage creates a new choose.delivery.package model and returns its id.
func (c *Client) CreateChooseDeliveryPackages(cdps []*ChooseDeliveryPackage) ([]int64, error) {
	var vv []interface{}
	for _, v := range cdps {
		vv = append(vv, v)
	}
	return c.Create(ChooseDeliveryPackageModel, vv, nil)
}

// UpdateChooseDeliveryPackage updates an existing choose.delivery.package record.
func (c *Client) UpdateChooseDeliveryPackage(cdp *ChooseDeliveryPackage) error {
	return c.UpdateChooseDeliveryPackages([]int64{cdp.Id.Get()}, cdp)
}

// UpdateChooseDeliveryPackages updates existing choose.delivery.package records.
// All records (represented by ids) will be updated by cdp values.
func (c *Client) UpdateChooseDeliveryPackages(ids []int64, cdp *ChooseDeliveryPackage) error {
	return c.Update(ChooseDeliveryPackageModel, ids, cdp, nil)
}

// DeleteChooseDeliveryPackage deletes an existing choose.delivery.package record.
func (c *Client) DeleteChooseDeliveryPackage(id int64) error {
	return c.DeleteChooseDeliveryPackages([]int64{id})
}

// DeleteChooseDeliveryPackages deletes existing choose.delivery.package records.
func (c *Client) DeleteChooseDeliveryPackages(ids []int64) error {
	return c.Delete(ChooseDeliveryPackageModel, ids)
}

// GetChooseDeliveryPackage gets choose.delivery.package existing record.
func (c *Client) GetChooseDeliveryPackage(id int64) (*ChooseDeliveryPackage, error) {
	cdps, err := c.GetChooseDeliveryPackages([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*cdps)[0]), nil
}

// GetChooseDeliveryPackages gets choose.delivery.package existing records.
func (c *Client) GetChooseDeliveryPackages(ids []int64) (*ChooseDeliveryPackages, error) {
	cdps := &ChooseDeliveryPackages{}
	if err := c.Read(ChooseDeliveryPackageModel, ids, nil, cdps); err != nil {
		return nil, err
	}
	return cdps, nil
}

// FindChooseDeliveryPackage finds choose.delivery.package record by querying it with criteria.
func (c *Client) FindChooseDeliveryPackage(criteria *Criteria) (*ChooseDeliveryPackage, error) {
	cdps := &ChooseDeliveryPackages{}
	if err := c.SearchRead(ChooseDeliveryPackageModel, criteria, NewOptions().Limit(1), cdps); err != nil {
		return nil, err
	}
	return &((*cdps)[0]), nil
}

// FindChooseDeliveryPackages finds choose.delivery.package records by querying it
// and filtering it with criteria and options.
func (c *Client) FindChooseDeliveryPackages(criteria *Criteria, options *Options) (*ChooseDeliveryPackages, error) {
	cdps := &ChooseDeliveryPackages{}
	if err := c.SearchRead(ChooseDeliveryPackageModel, criteria, options, cdps); err != nil {
		return nil, err
	}
	return cdps, nil
}

// FindChooseDeliveryPackageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindChooseDeliveryPackageIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ChooseDeliveryPackageModel, criteria, options)
}

// FindChooseDeliveryPackageId finds record id by querying it with criteria.
func (c *Client) FindChooseDeliveryPackageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ChooseDeliveryPackageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
