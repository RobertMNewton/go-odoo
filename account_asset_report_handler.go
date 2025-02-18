package odoo

// AccountAssetReportHandler represents account.asset.report.handler model.
type AccountAssetReportHandler struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountAssetReportHandlers represents array of account.asset.report.handler model.
type AccountAssetReportHandlers []AccountAssetReportHandler

// AccountAssetReportHandlerModel is the odoo model name.
const AccountAssetReportHandlerModel = "account.asset.report.handler"

// Many2One convert AccountAssetReportHandler to *Many2One.
func (aarh *AccountAssetReportHandler) Many2One() *Many2One {
	return NewMany2One(aarh.Id.Get(), "")
}

// CreateAccountAssetReportHandler creates a new account.asset.report.handler model and returns its id.
func (c *Client) CreateAccountAssetReportHandler(aarh *AccountAssetReportHandler) (int64, error) {
	ids, err := c.CreateAccountAssetReportHandlers([]*AccountAssetReportHandler{aarh})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountAssetReportHandler creates a new account.asset.report.handler model and returns its id.
func (c *Client) CreateAccountAssetReportHandlers(aarhs []*AccountAssetReportHandler) ([]int64, error) {
	var vv []interface{}
	for _, v := range aarhs {
		vv = append(vv, v)
	}
	return c.Create(AccountAssetReportHandlerModel, vv, nil)
}

// UpdateAccountAssetReportHandler updates an existing account.asset.report.handler record.
func (c *Client) UpdateAccountAssetReportHandler(aarh *AccountAssetReportHandler) error {
	return c.UpdateAccountAssetReportHandlers([]int64{aarh.Id.Get()}, aarh)
}

// UpdateAccountAssetReportHandlers updates existing account.asset.report.handler records.
// All records (represented by ids) will be updated by aarh values.
func (c *Client) UpdateAccountAssetReportHandlers(ids []int64, aarh *AccountAssetReportHandler) error {
	return c.Update(AccountAssetReportHandlerModel, ids, aarh, nil)
}

// DeleteAccountAssetReportHandler deletes an existing account.asset.report.handler record.
func (c *Client) DeleteAccountAssetReportHandler(id int64) error {
	return c.DeleteAccountAssetReportHandlers([]int64{id})
}

// DeleteAccountAssetReportHandlers deletes existing account.asset.report.handler records.
func (c *Client) DeleteAccountAssetReportHandlers(ids []int64) error {
	return c.Delete(AccountAssetReportHandlerModel, ids)
}

// GetAccountAssetReportHandler gets account.asset.report.handler existing record.
func (c *Client) GetAccountAssetReportHandler(id int64) (*AccountAssetReportHandler, error) {
	aarhs, err := c.GetAccountAssetReportHandlers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aarhs)[0]), nil
}

// GetAccountAssetReportHandlers gets account.asset.report.handler existing records.
func (c *Client) GetAccountAssetReportHandlers(ids []int64) (*AccountAssetReportHandlers, error) {
	aarhs := &AccountAssetReportHandlers{}
	if err := c.Read(AccountAssetReportHandlerModel, ids, nil, aarhs); err != nil {
		return nil, err
	}
	return aarhs, nil
}

// FindAccountAssetReportHandler finds account.asset.report.handler record by querying it with criteria.
func (c *Client) FindAccountAssetReportHandler(criteria *Criteria) (*AccountAssetReportHandler, error) {
	aarhs := &AccountAssetReportHandlers{}
	if err := c.SearchRead(AccountAssetReportHandlerModel, criteria, NewOptions().Limit(1), aarhs); err != nil {
		return nil, err
	}
	return &((*aarhs)[0]), nil
}

// FindAccountAssetReportHandlers finds account.asset.report.handler records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAssetReportHandlers(criteria *Criteria, options *Options) (*AccountAssetReportHandlers, error) {
	aarhs := &AccountAssetReportHandlers{}
	if err := c.SearchRead(AccountAssetReportHandlerModel, criteria, options, aarhs); err != nil {
		return nil, err
	}
	return aarhs, nil
}

// FindAccountAssetReportHandlerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAssetReportHandlerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountAssetReportHandlerModel, criteria, options)
}

// FindAccountAssetReportHandlerId finds record id by querying it with criteria.
func (c *Client) FindAccountAssetReportHandlerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAssetReportHandlerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
