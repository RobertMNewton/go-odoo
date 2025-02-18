package odoo

// CalendarBookingLine represents calendar.booking.line model.
type CalendarBookingLine struct {
	AppointmentResourceId *Many2One `xmlrpc:"appointment_resource_id,omitempty"`
	CalendarBookingId     *Many2One `xmlrpc:"calendar_booking_id,omitempty"`
	CapacityReserved      *Int      `xmlrpc:"capacity_reserved,omitempty"`
	CapacityUsed          *Int      `xmlrpc:"capacity_used,omitempty"`
	CreateDate            *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid             *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName           *String   `xmlrpc:"display_name,omitempty"`
	Id                    *Int      `xmlrpc:"id,omitempty"`
	WriteDate             *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid              *Many2One `xmlrpc:"write_uid,omitempty"`
}

// CalendarBookingLines represents array of calendar.booking.line model.
type CalendarBookingLines []CalendarBookingLine

// CalendarBookingLineModel is the odoo model name.
const CalendarBookingLineModel = "calendar.booking.line"

// Many2One convert CalendarBookingLine to *Many2One.
func (cbl *CalendarBookingLine) Many2One() *Many2One {
	return NewMany2One(cbl.Id.Get(), "")
}

// CreateCalendarBookingLine creates a new calendar.booking.line model and returns its id.
func (c *Client) CreateCalendarBookingLine(cbl *CalendarBookingLine) (int64, error) {
	ids, err := c.CreateCalendarBookingLines([]*CalendarBookingLine{cbl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCalendarBookingLine creates a new calendar.booking.line model and returns its id.
func (c *Client) CreateCalendarBookingLines(cbls []*CalendarBookingLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range cbls {
		vv = append(vv, v)
	}
	return c.Create(CalendarBookingLineModel, vv, nil)
}

// UpdateCalendarBookingLine updates an existing calendar.booking.line record.
func (c *Client) UpdateCalendarBookingLine(cbl *CalendarBookingLine) error {
	return c.UpdateCalendarBookingLines([]int64{cbl.Id.Get()}, cbl)
}

// UpdateCalendarBookingLines updates existing calendar.booking.line records.
// All records (represented by ids) will be updated by cbl values.
func (c *Client) UpdateCalendarBookingLines(ids []int64, cbl *CalendarBookingLine) error {
	return c.Update(CalendarBookingLineModel, ids, cbl, nil)
}

// DeleteCalendarBookingLine deletes an existing calendar.booking.line record.
func (c *Client) DeleteCalendarBookingLine(id int64) error {
	return c.DeleteCalendarBookingLines([]int64{id})
}

// DeleteCalendarBookingLines deletes existing calendar.booking.line records.
func (c *Client) DeleteCalendarBookingLines(ids []int64) error {
	return c.Delete(CalendarBookingLineModel, ids)
}

// GetCalendarBookingLine gets calendar.booking.line existing record.
func (c *Client) GetCalendarBookingLine(id int64) (*CalendarBookingLine, error) {
	cbls, err := c.GetCalendarBookingLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*cbls)[0]), nil
}

// GetCalendarBookingLines gets calendar.booking.line existing records.
func (c *Client) GetCalendarBookingLines(ids []int64) (*CalendarBookingLines, error) {
	cbls := &CalendarBookingLines{}
	if err := c.Read(CalendarBookingLineModel, ids, nil, cbls); err != nil {
		return nil, err
	}
	return cbls, nil
}

// FindCalendarBookingLine finds calendar.booking.line record by querying it with criteria.
func (c *Client) FindCalendarBookingLine(criteria *Criteria) (*CalendarBookingLine, error) {
	cbls := &CalendarBookingLines{}
	if err := c.SearchRead(CalendarBookingLineModel, criteria, NewOptions().Limit(1), cbls); err != nil {
		return nil, err
	}
	return &((*cbls)[0]), nil
}

// FindCalendarBookingLines finds calendar.booking.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarBookingLines(criteria *Criteria, options *Options) (*CalendarBookingLines, error) {
	cbls := &CalendarBookingLines{}
	if err := c.SearchRead(CalendarBookingLineModel, criteria, options, cbls); err != nil {
		return nil, err
	}
	return cbls, nil
}

// FindCalendarBookingLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarBookingLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CalendarBookingLineModel, criteria, options)
}

// FindCalendarBookingLineId finds record id by querying it with criteria.
func (c *Client) FindCalendarBookingLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CalendarBookingLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
