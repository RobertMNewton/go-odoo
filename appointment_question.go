package odoo

// AppointmentQuestion represents appointment.question model.
type AppointmentQuestion struct {
	AnswerIds         *Relation  `xmlrpc:"answer_ids,omitempty"`
	AnswerInputIds    *Relation  `xmlrpc:"answer_input_ids,omitempty"`
	AppointmentTypeId *Many2One  `xmlrpc:"appointment_type_id,omitempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName       *String    `xmlrpc:"display_name,omitempty"`
	Id                *Int       `xmlrpc:"id,omitempty"`
	Name              *String    `xmlrpc:"name,omitempty"`
	Placeholder       *String    `xmlrpc:"placeholder,omitempty"`
	QuestionRequired  *Bool      `xmlrpc:"question_required,omitempty"`
	QuestionType      *Selection `xmlrpc:"question_type,omitempty"`
	Sequence          *Int       `xmlrpc:"sequence,omitempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AppointmentQuestions represents array of appointment.question model.
type AppointmentQuestions []AppointmentQuestion

// AppointmentQuestionModel is the odoo model name.
const AppointmentQuestionModel = "appointment.question"

// Many2One convert AppointmentQuestion to *Many2One.
func (aq *AppointmentQuestion) Many2One() *Many2One {
	return NewMany2One(aq.Id.Get(), "")
}

// CreateAppointmentQuestion creates a new appointment.question model and returns its id.
func (c *Client) CreateAppointmentQuestion(aq *AppointmentQuestion) (int64, error) {
	ids, err := c.CreateAppointmentQuestions([]*AppointmentQuestion{aq})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAppointmentQuestion creates a new appointment.question model and returns its id.
func (c *Client) CreateAppointmentQuestions(aqs []*AppointmentQuestion) ([]int64, error) {
	var vv []interface{}
	for _, v := range aqs {
		vv = append(vv, v)
	}
	return c.Create(AppointmentQuestionModel, vv, nil)
}

// UpdateAppointmentQuestion updates an existing appointment.question record.
func (c *Client) UpdateAppointmentQuestion(aq *AppointmentQuestion) error {
	return c.UpdateAppointmentQuestions([]int64{aq.Id.Get()}, aq)
}

// UpdateAppointmentQuestions updates existing appointment.question records.
// All records (represented by ids) will be updated by aq values.
func (c *Client) UpdateAppointmentQuestions(ids []int64, aq *AppointmentQuestion) error {
	return c.Update(AppointmentQuestionModel, ids, aq, nil)
}

// DeleteAppointmentQuestion deletes an existing appointment.question record.
func (c *Client) DeleteAppointmentQuestion(id int64) error {
	return c.DeleteAppointmentQuestions([]int64{id})
}

// DeleteAppointmentQuestions deletes existing appointment.question records.
func (c *Client) DeleteAppointmentQuestions(ids []int64) error {
	return c.Delete(AppointmentQuestionModel, ids)
}

// GetAppointmentQuestion gets appointment.question existing record.
func (c *Client) GetAppointmentQuestion(id int64) (*AppointmentQuestion, error) {
	aqs, err := c.GetAppointmentQuestions([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aqs)[0]), nil
}

// GetAppointmentQuestions gets appointment.question existing records.
func (c *Client) GetAppointmentQuestions(ids []int64) (*AppointmentQuestions, error) {
	aqs := &AppointmentQuestions{}
	if err := c.Read(AppointmentQuestionModel, ids, nil, aqs); err != nil {
		return nil, err
	}
	return aqs, nil
}

// FindAppointmentQuestion finds appointment.question record by querying it with criteria.
func (c *Client) FindAppointmentQuestion(criteria *Criteria) (*AppointmentQuestion, error) {
	aqs := &AppointmentQuestions{}
	if err := c.SearchRead(AppointmentQuestionModel, criteria, NewOptions().Limit(1), aqs); err != nil {
		return nil, err
	}
	return &((*aqs)[0]), nil
}

// FindAppointmentQuestions finds appointment.question records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAppointmentQuestions(criteria *Criteria, options *Options) (*AppointmentQuestions, error) {
	aqs := &AppointmentQuestions{}
	if err := c.SearchRead(AppointmentQuestionModel, criteria, options, aqs); err != nil {
		return nil, err
	}
	return aqs, nil
}

// FindAppointmentQuestionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAppointmentQuestionIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AppointmentQuestionModel, criteria, options)
}

// FindAppointmentQuestionId finds record id by querying it with criteria.
func (c *Client) FindAppointmentQuestionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AppointmentQuestionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
