package odoo

// SurveyUserInput represents survey.user_input model.
type SurveyUserInput struct {
	AccessToken                 *String    `xmlrpc:"access_token,omitempty"`
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
	AttemptsCount               *Int       `xmlrpc:"attempts_count,omitempty"`
	AttemptsLimit               *Int       `xmlrpc:"attempts_limit,omitempty"`
	AttemptsNumber              *Int       `xmlrpc:"attempts_number,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	Deadline                    *Time      `xmlrpc:"deadline,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	Email                       *String    `xmlrpc:"email,omitempty"`
	EndDatetime                 *Time      `xmlrpc:"end_datetime,omitempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	InviteToken                 *String    `xmlrpc:"invite_token,omitempty"`
	IsAttemptsLimited           *Bool      `xmlrpc:"is_attempts_limited,omitempty"`
	IsSessionAnswer             *Bool      `xmlrpc:"is_session_answer,omitempty"`
	LastDisplayedPageId         *Many2One  `xmlrpc:"last_displayed_page_id,omitempty"`
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
	Nickname                    *String    `xmlrpc:"nickname,omitempty"`
	PartnerId                   *Many2One  `xmlrpc:"partner_id,omitempty"`
	PredefinedQuestionIds       *Relation  `xmlrpc:"predefined_question_ids,omitempty"`
	QuestionTimeLimitReached    *Bool      `xmlrpc:"question_time_limit_reached,omitempty"`
	RatingIds                   *Relation  `xmlrpc:"rating_ids,omitempty"`
	ScoringPercentage           *Float     `xmlrpc:"scoring_percentage,omitempty"`
	ScoringSuccess              *Bool      `xmlrpc:"scoring_success,omitempty"`
	ScoringTotal                *Float     `xmlrpc:"scoring_total,omitempty"`
	ScoringType                 *Selection `xmlrpc:"scoring_type,omitempty"`
	StartDatetime               *Time      `xmlrpc:"start_datetime,omitempty"`
	State                       *Selection `xmlrpc:"state,omitempty"`
	SurveyFirstSubmitted        *Bool      `xmlrpc:"survey_first_submitted,omitempty"`
	SurveyId                    *Many2One  `xmlrpc:"survey_id,omitempty"`
	SurveyTimeLimitReached      *Bool      `xmlrpc:"survey_time_limit_reached,omitempty"`
	TestEntry                   *Bool      `xmlrpc:"test_entry,omitempty"`
	UserInputLineIds            *Relation  `xmlrpc:"user_input_line_ids,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// SurveyUserInputs represents array of survey.user_input model.
type SurveyUserInputs []SurveyUserInput

// SurveyUserInputModel is the odoo model name.
const SurveyUserInputModel = "survey.user_input"

// Many2One convert SurveyUserInput to *Many2One.
func (su *SurveyUserInput) Many2One() *Many2One {
	return NewMany2One(su.Id.Get(), "")
}

// CreateSurveyUserInput creates a new survey.user_input model and returns its id.
func (c *Client) CreateSurveyUserInput(su *SurveyUserInput) (int64, error) {
	ids, err := c.CreateSurveyUserInputs([]*SurveyUserInput{su})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSurveyUserInput creates a new survey.user_input model and returns its id.
func (c *Client) CreateSurveyUserInputs(sus []*SurveyUserInput) ([]int64, error) {
	var vv []interface{}
	for _, v := range sus {
		vv = append(vv, v)
	}
	return c.Create(SurveyUserInputModel, vv, nil)
}

// UpdateSurveyUserInput updates an existing survey.user_input record.
func (c *Client) UpdateSurveyUserInput(su *SurveyUserInput) error {
	return c.UpdateSurveyUserInputs([]int64{su.Id.Get()}, su)
}

// UpdateSurveyUserInputs updates existing survey.user_input records.
// All records (represented by ids) will be updated by su values.
func (c *Client) UpdateSurveyUserInputs(ids []int64, su *SurveyUserInput) error {
	return c.Update(SurveyUserInputModel, ids, su, nil)
}

// DeleteSurveyUserInput deletes an existing survey.user_input record.
func (c *Client) DeleteSurveyUserInput(id int64) error {
	return c.DeleteSurveyUserInputs([]int64{id})
}

// DeleteSurveyUserInputs deletes existing survey.user_input records.
func (c *Client) DeleteSurveyUserInputs(ids []int64) error {
	return c.Delete(SurveyUserInputModel, ids)
}

// GetSurveyUserInput gets survey.user_input existing record.
func (c *Client) GetSurveyUserInput(id int64) (*SurveyUserInput, error) {
	sus, err := c.GetSurveyUserInputs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sus)[0]), nil
}

// GetSurveyUserInputs gets survey.user_input existing records.
func (c *Client) GetSurveyUserInputs(ids []int64) (*SurveyUserInputs, error) {
	sus := &SurveyUserInputs{}
	if err := c.Read(SurveyUserInputModel, ids, nil, sus); err != nil {
		return nil, err
	}
	return sus, nil
}

// FindSurveyUserInput finds survey.user_input record by querying it with criteria.
func (c *Client) FindSurveyUserInput(criteria *Criteria) (*SurveyUserInput, error) {
	sus := &SurveyUserInputs{}
	if err := c.SearchRead(SurveyUserInputModel, criteria, NewOptions().Limit(1), sus); err != nil {
		return nil, err
	}
	return &((*sus)[0]), nil
}

// FindSurveyUserInputs finds survey.user_input records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSurveyUserInputs(criteria *Criteria, options *Options) (*SurveyUserInputs, error) {
	sus := &SurveyUserInputs{}
	if err := c.SearchRead(SurveyUserInputModel, criteria, options, sus); err != nil {
		return nil, err
	}
	return sus, nil
}

// FindSurveyUserInputIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSurveyUserInputIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SurveyUserInputModel, criteria, options)
}

// FindSurveyUserInputId finds record id by querying it with criteria.
func (c *Client) FindSurveyUserInputId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SurveyUserInputModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
