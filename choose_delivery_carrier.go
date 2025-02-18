package odoo

// ChooseDeliveryCarrier represents choose.delivery.carrier model.
type ChooseDeliveryCarrier struct {
	AvailableCarrierIds *Relation  `xmlrpc:"available_carrier_ids,omitempty"`
	CarrierId           *Many2One  `xmlrpc:"carrier_id,omitempty"`
	CompanyId           *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate          *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid           *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId          *Many2One  `xmlrpc:"currency_id,omitempty"`
	DeliveryMessage     *String    `xmlrpc:"delivery_message,omitempty"`
	DeliveryPrice       *Float     `xmlrpc:"delivery_price,omitempty"`
	DeliveryType        *Selection `xmlrpc:"delivery_type,omitempty"`
	DisplayName         *String    `xmlrpc:"display_name,omitempty"`
	DisplayPrice        *Float     `xmlrpc:"display_price,omitempty"`
	Id                  *Int       `xmlrpc:"id,omitempty"`
	InvoicingMessage    *String    `xmlrpc:"invoicing_message,omitempty"`
	OrderId             *Many2One  `xmlrpc:"order_id,omitempty"`
	PartnerId           *Many2One  `xmlrpc:"partner_id,omitempty"`
	TotalWeight         *Float     `xmlrpc:"total_weight,omitempty"`
	WeightUomName       *String    `xmlrpc:"weight_uom_name,omitempty"`
	WriteDate           *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid            *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ChooseDeliveryCarriers represents array of choose.delivery.carrier model.
type ChooseDeliveryCarriers []ChooseDeliveryCarrier

// ChooseDeliveryCarrierModel is the odoo model name.
const ChooseDeliveryCarrierModel = "choose.delivery.carrier"

// Many2One convert ChooseDeliveryCarrier to *Many2One.
func (cdc *ChooseDeliveryCarrier) Many2One() *Many2One {
	return NewMany2One(cdc.Id.Get(), "")
}

// CreateChooseDeliveryCarrier creates a new choose.delivery.carrier model and returns its id.
func (c *Client) CreateChooseDeliveryCarrier(cdc *ChooseDeliveryCarrier) (int64, error) {
	ids, err := c.CreateChooseDeliveryCarriers([]*ChooseDeliveryCarrier{cdc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateChooseDeliveryCarrier creates a new choose.delivery.carrier model and returns its id.
func (c *Client) CreateChooseDeliveryCarriers(cdcs []*ChooseDeliveryCarrier) ([]int64, error) {
	var vv []interface{}
	for _, v := range cdcs {
		vv = append(vv, v)
	}
	return c.Create(ChooseDeliveryCarrierModel, vv, nil)
}

// UpdateChooseDeliveryCarrier updates an existing choose.delivery.carrier record.
func (c *Client) UpdateChooseDeliveryCarrier(cdc *ChooseDeliveryCarrier) error {
	return c.UpdateChooseDeliveryCarriers([]int64{cdc.Id.Get()}, cdc)
}

// UpdateChooseDeliveryCarriers updates existing choose.delivery.carrier records.
// All records (represented by ids) will be updated by cdc values.
func (c *Client) UpdateChooseDeliveryCarriers(ids []int64, cdc *ChooseDeliveryCarrier) error {
	return c.Update(ChooseDeliveryCarrierModel, ids, cdc, nil)
}

// DeleteChooseDeliveryCarrier deletes an existing choose.delivery.carrier record.
func (c *Client) DeleteChooseDeliveryCarrier(id int64) error {
	return c.DeleteChooseDeliveryCarriers([]int64{id})
}

// DeleteChooseDeliveryCarriers deletes existing choose.delivery.carrier records.
func (c *Client) DeleteChooseDeliveryCarriers(ids []int64) error {
	return c.Delete(ChooseDeliveryCarrierModel, ids)
}

// GetChooseDeliveryCarrier gets choose.delivery.carrier existing record.
func (c *Client) GetChooseDeliveryCarrier(id int64) (*ChooseDeliveryCarrier, error) {
	cdcs, err := c.GetChooseDeliveryCarriers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*cdcs)[0]), nil
}

// GetChooseDeliveryCarriers gets choose.delivery.carrier existing records.
func (c *Client) GetChooseDeliveryCarriers(ids []int64) (*ChooseDeliveryCarriers, error) {
	cdcs := &ChooseDeliveryCarriers{}
	if err := c.Read(ChooseDeliveryCarrierModel, ids, nil, cdcs); err != nil {
		return nil, err
	}
	return cdcs, nil
}

// FindChooseDeliveryCarrier finds choose.delivery.carrier record by querying it with criteria.
func (c *Client) FindChooseDeliveryCarrier(criteria *Criteria) (*ChooseDeliveryCarrier, error) {
	cdcs := &ChooseDeliveryCarriers{}
	if err := c.SearchRead(ChooseDeliveryCarrierModel, criteria, NewOptions().Limit(1), cdcs); err != nil {
		return nil, err
	}
	return &((*cdcs)[0]), nil
}

// FindChooseDeliveryCarriers finds choose.delivery.carrier records by querying it
// and filtering it with criteria and options.
func (c *Client) FindChooseDeliveryCarriers(criteria *Criteria, options *Options) (*ChooseDeliveryCarriers, error) {
	cdcs := &ChooseDeliveryCarriers{}
	if err := c.SearchRead(ChooseDeliveryCarrierModel, criteria, options, cdcs); err != nil {
		return nil, err
	}
	return cdcs, nil
}

// FindChooseDeliveryCarrierIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindChooseDeliveryCarrierIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ChooseDeliveryCarrierModel, criteria, options)
}

// FindChooseDeliveryCarrierId finds record id by querying it with criteria.
func (c *Client) FindChooseDeliveryCarrierId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ChooseDeliveryCarrierModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
