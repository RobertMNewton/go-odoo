package odoo

// SaleEdiCommon represents sale.edi.common model.
type SaleEdiCommon struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// SaleEdiCommons represents array of sale.edi.common model.
type SaleEdiCommons []SaleEdiCommon

// SaleEdiCommonModel is the odoo model name.
const SaleEdiCommonModel = "sale.edi.common"

// Many2One convert SaleEdiCommon to *Many2One.
func (sec *SaleEdiCommon) Many2One() *Many2One {
	return NewMany2One(sec.Id.Get(), "")
}

// CreateSaleEdiCommon creates a new sale.edi.common model and returns its id.
func (c *Client) CreateSaleEdiCommon(sec *SaleEdiCommon) (int64, error) {
	ids, err := c.CreateSaleEdiCommons([]*SaleEdiCommon{sec})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleEdiCommon creates a new sale.edi.common model and returns its id.
func (c *Client) CreateSaleEdiCommons(secs []*SaleEdiCommon) ([]int64, error) {
	var vv []interface{}
	for _, v := range secs {
		vv = append(vv, v)
	}
	return c.Create(SaleEdiCommonModel, vv, nil)
}

// UpdateSaleEdiCommon updates an existing sale.edi.common record.
func (c *Client) UpdateSaleEdiCommon(sec *SaleEdiCommon) error {
	return c.UpdateSaleEdiCommons([]int64{sec.Id.Get()}, sec)
}

// UpdateSaleEdiCommons updates existing sale.edi.common records.
// All records (represented by ids) will be updated by sec values.
func (c *Client) UpdateSaleEdiCommons(ids []int64, sec *SaleEdiCommon) error {
	return c.Update(SaleEdiCommonModel, ids, sec, nil)
}

// DeleteSaleEdiCommon deletes an existing sale.edi.common record.
func (c *Client) DeleteSaleEdiCommon(id int64) error {
	return c.DeleteSaleEdiCommons([]int64{id})
}

// DeleteSaleEdiCommons deletes existing sale.edi.common records.
func (c *Client) DeleteSaleEdiCommons(ids []int64) error {
	return c.Delete(SaleEdiCommonModel, ids)
}

// GetSaleEdiCommon gets sale.edi.common existing record.
func (c *Client) GetSaleEdiCommon(id int64) (*SaleEdiCommon, error) {
	secs, err := c.GetSaleEdiCommons([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*secs)[0]), nil
}

// GetSaleEdiCommons gets sale.edi.common existing records.
func (c *Client) GetSaleEdiCommons(ids []int64) (*SaleEdiCommons, error) {
	secs := &SaleEdiCommons{}
	if err := c.Read(SaleEdiCommonModel, ids, nil, secs); err != nil {
		return nil, err
	}
	return secs, nil
}

// FindSaleEdiCommon finds sale.edi.common record by querying it with criteria.
func (c *Client) FindSaleEdiCommon(criteria *Criteria) (*SaleEdiCommon, error) {
	secs := &SaleEdiCommons{}
	if err := c.SearchRead(SaleEdiCommonModel, criteria, NewOptions().Limit(1), secs); err != nil {
		return nil, err
	}
	return &((*secs)[0]), nil
}

// FindSaleEdiCommons finds sale.edi.common records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleEdiCommons(criteria *Criteria, options *Options) (*SaleEdiCommons, error) {
	secs := &SaleEdiCommons{}
	if err := c.SearchRead(SaleEdiCommonModel, criteria, options, secs); err != nil {
		return nil, err
	}
	return secs, nil
}

// FindSaleEdiCommonIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleEdiCommonIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SaleEdiCommonModel, criteria, options)
}

// FindSaleEdiCommonId finds record id by querying it with criteria.
func (c *Client) FindSaleEdiCommonId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleEdiCommonModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
