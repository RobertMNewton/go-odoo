package odoo

// AppointmentInvite represents appointment.invite model.
type AppointmentInvite struct {
	AccessToken               *String    `xmlrpc:"access_token,omitempty"`
	AppointmentTypeCount      *Int       `xmlrpc:"appointment_type_count,omitempty"`
	AppointmentTypeIds        *Relation  `xmlrpc:"appointment_type_ids,omitempty"`
	AppointmentTypeInfoMsg    *String    `xmlrpc:"appointment_type_info_msg,omitempty"`
	AppointmentTypeWarningMsg *String    `xmlrpc:"appointment_type_warning_msg,omitempty"`
	BaseBookUrl               *String    `xmlrpc:"base_book_url,omitempty"`
	BookUrl                   *String    `xmlrpc:"book_url,omitempty"`
	CalendarEventCount        *Int       `xmlrpc:"calendar_event_count,omitempty"`
	CalendarEventIds          *Relation  `xmlrpc:"calendar_event_ids,omitempty"`
	CanPublish                *Bool      `xmlrpc:"can_publish,omitempty"`
	CreateDate                *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                 *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisableSaveButton         *Bool      `xmlrpc:"disable_save_button,omitempty"`
	DisplayName               *String    `xmlrpc:"display_name,omitempty"`
	Id                        *Int       `xmlrpc:"id,omitempty"`
	IsPublished               *Bool      `xmlrpc:"is_published,omitempty"`
	OpportunityId             *Many2One  `xmlrpc:"opportunity_id,omitempty"`
	RedirectUrl               *String    `xmlrpc:"redirect_url,omitempty"`
	ResourceIds               *Relation  `xmlrpc:"resource_ids,omitempty"`
	ResourcesChoice           *Selection `xmlrpc:"resources_choice,omitempty"`
	ScheduleBasedOn           *String    `xmlrpc:"schedule_based_on,omitempty"`
	ShortCode                 *String    `xmlrpc:"short_code,omitempty"`
	ShortCodeFormatWarning    *Bool      `xmlrpc:"short_code_format_warning,omitempty"`
	ShortCodeUniqueWarning    *Bool      `xmlrpc:"short_code_unique_warning,omitempty"`
	StaffUserIds              *Relation  `xmlrpc:"staff_user_ids,omitempty"`
	SuggestedResourceCount    *Int       `xmlrpc:"suggested_resource_count,omitempty"`
	SuggestedResourceIds      *Relation  `xmlrpc:"suggested_resource_ids,omitempty"`
	SuggestedStaffUserCount   *Int       `xmlrpc:"suggested_staff_user_count,omitempty"`
	SuggestedStaffUserIds     *Relation  `xmlrpc:"suggested_staff_user_ids,omitempty"`
	WebsiteId                 *Many2One  `xmlrpc:"website_id,omitempty"`
	WebsitePublished          *Bool      `xmlrpc:"website_published,omitempty"`
	WebsiteUrl                *String    `xmlrpc:"website_url,omitempty"`
	WriteDate                 *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                  *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AppointmentInvites represents array of appointment.invite model.
type AppointmentInvites []AppointmentInvite

// AppointmentInviteModel is the odoo model name.
const AppointmentInviteModel = "appointment.invite"

// Many2One convert AppointmentInvite to *Many2One.
func (ai *AppointmentInvite) Many2One() *Many2One {
	return NewMany2One(ai.Id.Get(), "")
}

// CreateAppointmentInvite creates a new appointment.invite model and returns its id.
func (c *Client) CreateAppointmentInvite(ai *AppointmentInvite) (int64, error) {
	ids, err := c.CreateAppointmentInvites([]*AppointmentInvite{ai})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAppointmentInvite creates a new appointment.invite model and returns its id.
func (c *Client) CreateAppointmentInvites(ais []*AppointmentInvite) ([]int64, error) {
	var vv []interface{}
	for _, v := range ais {
		vv = append(vv, v)
	}
	return c.Create(AppointmentInviteModel, vv, nil)
}

// UpdateAppointmentInvite updates an existing appointment.invite record.
func (c *Client) UpdateAppointmentInvite(ai *AppointmentInvite) error {
	return c.UpdateAppointmentInvites([]int64{ai.Id.Get()}, ai)
}

// UpdateAppointmentInvites updates existing appointment.invite records.
// All records (represented by ids) will be updated by ai values.
func (c *Client) UpdateAppointmentInvites(ids []int64, ai *AppointmentInvite) error {
	return c.Update(AppointmentInviteModel, ids, ai, nil)
}

// DeleteAppointmentInvite deletes an existing appointment.invite record.
func (c *Client) DeleteAppointmentInvite(id int64) error {
	return c.DeleteAppointmentInvites([]int64{id})
}

// DeleteAppointmentInvites deletes existing appointment.invite records.
func (c *Client) DeleteAppointmentInvites(ids []int64) error {
	return c.Delete(AppointmentInviteModel, ids)
}

// GetAppointmentInvite gets appointment.invite existing record.
func (c *Client) GetAppointmentInvite(id int64) (*AppointmentInvite, error) {
	ais, err := c.GetAppointmentInvites([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ais)[0]), nil
}

// GetAppointmentInvites gets appointment.invite existing records.
func (c *Client) GetAppointmentInvites(ids []int64) (*AppointmentInvites, error) {
	ais := &AppointmentInvites{}
	if err := c.Read(AppointmentInviteModel, ids, nil, ais); err != nil {
		return nil, err
	}
	return ais, nil
}

// FindAppointmentInvite finds appointment.invite record by querying it with criteria.
func (c *Client) FindAppointmentInvite(criteria *Criteria) (*AppointmentInvite, error) {
	ais := &AppointmentInvites{}
	if err := c.SearchRead(AppointmentInviteModel, criteria, NewOptions().Limit(1), ais); err != nil {
		return nil, err
	}
	return &((*ais)[0]), nil
}

// FindAppointmentInvites finds appointment.invite records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAppointmentInvites(criteria *Criteria, options *Options) (*AppointmentInvites, error) {
	ais := &AppointmentInvites{}
	if err := c.SearchRead(AppointmentInviteModel, criteria, options, ais); err != nil {
		return nil, err
	}
	return ais, nil
}

// FindAppointmentInviteIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAppointmentInviteIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AppointmentInviteModel, criteria, options)
}

// FindAppointmentInviteId finds record id by querying it with criteria.
func (c *Client) FindAppointmentInviteId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AppointmentInviteModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
