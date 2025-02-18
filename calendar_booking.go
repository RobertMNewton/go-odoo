package odoo

// CalendarBooking represents calendar.booking model.
type CalendarBooking struct {
	AccountMoveId             *Many2One `xmlrpc:"account_move_id,omitempty"`
	AppointmentAnswerInputIds *Relation `xmlrpc:"appointment_answer_input_ids,omitempty"`
	AppointmentInviteId       *Many2One `xmlrpc:"appointment_invite_id,omitempty"`
	AppointmentTypeId         *Many2One `xmlrpc:"appointment_type_id,omitempty"`
	AskedCapacity             *Int      `xmlrpc:"asked_capacity,omitempty"`
	BookingLineIds            *Relation `xmlrpc:"booking_line_ids,omitempty"`
	BookingToken              *String   `xmlrpc:"booking_token,omitempty"`
	CalendarEventId           *Many2One `xmlrpc:"calendar_event_id,omitempty"`
	CreateDate                *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                 *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName               *String   `xmlrpc:"display_name,omitempty"`
	Duration                  *Float    `xmlrpc:"duration,omitempty"`
	GuestIds                  *Relation `xmlrpc:"guest_ids,omitempty"`
	Id                        *Int      `xmlrpc:"id,omitempty"`
	Name                      *String   `xmlrpc:"name,omitempty"`
	NotAvailable              *Bool     `xmlrpc:"not_available,omitempty"`
	PartnerId                 *Many2One `xmlrpc:"partner_id,omitempty"`
	ProductId                 *Many2One `xmlrpc:"product_id,omitempty"`
	StaffUserId               *Many2One `xmlrpc:"staff_user_id,omitempty"`
	Start                     *Time     `xmlrpc:"start,omitempty"`
	Stop                      *Time     `xmlrpc:"stop,omitempty"`
	WriteDate                 *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                  *Many2One `xmlrpc:"write_uid,omitempty"`
}

// CalendarBookings represents array of calendar.booking model.
type CalendarBookings []CalendarBooking

// CalendarBookingModel is the odoo model name.
const CalendarBookingModel = "calendar.booking"

// Many2One convert CalendarBooking to *Many2One.
func (cb *CalendarBooking) Many2One() *Many2One {
	return NewMany2One(cb.Id.Get(), "")
}

// CreateCalendarBooking creates a new calendar.booking model and returns its id.
func (c *Client) CreateCalendarBooking(cb *CalendarBooking) (int64, error) {
	ids, err := c.CreateCalendarBookings([]*CalendarBooking{cb})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCalendarBooking creates a new calendar.booking model and returns its id.
func (c *Client) CreateCalendarBookings(cbs []*CalendarBooking) ([]int64, error) {
	var vv []interface{}
	for _, v := range cbs {
		vv = append(vv, v)
	}
	return c.Create(CalendarBookingModel, vv, nil)
}

// UpdateCalendarBooking updates an existing calendar.booking record.
func (c *Client) UpdateCalendarBooking(cb *CalendarBooking) error {
	return c.UpdateCalendarBookings([]int64{cb.Id.Get()}, cb)
}

// UpdateCalendarBookings updates existing calendar.booking records.
// All records (represented by ids) will be updated by cb values.
func (c *Client) UpdateCalendarBookings(ids []int64, cb *CalendarBooking) error {
	return c.Update(CalendarBookingModel, ids, cb, nil)
}

// DeleteCalendarBooking deletes an existing calendar.booking record.
func (c *Client) DeleteCalendarBooking(id int64) error {
	return c.DeleteCalendarBookings([]int64{id})
}

// DeleteCalendarBookings deletes existing calendar.booking records.
func (c *Client) DeleteCalendarBookings(ids []int64) error {
	return c.Delete(CalendarBookingModel, ids)
}

// GetCalendarBooking gets calendar.booking existing record.
func (c *Client) GetCalendarBooking(id int64) (*CalendarBooking, error) {
	cbs, err := c.GetCalendarBookings([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*cbs)[0]), nil
}

// GetCalendarBookings gets calendar.booking existing records.
func (c *Client) GetCalendarBookings(ids []int64) (*CalendarBookings, error) {
	cbs := &CalendarBookings{}
	if err := c.Read(CalendarBookingModel, ids, nil, cbs); err != nil {
		return nil, err
	}
	return cbs, nil
}

// FindCalendarBooking finds calendar.booking record by querying it with criteria.
func (c *Client) FindCalendarBooking(criteria *Criteria) (*CalendarBooking, error) {
	cbs := &CalendarBookings{}
	if err := c.SearchRead(CalendarBookingModel, criteria, NewOptions().Limit(1), cbs); err != nil {
		return nil, err
	}
	return &((*cbs)[0]), nil
}

// FindCalendarBookings finds calendar.booking records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarBookings(criteria *Criteria, options *Options) (*CalendarBookings, error) {
	cbs := &CalendarBookings{}
	if err := c.SearchRead(CalendarBookingModel, criteria, options, cbs); err != nil {
		return nil, err
	}
	return cbs, nil
}

// FindCalendarBookingIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarBookingIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CalendarBookingModel, criteria, options)
}

// FindCalendarBookingId finds record id by querying it with criteria.
func (c *Client) FindCalendarBookingId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CalendarBookingModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
