package odoo

// SurveyQuestion represents survey.question model.
type SurveyQuestion struct {
	AllowedTriggeringQuestionIds *Relation  `xmlrpc:"allowed_triggering_question_ids,omitempty"`
	AnswerDate                   *Time      `xmlrpc:"answer_date,omitempty"`
	AnswerDatetime               *Time      `xmlrpc:"answer_datetime,omitempty"`
	AnswerNumericalBox           *Float     `xmlrpc:"answer_numerical_box,omitempty"`
	AnswerScore                  *Float     `xmlrpc:"answer_score,omitempty"`
	BackgroundImage              *String    `xmlrpc:"background_image,omitempty"`
	BackgroundImageUrl           *String    `xmlrpc:"background_image_url,omitempty"`
	CommentCountAsAnswer         *Bool      `xmlrpc:"comment_count_as_answer,omitempty"`
	CommentsAllowed              *Bool      `xmlrpc:"comments_allowed,omitempty"`
	CommentsMessage              *String    `xmlrpc:"comments_message,omitempty"`
	ConstrErrorMsg               *String    `xmlrpc:"constr_error_msg,omitempty"`
	ConstrMandatory              *Bool      `xmlrpc:"constr_mandatory,omitempty"`
	CreateDate                   *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                    *Many2One  `xmlrpc:"create_uid,omitempty"`
	Description                  *String    `xmlrpc:"description,omitempty"`
	DisplayName                  *String    `xmlrpc:"display_name,omitempty"`
	HasImageOnlySuggestedAnswer  *Bool      `xmlrpc:"has_image_only_suggested_answer,omitempty"`
	Id                           *Int       `xmlrpc:"id,omitempty"`
	IsPage                       *Bool      `xmlrpc:"is_page,omitempty"`
	IsPlacedBeforeTrigger        *Bool      `xmlrpc:"is_placed_before_trigger,omitempty"`
	IsScoredQuestion             *Bool      `xmlrpc:"is_scored_question,omitempty"`
	IsTimeCustomized             *Bool      `xmlrpc:"is_time_customized,omitempty"`
	IsTimeLimited                *Bool      `xmlrpc:"is_time_limited,omitempty"`
	MatrixRowIds                 *Relation  `xmlrpc:"matrix_row_ids,omitempty"`
	MatrixSubtype                *Selection `xmlrpc:"matrix_subtype,omitempty"`
	PageId                       *Many2One  `xmlrpc:"page_id,omitempty"`
	QuestionIds                  *Relation  `xmlrpc:"question_ids,omitempty"`
	QuestionPlaceholder          *String    `xmlrpc:"question_placeholder,omitempty"`
	QuestionType                 *Selection `xmlrpc:"question_type,omitempty"`
	QuestionsSelection           *Selection `xmlrpc:"questions_selection,omitempty"`
	RandomQuestionsCount         *Int       `xmlrpc:"random_questions_count,omitempty"`
	SaveAsEmail                  *Bool      `xmlrpc:"save_as_email,omitempty"`
	SaveAsNickname               *Bool      `xmlrpc:"save_as_nickname,omitempty"`
	ScaleMax                     *Int       `xmlrpc:"scale_max,omitempty"`
	ScaleMaxLabel                *String    `xmlrpc:"scale_max_label,omitempty"`
	ScaleMidLabel                *String    `xmlrpc:"scale_mid_label,omitempty"`
	ScaleMin                     *Int       `xmlrpc:"scale_min,omitempty"`
	ScaleMinLabel                *String    `xmlrpc:"scale_min_label,omitempty"`
	ScoringType                  *Selection `xmlrpc:"scoring_type,omitempty"`
	Sequence                     *Int       `xmlrpc:"sequence,omitempty"`
	SessionAvailable             *Bool      `xmlrpc:"session_available,omitempty"`
	SuggestedAnswerIds           *Relation  `xmlrpc:"suggested_answer_ids,omitempty"`
	SurveyId                     *Many2One  `xmlrpc:"survey_id,omitempty"`
	TimeLimit                    *Int       `xmlrpc:"time_limit,omitempty"`
	Title                        *String    `xmlrpc:"title,omitempty"`
	TriggeringAnswerIds          *Relation  `xmlrpc:"triggering_answer_ids,omitempty"`
	TriggeringQuestionIds        *Relation  `xmlrpc:"triggering_question_ids,omitempty"`
	UserInputLineIds             *Relation  `xmlrpc:"user_input_line_ids,omitempty"`
	ValidationEmail              *Bool      `xmlrpc:"validation_email,omitempty"`
	ValidationErrorMsg           *String    `xmlrpc:"validation_error_msg,omitempty"`
	ValidationLengthMax          *Int       `xmlrpc:"validation_length_max,omitempty"`
	ValidationLengthMin          *Int       `xmlrpc:"validation_length_min,omitempty"`
	ValidationMaxDate            *Time      `xmlrpc:"validation_max_date,omitempty"`
	ValidationMaxDatetime        *Time      `xmlrpc:"validation_max_datetime,omitempty"`
	ValidationMaxFloatValue      *Float     `xmlrpc:"validation_max_float_value,omitempty"`
	ValidationMinDate            *Time      `xmlrpc:"validation_min_date,omitempty"`
	ValidationMinDatetime        *Time      `xmlrpc:"validation_min_datetime,omitempty"`
	ValidationMinFloatValue      *Float     `xmlrpc:"validation_min_float_value,omitempty"`
	ValidationRequired           *Bool      `xmlrpc:"validation_required,omitempty"`
	WriteDate                    *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                     *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// SurveyQuestions represents array of survey.question model.
type SurveyQuestions []SurveyQuestion

// SurveyQuestionModel is the odoo model name.
const SurveyQuestionModel = "survey.question"

// Many2One convert SurveyQuestion to *Many2One.
func (sq *SurveyQuestion) Many2One() *Many2One {
	return NewMany2One(sq.Id.Get(), "")
}

// CreateSurveyQuestion creates a new survey.question model and returns its id.
func (c *Client) CreateSurveyQuestion(sq *SurveyQuestion) (int64, error) {
	ids, err := c.CreateSurveyQuestions([]*SurveyQuestion{sq})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSurveyQuestion creates a new survey.question model and returns its id.
func (c *Client) CreateSurveyQuestions(sqs []*SurveyQuestion) ([]int64, error) {
	var vv []interface{}
	for _, v := range sqs {
		vv = append(vv, v)
	}
	return c.Create(SurveyQuestionModel, vv, nil)
}

// UpdateSurveyQuestion updates an existing survey.question record.
func (c *Client) UpdateSurveyQuestion(sq *SurveyQuestion) error {
	return c.UpdateSurveyQuestions([]int64{sq.Id.Get()}, sq)
}

// UpdateSurveyQuestions updates existing survey.question records.
// All records (represented by ids) will be updated by sq values.
func (c *Client) UpdateSurveyQuestions(ids []int64, sq *SurveyQuestion) error {
	return c.Update(SurveyQuestionModel, ids, sq, nil)
}

// DeleteSurveyQuestion deletes an existing survey.question record.
func (c *Client) DeleteSurveyQuestion(id int64) error {
	return c.DeleteSurveyQuestions([]int64{id})
}

// DeleteSurveyQuestions deletes existing survey.question records.
func (c *Client) DeleteSurveyQuestions(ids []int64) error {
	return c.Delete(SurveyQuestionModel, ids)
}

// GetSurveyQuestion gets survey.question existing record.
func (c *Client) GetSurveyQuestion(id int64) (*SurveyQuestion, error) {
	sqs, err := c.GetSurveyQuestions([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sqs)[0]), nil
}

// GetSurveyQuestions gets survey.question existing records.
func (c *Client) GetSurveyQuestions(ids []int64) (*SurveyQuestions, error) {
	sqs := &SurveyQuestions{}
	if err := c.Read(SurveyQuestionModel, ids, nil, sqs); err != nil {
		return nil, err
	}
	return sqs, nil
}

// FindSurveyQuestion finds survey.question record by querying it with criteria.
func (c *Client) FindSurveyQuestion(criteria *Criteria) (*SurveyQuestion, error) {
	sqs := &SurveyQuestions{}
	if err := c.SearchRead(SurveyQuestionModel, criteria, NewOptions().Limit(1), sqs); err != nil {
		return nil, err
	}
	return &((*sqs)[0]), nil
}

// FindSurveyQuestions finds survey.question records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSurveyQuestions(criteria *Criteria, options *Options) (*SurveyQuestions, error) {
	sqs := &SurveyQuestions{}
	if err := c.SearchRead(SurveyQuestionModel, criteria, options, sqs); err != nil {
		return nil, err
	}
	return sqs, nil
}

// FindSurveyQuestionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSurveyQuestionIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SurveyQuestionModel, criteria, options)
}

// FindSurveyQuestionId finds record id by querying it with criteria.
func (c *Client) FindSurveyQuestionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SurveyQuestionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
