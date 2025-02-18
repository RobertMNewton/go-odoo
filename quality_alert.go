package odoo

// QualityAlert represents quality.alert model.
type QualityAlert struct {
	ActionCorrective            *String    `xmlrpc:"action_corrective,omitempty"`
	ActionPreventive            *String    `xmlrpc:"action_preventive,omitempty"`
	ActivityCalendarEventId     *Many2One  `xmlrpc:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon            *String    `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	CheckId                     *Many2One  `xmlrpc:"check_id,omitempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DateAssign                  *Time      `xmlrpc:"date_assign,omitempty"`
	DateClose                   *Time      `xmlrpc:"date_close,omitempty"`
	Description                 *String    `xmlrpc:"description,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	EmailCc                     *String    `xmlrpc:"email_cc,omitempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	LotId                       *Many2One  `xmlrpc:"lot_id,omitempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MyActivityDateDeadline      *Time      `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                        *String    `xmlrpc:"name,omitempty"`
	PartnerId                   *Many2One  `xmlrpc:"partner_id,omitempty"`
	PickingId                   *Many2One  `xmlrpc:"picking_id,omitempty"`
	Priority                    *Selection `xmlrpc:"priority,omitempty"`
	ProductId                   *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductTmplId               *Many2One  `xmlrpc:"product_tmpl_id,omitempty"`
	ProductionId                *Many2One  `xmlrpc:"production_id,omitempty"`
	RatingIds                   *Relation  `xmlrpc:"rating_ids,omitempty"`
	ReasonId                    *Many2One  `xmlrpc:"reason_id,omitempty"`
	StageId                     *Many2One  `xmlrpc:"stage_id,omitempty"`
	TagIds                      *Relation  `xmlrpc:"tag_ids,omitempty"`
	TeamId                      *Many2One  `xmlrpc:"team_id,omitempty"`
	Title                       *String    `xmlrpc:"title,omitempty"`
	UserId                      *Many2One  `xmlrpc:"user_id,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WorkcenterId                *Many2One  `xmlrpc:"workcenter_id,omitempty"`
	WorkorderId                 *Many2One  `xmlrpc:"workorder_id,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// QualityAlerts represents array of quality.alert model.
type QualityAlerts []QualityAlert

// QualityAlertModel is the odoo model name.
const QualityAlertModel = "quality.alert"

// Many2One convert QualityAlert to *Many2One.
func (qa *QualityAlert) Many2One() *Many2One {
	return NewMany2One(qa.Id.Get(), "")
}

// CreateQualityAlert creates a new quality.alert model and returns its id.
func (c *Client) CreateQualityAlert(qa *QualityAlert) (int64, error) {
	ids, err := c.CreateQualityAlerts([]*QualityAlert{qa})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateQualityAlert creates a new quality.alert model and returns its id.
func (c *Client) CreateQualityAlerts(qas []*QualityAlert) ([]int64, error) {
	var vv []interface{}
	for _, v := range qas {
		vv = append(vv, v)
	}
	return c.Create(QualityAlertModel, vv, nil)
}

// UpdateQualityAlert updates an existing quality.alert record.
func (c *Client) UpdateQualityAlert(qa *QualityAlert) error {
	return c.UpdateQualityAlerts([]int64{qa.Id.Get()}, qa)
}

// UpdateQualityAlerts updates existing quality.alert records.
// All records (represented by ids) will be updated by qa values.
func (c *Client) UpdateQualityAlerts(ids []int64, qa *QualityAlert) error {
	return c.Update(QualityAlertModel, ids, qa, nil)
}

// DeleteQualityAlert deletes an existing quality.alert record.
func (c *Client) DeleteQualityAlert(id int64) error {
	return c.DeleteQualityAlerts([]int64{id})
}

// DeleteQualityAlerts deletes existing quality.alert records.
func (c *Client) DeleteQualityAlerts(ids []int64) error {
	return c.Delete(QualityAlertModel, ids)
}

// GetQualityAlert gets quality.alert existing record.
func (c *Client) GetQualityAlert(id int64) (*QualityAlert, error) {
	qas, err := c.GetQualityAlerts([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*qas)[0]), nil
}

// GetQualityAlerts gets quality.alert existing records.
func (c *Client) GetQualityAlerts(ids []int64) (*QualityAlerts, error) {
	qas := &QualityAlerts{}
	if err := c.Read(QualityAlertModel, ids, nil, qas); err != nil {
		return nil, err
	}
	return qas, nil
}

// FindQualityAlert finds quality.alert record by querying it with criteria.
func (c *Client) FindQualityAlert(criteria *Criteria) (*QualityAlert, error) {
	qas := &QualityAlerts{}
	if err := c.SearchRead(QualityAlertModel, criteria, NewOptions().Limit(1), qas); err != nil {
		return nil, err
	}
	return &((*qas)[0]), nil
}

// FindQualityAlerts finds quality.alert records by querying it
// and filtering it with criteria and options.
func (c *Client) FindQualityAlerts(criteria *Criteria, options *Options) (*QualityAlerts, error) {
	qas := &QualityAlerts{}
	if err := c.SearchRead(QualityAlertModel, criteria, options, qas); err != nil {
		return nil, err
	}
	return qas, nil
}

// FindQualityAlertIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindQualityAlertIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(QualityAlertModel, criteria, options)
}

// FindQualityAlertId finds record id by querying it with criteria.
func (c *Client) FindQualityAlertId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(QualityAlertModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
