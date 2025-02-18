package odoo

// AppointmentAnswerInput represents appointment.answer.input model.
type AppointmentAnswerInput struct {
	AppointmentTypeId *Many2One  `xmlrpc:"appointment_type_id,omitempty"`
	CalendarBookingId *Many2One  `xmlrpc:"calendar_booking_id,omitempty"`
	CalendarEventId   *Many2One  `xmlrpc:"calendar_event_id,omitempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName       *String    `xmlrpc:"display_name,omitempty"`
	Id                *Int       `xmlrpc:"id,omitempty"`
	PartnerId         *Many2One  `xmlrpc:"partner_id,omitempty"`
	QuestionId        *Many2One  `xmlrpc:"question_id,omitempty"`
	QuestionType      *Selection `xmlrpc:"question_type,omitempty"`
	ValueAnswerId     *Many2One  `xmlrpc:"value_answer_id,omitempty"`
	ValueTextBox      *String    `xmlrpc:"value_text_box,omitempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AppointmentAnswerInputs represents array of appointment.answer.input model.
type AppointmentAnswerInputs []AppointmentAnswerInput

// AppointmentAnswerInputModel is the odoo model name.
const AppointmentAnswerInputModel = "appointment.answer.input"

// Many2One convert AppointmentAnswerInput to *Many2One.
func (aai *AppointmentAnswerInput) Many2One() *Many2One {
	return NewMany2One(aai.Id.Get(), "")
}

// CreateAppointmentAnswerInput creates a new appointment.answer.input model and returns its id.
func (c *Client) CreateAppointmentAnswerInput(aai *AppointmentAnswerInput) (int64, error) {
	ids, err := c.CreateAppointmentAnswerInputs([]*AppointmentAnswerInput{aai})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAppointmentAnswerInput creates a new appointment.answer.input model and returns its id.
func (c *Client) CreateAppointmentAnswerInputs(aais []*AppointmentAnswerInput) ([]int64, error) {
	var vv []interface{}
	for _, v := range aais {
		vv = append(vv, v)
	}
	return c.Create(AppointmentAnswerInputModel, vv, nil)
}

// UpdateAppointmentAnswerInput updates an existing appointment.answer.input record.
func (c *Client) UpdateAppointmentAnswerInput(aai *AppointmentAnswerInput) error {
	return c.UpdateAppointmentAnswerInputs([]int64{aai.Id.Get()}, aai)
}

// UpdateAppointmentAnswerInputs updates existing appointment.answer.input records.
// All records (represented by ids) will be updated by aai values.
func (c *Client) UpdateAppointmentAnswerInputs(ids []int64, aai *AppointmentAnswerInput) error {
	return c.Update(AppointmentAnswerInputModel, ids, aai, nil)
}

// DeleteAppointmentAnswerInput deletes an existing appointment.answer.input record.
func (c *Client) DeleteAppointmentAnswerInput(id int64) error {
	return c.DeleteAppointmentAnswerInputs([]int64{id})
}

// DeleteAppointmentAnswerInputs deletes existing appointment.answer.input records.
func (c *Client) DeleteAppointmentAnswerInputs(ids []int64) error {
	return c.Delete(AppointmentAnswerInputModel, ids)
}

// GetAppointmentAnswerInput gets appointment.answer.input existing record.
func (c *Client) GetAppointmentAnswerInput(id int64) (*AppointmentAnswerInput, error) {
	aais, err := c.GetAppointmentAnswerInputs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aais)[0]), nil
}

// GetAppointmentAnswerInputs gets appointment.answer.input existing records.
func (c *Client) GetAppointmentAnswerInputs(ids []int64) (*AppointmentAnswerInputs, error) {
	aais := &AppointmentAnswerInputs{}
	if err := c.Read(AppointmentAnswerInputModel, ids, nil, aais); err != nil {
		return nil, err
	}
	return aais, nil
}

// FindAppointmentAnswerInput finds appointment.answer.input record by querying it with criteria.
func (c *Client) FindAppointmentAnswerInput(criteria *Criteria) (*AppointmentAnswerInput, error) {
	aais := &AppointmentAnswerInputs{}
	if err := c.SearchRead(AppointmentAnswerInputModel, criteria, NewOptions().Limit(1), aais); err != nil {
		return nil, err
	}
	return &((*aais)[0]), nil
}

// FindAppointmentAnswerInputs finds appointment.answer.input records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAppointmentAnswerInputs(criteria *Criteria, options *Options) (*AppointmentAnswerInputs, error) {
	aais := &AppointmentAnswerInputs{}
	if err := c.SearchRead(AppointmentAnswerInputModel, criteria, options, aais); err != nil {
		return nil, err
	}
	return aais, nil
}

// FindAppointmentAnswerInputIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAppointmentAnswerInputIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AppointmentAnswerInputModel, criteria, options)
}

// FindAppointmentAnswerInputId finds record id by querying it with criteria.
func (c *Client) FindAppointmentAnswerInputId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AppointmentAnswerInputModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
