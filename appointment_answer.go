package odoo

// AppointmentAnswer represents appointment.answer model.
type AppointmentAnswer struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	QuestionId  *Many2One `xmlrpc:"question_id,omitempty"`
	Sequence    *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AppointmentAnswers represents array of appointment.answer model.
type AppointmentAnswers []AppointmentAnswer

// AppointmentAnswerModel is the odoo model name.
const AppointmentAnswerModel = "appointment.answer"

// Many2One convert AppointmentAnswer to *Many2One.
func (aa *AppointmentAnswer) Many2One() *Many2One {
	return NewMany2One(aa.Id.Get(), "")
}

// CreateAppointmentAnswer creates a new appointment.answer model and returns its id.
func (c *Client) CreateAppointmentAnswer(aa *AppointmentAnswer) (int64, error) {
	ids, err := c.CreateAppointmentAnswers([]*AppointmentAnswer{aa})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAppointmentAnswer creates a new appointment.answer model and returns its id.
func (c *Client) CreateAppointmentAnswers(aas []*AppointmentAnswer) ([]int64, error) {
	var vv []interface{}
	for _, v := range aas {
		vv = append(vv, v)
	}
	return c.Create(AppointmentAnswerModel, vv, nil)
}

// UpdateAppointmentAnswer updates an existing appointment.answer record.
func (c *Client) UpdateAppointmentAnswer(aa *AppointmentAnswer) error {
	return c.UpdateAppointmentAnswers([]int64{aa.Id.Get()}, aa)
}

// UpdateAppointmentAnswers updates existing appointment.answer records.
// All records (represented by ids) will be updated by aa values.
func (c *Client) UpdateAppointmentAnswers(ids []int64, aa *AppointmentAnswer) error {
	return c.Update(AppointmentAnswerModel, ids, aa, nil)
}

// DeleteAppointmentAnswer deletes an existing appointment.answer record.
func (c *Client) DeleteAppointmentAnswer(id int64) error {
	return c.DeleteAppointmentAnswers([]int64{id})
}

// DeleteAppointmentAnswers deletes existing appointment.answer records.
func (c *Client) DeleteAppointmentAnswers(ids []int64) error {
	return c.Delete(AppointmentAnswerModel, ids)
}

// GetAppointmentAnswer gets appointment.answer existing record.
func (c *Client) GetAppointmentAnswer(id int64) (*AppointmentAnswer, error) {
	aas, err := c.GetAppointmentAnswers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aas)[0]), nil
}

// GetAppointmentAnswers gets appointment.answer existing records.
func (c *Client) GetAppointmentAnswers(ids []int64) (*AppointmentAnswers, error) {
	aas := &AppointmentAnswers{}
	if err := c.Read(AppointmentAnswerModel, ids, nil, aas); err != nil {
		return nil, err
	}
	return aas, nil
}

// FindAppointmentAnswer finds appointment.answer record by querying it with criteria.
func (c *Client) FindAppointmentAnswer(criteria *Criteria) (*AppointmentAnswer, error) {
	aas := &AppointmentAnswers{}
	if err := c.SearchRead(AppointmentAnswerModel, criteria, NewOptions().Limit(1), aas); err != nil {
		return nil, err
	}
	return &((*aas)[0]), nil
}

// FindAppointmentAnswers finds appointment.answer records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAppointmentAnswers(criteria *Criteria, options *Options) (*AppointmentAnswers, error) {
	aas := &AppointmentAnswers{}
	if err := c.SearchRead(AppointmentAnswerModel, criteria, options, aas); err != nil {
		return nil, err
	}
	return aas, nil
}

// FindAppointmentAnswerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAppointmentAnswerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AppointmentAnswerModel, criteria, options)
}

// FindAppointmentAnswerId finds record id by querying it with criteria.
func (c *Client) FindAppointmentAnswerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AppointmentAnswerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
