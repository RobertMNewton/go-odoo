package odoo

// DeliveryZipPrefix represents delivery.zip.prefix model.
type DeliveryZipPrefix struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// DeliveryZipPrefixs represents array of delivery.zip.prefix model.
type DeliveryZipPrefixs []DeliveryZipPrefix

// DeliveryZipPrefixModel is the odoo model name.
const DeliveryZipPrefixModel = "delivery.zip.prefix"

// Many2One convert DeliveryZipPrefix to *Many2One.
func (dzp *DeliveryZipPrefix) Many2One() *Many2One {
	return NewMany2One(dzp.Id.Get(), "")
}

// CreateDeliveryZipPrefix creates a new delivery.zip.prefix model and returns its id.
func (c *Client) CreateDeliveryZipPrefix(dzp *DeliveryZipPrefix) (int64, error) {
	ids, err := c.CreateDeliveryZipPrefixs([]*DeliveryZipPrefix{dzp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDeliveryZipPrefix creates a new delivery.zip.prefix model and returns its id.
func (c *Client) CreateDeliveryZipPrefixs(dzps []*DeliveryZipPrefix) ([]int64, error) {
	var vv []interface{}
	for _, v := range dzps {
		vv = append(vv, v)
	}
	return c.Create(DeliveryZipPrefixModel, vv, nil)
}

// UpdateDeliveryZipPrefix updates an existing delivery.zip.prefix record.
func (c *Client) UpdateDeliveryZipPrefix(dzp *DeliveryZipPrefix) error {
	return c.UpdateDeliveryZipPrefixs([]int64{dzp.Id.Get()}, dzp)
}

// UpdateDeliveryZipPrefixs updates existing delivery.zip.prefix records.
// All records (represented by ids) will be updated by dzp values.
func (c *Client) UpdateDeliveryZipPrefixs(ids []int64, dzp *DeliveryZipPrefix) error {
	return c.Update(DeliveryZipPrefixModel, ids, dzp, nil)
}

// DeleteDeliveryZipPrefix deletes an existing delivery.zip.prefix record.
func (c *Client) DeleteDeliveryZipPrefix(id int64) error {
	return c.DeleteDeliveryZipPrefixs([]int64{id})
}

// DeleteDeliveryZipPrefixs deletes existing delivery.zip.prefix records.
func (c *Client) DeleteDeliveryZipPrefixs(ids []int64) error {
	return c.Delete(DeliveryZipPrefixModel, ids)
}

// GetDeliveryZipPrefix gets delivery.zip.prefix existing record.
func (c *Client) GetDeliveryZipPrefix(id int64) (*DeliveryZipPrefix, error) {
	dzps, err := c.GetDeliveryZipPrefixs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*dzps)[0]), nil
}

// GetDeliveryZipPrefixs gets delivery.zip.prefix existing records.
func (c *Client) GetDeliveryZipPrefixs(ids []int64) (*DeliveryZipPrefixs, error) {
	dzps := &DeliveryZipPrefixs{}
	if err := c.Read(DeliveryZipPrefixModel, ids, nil, dzps); err != nil {
		return nil, err
	}
	return dzps, nil
}

// FindDeliveryZipPrefix finds delivery.zip.prefix record by querying it with criteria.
func (c *Client) FindDeliveryZipPrefix(criteria *Criteria) (*DeliveryZipPrefix, error) {
	dzps := &DeliveryZipPrefixs{}
	if err := c.SearchRead(DeliveryZipPrefixModel, criteria, NewOptions().Limit(1), dzps); err != nil {
		return nil, err
	}
	return &((*dzps)[0]), nil
}

// FindDeliveryZipPrefixs finds delivery.zip.prefix records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDeliveryZipPrefixs(criteria *Criteria, options *Options) (*DeliveryZipPrefixs, error) {
	dzps := &DeliveryZipPrefixs{}
	if err := c.SearchRead(DeliveryZipPrefixModel, criteria, options, dzps); err != nil {
		return nil, err
	}
	return dzps, nil
}

// FindDeliveryZipPrefixIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDeliveryZipPrefixIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(DeliveryZipPrefixModel, criteria, options)
}

// FindDeliveryZipPrefixId finds record id by querying it with criteria.
func (c *Client) FindDeliveryZipPrefixId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DeliveryZipPrefixModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
