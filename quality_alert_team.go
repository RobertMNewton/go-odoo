package odoo

// QualityAlertTeam represents quality.alert.team model.
type QualityAlertTeam struct {
	AlertCount               *Int       `xmlrpc:"alert_count,omitempty"`
	AliasBouncedContent      *String    `xmlrpc:"alias_bounced_content,omitempty"`
	AliasContact             *Selection `xmlrpc:"alias_contact,omitempty"`
	AliasDefaults            *String    `xmlrpc:"alias_defaults,omitempty"`
	AliasDomain              *String    `xmlrpc:"alias_domain,omitempty"`
	AliasDomainId            *Many2One  `xmlrpc:"alias_domain_id,omitempty"`
	AliasEmail               *String    `xmlrpc:"alias_email,omitempty"`
	AliasForceThreadId       *Int       `xmlrpc:"alias_force_thread_id,omitempty"`
	AliasFullName            *String    `xmlrpc:"alias_full_name,omitempty"`
	AliasId                  *Many2One  `xmlrpc:"alias_id,omitempty"`
	AliasIncomingLocal       *Bool      `xmlrpc:"alias_incoming_local,omitempty"`
	AliasModelId             *Many2One  `xmlrpc:"alias_model_id,omitempty"`
	AliasName                *String    `xmlrpc:"alias_name,omitempty"`
	AliasParentModelId       *Many2One  `xmlrpc:"alias_parent_model_id,omitempty"`
	AliasParentThreadId      *Int       `xmlrpc:"alias_parent_thread_id,omitempty"`
	AliasStatus              *Selection `xmlrpc:"alias_status,omitempty"`
	CheckCount               *Int       `xmlrpc:"check_count,omitempty"`
	Color                    *Int       `xmlrpc:"color,omitempty"`
	CompanyId                *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty"`
	HasMessage               *Bool      `xmlrpc:"has_message,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty"`
	MessageAttachmentCount   *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds       *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError          *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter   *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError       *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds               *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower        *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction        *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	Name                     *String    `xmlrpc:"name,omitempty"`
	RatingIds                *Relation  `xmlrpc:"rating_ids,omitempty"`
	Sequence                 *Int       `xmlrpc:"sequence,omitempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// QualityAlertTeams represents array of quality.alert.team model.
type QualityAlertTeams []QualityAlertTeam

// QualityAlertTeamModel is the odoo model name.
const QualityAlertTeamModel = "quality.alert.team"

// Many2One convert QualityAlertTeam to *Many2One.
func (qat *QualityAlertTeam) Many2One() *Many2One {
	return NewMany2One(qat.Id.Get(), "")
}

// CreateQualityAlertTeam creates a new quality.alert.team model and returns its id.
func (c *Client) CreateQualityAlertTeam(qat *QualityAlertTeam) (int64, error) {
	ids, err := c.CreateQualityAlertTeams([]*QualityAlertTeam{qat})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateQualityAlertTeam creates a new quality.alert.team model and returns its id.
func (c *Client) CreateQualityAlertTeams(qats []*QualityAlertTeam) ([]int64, error) {
	var vv []interface{}
	for _, v := range qats {
		vv = append(vv, v)
	}
	return c.Create(QualityAlertTeamModel, vv, nil)
}

// UpdateQualityAlertTeam updates an existing quality.alert.team record.
func (c *Client) UpdateQualityAlertTeam(qat *QualityAlertTeam) error {
	return c.UpdateQualityAlertTeams([]int64{qat.Id.Get()}, qat)
}

// UpdateQualityAlertTeams updates existing quality.alert.team records.
// All records (represented by ids) will be updated by qat values.
func (c *Client) UpdateQualityAlertTeams(ids []int64, qat *QualityAlertTeam) error {
	return c.Update(QualityAlertTeamModel, ids, qat, nil)
}

// DeleteQualityAlertTeam deletes an existing quality.alert.team record.
func (c *Client) DeleteQualityAlertTeam(id int64) error {
	return c.DeleteQualityAlertTeams([]int64{id})
}

// DeleteQualityAlertTeams deletes existing quality.alert.team records.
func (c *Client) DeleteQualityAlertTeams(ids []int64) error {
	return c.Delete(QualityAlertTeamModel, ids)
}

// GetQualityAlertTeam gets quality.alert.team existing record.
func (c *Client) GetQualityAlertTeam(id int64) (*QualityAlertTeam, error) {
	qats, err := c.GetQualityAlertTeams([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*qats)[0]), nil
}

// GetQualityAlertTeams gets quality.alert.team existing records.
func (c *Client) GetQualityAlertTeams(ids []int64) (*QualityAlertTeams, error) {
	qats := &QualityAlertTeams{}
	if err := c.Read(QualityAlertTeamModel, ids, nil, qats); err != nil {
		return nil, err
	}
	return qats, nil
}

// FindQualityAlertTeam finds quality.alert.team record by querying it with criteria.
func (c *Client) FindQualityAlertTeam(criteria *Criteria) (*QualityAlertTeam, error) {
	qats := &QualityAlertTeams{}
	if err := c.SearchRead(QualityAlertTeamModel, criteria, NewOptions().Limit(1), qats); err != nil {
		return nil, err
	}
	return &((*qats)[0]), nil
}

// FindQualityAlertTeams finds quality.alert.team records by querying it
// and filtering it with criteria and options.
func (c *Client) FindQualityAlertTeams(criteria *Criteria, options *Options) (*QualityAlertTeams, error) {
	qats := &QualityAlertTeams{}
	if err := c.SearchRead(QualityAlertTeamModel, criteria, options, qats); err != nil {
		return nil, err
	}
	return qats, nil
}

// FindQualityAlertTeamIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindQualityAlertTeamIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(QualityAlertTeamModel, criteria, options)
}

// FindQualityAlertTeamId finds record id by querying it with criteria.
func (c *Client) FindQualityAlertTeamId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(QualityAlertTeamModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
