package odoo

// DeliveryCarrier represents delivery.carrier model.
type DeliveryCarrier struct {
	Active                    *Bool      `xmlrpc:"active,omitempty"`
	Amount                    *Float     `xmlrpc:"amount,omitempty"`
	CanGenerateReturn         *Bool      `xmlrpc:"can_generate_return,omitempty"`
	CarrierDescription        *String    `xmlrpc:"carrier_description,omitempty"`
	CompanyId                 *Many2One  `xmlrpc:"company_id,omitempty"`
	CountryIds                *Relation  `xmlrpc:"country_ids,omitempty"`
	CreateDate                *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                 *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                *Many2One  `xmlrpc:"currency_id,omitempty"`
	DebugLogging              *Bool      `xmlrpc:"debug_logging,omitempty"`
	DeliveryType              *Selection `xmlrpc:"delivery_type,omitempty"`
	DisplayName               *String    `xmlrpc:"display_name,omitempty"`
	ExcludedTagIds            *Relation  `xmlrpc:"excluded_tag_ids,omitempty"`
	FixedMargin               *Float     `xmlrpc:"fixed_margin,omitempty"`
	FixedPrice                *Float     `xmlrpc:"fixed_price,omitempty"`
	FreeOver                  *Bool      `xmlrpc:"free_over,omitempty"`
	GetReturnLabelFromPortal  *Bool      `xmlrpc:"get_return_label_from_portal,omitempty"`
	Id                        *Int       `xmlrpc:"id,omitempty"`
	IntegrationLevel          *Selection `xmlrpc:"integration_level,omitempty"`
	InvoicePolicy             *Selection `xmlrpc:"invoice_policy,omitempty"`
	Margin                    *Float     `xmlrpc:"margin,omitempty"`
	MaxVolume                 *Float     `xmlrpc:"max_volume,omitempty"`
	MaxWeight                 *Float     `xmlrpc:"max_weight,omitempty"`
	MustHaveTagIds            *Relation  `xmlrpc:"must_have_tag_ids,omitempty"`
	Name                      *String    `xmlrpc:"name,omitempty"`
	PriceRuleIds              *Relation  `xmlrpc:"price_rule_ids,omitempty"`
	ProdEnvironment           *Bool      `xmlrpc:"prod_environment,omitempty"`
	ProductId                 *Many2One  `xmlrpc:"product_id,omitempty"`
	ReturnLabelOnDelivery     *Bool      `xmlrpc:"return_label_on_delivery,omitempty"`
	RouteIds                  *Relation  `xmlrpc:"route_ids,omitempty"`
	Sequence                  *Int       `xmlrpc:"sequence,omitempty"`
	ShippingInsurance         *Int       `xmlrpc:"shipping_insurance,omitempty"`
	StateIds                  *Relation  `xmlrpc:"state_ids,omitempty"`
	SupportsShippingInsurance *Bool      `xmlrpc:"supports_shipping_insurance,omitempty"`
	TrackingUrl               *String    `xmlrpc:"tracking_url,omitempty"`
	VolumeUomName             *String    `xmlrpc:"volume_uom_name,omitempty"`
	WeightUomName             *String    `xmlrpc:"weight_uom_name,omitempty"`
	WriteDate                 *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                  *Many2One  `xmlrpc:"write_uid,omitempty"`
	ZipPrefixIds              *Relation  `xmlrpc:"zip_prefix_ids,omitempty"`
}

// DeliveryCarriers represents array of delivery.carrier model.
type DeliveryCarriers []DeliveryCarrier

// DeliveryCarrierModel is the odoo model name.
const DeliveryCarrierModel = "delivery.carrier"

// Many2One convert DeliveryCarrier to *Many2One.
func (dc *DeliveryCarrier) Many2One() *Many2One {
	return NewMany2One(dc.Id.Get(), "")
}

// CreateDeliveryCarrier creates a new delivery.carrier model and returns its id.
func (c *Client) CreateDeliveryCarrier(dc *DeliveryCarrier) (int64, error) {
	ids, err := c.CreateDeliveryCarriers([]*DeliveryCarrier{dc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDeliveryCarrier creates a new delivery.carrier model and returns its id.
func (c *Client) CreateDeliveryCarriers(dcs []*DeliveryCarrier) ([]int64, error) {
	var vv []interface{}
	for _, v := range dcs {
		vv = append(vv, v)
	}
	return c.Create(DeliveryCarrierModel, vv, nil)
}

// UpdateDeliveryCarrier updates an existing delivery.carrier record.
func (c *Client) UpdateDeliveryCarrier(dc *DeliveryCarrier) error {
	return c.UpdateDeliveryCarriers([]int64{dc.Id.Get()}, dc)
}

// UpdateDeliveryCarriers updates existing delivery.carrier records.
// All records (represented by ids) will be updated by dc values.
func (c *Client) UpdateDeliveryCarriers(ids []int64, dc *DeliveryCarrier) error {
	return c.Update(DeliveryCarrierModel, ids, dc, nil)
}

// DeleteDeliveryCarrier deletes an existing delivery.carrier record.
func (c *Client) DeleteDeliveryCarrier(id int64) error {
	return c.DeleteDeliveryCarriers([]int64{id})
}

// DeleteDeliveryCarriers deletes existing delivery.carrier records.
func (c *Client) DeleteDeliveryCarriers(ids []int64) error {
	return c.Delete(DeliveryCarrierModel, ids)
}

// GetDeliveryCarrier gets delivery.carrier existing record.
func (c *Client) GetDeliveryCarrier(id int64) (*DeliveryCarrier, error) {
	dcs, err := c.GetDeliveryCarriers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*dcs)[0]), nil
}

// GetDeliveryCarriers gets delivery.carrier existing records.
func (c *Client) GetDeliveryCarriers(ids []int64) (*DeliveryCarriers, error) {
	dcs := &DeliveryCarriers{}
	if err := c.Read(DeliveryCarrierModel, ids, nil, dcs); err != nil {
		return nil, err
	}
	return dcs, nil
}

// FindDeliveryCarrier finds delivery.carrier record by querying it with criteria.
func (c *Client) FindDeliveryCarrier(criteria *Criteria) (*DeliveryCarrier, error) {
	dcs := &DeliveryCarriers{}
	if err := c.SearchRead(DeliveryCarrierModel, criteria, NewOptions().Limit(1), dcs); err != nil {
		return nil, err
	}
	return &((*dcs)[0]), nil
}

// FindDeliveryCarriers finds delivery.carrier records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDeliveryCarriers(criteria *Criteria, options *Options) (*DeliveryCarriers, error) {
	dcs := &DeliveryCarriers{}
	if err := c.SearchRead(DeliveryCarrierModel, criteria, options, dcs); err != nil {
		return nil, err
	}
	return dcs, nil
}

// FindDeliveryCarrierIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDeliveryCarrierIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(DeliveryCarrierModel, criteria, options)
}

// FindDeliveryCarrierId finds record id by querying it with criteria.
func (c *Client) FindDeliveryCarrierId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DeliveryCarrierModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
