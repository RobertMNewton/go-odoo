package odoo

// AssetModify represents asset.modify model.
type AssetModify struct {
	AccountAssetCounterpartId    *Many2One  `xmlrpc:"account_asset_counterpart_id,omitempty"`
	AccountAssetId               *Many2One  `xmlrpc:"account_asset_id,omitempty"`
	AccountDepreciationExpenseId *Many2One  `xmlrpc:"account_depreciation_expense_id,omitempty"`
	AccountDepreciationId        *Many2One  `xmlrpc:"account_depreciation_id,omitempty"`
	AssetId                      *Many2One  `xmlrpc:"asset_id,omitempty"`
	CompanyId                    *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                   *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                    *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                   *Many2One  `xmlrpc:"currency_id,omitempty"`
	Date                         *Time      `xmlrpc:"date,omitempty"`
	DisplayName                  *String    `xmlrpc:"display_name,omitempty"`
	GainAccountId                *Many2One  `xmlrpc:"gain_account_id,omitempty"`
	GainOrLoss                   *Selection `xmlrpc:"gain_or_loss,omitempty"`
	GainValue                    *Bool      `xmlrpc:"gain_value,omitempty"`
	Id                           *Int       `xmlrpc:"id,omitempty"`
	InformationalText            *String    `xmlrpc:"informational_text,omitempty"`
	InvoiceIds                   *Relation  `xmlrpc:"invoice_ids,omitempty"`
	InvoiceLineIds               *Relation  `xmlrpc:"invoice_line_ids,omitempty"`
	LossAccountId                *Many2One  `xmlrpc:"loss_account_id,omitempty"`
	MethodNumber                 *Int       `xmlrpc:"method_number,omitempty"`
	MethodPeriod                 *Selection `xmlrpc:"method_period,omitempty"`
	ModifyAction                 *Selection `xmlrpc:"modify_action,omitempty"`
	Name                         *String    `xmlrpc:"name,omitempty"`
	SalvageValue                 *Float     `xmlrpc:"salvage_value,omitempty"`
	SelectInvoiceLineId          *Bool      `xmlrpc:"select_invoice_line_id,omitempty"`
	ValueResidual                *Float     `xmlrpc:"value_residual,omitempty"`
	WriteDate                    *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                     *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AssetModifys represents array of asset.modify model.
type AssetModifys []AssetModify

// AssetModifyModel is the odoo model name.
const AssetModifyModel = "asset.modify"

// Many2One convert AssetModify to *Many2One.
func (am *AssetModify) Many2One() *Many2One {
	return NewMany2One(am.Id.Get(), "")
}

// CreateAssetModify creates a new asset.modify model and returns its id.
func (c *Client) CreateAssetModify(am *AssetModify) (int64, error) {
	ids, err := c.CreateAssetModifys([]*AssetModify{am})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAssetModify creates a new asset.modify model and returns its id.
func (c *Client) CreateAssetModifys(ams []*AssetModify) ([]int64, error) {
	var vv []interface{}
	for _, v := range ams {
		vv = append(vv, v)
	}
	return c.Create(AssetModifyModel, vv, nil)
}

// UpdateAssetModify updates an existing asset.modify record.
func (c *Client) UpdateAssetModify(am *AssetModify) error {
	return c.UpdateAssetModifys([]int64{am.Id.Get()}, am)
}

// UpdateAssetModifys updates existing asset.modify records.
// All records (represented by ids) will be updated by am values.
func (c *Client) UpdateAssetModifys(ids []int64, am *AssetModify) error {
	return c.Update(AssetModifyModel, ids, am, nil)
}

// DeleteAssetModify deletes an existing asset.modify record.
func (c *Client) DeleteAssetModify(id int64) error {
	return c.DeleteAssetModifys([]int64{id})
}

// DeleteAssetModifys deletes existing asset.modify records.
func (c *Client) DeleteAssetModifys(ids []int64) error {
	return c.Delete(AssetModifyModel, ids)
}

// GetAssetModify gets asset.modify existing record.
func (c *Client) GetAssetModify(id int64) (*AssetModify, error) {
	ams, err := c.GetAssetModifys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ams)[0]), nil
}

// GetAssetModifys gets asset.modify existing records.
func (c *Client) GetAssetModifys(ids []int64) (*AssetModifys, error) {
	ams := &AssetModifys{}
	if err := c.Read(AssetModifyModel, ids, nil, ams); err != nil {
		return nil, err
	}
	return ams, nil
}

// FindAssetModify finds asset.modify record by querying it with criteria.
func (c *Client) FindAssetModify(criteria *Criteria) (*AssetModify, error) {
	ams := &AssetModifys{}
	if err := c.SearchRead(AssetModifyModel, criteria, NewOptions().Limit(1), ams); err != nil {
		return nil, err
	}
	return &((*ams)[0]), nil
}

// FindAssetModifys finds asset.modify records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAssetModifys(criteria *Criteria, options *Options) (*AssetModifys, error) {
	ams := &AssetModifys{}
	if err := c.SearchRead(AssetModifyModel, criteria, options, ams); err != nil {
		return nil, err
	}
	return ams, nil
}

// FindAssetModifyIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAssetModifyIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AssetModifyModel, criteria, options)
}

// FindAssetModifyId finds record id by querying it with criteria.
func (c *Client) FindAssetModifyId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AssetModifyModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
