package odoo

// AccountFollowupFollowupLine represents account_followup.followup.line model.
type AccountFollowupFollowupLine struct {
	ActivityDefaultResponsibleType *Selection `xmlrpc:"activity_default_responsible_type,omitempty"`
	ActivityNote                   *String    `xmlrpc:"activity_note,omitempty"`
	ActivitySummary                *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeId                 *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	AdditionalFollowerIds          *Relation  `xmlrpc:"additional_follower_ids,omitempty"`
	AutoExecute                    *Bool      `xmlrpc:"auto_execute,omitempty"`
	CompanyId                      *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateActivity                 *Bool      `xmlrpc:"create_activity,omitempty"`
	CreateDate                     *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                      *Many2One  `xmlrpc:"create_uid,omitempty"`
	Delay                          *Int       `xmlrpc:"delay,omitempty"`
	DisplayName                    *String    `xmlrpc:"display_name,omitempty"`
	Id                             *Int       `xmlrpc:"id,omitempty"`
	JoinInvoices                   *Bool      `xmlrpc:"join_invoices,omitempty"`
	MailTemplateId                 *Many2One  `xmlrpc:"mail_template_id,omitempty"`
	Name                           *String    `xmlrpc:"name,omitempty"`
	SendEmail                      *Bool      `xmlrpc:"send_email,omitempty"`
	SendSms                        *Bool      `xmlrpc:"send_sms,omitempty"`
	SmsTemplateId                  *Many2One  `xmlrpc:"sms_template_id,omitempty"`
	WriteDate                      *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                       *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountFollowupFollowupLines represents array of account_followup.followup.line model.
type AccountFollowupFollowupLines []AccountFollowupFollowupLine

// AccountFollowupFollowupLineModel is the odoo model name.
const AccountFollowupFollowupLineModel = "account_followup.followup.line"

// Many2One convert AccountFollowupFollowupLine to *Many2One.
func (afl *AccountFollowupFollowupLine) Many2One() *Many2One {
	return NewMany2One(afl.Id.Get(), "")
}

// CreateAccountFollowupFollowupLine creates a new account_followup.followup.line model and returns its id.
func (c *Client) CreateAccountFollowupFollowupLine(afl *AccountFollowupFollowupLine) (int64, error) {
	ids, err := c.CreateAccountFollowupFollowupLines([]*AccountFollowupFollowupLine{afl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountFollowupFollowupLine creates a new account_followup.followup.line model and returns its id.
func (c *Client) CreateAccountFollowupFollowupLines(afls []*AccountFollowupFollowupLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range afls {
		vv = append(vv, v)
	}
	return c.Create(AccountFollowupFollowupLineModel, vv, nil)
}

// UpdateAccountFollowupFollowupLine updates an existing account_followup.followup.line record.
func (c *Client) UpdateAccountFollowupFollowupLine(afl *AccountFollowupFollowupLine) error {
	return c.UpdateAccountFollowupFollowupLines([]int64{afl.Id.Get()}, afl)
}

// UpdateAccountFollowupFollowupLines updates existing account_followup.followup.line records.
// All records (represented by ids) will be updated by afl values.
func (c *Client) UpdateAccountFollowupFollowupLines(ids []int64, afl *AccountFollowupFollowupLine) error {
	return c.Update(AccountFollowupFollowupLineModel, ids, afl, nil)
}

// DeleteAccountFollowupFollowupLine deletes an existing account_followup.followup.line record.
func (c *Client) DeleteAccountFollowupFollowupLine(id int64) error {
	return c.DeleteAccountFollowupFollowupLines([]int64{id})
}

// DeleteAccountFollowupFollowupLines deletes existing account_followup.followup.line records.
func (c *Client) DeleteAccountFollowupFollowupLines(ids []int64) error {
	return c.Delete(AccountFollowupFollowupLineModel, ids)
}

// GetAccountFollowupFollowupLine gets account_followup.followup.line existing record.
func (c *Client) GetAccountFollowupFollowupLine(id int64) (*AccountFollowupFollowupLine, error) {
	afls, err := c.GetAccountFollowupFollowupLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*afls)[0]), nil
}

// GetAccountFollowupFollowupLines gets account_followup.followup.line existing records.
func (c *Client) GetAccountFollowupFollowupLines(ids []int64) (*AccountFollowupFollowupLines, error) {
	afls := &AccountFollowupFollowupLines{}
	if err := c.Read(AccountFollowupFollowupLineModel, ids, nil, afls); err != nil {
		return nil, err
	}
	return afls, nil
}

// FindAccountFollowupFollowupLine finds account_followup.followup.line record by querying it with criteria.
func (c *Client) FindAccountFollowupFollowupLine(criteria *Criteria) (*AccountFollowupFollowupLine, error) {
	afls := &AccountFollowupFollowupLines{}
	if err := c.SearchRead(AccountFollowupFollowupLineModel, criteria, NewOptions().Limit(1), afls); err != nil {
		return nil, err
	}
	return &((*afls)[0]), nil
}

// FindAccountFollowupFollowupLines finds account_followup.followup.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFollowupFollowupLines(criteria *Criteria, options *Options) (*AccountFollowupFollowupLines, error) {
	afls := &AccountFollowupFollowupLines{}
	if err := c.SearchRead(AccountFollowupFollowupLineModel, criteria, options, afls); err != nil {
		return nil, err
	}
	return afls, nil
}

// FindAccountFollowupFollowupLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFollowupFollowupLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountFollowupFollowupLineModel, criteria, options)
}

// FindAccountFollowupFollowupLineId finds record id by querying it with criteria.
func (c *Client) FindAccountFollowupFollowupLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountFollowupFollowupLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
