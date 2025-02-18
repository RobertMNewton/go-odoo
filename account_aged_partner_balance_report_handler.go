package odoo

// AccountAgedPartnerBalanceReportHandler represents account.aged.partner.balance.report.handler model.
type AccountAgedPartnerBalanceReportHandler struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// AccountAgedPartnerBalanceReportHandlers represents array of account.aged.partner.balance.report.handler model.
type AccountAgedPartnerBalanceReportHandlers []AccountAgedPartnerBalanceReportHandler

// AccountAgedPartnerBalanceReportHandlerModel is the odoo model name.
const AccountAgedPartnerBalanceReportHandlerModel = "account.aged.partner.balance.report.handler"

// Many2One convert AccountAgedPartnerBalanceReportHandler to *Many2One.
func (aapbrh *AccountAgedPartnerBalanceReportHandler) Many2One() *Many2One {
	return NewMany2One(aapbrh.Id.Get(), "")
}

// CreateAccountAgedPartnerBalanceReportHandler creates a new account.aged.partner.balance.report.handler model and returns its id.
func (c *Client) CreateAccountAgedPartnerBalanceReportHandler(aapbrh *AccountAgedPartnerBalanceReportHandler) (int64, error) {
	ids, err := c.CreateAccountAgedPartnerBalanceReportHandlers([]*AccountAgedPartnerBalanceReportHandler{aapbrh})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountAgedPartnerBalanceReportHandler creates a new account.aged.partner.balance.report.handler model and returns its id.
func (c *Client) CreateAccountAgedPartnerBalanceReportHandlers(aapbrhs []*AccountAgedPartnerBalanceReportHandler) ([]int64, error) {
	var vv []interface{}
	for _, v := range aapbrhs {
		vv = append(vv, v)
	}
	return c.Create(AccountAgedPartnerBalanceReportHandlerModel, vv, nil)
}

// UpdateAccountAgedPartnerBalanceReportHandler updates an existing account.aged.partner.balance.report.handler record.
func (c *Client) UpdateAccountAgedPartnerBalanceReportHandler(aapbrh *AccountAgedPartnerBalanceReportHandler) error {
	return c.UpdateAccountAgedPartnerBalanceReportHandlers([]int64{aapbrh.Id.Get()}, aapbrh)
}

// UpdateAccountAgedPartnerBalanceReportHandlers updates existing account.aged.partner.balance.report.handler records.
// All records (represented by ids) will be updated by aapbrh values.
func (c *Client) UpdateAccountAgedPartnerBalanceReportHandlers(ids []int64, aapbrh *AccountAgedPartnerBalanceReportHandler) error {
	return c.Update(AccountAgedPartnerBalanceReportHandlerModel, ids, aapbrh, nil)
}

// DeleteAccountAgedPartnerBalanceReportHandler deletes an existing account.aged.partner.balance.report.handler record.
func (c *Client) DeleteAccountAgedPartnerBalanceReportHandler(id int64) error {
	return c.DeleteAccountAgedPartnerBalanceReportHandlers([]int64{id})
}

// DeleteAccountAgedPartnerBalanceReportHandlers deletes existing account.aged.partner.balance.report.handler records.
func (c *Client) DeleteAccountAgedPartnerBalanceReportHandlers(ids []int64) error {
	return c.Delete(AccountAgedPartnerBalanceReportHandlerModel, ids)
}

// GetAccountAgedPartnerBalanceReportHandler gets account.aged.partner.balance.report.handler existing record.
func (c *Client) GetAccountAgedPartnerBalanceReportHandler(id int64) (*AccountAgedPartnerBalanceReportHandler, error) {
	aapbrhs, err := c.GetAccountAgedPartnerBalanceReportHandlers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aapbrhs)[0]), nil
}

// GetAccountAgedPartnerBalanceReportHandlers gets account.aged.partner.balance.report.handler existing records.
func (c *Client) GetAccountAgedPartnerBalanceReportHandlers(ids []int64) (*AccountAgedPartnerBalanceReportHandlers, error) {
	aapbrhs := &AccountAgedPartnerBalanceReportHandlers{}
	if err := c.Read(AccountAgedPartnerBalanceReportHandlerModel, ids, nil, aapbrhs); err != nil {
		return nil, err
	}
	return aapbrhs, nil
}

// FindAccountAgedPartnerBalanceReportHandler finds account.aged.partner.balance.report.handler record by querying it with criteria.
func (c *Client) FindAccountAgedPartnerBalanceReportHandler(criteria *Criteria) (*AccountAgedPartnerBalanceReportHandler, error) {
	aapbrhs := &AccountAgedPartnerBalanceReportHandlers{}
	if err := c.SearchRead(AccountAgedPartnerBalanceReportHandlerModel, criteria, NewOptions().Limit(1), aapbrhs); err != nil {
		return nil, err
	}
	return &((*aapbrhs)[0]), nil
}

// FindAccountAgedPartnerBalanceReportHandlers finds account.aged.partner.balance.report.handler records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAgedPartnerBalanceReportHandlers(criteria *Criteria, options *Options) (*AccountAgedPartnerBalanceReportHandlers, error) {
	aapbrhs := &AccountAgedPartnerBalanceReportHandlers{}
	if err := c.SearchRead(AccountAgedPartnerBalanceReportHandlerModel, criteria, options, aapbrhs); err != nil {
		return nil, err
	}
	return aapbrhs, nil
}

// FindAccountAgedPartnerBalanceReportHandlerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAgedPartnerBalanceReportHandlerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountAgedPartnerBalanceReportHandlerModel, criteria, options)
}

// FindAccountAgedPartnerBalanceReportHandlerId finds record id by querying it with criteria.
func (c *Client) FindAccountAgedPartnerBalanceReportHandlerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAgedPartnerBalanceReportHandlerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
