package odoo

// AppointmentSlot represents appointment.slot model.
type AppointmentSlot struct {
	Allday                *Bool      `xmlrpc:"allday,omitempty"`
	AppointmentTypeId     *Many2One  `xmlrpc:"appointment_type_id,omitempty"`
	CreateDate            *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid             *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName           *String    `xmlrpc:"display_name,omitempty"`
	Duration              *Float     `xmlrpc:"duration,omitempty"`
	EndDatetime           *Time      `xmlrpc:"end_datetime,omitempty"`
	EndHour               *Float     `xmlrpc:"end_hour,omitempty"`
	Id                    *Int       `xmlrpc:"id,omitempty"`
	RestrictToResourceIds *Relation  `xmlrpc:"restrict_to_resource_ids,omitempty"`
	RestrictToUserIds     *Relation  `xmlrpc:"restrict_to_user_ids,omitempty"`
	ScheduleBasedOn       *Selection `xmlrpc:"schedule_based_on,omitempty"`
	SlotType              *Selection `xmlrpc:"slot_type,omitempty"`
	StartDatetime         *Time      `xmlrpc:"start_datetime,omitempty"`
	StartHour             *Float     `xmlrpc:"start_hour,omitempty"`
	Weekday               *Selection `xmlrpc:"weekday,omitempty"`
	WriteDate             *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid              *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AppointmentSlots represents array of appointment.slot model.
type AppointmentSlots []AppointmentSlot

// AppointmentSlotModel is the odoo model name.
const AppointmentSlotModel = "appointment.slot"

// Many2One convert AppointmentSlot to *Many2One.
func (as *AppointmentSlot) Many2One() *Many2One {
	return NewMany2One(as.Id.Get(), "")
}

// CreateAppointmentSlot creates a new appointment.slot model and returns its id.
func (c *Client) CreateAppointmentSlot(as *AppointmentSlot) (int64, error) {
	ids, err := c.CreateAppointmentSlots([]*AppointmentSlot{as})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAppointmentSlot creates a new appointment.slot model and returns its id.
func (c *Client) CreateAppointmentSlots(ass []*AppointmentSlot) ([]int64, error) {
	var vv []interface{}
	for _, v := range ass {
		vv = append(vv, v)
	}
	return c.Create(AppointmentSlotModel, vv, nil)
}

// UpdateAppointmentSlot updates an existing appointment.slot record.
func (c *Client) UpdateAppointmentSlot(as *AppointmentSlot) error {
	return c.UpdateAppointmentSlots([]int64{as.Id.Get()}, as)
}

// UpdateAppointmentSlots updates existing appointment.slot records.
// All records (represented by ids) will be updated by as values.
func (c *Client) UpdateAppointmentSlots(ids []int64, as *AppointmentSlot) error {
	return c.Update(AppointmentSlotModel, ids, as, nil)
}

// DeleteAppointmentSlot deletes an existing appointment.slot record.
func (c *Client) DeleteAppointmentSlot(id int64) error {
	return c.DeleteAppointmentSlots([]int64{id})
}

// DeleteAppointmentSlots deletes existing appointment.slot records.
func (c *Client) DeleteAppointmentSlots(ids []int64) error {
	return c.Delete(AppointmentSlotModel, ids)
}

// GetAppointmentSlot gets appointment.slot existing record.
func (c *Client) GetAppointmentSlot(id int64) (*AppointmentSlot, error) {
	ass, err := c.GetAppointmentSlots([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ass)[0]), nil
}

// GetAppointmentSlots gets appointment.slot existing records.
func (c *Client) GetAppointmentSlots(ids []int64) (*AppointmentSlots, error) {
	ass := &AppointmentSlots{}
	if err := c.Read(AppointmentSlotModel, ids, nil, ass); err != nil {
		return nil, err
	}
	return ass, nil
}

// FindAppointmentSlot finds appointment.slot record by querying it with criteria.
func (c *Client) FindAppointmentSlot(criteria *Criteria) (*AppointmentSlot, error) {
	ass := &AppointmentSlots{}
	if err := c.SearchRead(AppointmentSlotModel, criteria, NewOptions().Limit(1), ass); err != nil {
		return nil, err
	}
	return &((*ass)[0]), nil
}

// FindAppointmentSlots finds appointment.slot records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAppointmentSlots(criteria *Criteria, options *Options) (*AppointmentSlots, error) {
	ass := &AppointmentSlots{}
	if err := c.SearchRead(AppointmentSlotModel, criteria, options, ass); err != nil {
		return nil, err
	}
	return ass, nil
}

// FindAppointmentSlotIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAppointmentSlotIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AppointmentSlotModel, criteria, options)
}

// FindAppointmentSlotId finds record id by querying it with criteria.
func (c *Client) FindAppointmentSlotId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AppointmentSlotModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
