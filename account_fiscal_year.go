package odoo

// AccountFiscalYear represents account.fiscal.year model.
type AccountFiscalYear struct {
	CompanyId   *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DateFrom    *Time     `xmlrpc:"date_from,omitempty"`
	DateTo      *Time     `xmlrpc:"date_to,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountFiscalYears represents array of account.fiscal.year model.
type AccountFiscalYears []AccountFiscalYear

// AccountFiscalYearModel is the odoo model name.
const AccountFiscalYearModel = "account.fiscal.year"

// Many2One convert AccountFiscalYear to *Many2One.
func (afy *AccountFiscalYear) Many2One() *Many2One {
	return NewMany2One(afy.Id.Get(), "")
}

// CreateAccountFiscalYear creates a new account.fiscal.year model and returns its id.
func (c *Client) CreateAccountFiscalYear(afy *AccountFiscalYear) (int64, error) {
	ids, err := c.CreateAccountFiscalYears([]*AccountFiscalYear{afy})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountFiscalYear creates a new account.fiscal.year model and returns its id.
func (c *Client) CreateAccountFiscalYears(afys []*AccountFiscalYear) ([]int64, error) {
	var vv []interface{}
	for _, v := range afys {
		vv = append(vv, v)
	}
	return c.Create(AccountFiscalYearModel, vv, nil)
}

// UpdateAccountFiscalYear updates an existing account.fiscal.year record.
func (c *Client) UpdateAccountFiscalYear(afy *AccountFiscalYear) error {
	return c.UpdateAccountFiscalYears([]int64{afy.Id.Get()}, afy)
}

// UpdateAccountFiscalYears updates existing account.fiscal.year records.
// All records (represented by ids) will be updated by afy values.
func (c *Client) UpdateAccountFiscalYears(ids []int64, afy *AccountFiscalYear) error {
	return c.Update(AccountFiscalYearModel, ids, afy, nil)
}

// DeleteAccountFiscalYear deletes an existing account.fiscal.year record.
func (c *Client) DeleteAccountFiscalYear(id int64) error {
	return c.DeleteAccountFiscalYears([]int64{id})
}

// DeleteAccountFiscalYears deletes existing account.fiscal.year records.
func (c *Client) DeleteAccountFiscalYears(ids []int64) error {
	return c.Delete(AccountFiscalYearModel, ids)
}

// GetAccountFiscalYear gets account.fiscal.year existing record.
func (c *Client) GetAccountFiscalYear(id int64) (*AccountFiscalYear, error) {
	afys, err := c.GetAccountFiscalYears([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*afys)[0]), nil
}

// GetAccountFiscalYears gets account.fiscal.year existing records.
func (c *Client) GetAccountFiscalYears(ids []int64) (*AccountFiscalYears, error) {
	afys := &AccountFiscalYears{}
	if err := c.Read(AccountFiscalYearModel, ids, nil, afys); err != nil {
		return nil, err
	}
	return afys, nil
}

// FindAccountFiscalYear finds account.fiscal.year record by querying it with criteria.
func (c *Client) FindAccountFiscalYear(criteria *Criteria) (*AccountFiscalYear, error) {
	afys := &AccountFiscalYears{}
	if err := c.SearchRead(AccountFiscalYearModel, criteria, NewOptions().Limit(1), afys); err != nil {
		return nil, err
	}
	return &((*afys)[0]), nil
}

// FindAccountFiscalYears finds account.fiscal.year records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFiscalYears(criteria *Criteria, options *Options) (*AccountFiscalYears, error) {
	afys := &AccountFiscalYears{}
	if err := c.SearchRead(AccountFiscalYearModel, criteria, options, afys); err != nil {
		return nil, err
	}
	return afys, nil
}

// FindAccountFiscalYearIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFiscalYearIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountFiscalYearModel, criteria, options)
}

// FindAccountFiscalYearId finds record id by querying it with criteria.
func (c *Client) FindAccountFiscalYearId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountFiscalYearModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
