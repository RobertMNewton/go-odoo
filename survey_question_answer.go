package odoo

// SurveyQuestionAnswer represents survey.question.answer model.
type SurveyQuestionAnswer struct {
	AnswerScore        *Float     `xmlrpc:"answer_score,omitempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName        *String    `xmlrpc:"display_name,omitempty"`
	Id                 *Int       `xmlrpc:"id,omitempty"`
	IsCorrect          *Bool      `xmlrpc:"is_correct,omitempty"`
	MatrixQuestionId   *Many2One  `xmlrpc:"matrix_question_id,omitempty"`
	QuestionId         *Many2One  `xmlrpc:"question_id,omitempty"`
	QuestionType       *Selection `xmlrpc:"question_type,omitempty"`
	ScoringType        *Selection `xmlrpc:"scoring_type,omitempty"`
	Sequence           *Int       `xmlrpc:"sequence,omitempty"`
	Value              *String    `xmlrpc:"value,omitempty"`
	ValueImage         *String    `xmlrpc:"value_image,omitempty"`
	ValueImageFilename *String    `xmlrpc:"value_image_filename,omitempty"`
	ValueLabel         *String    `xmlrpc:"value_label,omitempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// SurveyQuestionAnswers represents array of survey.question.answer model.
type SurveyQuestionAnswers []SurveyQuestionAnswer

// SurveyQuestionAnswerModel is the odoo model name.
const SurveyQuestionAnswerModel = "survey.question.answer"

// Many2One convert SurveyQuestionAnswer to *Many2One.
func (sqa *SurveyQuestionAnswer) Many2One() *Many2One {
	return NewMany2One(sqa.Id.Get(), "")
}

// CreateSurveyQuestionAnswer creates a new survey.question.answer model and returns its id.
func (c *Client) CreateSurveyQuestionAnswer(sqa *SurveyQuestionAnswer) (int64, error) {
	ids, err := c.CreateSurveyQuestionAnswers([]*SurveyQuestionAnswer{sqa})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSurveyQuestionAnswer creates a new survey.question.answer model and returns its id.
func (c *Client) CreateSurveyQuestionAnswers(sqas []*SurveyQuestionAnswer) ([]int64, error) {
	var vv []interface{}
	for _, v := range sqas {
		vv = append(vv, v)
	}
	return c.Create(SurveyQuestionAnswerModel, vv, nil)
}

// UpdateSurveyQuestionAnswer updates an existing survey.question.answer record.
func (c *Client) UpdateSurveyQuestionAnswer(sqa *SurveyQuestionAnswer) error {
	return c.UpdateSurveyQuestionAnswers([]int64{sqa.Id.Get()}, sqa)
}

// UpdateSurveyQuestionAnswers updates existing survey.question.answer records.
// All records (represented by ids) will be updated by sqa values.
func (c *Client) UpdateSurveyQuestionAnswers(ids []int64, sqa *SurveyQuestionAnswer) error {
	return c.Update(SurveyQuestionAnswerModel, ids, sqa, nil)
}

// DeleteSurveyQuestionAnswer deletes an existing survey.question.answer record.
func (c *Client) DeleteSurveyQuestionAnswer(id int64) error {
	return c.DeleteSurveyQuestionAnswers([]int64{id})
}

// DeleteSurveyQuestionAnswers deletes existing survey.question.answer records.
func (c *Client) DeleteSurveyQuestionAnswers(ids []int64) error {
	return c.Delete(SurveyQuestionAnswerModel, ids)
}

// GetSurveyQuestionAnswer gets survey.question.answer existing record.
func (c *Client) GetSurveyQuestionAnswer(id int64) (*SurveyQuestionAnswer, error) {
	sqas, err := c.GetSurveyQuestionAnswers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sqas)[0]), nil
}

// GetSurveyQuestionAnswers gets survey.question.answer existing records.
func (c *Client) GetSurveyQuestionAnswers(ids []int64) (*SurveyQuestionAnswers, error) {
	sqas := &SurveyQuestionAnswers{}
	if err := c.Read(SurveyQuestionAnswerModel, ids, nil, sqas); err != nil {
		return nil, err
	}
	return sqas, nil
}

// FindSurveyQuestionAnswer finds survey.question.answer record by querying it with criteria.
func (c *Client) FindSurveyQuestionAnswer(criteria *Criteria) (*SurveyQuestionAnswer, error) {
	sqas := &SurveyQuestionAnswers{}
	if err := c.SearchRead(SurveyQuestionAnswerModel, criteria, NewOptions().Limit(1), sqas); err != nil {
		return nil, err
	}
	return &((*sqas)[0]), nil
}

// FindSurveyQuestionAnswers finds survey.question.answer records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSurveyQuestionAnswers(criteria *Criteria, options *Options) (*SurveyQuestionAnswers, error) {
	sqas := &SurveyQuestionAnswers{}
	if err := c.SearchRead(SurveyQuestionAnswerModel, criteria, options, sqas); err != nil {
		return nil, err
	}
	return sqas, nil
}

// FindSurveyQuestionAnswerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSurveyQuestionAnswerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SurveyQuestionAnswerModel, criteria, options)
}

// FindSurveyQuestionAnswerId finds record id by querying it with criteria.
func (c *Client) FindSurveyQuestionAnswerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SurveyQuestionAnswerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
