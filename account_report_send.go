package odoo

// AccountReportSend represents account.report.send model.
type AccountReportSend struct {
	AccountReportId       *Many2One   `xmlrpc:"account_report_id,omitempty"`
	CheckboxDownload      *Bool       `xmlrpc:"checkbox_download,omitempty"`
	CheckboxSendMail      *Bool       `xmlrpc:"checkbox_send_mail,omitempty"`
	CreateDate            *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid             *Many2One   `xmlrpc:"create_uid,omitempty"`
	DisplayMailComposer   *Bool       `xmlrpc:"display_mail_composer,omitempty"`
	DisplayName           *String     `xmlrpc:"display_name,omitempty"`
	EnableDownload        *Bool       `xmlrpc:"enable_download,omitempty"`
	EnableSendMail        *Bool       `xmlrpc:"enable_send_mail,omitempty"`
	Id                    *Int        `xmlrpc:"id,omitempty"`
	MailAttachmentsWidget interface{} `xmlrpc:"mail_attachments_widget,omitempty"`
	MailBody              *String     `xmlrpc:"mail_body,omitempty"`
	MailLang              *String     `xmlrpc:"mail_lang,omitempty"`
	MailPartnerIds        *Relation   `xmlrpc:"mail_partner_ids,omitempty"`
	MailSubject           *String     `xmlrpc:"mail_subject,omitempty"`
	MailTemplateId        *Many2One   `xmlrpc:"mail_template_id,omitempty"`
	Mode                  *Selection  `xmlrpc:"mode,omitempty"`
	PartnerIds            *Relation   `xmlrpc:"partner_ids,omitempty"`
	ReportOptions         interface{} `xmlrpc:"report_options,omitempty"`
	SendMailReadonly      *Bool       `xmlrpc:"send_mail_readonly,omitempty"`
	Warnings              interface{} `xmlrpc:"warnings,omitempty"`
	WriteDate             *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid              *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// AccountReportSends represents array of account.report.send model.
type AccountReportSends []AccountReportSend

// AccountReportSendModel is the odoo model name.
const AccountReportSendModel = "account.report.send"

// Many2One convert AccountReportSend to *Many2One.
func (ars *AccountReportSend) Many2One() *Many2One {
	return NewMany2One(ars.Id.Get(), "")
}

// CreateAccountReportSend creates a new account.report.send model and returns its id.
func (c *Client) CreateAccountReportSend(ars *AccountReportSend) (int64, error) {
	ids, err := c.CreateAccountReportSends([]*AccountReportSend{ars})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountReportSend creates a new account.report.send model and returns its id.
func (c *Client) CreateAccountReportSends(arss []*AccountReportSend) ([]int64, error) {
	var vv []interface{}
	for _, v := range arss {
		vv = append(vv, v)
	}
	return c.Create(AccountReportSendModel, vv, nil)
}

// UpdateAccountReportSend updates an existing account.report.send record.
func (c *Client) UpdateAccountReportSend(ars *AccountReportSend) error {
	return c.UpdateAccountReportSends([]int64{ars.Id.Get()}, ars)
}

// UpdateAccountReportSends updates existing account.report.send records.
// All records (represented by ids) will be updated by ars values.
func (c *Client) UpdateAccountReportSends(ids []int64, ars *AccountReportSend) error {
	return c.Update(AccountReportSendModel, ids, ars, nil)
}

// DeleteAccountReportSend deletes an existing account.report.send record.
func (c *Client) DeleteAccountReportSend(id int64) error {
	return c.DeleteAccountReportSends([]int64{id})
}

// DeleteAccountReportSends deletes existing account.report.send records.
func (c *Client) DeleteAccountReportSends(ids []int64) error {
	return c.Delete(AccountReportSendModel, ids)
}

// GetAccountReportSend gets account.report.send existing record.
func (c *Client) GetAccountReportSend(id int64) (*AccountReportSend, error) {
	arss, err := c.GetAccountReportSends([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*arss)[0]), nil
}

// GetAccountReportSends gets account.report.send existing records.
func (c *Client) GetAccountReportSends(ids []int64) (*AccountReportSends, error) {
	arss := &AccountReportSends{}
	if err := c.Read(AccountReportSendModel, ids, nil, arss); err != nil {
		return nil, err
	}
	return arss, nil
}

// FindAccountReportSend finds account.report.send record by querying it with criteria.
func (c *Client) FindAccountReportSend(criteria *Criteria) (*AccountReportSend, error) {
	arss := &AccountReportSends{}
	if err := c.SearchRead(AccountReportSendModel, criteria, NewOptions().Limit(1), arss); err != nil {
		return nil, err
	}
	return &((*arss)[0]), nil
}

// FindAccountReportSends finds account.report.send records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportSends(criteria *Criteria, options *Options) (*AccountReportSends, error) {
	arss := &AccountReportSends{}
	if err := c.SearchRead(AccountReportSendModel, criteria, options, arss); err != nil {
		return nil, err
	}
	return arss, nil
}

// FindAccountReportSendIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportSendIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountReportSendModel, criteria, options)
}

// FindAccountReportSendId finds record id by querying it with criteria.
func (c *Client) FindAccountReportSendId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReportSendModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
