package odoo

// AccountChangeLockDate represents account.change.lock.date model.
type AccountChangeLockDate struct {
	CompanyId                                   *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrentHardLockDate                         *Time      `xmlrpc:"current_hard_lock_date,omitempty"`
	DisplayName                                 *String    `xmlrpc:"display_name,omitempty"`
	ExceptionAppliesTo                          *Selection `xmlrpc:"exception_applies_to,omitempty"`
	ExceptionDuration                           *Selection `xmlrpc:"exception_duration,omitempty"`
	ExceptionNeededFields                       *String    `xmlrpc:"exception_needed_fields,omitempty"`
	ExceptionReason                             *String    `xmlrpc:"exception_reason,omitempty"`
	FiscalyearLockDate                          *Time      `xmlrpc:"fiscalyear_lock_date,omitempty"`
	FiscalyearLockDateForEveryone               *Time      `xmlrpc:"fiscalyear_lock_date_for_everyone,omitempty"`
	FiscalyearLockDateForMe                     *Time      `xmlrpc:"fiscalyear_lock_date_for_me,omitempty"`
	HardLockDate                                *Time      `xmlrpc:"hard_lock_date,omitempty"`
	Id                                          *Int       `xmlrpc:"id,omitempty"`
	MinFiscalyearLockDateExceptionForEveryoneId *Many2One  `xmlrpc:"min_fiscalyear_lock_date_exception_for_everyone_id,omitempty"`
	MinFiscalyearLockDateExceptionForMeId       *Many2One  `xmlrpc:"min_fiscalyear_lock_date_exception_for_me_id,omitempty"`
	MinPurchaseLockDateExceptionForEveryoneId   *Many2One  `xmlrpc:"min_purchase_lock_date_exception_for_everyone_id,omitempty"`
	MinPurchaseLockDateExceptionForMeId         *Many2One  `xmlrpc:"min_purchase_lock_date_exception_for_me_id,omitempty"`
	MinSaleLockDateExceptionForEveryoneId       *Many2One  `xmlrpc:"min_sale_lock_date_exception_for_everyone_id,omitempty"`
	MinSaleLockDateExceptionForMeId             *Many2One  `xmlrpc:"min_sale_lock_date_exception_for_me_id,omitempty"`
	MinTaxLockDateExceptionForEveryoneId        *Many2One  `xmlrpc:"min_tax_lock_date_exception_for_everyone_id,omitempty"`
	MinTaxLockDateExceptionForMeId              *Many2One  `xmlrpc:"min_tax_lock_date_exception_for_me_id,omitempty"`
	PurchaseLockDate                            *Time      `xmlrpc:"purchase_lock_date,omitempty"`
	PurchaseLockDateForEveryone                 *Time      `xmlrpc:"purchase_lock_date_for_everyone,omitempty"`
	PurchaseLockDateForMe                       *Time      `xmlrpc:"purchase_lock_date_for_me,omitempty"`
	SaleLockDate                                *Time      `xmlrpc:"sale_lock_date,omitempty"`
	SaleLockDateForEveryone                     *Time      `xmlrpc:"sale_lock_date_for_everyone,omitempty"`
	SaleLockDateForMe                           *Time      `xmlrpc:"sale_lock_date_for_me,omitempty"`
	ShowDraftEntriesWarning                     *Bool      `xmlrpc:"show_draft_entries_warning,omitempty"`
	TaxLockDate                                 *Time      `xmlrpc:"tax_lock_date,omitempty"`
	TaxLockDateForEveryone                      *Time      `xmlrpc:"tax_lock_date_for_everyone,omitempty"`
	TaxLockDateForMe                            *Time      `xmlrpc:"tax_lock_date_for_me,omitempty"`
	WriteDate                                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountChangeLockDates represents array of account.change.lock.date model.
type AccountChangeLockDates []AccountChangeLockDate

// AccountChangeLockDateModel is the odoo model name.
const AccountChangeLockDateModel = "account.change.lock.date"

// Many2One convert AccountChangeLockDate to *Many2One.
func (acld *AccountChangeLockDate) Many2One() *Many2One {
	return NewMany2One(acld.Id.Get(), "")
}

// CreateAccountChangeLockDate creates a new account.change.lock.date model and returns its id.
func (c *Client) CreateAccountChangeLockDate(acld *AccountChangeLockDate) (int64, error) {
	ids, err := c.CreateAccountChangeLockDates([]*AccountChangeLockDate{acld})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountChangeLockDate creates a new account.change.lock.date model and returns its id.
func (c *Client) CreateAccountChangeLockDates(aclds []*AccountChangeLockDate) ([]int64, error) {
	var vv []interface{}
	for _, v := range aclds {
		vv = append(vv, v)
	}
	return c.Create(AccountChangeLockDateModel, vv, nil)
}

// UpdateAccountChangeLockDate updates an existing account.change.lock.date record.
func (c *Client) UpdateAccountChangeLockDate(acld *AccountChangeLockDate) error {
	return c.UpdateAccountChangeLockDates([]int64{acld.Id.Get()}, acld)
}

// UpdateAccountChangeLockDates updates existing account.change.lock.date records.
// All records (represented by ids) will be updated by acld values.
func (c *Client) UpdateAccountChangeLockDates(ids []int64, acld *AccountChangeLockDate) error {
	return c.Update(AccountChangeLockDateModel, ids, acld, nil)
}

// DeleteAccountChangeLockDate deletes an existing account.change.lock.date record.
func (c *Client) DeleteAccountChangeLockDate(id int64) error {
	return c.DeleteAccountChangeLockDates([]int64{id})
}

// DeleteAccountChangeLockDates deletes existing account.change.lock.date records.
func (c *Client) DeleteAccountChangeLockDates(ids []int64) error {
	return c.Delete(AccountChangeLockDateModel, ids)
}

// GetAccountChangeLockDate gets account.change.lock.date existing record.
func (c *Client) GetAccountChangeLockDate(id int64) (*AccountChangeLockDate, error) {
	aclds, err := c.GetAccountChangeLockDates([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aclds)[0]), nil
}

// GetAccountChangeLockDates gets account.change.lock.date existing records.
func (c *Client) GetAccountChangeLockDates(ids []int64) (*AccountChangeLockDates, error) {
	aclds := &AccountChangeLockDates{}
	if err := c.Read(AccountChangeLockDateModel, ids, nil, aclds); err != nil {
		return nil, err
	}
	return aclds, nil
}

// FindAccountChangeLockDate finds account.change.lock.date record by querying it with criteria.
func (c *Client) FindAccountChangeLockDate(criteria *Criteria) (*AccountChangeLockDate, error) {
	aclds := &AccountChangeLockDates{}
	if err := c.SearchRead(AccountChangeLockDateModel, criteria, NewOptions().Limit(1), aclds); err != nil {
		return nil, err
	}
	return &((*aclds)[0]), nil
}

// FindAccountChangeLockDates finds account.change.lock.date records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountChangeLockDates(criteria *Criteria, options *Options) (*AccountChangeLockDates, error) {
	aclds := &AccountChangeLockDates{}
	if err := c.SearchRead(AccountChangeLockDateModel, criteria, options, aclds); err != nil {
		return nil, err
	}
	return aclds, nil
}

// FindAccountChangeLockDateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountChangeLockDateIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountChangeLockDateModel, criteria, options)
}

// FindAccountChangeLockDateId finds record id by querying it with criteria.
func (c *Client) FindAccountChangeLockDateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountChangeLockDateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
