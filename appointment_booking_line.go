package odoo

// AppointmentBookingLine represents appointment.booking.line model.
type AppointmentBookingLine struct {
	Active                *Bool     `xmlrpc:"active,omitempty"`
	AppointmentResourceId *Many2One `xmlrpc:"appointment_resource_id,omitempty"`
	AppointmentTypeId     *Many2One `xmlrpc:"appointment_type_id,omitempty"`
	CalendarEventId       *Many2One `xmlrpc:"calendar_event_id,omitempty"`
	CapacityReserved      *Int      `xmlrpc:"capacity_reserved,omitempty"`
	CapacityUsed          *Int      `xmlrpc:"capacity_used,omitempty"`
	CreateDate            *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid             *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName           *String   `xmlrpc:"display_name,omitempty"`
	EventStart            *Time     `xmlrpc:"event_start,omitempty"`
	EventStop             *Time     `xmlrpc:"event_stop,omitempty"`
	Id                    *Int      `xmlrpc:"id,omitempty"`
	WriteDate             *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid              *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AppointmentBookingLines represents array of appointment.booking.line model.
type AppointmentBookingLines []AppointmentBookingLine

// AppointmentBookingLineModel is the odoo model name.
const AppointmentBookingLineModel = "appointment.booking.line"

// Many2One convert AppointmentBookingLine to *Many2One.
func (abl *AppointmentBookingLine) Many2One() *Many2One {
	return NewMany2One(abl.Id.Get(), "")
}

// CreateAppointmentBookingLine creates a new appointment.booking.line model and returns its id.
func (c *Client) CreateAppointmentBookingLine(abl *AppointmentBookingLine) (int64, error) {
	ids, err := c.CreateAppointmentBookingLines([]*AppointmentBookingLine{abl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAppointmentBookingLine creates a new appointment.booking.line model and returns its id.
func (c *Client) CreateAppointmentBookingLines(abls []*AppointmentBookingLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range abls {
		vv = append(vv, v)
	}
	return c.Create(AppointmentBookingLineModel, vv, nil)
}

// UpdateAppointmentBookingLine updates an existing appointment.booking.line record.
func (c *Client) UpdateAppointmentBookingLine(abl *AppointmentBookingLine) error {
	return c.UpdateAppointmentBookingLines([]int64{abl.Id.Get()}, abl)
}

// UpdateAppointmentBookingLines updates existing appointment.booking.line records.
// All records (represented by ids) will be updated by abl values.
func (c *Client) UpdateAppointmentBookingLines(ids []int64, abl *AppointmentBookingLine) error {
	return c.Update(AppointmentBookingLineModel, ids, abl, nil)
}

// DeleteAppointmentBookingLine deletes an existing appointment.booking.line record.
func (c *Client) DeleteAppointmentBookingLine(id int64) error {
	return c.DeleteAppointmentBookingLines([]int64{id})
}

// DeleteAppointmentBookingLines deletes existing appointment.booking.line records.
func (c *Client) DeleteAppointmentBookingLines(ids []int64) error {
	return c.Delete(AppointmentBookingLineModel, ids)
}

// GetAppointmentBookingLine gets appointment.booking.line existing record.
func (c *Client) GetAppointmentBookingLine(id int64) (*AppointmentBookingLine, error) {
	abls, err := c.GetAppointmentBookingLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*abls)[0]), nil
}

// GetAppointmentBookingLines gets appointment.booking.line existing records.
func (c *Client) GetAppointmentBookingLines(ids []int64) (*AppointmentBookingLines, error) {
	abls := &AppointmentBookingLines{}
	if err := c.Read(AppointmentBookingLineModel, ids, nil, abls); err != nil {
		return nil, err
	}
	return abls, nil
}

// FindAppointmentBookingLine finds appointment.booking.line record by querying it with criteria.
func (c *Client) FindAppointmentBookingLine(criteria *Criteria) (*AppointmentBookingLine, error) {
	abls := &AppointmentBookingLines{}
	if err := c.SearchRead(AppointmentBookingLineModel, criteria, NewOptions().Limit(1), abls); err != nil {
		return nil, err
	}
	return &((*abls)[0]), nil
}

// FindAppointmentBookingLines finds appointment.booking.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAppointmentBookingLines(criteria *Criteria, options *Options) (*AppointmentBookingLines, error) {
	abls := &AppointmentBookingLines{}
	if err := c.SearchRead(AppointmentBookingLineModel, criteria, options, abls); err != nil {
		return nil, err
	}
	return abls, nil
}

// FindAppointmentBookingLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAppointmentBookingLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AppointmentBookingLineModel, criteria, options)
}

// FindAppointmentBookingLineId finds record id by querying it with criteria.
func (c *Client) FindAppointmentBookingLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AppointmentBookingLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
