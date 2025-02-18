package odoo

// AccountTransferModelLine represents account.transfer.model.line model.
type AccountTransferModelLine struct {
	AccountId          *Many2One `xmlrpc:"account_id,omitempty"`
	AnalyticAccountIds *Relation `xmlrpc:"analytic_account_ids,omitempty"`
	CreateDate         *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName        *String   `xmlrpc:"display_name,omitempty"`
	Id                 *Int      `xmlrpc:"id,omitempty"`
	PartnerIds         *Relation `xmlrpc:"partner_ids,omitempty"`
	Percent            *Float    `xmlrpc:"percent,omitempty"`
	PercentIsReadonly  *Bool     `xmlrpc:"percent_is_readonly,omitempty"`
	Sequence           *Int      `xmlrpc:"sequence,omitempty"`
	TransferModelId    *Many2One `xmlrpc:"transfer_model_id,omitempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountTransferModelLines represents array of account.transfer.model.line model.
type AccountTransferModelLines []AccountTransferModelLine

// AccountTransferModelLineModel is the odoo model name.
const AccountTransferModelLineModel = "account.transfer.model.line"

// Many2One convert AccountTransferModelLine to *Many2One.
func (atml *AccountTransferModelLine) Many2One() *Many2One {
	return NewMany2One(atml.Id.Get(), "")
}

// CreateAccountTransferModelLine creates a new account.transfer.model.line model and returns its id.
func (c *Client) CreateAccountTransferModelLine(atml *AccountTransferModelLine) (int64, error) {
	ids, err := c.CreateAccountTransferModelLines([]*AccountTransferModelLine{atml})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountTransferModelLine creates a new account.transfer.model.line model and returns its id.
func (c *Client) CreateAccountTransferModelLines(atmls []*AccountTransferModelLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range atmls {
		vv = append(vv, v)
	}
	return c.Create(AccountTransferModelLineModel, vv, nil)
}

// UpdateAccountTransferModelLine updates an existing account.transfer.model.line record.
func (c *Client) UpdateAccountTransferModelLine(atml *AccountTransferModelLine) error {
	return c.UpdateAccountTransferModelLines([]int64{atml.Id.Get()}, atml)
}

// UpdateAccountTransferModelLines updates existing account.transfer.model.line records.
// All records (represented by ids) will be updated by atml values.
func (c *Client) UpdateAccountTransferModelLines(ids []int64, atml *AccountTransferModelLine) error {
	return c.Update(AccountTransferModelLineModel, ids, atml, nil)
}

// DeleteAccountTransferModelLine deletes an existing account.transfer.model.line record.
func (c *Client) DeleteAccountTransferModelLine(id int64) error {
	return c.DeleteAccountTransferModelLines([]int64{id})
}

// DeleteAccountTransferModelLines deletes existing account.transfer.model.line records.
func (c *Client) DeleteAccountTransferModelLines(ids []int64) error {
	return c.Delete(AccountTransferModelLineModel, ids)
}

// GetAccountTransferModelLine gets account.transfer.model.line existing record.
func (c *Client) GetAccountTransferModelLine(id int64) (*AccountTransferModelLine, error) {
	atmls, err := c.GetAccountTransferModelLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*atmls)[0]), nil
}

// GetAccountTransferModelLines gets account.transfer.model.line existing records.
func (c *Client) GetAccountTransferModelLines(ids []int64) (*AccountTransferModelLines, error) {
	atmls := &AccountTransferModelLines{}
	if err := c.Read(AccountTransferModelLineModel, ids, nil, atmls); err != nil {
		return nil, err
	}
	return atmls, nil
}

// FindAccountTransferModelLine finds account.transfer.model.line record by querying it with criteria.
func (c *Client) FindAccountTransferModelLine(criteria *Criteria) (*AccountTransferModelLine, error) {
	atmls := &AccountTransferModelLines{}
	if err := c.SearchRead(AccountTransferModelLineModel, criteria, NewOptions().Limit(1), atmls); err != nil {
		return nil, err
	}
	return &((*atmls)[0]), nil
}

// FindAccountTransferModelLines finds account.transfer.model.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTransferModelLines(criteria *Criteria, options *Options) (*AccountTransferModelLines, error) {
	atmls := &AccountTransferModelLines{}
	if err := c.SearchRead(AccountTransferModelLineModel, criteria, options, atmls); err != nil {
		return nil, err
	}
	return atmls, nil
}

// FindAccountTransferModelLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTransferModelLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountTransferModelLineModel, criteria, options)
}

// FindAccountTransferModelLineId finds record id by querying it with criteria.
func (c *Client) FindAccountTransferModelLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountTransferModelLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
