package odoo

// GamificationChallenge represents gamification.challenge model.
type GamificationChallenge struct {
	ChallengeCategory        *Selection `xmlrpc:"challenge_category,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty"`
	Description              *String    `xmlrpc:"description,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty"`
	EndDate                  *Time      `xmlrpc:"end_date,omitempty"`
	HasMessage               *Bool      `xmlrpc:"has_message,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty"`
	InvitedUserIds           *Relation  `xmlrpc:"invited_user_ids,omitempty"`
	LastReportDate           *Time      `xmlrpc:"last_report_date,omitempty"`
	LineIds                  *Relation  `xmlrpc:"line_ids,omitempty"`
	ManagerId                *Many2One  `xmlrpc:"manager_id,omitempty"`
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
	NextReportDate           *Time      `xmlrpc:"next_report_date,omitempty"`
	Period                   *Selection `xmlrpc:"period,omitempty"`
	RatingIds                *Relation  `xmlrpc:"rating_ids,omitempty"`
	RemindUpdateDelay        *Int       `xmlrpc:"remind_update_delay,omitempty"`
	ReportMessageFrequency   *Selection `xmlrpc:"report_message_frequency,omitempty"`
	ReportMessageGroupId     *Many2One  `xmlrpc:"report_message_group_id,omitempty"`
	ReportTemplateId         *Many2One  `xmlrpc:"report_template_id,omitempty"`
	RewardFailure            *Bool      `xmlrpc:"reward_failure,omitempty"`
	RewardFirstId            *Many2One  `xmlrpc:"reward_first_id,omitempty"`
	RewardId                 *Many2One  `xmlrpc:"reward_id,omitempty"`
	RewardRealtime           *Bool      `xmlrpc:"reward_realtime,omitempty"`
	RewardSecondId           *Many2One  `xmlrpc:"reward_second_id,omitempty"`
	RewardThirdId            *Many2One  `xmlrpc:"reward_third_id,omitempty"`
	StartDate                *Time      `xmlrpc:"start_date,omitempty"`
	State                    *Selection `xmlrpc:"state,omitempty"`
	UserCount                *Int       `xmlrpc:"user_count,omitempty"`
	UserDomain               *String    `xmlrpc:"user_domain,omitempty"`
	UserIds                  *Relation  `xmlrpc:"user_ids,omitempty"`
	VisibilityMode           *Selection `xmlrpc:"visibility_mode,omitempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// GamificationChallenges represents array of gamification.challenge model.
type GamificationChallenges []GamificationChallenge

// GamificationChallengeModel is the odoo model name.
const GamificationChallengeModel = "gamification.challenge"

// Many2One convert GamificationChallenge to *Many2One.
func (gc *GamificationChallenge) Many2One() *Many2One {
	return NewMany2One(gc.Id.Get(), "")
}

// CreateGamificationChallenge creates a new gamification.challenge model and returns its id.
func (c *Client) CreateGamificationChallenge(gc *GamificationChallenge) (int64, error) {
	ids, err := c.CreateGamificationChallenges([]*GamificationChallenge{gc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateGamificationChallenge creates a new gamification.challenge model and returns its id.
func (c *Client) CreateGamificationChallenges(gcs []*GamificationChallenge) ([]int64, error) {
	var vv []interface{}
	for _, v := range gcs {
		vv = append(vv, v)
	}
	return c.Create(GamificationChallengeModel, vv, nil)
}

// UpdateGamificationChallenge updates an existing gamification.challenge record.
func (c *Client) UpdateGamificationChallenge(gc *GamificationChallenge) error {
	return c.UpdateGamificationChallenges([]int64{gc.Id.Get()}, gc)
}

// UpdateGamificationChallenges updates existing gamification.challenge records.
// All records (represented by ids) will be updated by gc values.
func (c *Client) UpdateGamificationChallenges(ids []int64, gc *GamificationChallenge) error {
	return c.Update(GamificationChallengeModel, ids, gc, nil)
}

// DeleteGamificationChallenge deletes an existing gamification.challenge record.
func (c *Client) DeleteGamificationChallenge(id int64) error {
	return c.DeleteGamificationChallenges([]int64{id})
}

// DeleteGamificationChallenges deletes existing gamification.challenge records.
func (c *Client) DeleteGamificationChallenges(ids []int64) error {
	return c.Delete(GamificationChallengeModel, ids)
}

// GetGamificationChallenge gets gamification.challenge existing record.
func (c *Client) GetGamificationChallenge(id int64) (*GamificationChallenge, error) {
	gcs, err := c.GetGamificationChallenges([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*gcs)[0]), nil
}

// GetGamificationChallenges gets gamification.challenge existing records.
func (c *Client) GetGamificationChallenges(ids []int64) (*GamificationChallenges, error) {
	gcs := &GamificationChallenges{}
	if err := c.Read(GamificationChallengeModel, ids, nil, gcs); err != nil {
		return nil, err
	}
	return gcs, nil
}

// FindGamificationChallenge finds gamification.challenge record by querying it with criteria.
func (c *Client) FindGamificationChallenge(criteria *Criteria) (*GamificationChallenge, error) {
	gcs := &GamificationChallenges{}
	if err := c.SearchRead(GamificationChallengeModel, criteria, NewOptions().Limit(1), gcs); err != nil {
		return nil, err
	}
	return &((*gcs)[0]), nil
}

// FindGamificationChallenges finds gamification.challenge records by querying it
// and filtering it with criteria and options.
func (c *Client) FindGamificationChallenges(criteria *Criteria, options *Options) (*GamificationChallenges, error) {
	gcs := &GamificationChallenges{}
	if err := c.SearchRead(GamificationChallengeModel, criteria, options, gcs); err != nil {
		return nil, err
	}
	return gcs, nil
}

// FindGamificationChallengeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindGamificationChallengeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(GamificationChallengeModel, criteria, options)
}

// FindGamificationChallengeId finds record id by querying it with criteria.
func (c *Client) FindGamificationChallengeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(GamificationChallengeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
