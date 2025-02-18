package odoo

// AppointmentType represents appointment.type model.
type AppointmentType struct {
	Active                               *Bool      `xmlrpc:"active,omitempty"`
	ActivityCalendarEventId              *Many2One  `xmlrpc:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline                 *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration          *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon                *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                          *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState                        *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary                      *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon                     *String    `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId                       *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId                       *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AllowGuests                          *Bool      `xmlrpc:"allow_guests,omitempty"`
	AppointmentCount                     *Int       `xmlrpc:"appointment_count,omitempty"`
	AppointmentCountRequest              *Int       `xmlrpc:"appointment_count_request,omitempty"`
	AppointmentCountUpcoming             *Int       `xmlrpc:"appointment_count_upcoming,omitempty"`
	AppointmentDuration                  *Float     `xmlrpc:"appointment_duration,omitempty"`
	AppointmentDurationFormatted         *String    `xmlrpc:"appointment_duration_formatted,omitempty"`
	AppointmentInviteCount               *Int       `xmlrpc:"appointment_invite_count,omitempty"`
	AppointmentInviteIds                 *Relation  `xmlrpc:"appointment_invite_ids,omitempty"`
	AppointmentManualConfirmation        *Bool      `xmlrpc:"appointment_manual_confirmation,omitempty"`
	AppointmentTz                        *Selection `xmlrpc:"appointment_tz,omitempty"`
	AssignMethod                         *Selection `xmlrpc:"assign_method,omitempty"`
	AvatarsDisplay                       *Selection `xmlrpc:"avatars_display,omitempty"`
	BookedMailTemplateId                 *Many2One  `xmlrpc:"booked_mail_template_id,omitempty"`
	CanPublish                           *Bool      `xmlrpc:"can_publish,omitempty"`
	CanceledMailTemplateId               *Many2One  `xmlrpc:"canceled_mail_template_id,omitempty"`
	Category                             *Selection `xmlrpc:"category,omitempty"`
	CategoryTimeDisplay                  *Selection `xmlrpc:"category_time_display,omitempty"`
	ConnectorsDisplayed                  *Bool      `xmlrpc:"connectors_displayed,omitempty"`
	CountryIds                           *Relation  `xmlrpc:"country_ids,omitempty"`
	CreateDate                           *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                            *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName                          *String    `xmlrpc:"display_name,omitempty"`
	EndDatetime                          *Time      `xmlrpc:"end_datetime,omitempty"`
	EventVideocallSource                 *Selection `xmlrpc:"event_videocall_source,omitempty"`
	HasMessage                           *Bool      `xmlrpc:"has_message,omitempty"`
	HasPaymentStep                       *Bool      `xmlrpc:"has_payment_step,omitempty"`
	HideDuration                         *Bool      `xmlrpc:"hide_duration,omitempty"`
	HideTimezone                         *Bool      `xmlrpc:"hide_timezone,omitempty"`
	Id                                   *Int       `xmlrpc:"id,omitempty"`
	Image1024                            *String    `xmlrpc:"image_1024,omitempty"`
	Image128                             *String    `xmlrpc:"image_128,omitempty"`
	Image1920                            *String    `xmlrpc:"image_1920,omitempty"`
	Image256                             *String    `xmlrpc:"image_256,omitempty"`
	Image512                             *String    `xmlrpc:"image_512,omitempty"`
	IsPublished                          *Bool      `xmlrpc:"is_published,omitempty"`
	IsSeoOptimized                       *Bool      `xmlrpc:"is_seo_optimized,omitempty"`
	LeadCount                            *Int       `xmlrpc:"lead_count,omitempty"`
	LeadCreate                           *Bool      `xmlrpc:"lead_create,omitempty"`
	LeadIds                              *Relation  `xmlrpc:"lead_ids,omitempty"`
	Location                             *String    `xmlrpc:"location,omitempty"`
	LocationId                           *Many2One  `xmlrpc:"location_id,omitempty"`
	MaxScheduleDays                      *Int       `xmlrpc:"max_schedule_days,omitempty"`
	MeetingIds                           *Relation  `xmlrpc:"meeting_ids,omitempty"`
	MessageAttachmentCount               *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageConfirmation                  *String    `xmlrpc:"message_confirmation,omitempty"`
	MessageFollowerIds                   *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError                      *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter               *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError                   *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                           *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIntro                         *String    `xmlrpc:"message_intro,omitempty"`
	MessageIsFollower                    *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction                    *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter             *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds                    *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MinCancellationHours                 *Float     `xmlrpc:"min_cancellation_hours,omitempty"`
	MinScheduleHours                     *Float     `xmlrpc:"min_schedule_hours,omitempty"`
	MyActivityDateDeadline               *Time      `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                                 *String    `xmlrpc:"name,omitempty"`
	ProductCurrencyId                    *Many2One  `xmlrpc:"product_currency_id,omitempty"`
	ProductId                            *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductLstPrice                      *Float     `xmlrpc:"product_lst_price,omitempty"`
	QuestionIds                          *Relation  `xmlrpc:"question_ids,omitempty"`
	RatingIds                            *Relation  `xmlrpc:"rating_ids,omitempty"`
	ReminderIds                          *Relation  `xmlrpc:"reminder_ids,omitempty"`
	ResourceCount                        *Int       `xmlrpc:"resource_count,omitempty"`
	ResourceIds                          *Relation  `xmlrpc:"resource_ids,omitempty"`
	ResourceManageCapacity               *Bool      `xmlrpc:"resource_manage_capacity,omitempty"`
	ResourceManualConfirmationPercentage *Float     `xmlrpc:"resource_manual_confirmation_percentage,omitempty"`
	ResourceTotalCapacity                *Int       `xmlrpc:"resource_total_capacity,omitempty"`
	ScheduleBasedOn                      *Selection `xmlrpc:"schedule_based_on,omitempty"`
	SeoName                              *String    `xmlrpc:"seo_name,omitempty"`
	Sequence                             *Int       `xmlrpc:"sequence,omitempty"`
	SlotCreationInterval                 *Float     `xmlrpc:"slot_creation_interval,omitempty"`
	SlotIds                              *Relation  `xmlrpc:"slot_ids,omitempty"`
	StaffUserCount                       *Int       `xmlrpc:"staff_user_count,omitempty"`
	StaffUserIds                         *Relation  `xmlrpc:"staff_user_ids,omitempty"`
	StartDatetime                        *Time      `xmlrpc:"start_datetime,omitempty"`
	WebsiteId                            *Many2One  `xmlrpc:"website_id,omitempty"`
	WebsiteMessageIds                    *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WebsiteMetaDescription               *String    `xmlrpc:"website_meta_description,omitempty"`
	WebsiteMetaKeywords                  *String    `xmlrpc:"website_meta_keywords,omitempty"`
	WebsiteMetaOgImg                     *String    `xmlrpc:"website_meta_og_img,omitempty"`
	WebsiteMetaTitle                     *String    `xmlrpc:"website_meta_title,omitempty"`
	WebsitePublished                     *Bool      `xmlrpc:"website_published,omitempty"`
	WebsiteUrl                           *String    `xmlrpc:"website_url,omitempty"`
	WorkHoursActivated                   *Bool      `xmlrpc:"work_hours_activated,omitempty"`
	WriteDate                            *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                             *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AppointmentTypes represents array of appointment.type model.
type AppointmentTypes []AppointmentType

// AppointmentTypeModel is the odoo model name.
const AppointmentTypeModel = "appointment.type"

// Many2One convert AppointmentType to *Many2One.
func (at *AppointmentType) Many2One() *Many2One {
	return NewMany2One(at.Id.Get(), "")
}

// CreateAppointmentType creates a new appointment.type model and returns its id.
func (c *Client) CreateAppointmentType(at *AppointmentType) (int64, error) {
	ids, err := c.CreateAppointmentTypes([]*AppointmentType{at})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAppointmentType creates a new appointment.type model and returns its id.
func (c *Client) CreateAppointmentTypes(ats []*AppointmentType) ([]int64, error) {
	var vv []interface{}
	for _, v := range ats {
		vv = append(vv, v)
	}
	return c.Create(AppointmentTypeModel, vv, nil)
}

// UpdateAppointmentType updates an existing appointment.type record.
func (c *Client) UpdateAppointmentType(at *AppointmentType) error {
	return c.UpdateAppointmentTypes([]int64{at.Id.Get()}, at)
}

// UpdateAppointmentTypes updates existing appointment.type records.
// All records (represented by ids) will be updated by at values.
func (c *Client) UpdateAppointmentTypes(ids []int64, at *AppointmentType) error {
	return c.Update(AppointmentTypeModel, ids, at, nil)
}

// DeleteAppointmentType deletes an existing appointment.type record.
func (c *Client) DeleteAppointmentType(id int64) error {
	return c.DeleteAppointmentTypes([]int64{id})
}

// DeleteAppointmentTypes deletes existing appointment.type records.
func (c *Client) DeleteAppointmentTypes(ids []int64) error {
	return c.Delete(AppointmentTypeModel, ids)
}

// GetAppointmentType gets appointment.type existing record.
func (c *Client) GetAppointmentType(id int64) (*AppointmentType, error) {
	ats, err := c.GetAppointmentTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ats)[0]), nil
}

// GetAppointmentTypes gets appointment.type existing records.
func (c *Client) GetAppointmentTypes(ids []int64) (*AppointmentTypes, error) {
	ats := &AppointmentTypes{}
	if err := c.Read(AppointmentTypeModel, ids, nil, ats); err != nil {
		return nil, err
	}
	return ats, nil
}

// FindAppointmentType finds appointment.type record by querying it with criteria.
func (c *Client) FindAppointmentType(criteria *Criteria) (*AppointmentType, error) {
	ats := &AppointmentTypes{}
	if err := c.SearchRead(AppointmentTypeModel, criteria, NewOptions().Limit(1), ats); err != nil {
		return nil, err
	}
	return &((*ats)[0]), nil
}

// FindAppointmentTypes finds appointment.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAppointmentTypes(criteria *Criteria, options *Options) (*AppointmentTypes, error) {
	ats := &AppointmentTypes{}
	if err := c.SearchRead(AppointmentTypeModel, criteria, options, ats); err != nil {
		return nil, err
	}
	return ats, nil
}

// FindAppointmentTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAppointmentTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AppointmentTypeModel, criteria, options)
}

// FindAppointmentTypeId finds record id by querying it with criteria.
func (c *Client) FindAppointmentTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AppointmentTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
