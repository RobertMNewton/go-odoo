package odoo

// AccountAssetGroup represents account.asset.group model.
type AccountAssetGroup struct {
	CompanyId         *Many2One `xmlrpc:"company_id,omitempty"`
	CountLinkedAssets *Int      `xmlrpc:"count_linked_assets,omitempty"`
	CountLinkedLoans  *Int      `xmlrpc:"count_linked_loans,omitempty"`
	CreateDate        *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName       *String   `xmlrpc:"display_name,omitempty"`
	Id                *Int      `xmlrpc:"id,omitempty"`
	LinkedAssetIds    *Relation `xmlrpc:"linked_asset_ids,omitempty"`
	LinkedLoanIds     *Relation `xmlrpc:"linked_loan_ids,omitempty"`
	Name              *String   `xmlrpc:"name,omitempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountAssetGroups represents array of account.asset.group model.
type AccountAssetGroups []AccountAssetGroup

// AccountAssetGroupModel is the odoo model name.
const AccountAssetGroupModel = "account.asset.group"

// Many2One convert AccountAssetGroup to *Many2One.
func (aag *AccountAssetGroup) Many2One() *Many2One {
	return NewMany2One(aag.Id.Get(), "")
}

// CreateAccountAssetGroup creates a new account.asset.group model and returns its id.
func (c *Client) CreateAccountAssetGroup(aag *AccountAssetGroup) (int64, error) {
	ids, err := c.CreateAccountAssetGroups([]*AccountAssetGroup{aag})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountAssetGroup creates a new account.asset.group model and returns its id.
func (c *Client) CreateAccountAssetGroups(aags []*AccountAssetGroup) ([]int64, error) {
	var vv []interface{}
	for _, v := range aags {
		vv = append(vv, v)
	}
	return c.Create(AccountAssetGroupModel, vv, nil)
}

// UpdateAccountAssetGroup updates an existing account.asset.group record.
func (c *Client) UpdateAccountAssetGroup(aag *AccountAssetGroup) error {
	return c.UpdateAccountAssetGroups([]int64{aag.Id.Get()}, aag)
}

// UpdateAccountAssetGroups updates existing account.asset.group records.
// All records (represented by ids) will be updated by aag values.
func (c *Client) UpdateAccountAssetGroups(ids []int64, aag *AccountAssetGroup) error {
	return c.Update(AccountAssetGroupModel, ids, aag, nil)
}

// DeleteAccountAssetGroup deletes an existing account.asset.group record.
func (c *Client) DeleteAccountAssetGroup(id int64) error {
	return c.DeleteAccountAssetGroups([]int64{id})
}

// DeleteAccountAssetGroups deletes existing account.asset.group records.
func (c *Client) DeleteAccountAssetGroups(ids []int64) error {
	return c.Delete(AccountAssetGroupModel, ids)
}

// GetAccountAssetGroup gets account.asset.group existing record.
func (c *Client) GetAccountAssetGroup(id int64) (*AccountAssetGroup, error) {
	aags, err := c.GetAccountAssetGroups([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aags)[0]), nil
}

// GetAccountAssetGroups gets account.asset.group existing records.
func (c *Client) GetAccountAssetGroups(ids []int64) (*AccountAssetGroups, error) {
	aags := &AccountAssetGroups{}
	if err := c.Read(AccountAssetGroupModel, ids, nil, aags); err != nil {
		return nil, err
	}
	return aags, nil
}

// FindAccountAssetGroup finds account.asset.group record by querying it with criteria.
func (c *Client) FindAccountAssetGroup(criteria *Criteria) (*AccountAssetGroup, error) {
	aags := &AccountAssetGroups{}
	if err := c.SearchRead(AccountAssetGroupModel, criteria, NewOptions().Limit(1), aags); err != nil {
		return nil, err
	}
	return &((*aags)[0]), nil
}

// FindAccountAssetGroups finds account.asset.group records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAssetGroups(criteria *Criteria, options *Options) (*AccountAssetGroups, error) {
	aags := &AccountAssetGroups{}
	if err := c.SearchRead(AccountAssetGroupModel, criteria, options, aags); err != nil {
		return nil, err
	}
	return aags, nil
}

// FindAccountAssetGroupIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAssetGroupIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountAssetGroupModel, criteria, options)
}

// FindAccountAssetGroupId finds record id by querying it with criteria.
func (c *Client) FindAccountAssetGroupId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAssetGroupModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
