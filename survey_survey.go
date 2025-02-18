package odoo

// SurveySurvey represents survey.survey model.
type SurveySurvey struct {
	AccessMode                  *Selection  `xmlrpc:"access_mode,omitempty"`
	AccessToken                 *String     `xmlrpc:"access_token,omitempty"`
	Active                      *Bool       `xmlrpc:"active,omitempty"`
	ActivityCalendarEventId     *Many2One   `xmlrpc:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline        *Time       `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection  `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String     `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation   `xmlrpc:"activity_ids,omitempty"`
	ActivityState               *Selection  `xmlrpc:"activity_state,omitempty"`
	ActivitySummary             *String     `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon            *String     `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId              *Many2One   `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One   `xmlrpc:"activity_user_id,omitempty"`
	AllowedSurveyTypes          interface{} `xmlrpc:"allowed_survey_types,omitempty"`
	AnswerCount                 *Int        `xmlrpc:"answer_count,omitempty"`
	AnswerDoneCount             *Int        `xmlrpc:"answer_done_count,omitempty"`
	AnswerDurationAvg           *Float      `xmlrpc:"answer_duration_avg,omitempty"`
	AnswerScoreAvg              *Float      `xmlrpc:"answer_score_avg,omitempty"`
	AttemptsLimit               *Int        `xmlrpc:"attempts_limit,omitempty"`
	BackgroundImage             *String     `xmlrpc:"background_image,omitempty"`
	BackgroundImageUrl          *String     `xmlrpc:"background_image_url,omitempty"`
	Certification               *Bool       `xmlrpc:"certification,omitempty"`
	CertificationBadgeId        *Many2One   `xmlrpc:"certification_badge_id,omitempty"`
	CertificationBadgeIdDummy   *Many2One   `xmlrpc:"certification_badge_id_dummy,omitempty"`
	CertificationGiveBadge      *Bool       `xmlrpc:"certification_give_badge,omitempty"`
	CertificationMailTemplateId *Many2One   `xmlrpc:"certification_mail_template_id,omitempty"`
	CertificationReportLayout   *Selection  `xmlrpc:"certification_report_layout,omitempty"`
	CertificationValidityMonths *Int        `xmlrpc:"certification_validity_months,omitempty"`
	Color                       *Int        `xmlrpc:"color,omitempty"`
	CreateDate                  *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One   `xmlrpc:"create_uid,omitempty"`
	Description                 *String     `xmlrpc:"description,omitempty"`
	DescriptionDone             *String     `xmlrpc:"description_done,omitempty"`
	DisplayName                 *String     `xmlrpc:"display_name,omitempty"`
	HasConditionalQuestions     *Bool       `xmlrpc:"has_conditional_questions,omitempty"`
	HasMessage                  *Bool       `xmlrpc:"has_message,omitempty"`
	Id                          *Int        `xmlrpc:"id,omitempty"`
	IsAttemptsLimited           *Bool       `xmlrpc:"is_attempts_limited,omitempty"`
	IsTimeLimited               *Bool       `xmlrpc:"is_time_limited,omitempty"`
	MessageAttachmentCount      *Int        `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds          *Relation   `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError             *Bool       `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter      *Int        `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError          *Bool       `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                  *Relation   `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower           *Bool       `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction           *Bool       `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter    *Int        `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds           *Relation   `xmlrpc:"message_partner_ids,omitempty"`
	MyActivityDateDeadline      *Time       `xmlrpc:"my_activity_date_deadline,omitempty"`
	PageIds                     *Relation   `xmlrpc:"page_ids,omitempty"`
	ProgressionMode             *Selection  `xmlrpc:"progression_mode,omitempty"`
	QuestionAndPageIds          *Relation   `xmlrpc:"question_and_page_ids,omitempty"`
	QuestionCount               *Int        `xmlrpc:"question_count,omitempty"`
	QuestionIds                 *Relation   `xmlrpc:"question_ids,omitempty"`
	QuestionsLayout             *Selection  `xmlrpc:"questions_layout,omitempty"`
	QuestionsSelection          *Selection  `xmlrpc:"questions_selection,omitempty"`
	RatingIds                   *Relation   `xmlrpc:"rating_ids,omitempty"`
	RestrictUserIds             *Relation   `xmlrpc:"restrict_user_ids,omitempty"`
	ScoringMaxObtainable        *Float      `xmlrpc:"scoring_max_obtainable,omitempty"`
	ScoringSuccessMin           *Float      `xmlrpc:"scoring_success_min,omitempty"`
	ScoringType                 *Selection  `xmlrpc:"scoring_type,omitempty"`
	SessionAnswerCount          *Int        `xmlrpc:"session_answer_count,omitempty"`
	SessionAvailable            *Bool       `xmlrpc:"session_available,omitempty"`
	SessionCode                 *String     `xmlrpc:"session_code,omitempty"`
	SessionLink                 *String     `xmlrpc:"session_link,omitempty"`
	SessionQuestionAnswerCount  *Int        `xmlrpc:"session_question_answer_count,omitempty"`
	SessionQuestionId           *Many2One   `xmlrpc:"session_question_id,omitempty"`
	SessionQuestionStartTime    *Time       `xmlrpc:"session_question_start_time,omitempty"`
	SessionShowLeaderboard      *Bool       `xmlrpc:"session_show_leaderboard,omitempty"`
	SessionSpeedRating          *Bool       `xmlrpc:"session_speed_rating,omitempty"`
	SessionSpeedRatingTimeLimit *Int        `xmlrpc:"session_speed_rating_time_limit,omitempty"`
	SessionStartTime            *Time       `xmlrpc:"session_start_time,omitempty"`
	SessionState                *Selection  `xmlrpc:"session_state,omitempty"`
	SuccessCount                *Int        `xmlrpc:"success_count,omitempty"`
	SuccessRatio                *Int        `xmlrpc:"success_ratio,omitempty"`
	SurveyType                  *Selection  `xmlrpc:"survey_type,omitempty"`
	TimeLimit                   *Float      `xmlrpc:"time_limit,omitempty"`
	Title                       *String     `xmlrpc:"title,omitempty"`
	UserId                      *Many2One   `xmlrpc:"user_id,omitempty"`
	UserInputIds                *Relation   `xmlrpc:"user_input_ids,omitempty"`
	UsersCanGoBack              *Bool       `xmlrpc:"users_can_go_back,omitempty"`
	UsersCanSignup              *Bool       `xmlrpc:"users_can_signup,omitempty"`
	UsersLoginRequired          *Bool       `xmlrpc:"users_login_required,omitempty"`
	WebsiteMessageIds           *Relation   `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                   *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// SurveySurveys represents array of survey.survey model.
type SurveySurveys []SurveySurvey

// SurveySurveyModel is the odoo model name.
const SurveySurveyModel = "survey.survey"

// Many2One convert SurveySurvey to *Many2One.
func (ss *SurveySurvey) Many2One() *Many2One {
	return NewMany2One(ss.Id.Get(), "")
}

// CreateSurveySurvey creates a new survey.survey model and returns its id.
func (c *Client) CreateSurveySurvey(ss *SurveySurvey) (int64, error) {
	ids, err := c.CreateSurveySurveys([]*SurveySurvey{ss})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSurveySurvey creates a new survey.survey model and returns its id.
func (c *Client) CreateSurveySurveys(sss []*SurveySurvey) ([]int64, error) {
	var vv []interface{}
	for _, v := range sss {
		vv = append(vv, v)
	}
	return c.Create(SurveySurveyModel, vv, nil)
}

// UpdateSurveySurvey updates an existing survey.survey record.
func (c *Client) UpdateSurveySurvey(ss *SurveySurvey) error {
	return c.UpdateSurveySurveys([]int64{ss.Id.Get()}, ss)
}

// UpdateSurveySurveys updates existing survey.survey records.
// All records (represented by ids) will be updated by ss values.
func (c *Client) UpdateSurveySurveys(ids []int64, ss *SurveySurvey) error {
	return c.Update(SurveySurveyModel, ids, ss, nil)
}

// DeleteSurveySurvey deletes an existing survey.survey record.
func (c *Client) DeleteSurveySurvey(id int64) error {
	return c.DeleteSurveySurveys([]int64{id})
}

// DeleteSurveySurveys deletes existing survey.survey records.
func (c *Client) DeleteSurveySurveys(ids []int64) error {
	return c.Delete(SurveySurveyModel, ids)
}

// GetSurveySurvey gets survey.survey existing record.
func (c *Client) GetSurveySurvey(id int64) (*SurveySurvey, error) {
	sss, err := c.GetSurveySurveys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sss)[0]), nil
}

// GetSurveySurveys gets survey.survey existing records.
func (c *Client) GetSurveySurveys(ids []int64) (*SurveySurveys, error) {
	sss := &SurveySurveys{}
	if err := c.Read(SurveySurveyModel, ids, nil, sss); err != nil {
		return nil, err
	}
	return sss, nil
}

// FindSurveySurvey finds survey.survey record by querying it with criteria.
func (c *Client) FindSurveySurvey(criteria *Criteria) (*SurveySurvey, error) {
	sss := &SurveySurveys{}
	if err := c.SearchRead(SurveySurveyModel, criteria, NewOptions().Limit(1), sss); err != nil {
		return nil, err
	}
	return &((*sss)[0]), nil
}

// FindSurveySurveys finds survey.survey records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSurveySurveys(criteria *Criteria, options *Options) (*SurveySurveys, error) {
	sss := &SurveySurveys{}
	if err := c.SearchRead(SurveySurveyModel, criteria, options, sss); err != nil {
		return nil, err
	}
	return sss, nil
}

// FindSurveySurveyIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSurveySurveyIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SurveySurveyModel, criteria, options)
}

// FindSurveySurveyId finds record id by querying it with criteria.
func (c *Client) FindSurveySurveyId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SurveySurveyModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
