package odoo

// AccountTaxUnit represents account.tax.unit model.
type AccountTaxUnit struct {
	CompanyIds    *Relation `xmlrpc:"company_ids,omitempty"`
	CountryId     *Many2One `xmlrpc:"country_id,omitempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	FposSynced    *Bool     `xmlrpc:"fpos_synced,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	MainCompanyId *Many2One `xmlrpc:"main_company_id,omitempty"`
	Name          *String   `xmlrpc:"name,omitempty"`
	Vat           *String   `xmlrpc:"vat,omitempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountTaxUnits represents array of account.tax.unit model.
type AccountTaxUnits []AccountTaxUnit

// AccountTaxUnitModel is the odoo model name.
const AccountTaxUnitModel = "account.tax.unit"

// Many2One convert AccountTaxUnit to *Many2One.
func (atu *AccountTaxUnit) Many2One() *Many2One {
	return NewMany2One(atu.Id.Get(), "")
}

// CreateAccountTaxUnit creates a new account.tax.unit model and returns its id.
func (c *Client) CreateAccountTaxUnit(atu *AccountTaxUnit) (int64, error) {
	ids, err := c.CreateAccountTaxUnits([]*AccountTaxUnit{atu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountTaxUnit creates a new account.tax.unit model and returns its id.
func (c *Client) CreateAccountTaxUnits(atus []*AccountTaxUnit) ([]int64, error) {
	var vv []interface{}
	for _, v := range atus {
		vv = append(vv, v)
	}
	return c.Create(AccountTaxUnitModel, vv, nil)
}

// UpdateAccountTaxUnit updates an existing account.tax.unit record.
func (c *Client) UpdateAccountTaxUnit(atu *AccountTaxUnit) error {
	return c.UpdateAccountTaxUnits([]int64{atu.Id.Get()}, atu)
}

// UpdateAccountTaxUnits updates existing account.tax.unit records.
// All records (represented by ids) will be updated by atu values.
func (c *Client) UpdateAccountTaxUnits(ids []int64, atu *AccountTaxUnit) error {
	return c.Update(AccountTaxUnitModel, ids, atu, nil)
}

// DeleteAccountTaxUnit deletes an existing account.tax.unit record.
func (c *Client) DeleteAccountTaxUnit(id int64) error {
	return c.DeleteAccountTaxUnits([]int64{id})
}

// DeleteAccountTaxUnits deletes existing account.tax.unit records.
func (c *Client) DeleteAccountTaxUnits(ids []int64) error {
	return c.Delete(AccountTaxUnitModel, ids)
}

// GetAccountTaxUnit gets account.tax.unit existing record.
func (c *Client) GetAccountTaxUnit(id int64) (*AccountTaxUnit, error) {
	atus, err := c.GetAccountTaxUnits([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*atus)[0]), nil
}

// GetAccountTaxUnits gets account.tax.unit existing records.
func (c *Client) GetAccountTaxUnits(ids []int64) (*AccountTaxUnits, error) {
	atus := &AccountTaxUnits{}
	if err := c.Read(AccountTaxUnitModel, ids, nil, atus); err != nil {
		return nil, err
	}
	return atus, nil
}

// FindAccountTaxUnit finds account.tax.unit record by querying it with criteria.
func (c *Client) FindAccountTaxUnit(criteria *Criteria) (*AccountTaxUnit, error) {
	atus := &AccountTaxUnits{}
	if err := c.SearchRead(AccountTaxUnitModel, criteria, NewOptions().Limit(1), atus); err != nil {
		return nil, err
	}
	return &((*atus)[0]), nil
}

// FindAccountTaxUnits finds account.tax.unit records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTaxUnits(criteria *Criteria, options *Options) (*AccountTaxUnits, error) {
	atus := &AccountTaxUnits{}
	if err := c.SearchRead(AccountTaxUnitModel, criteria, options, atus); err != nil {
		return nil, err
	}
	return atus, nil
}

// FindAccountTaxUnitIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTaxUnitIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountTaxUnitModel, criteria, options)
}

// FindAccountTaxUnitId finds record id by querying it with criteria.
func (c *Client) FindAccountTaxUnitId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountTaxUnitModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
