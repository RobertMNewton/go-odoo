package odoo

// DeliveryPriceRule represents delivery.price.rule model.
type DeliveryPriceRule struct {
	CarrierId      *Many2One  `xmlrpc:"carrier_id,omitempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId     *Many2One  `xmlrpc:"currency_id,omitempty"`
	DisplayName    *String    `xmlrpc:"display_name,omitempty"`
	Id             *Int       `xmlrpc:"id,omitempty"`
	ListBasePrice  *Float     `xmlrpc:"list_base_price,omitempty"`
	ListPrice      *Float     `xmlrpc:"list_price,omitempty"`
	MaxValue       *Float     `xmlrpc:"max_value,omitempty"`
	Name           *String    `xmlrpc:"name,omitempty"`
	Operator       *Selection `xmlrpc:"operator,omitempty"`
	Sequence       *Int       `xmlrpc:"sequence,omitempty"`
	Variable       *Selection `xmlrpc:"variable,omitempty"`
	VariableFactor *Selection `xmlrpc:"variable_factor,omitempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// DeliveryPriceRules represents array of delivery.price.rule model.
type DeliveryPriceRules []DeliveryPriceRule

// DeliveryPriceRuleModel is the odoo model name.
const DeliveryPriceRuleModel = "delivery.price.rule"

// Many2One convert DeliveryPriceRule to *Many2One.
func (dpr *DeliveryPriceRule) Many2One() *Many2One {
	return NewMany2One(dpr.Id.Get(), "")
}

// CreateDeliveryPriceRule creates a new delivery.price.rule model and returns its id.
func (c *Client) CreateDeliveryPriceRule(dpr *DeliveryPriceRule) (int64, error) {
	ids, err := c.CreateDeliveryPriceRules([]*DeliveryPriceRule{dpr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDeliveryPriceRule creates a new delivery.price.rule model and returns its id.
func (c *Client) CreateDeliveryPriceRules(dprs []*DeliveryPriceRule) ([]int64, error) {
	var vv []interface{}
	for _, v := range dprs {
		vv = append(vv, v)
	}
	return c.Create(DeliveryPriceRuleModel, vv, nil)
}

// UpdateDeliveryPriceRule updates an existing delivery.price.rule record.
func (c *Client) UpdateDeliveryPriceRule(dpr *DeliveryPriceRule) error {
	return c.UpdateDeliveryPriceRules([]int64{dpr.Id.Get()}, dpr)
}

// UpdateDeliveryPriceRules updates existing delivery.price.rule records.
// All records (represented by ids) will be updated by dpr values.
func (c *Client) UpdateDeliveryPriceRules(ids []int64, dpr *DeliveryPriceRule) error {
	return c.Update(DeliveryPriceRuleModel, ids, dpr, nil)
}

// DeleteDeliveryPriceRule deletes an existing delivery.price.rule record.
func (c *Client) DeleteDeliveryPriceRule(id int64) error {
	return c.DeleteDeliveryPriceRules([]int64{id})
}

// DeleteDeliveryPriceRules deletes existing delivery.price.rule records.
func (c *Client) DeleteDeliveryPriceRules(ids []int64) error {
	return c.Delete(DeliveryPriceRuleModel, ids)
}

// GetDeliveryPriceRule gets delivery.price.rule existing record.
func (c *Client) GetDeliveryPriceRule(id int64) (*DeliveryPriceRule, error) {
	dprs, err := c.GetDeliveryPriceRules([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*dprs)[0]), nil
}

// GetDeliveryPriceRules gets delivery.price.rule existing records.
func (c *Client) GetDeliveryPriceRules(ids []int64) (*DeliveryPriceRules, error) {
	dprs := &DeliveryPriceRules{}
	if err := c.Read(DeliveryPriceRuleModel, ids, nil, dprs); err != nil {
		return nil, err
	}
	return dprs, nil
}

// FindDeliveryPriceRule finds delivery.price.rule record by querying it with criteria.
func (c *Client) FindDeliveryPriceRule(criteria *Criteria) (*DeliveryPriceRule, error) {
	dprs := &DeliveryPriceRules{}
	if err := c.SearchRead(DeliveryPriceRuleModel, criteria, NewOptions().Limit(1), dprs); err != nil {
		return nil, err
	}
	return &((*dprs)[0]), nil
}

// FindDeliveryPriceRules finds delivery.price.rule records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDeliveryPriceRules(criteria *Criteria, options *Options) (*DeliveryPriceRules, error) {
	dprs := &DeliveryPriceRules{}
	if err := c.SearchRead(DeliveryPriceRuleModel, criteria, options, dprs); err != nil {
		return nil, err
	}
	return dprs, nil
}

// FindDeliveryPriceRuleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDeliveryPriceRuleIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(DeliveryPriceRuleModel, criteria, options)
}

// FindDeliveryPriceRuleId finds record id by querying it with criteria.
func (c *Client) FindDeliveryPriceRuleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DeliveryPriceRuleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
