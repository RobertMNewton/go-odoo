package odoo

// SurveyInvite represents survey.invite model.
type SurveyInvite struct {
	AttachmentIds            *Relation  `xmlrpc:"attachment_ids,omitempty"`
	AuthorId                 *Many2One  `xmlrpc:"author_id,omitempty"`
	Body                     *String    `xmlrpc:"body,omitempty"`
	BodyHasTemplateValue     *Bool      `xmlrpc:"body_has_template_value,omitempty"`
	CanEditBody              *Bool      `xmlrpc:"can_edit_body,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty"`
	Deadline                 *Time      `xmlrpc:"deadline,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty"`
	Emails                   *String    `xmlrpc:"emails,omitempty"`
	ExistingEmails           *String    `xmlrpc:"existing_emails,omitempty"`
	ExistingMode             *Selection `xmlrpc:"existing_mode,omitempty"`
	ExistingPartnerIds       *Relation  `xmlrpc:"existing_partner_ids,omitempty"`
	ExistingText             *String    `xmlrpc:"existing_text,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty"`
	IsMailTemplateEditor     *Bool      `xmlrpc:"is_mail_template_editor,omitempty"`
	Lang                     *String    `xmlrpc:"lang,omitempty"`
	MailServerId             *Many2One  `xmlrpc:"mail_server_id,omitempty"`
	PartnerIds               *Relation  `xmlrpc:"partner_ids,omitempty"`
	RenderModel              *String    `xmlrpc:"render_model,omitempty"`
	SendEmail                *Bool      `xmlrpc:"send_email,omitempty"`
	Subject                  *String    `xmlrpc:"subject,omitempty"`
	SurveyAccessMode         *Selection `xmlrpc:"survey_access_mode,omitempty"`
	SurveyId                 *Many2One  `xmlrpc:"survey_id,omitempty"`
	SurveyStartUrl           *String    `xmlrpc:"survey_start_url,omitempty"`
	SurveyUsersCanSignup     *Bool      `xmlrpc:"survey_users_can_signup,omitempty"`
	SurveyUsersLoginRequired *Bool      `xmlrpc:"survey_users_login_required,omitempty"`
	TemplateId               *Many2One  `xmlrpc:"template_id,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// SurveyInvites represents array of survey.invite model.
type SurveyInvites []SurveyInvite

// SurveyInviteModel is the odoo model name.
const SurveyInviteModel = "survey.invite"

// Many2One convert SurveyInvite to *Many2One.
func (si *SurveyInvite) Many2One() *Many2One {
	return NewMany2One(si.Id.Get(), "")
}

// CreateSurveyInvite creates a new survey.invite model and returns its id.
func (c *Client) CreateSurveyInvite(si *SurveyInvite) (int64, error) {
	ids, err := c.CreateSurveyInvites([]*SurveyInvite{si})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSurveyInvite creates a new survey.invite model and returns its id.
func (c *Client) CreateSurveyInvites(sis []*SurveyInvite) ([]int64, error) {
	var vv []interface{}
	for _, v := range sis {
		vv = append(vv, v)
	}
	return c.Create(SurveyInviteModel, vv, nil)
}

// UpdateSurveyInvite updates an existing survey.invite record.
func (c *Client) UpdateSurveyInvite(si *SurveyInvite) error {
	return c.UpdateSurveyInvites([]int64{si.Id.Get()}, si)
}

// UpdateSurveyInvites updates existing survey.invite records.
// All records (represented by ids) will be updated by si values.
func (c *Client) UpdateSurveyInvites(ids []int64, si *SurveyInvite) error {
	return c.Update(SurveyInviteModel, ids, si, nil)
}

// DeleteSurveyInvite deletes an existing survey.invite record.
func (c *Client) DeleteSurveyInvite(id int64) error {
	return c.DeleteSurveyInvites([]int64{id})
}

// DeleteSurveyInvites deletes existing survey.invite records.
func (c *Client) DeleteSurveyInvites(ids []int64) error {
	return c.Delete(SurveyInviteModel, ids)
}

// GetSurveyInvite gets survey.invite existing record.
func (c *Client) GetSurveyInvite(id int64) (*SurveyInvite, error) {
	sis, err := c.GetSurveyInvites([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sis)[0]), nil
}

// GetSurveyInvites gets survey.invite existing records.
func (c *Client) GetSurveyInvites(ids []int64) (*SurveyInvites, error) {
	sis := &SurveyInvites{}
	if err := c.Read(SurveyInviteModel, ids, nil, sis); err != nil {
		return nil, err
	}
	return sis, nil
}

// FindSurveyInvite finds survey.invite record by querying it with criteria.
func (c *Client) FindSurveyInvite(criteria *Criteria) (*SurveyInvite, error) {
	sis := &SurveyInvites{}
	if err := c.SearchRead(SurveyInviteModel, criteria, NewOptions().Limit(1), sis); err != nil {
		return nil, err
	}
	return &((*sis)[0]), nil
}

// FindSurveyInvites finds survey.invite records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSurveyInvites(criteria *Criteria, options *Options) (*SurveyInvites, error) {
	sis := &SurveyInvites{}
	if err := c.SearchRead(SurveyInviteModel, criteria, options, sis); err != nil {
		return nil, err
	}
	return sis, nil
}

// FindSurveyInviteIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSurveyInviteIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SurveyInviteModel, criteria, options)
}

// FindSurveyInviteId finds record id by querying it with criteria.
func (c *Client) FindSurveyInviteId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SurveyInviteModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
