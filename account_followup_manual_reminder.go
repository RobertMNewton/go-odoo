package odoo

// AccountFollowupManualReminder represents account_followup.manual_reminder model.
type AccountFollowupManualReminder struct {
	AttachmentIds        *Relation `xmlrpc:"attachment_ids,omitempty"`
	Body                 *String   `xmlrpc:"body,omitempty"`
	BodyHasTemplateValue *Bool     `xmlrpc:"body_has_template_value,omitempty"`
	CanEditBody          *Bool     `xmlrpc:"can_edit_body,omitempty"`
	CreateDate           *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid            *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName          *String   `xmlrpc:"display_name,omitempty"`
	Email                *Bool     `xmlrpc:"email,omitempty"`
	EmailRecipientIds    *Relation `xmlrpc:"email_recipient_ids,omitempty"`
	Id                   *Int      `xmlrpc:"id,omitempty"`
	IsMailTemplateEditor *Bool     `xmlrpc:"is_mail_template_editor,omitempty"`
	JoinInvoices         *Bool     `xmlrpc:"join_invoices,omitempty"`
	Lang                 *String   `xmlrpc:"lang,omitempty"`
	PartnerId            *Many2One `xmlrpc:"partner_id,omitempty"`
	Print                *Bool     `xmlrpc:"print,omitempty"`
	RenderModel          *String   `xmlrpc:"render_model,omitempty"`
	ShowPrintButton      *Bool     `xmlrpc:"show_print_button,omitempty"`
	ShowSendButton       *Bool     `xmlrpc:"show_send_button,omitempty"`
	Sms                  *Bool     `xmlrpc:"sms,omitempty"`
	SmsBody              *String   `xmlrpc:"sms_body,omitempty"`
	SmsTemplateId        *Many2One `xmlrpc:"sms_template_id,omitempty"`
	Subject              *String   `xmlrpc:"subject,omitempty"`
	TemplateId           *Many2One `xmlrpc:"template_id,omitempty"`
	WriteDate            *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid             *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountFollowupManualReminders represents array of account_followup.manual_reminder model.
type AccountFollowupManualReminders []AccountFollowupManualReminder

// AccountFollowupManualReminderModel is the odoo model name.
const AccountFollowupManualReminderModel = "account_followup.manual_reminder"

// Many2One convert AccountFollowupManualReminder to *Many2One.
func (am *AccountFollowupManualReminder) Many2One() *Many2One {
	return NewMany2One(am.Id.Get(), "")
}

// CreateAccountFollowupManualReminder creates a new account_followup.manual_reminder model and returns its id.
func (c *Client) CreateAccountFollowupManualReminder(am *AccountFollowupManualReminder) (int64, error) {
	ids, err := c.CreateAccountFollowupManualReminders([]*AccountFollowupManualReminder{am})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountFollowupManualReminder creates a new account_followup.manual_reminder model and returns its id.
func (c *Client) CreateAccountFollowupManualReminders(ams []*AccountFollowupManualReminder) ([]int64, error) {
	var vv []interface{}
	for _, v := range ams {
		vv = append(vv, v)
	}
	return c.Create(AccountFollowupManualReminderModel, vv, nil)
}

// UpdateAccountFollowupManualReminder updates an existing account_followup.manual_reminder record.
func (c *Client) UpdateAccountFollowupManualReminder(am *AccountFollowupManualReminder) error {
	return c.UpdateAccountFollowupManualReminders([]int64{am.Id.Get()}, am)
}

// UpdateAccountFollowupManualReminders updates existing account_followup.manual_reminder records.
// All records (represented by ids) will be updated by am values.
func (c *Client) UpdateAccountFollowupManualReminders(ids []int64, am *AccountFollowupManualReminder) error {
	return c.Update(AccountFollowupManualReminderModel, ids, am, nil)
}

// DeleteAccountFollowupManualReminder deletes an existing account_followup.manual_reminder record.
func (c *Client) DeleteAccountFollowupManualReminder(id int64) error {
	return c.DeleteAccountFollowupManualReminders([]int64{id})
}

// DeleteAccountFollowupManualReminders deletes existing account_followup.manual_reminder records.
func (c *Client) DeleteAccountFollowupManualReminders(ids []int64) error {
	return c.Delete(AccountFollowupManualReminderModel, ids)
}

// GetAccountFollowupManualReminder gets account_followup.manual_reminder existing record.
func (c *Client) GetAccountFollowupManualReminder(id int64) (*AccountFollowupManualReminder, error) {
	ams, err := c.GetAccountFollowupManualReminders([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ams)[0]), nil
}

// GetAccountFollowupManualReminders gets account_followup.manual_reminder existing records.
func (c *Client) GetAccountFollowupManualReminders(ids []int64) (*AccountFollowupManualReminders, error) {
	ams := &AccountFollowupManualReminders{}
	if err := c.Read(AccountFollowupManualReminderModel, ids, nil, ams); err != nil {
		return nil, err
	}
	return ams, nil
}

// FindAccountFollowupManualReminder finds account_followup.manual_reminder record by querying it with criteria.
func (c *Client) FindAccountFollowupManualReminder(criteria *Criteria) (*AccountFollowupManualReminder, error) {
	ams := &AccountFollowupManualReminders{}
	if err := c.SearchRead(AccountFollowupManualReminderModel, criteria, NewOptions().Limit(1), ams); err != nil {
		return nil, err
	}
	return &((*ams)[0]), nil
}

// FindAccountFollowupManualReminders finds account_followup.manual_reminder records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFollowupManualReminders(criteria *Criteria, options *Options) (*AccountFollowupManualReminders, error) {
	ams := &AccountFollowupManualReminders{}
	if err := c.SearchRead(AccountFollowupManualReminderModel, criteria, options, ams); err != nil {
		return nil, err
	}
	return ams, nil
}

// FindAccountFollowupManualReminderIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFollowupManualReminderIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountFollowupManualReminderModel, criteria, options)
}

// FindAccountFollowupManualReminderId finds record id by querying it with criteria.
func (c *Client) FindAccountFollowupManualReminderId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountFollowupManualReminderModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
