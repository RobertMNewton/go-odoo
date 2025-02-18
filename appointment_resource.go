package odoo

// AppointmentResource represents appointment.resource model.
type AppointmentResource struct {
	Active                 *Bool      `xmlrpc:"active,omitempty"`
	AppointmentTypeIds     *Relation  `xmlrpc:"appointment_type_ids,omitempty"`
	Avatar1024             *String    `xmlrpc:"avatar_1024,omitempty"`
	Avatar128              *String    `xmlrpc:"avatar_128,omitempty"`
	Avatar1920             *String    `xmlrpc:"avatar_1920,omitempty"`
	Avatar256              *String    `xmlrpc:"avatar_256,omitempty"`
	Avatar512              *String    `xmlrpc:"avatar_512,omitempty"`
	Capacity               *Int       `xmlrpc:"capacity,omitempty"`
	CompanyId              *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid              *Many2One  `xmlrpc:"create_uid,omitempty"`
	Description            *String    `xmlrpc:"description,omitempty"`
	DestinationResourceIds *Relation  `xmlrpc:"destination_resource_ids,omitempty"`
	DisplayName            *String    `xmlrpc:"display_name,omitempty"`
	Id                     *Int       `xmlrpc:"id,omitempty"`
	Image1024              *String    `xmlrpc:"image_1024,omitempty"`
	Image128               *String    `xmlrpc:"image_128,omitempty"`
	Image1920              *String    `xmlrpc:"image_1920,omitempty"`
	Image256               *String    `xmlrpc:"image_256,omitempty"`
	Image512               *String    `xmlrpc:"image_512,omitempty"`
	LinkedResourceIds      *Relation  `xmlrpc:"linked_resource_ids,omitempty"`
	Name                   *String    `xmlrpc:"name,omitempty"`
	ResourceCalendarId     *Many2One  `xmlrpc:"resource_calendar_id,omitempty"`
	ResourceId             *Many2One  `xmlrpc:"resource_id,omitempty"`
	Sequence               *Int       `xmlrpc:"sequence,omitempty"`
	Shareable              *Bool      `xmlrpc:"shareable,omitempty"`
	SourceResourceIds      *Relation  `xmlrpc:"source_resource_ids,omitempty"`
	Tz                     *Selection `xmlrpc:"tz,omitempty"`
	WriteDate              *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid               *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AppointmentResources represents array of appointment.resource model.
type AppointmentResources []AppointmentResource

// AppointmentResourceModel is the odoo model name.
const AppointmentResourceModel = "appointment.resource"

// Many2One convert AppointmentResource to *Many2One.
func (ar *AppointmentResource) Many2One() *Many2One {
	return NewMany2One(ar.Id.Get(), "")
}

// CreateAppointmentResource creates a new appointment.resource model and returns its id.
func (c *Client) CreateAppointmentResource(ar *AppointmentResource) (int64, error) {
	ids, err := c.CreateAppointmentResources([]*AppointmentResource{ar})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAppointmentResource creates a new appointment.resource model and returns its id.
func (c *Client) CreateAppointmentResources(ars []*AppointmentResource) ([]int64, error) {
	var vv []interface{}
	for _, v := range ars {
		vv = append(vv, v)
	}
	return c.Create(AppointmentResourceModel, vv, nil)
}

// UpdateAppointmentResource updates an existing appointment.resource record.
func (c *Client) UpdateAppointmentResource(ar *AppointmentResource) error {
	return c.UpdateAppointmentResources([]int64{ar.Id.Get()}, ar)
}

// UpdateAppointmentResources updates existing appointment.resource records.
// All records (represented by ids) will be updated by ar values.
func (c *Client) UpdateAppointmentResources(ids []int64, ar *AppointmentResource) error {
	return c.Update(AppointmentResourceModel, ids, ar, nil)
}

// DeleteAppointmentResource deletes an existing appointment.resource record.
func (c *Client) DeleteAppointmentResource(id int64) error {
	return c.DeleteAppointmentResources([]int64{id})
}

// DeleteAppointmentResources deletes existing appointment.resource records.
func (c *Client) DeleteAppointmentResources(ids []int64) error {
	return c.Delete(AppointmentResourceModel, ids)
}

// GetAppointmentResource gets appointment.resource existing record.
func (c *Client) GetAppointmentResource(id int64) (*AppointmentResource, error) {
	ars, err := c.GetAppointmentResources([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ars)[0]), nil
}

// GetAppointmentResources gets appointment.resource existing records.
func (c *Client) GetAppointmentResources(ids []int64) (*AppointmentResources, error) {
	ars := &AppointmentResources{}
	if err := c.Read(AppointmentResourceModel, ids, nil, ars); err != nil {
		return nil, err
	}
	return ars, nil
}

// FindAppointmentResource finds appointment.resource record by querying it with criteria.
func (c *Client) FindAppointmentResource(criteria *Criteria) (*AppointmentResource, error) {
	ars := &AppointmentResources{}
	if err := c.SearchRead(AppointmentResourceModel, criteria, NewOptions().Limit(1), ars); err != nil {
		return nil, err
	}
	return &((*ars)[0]), nil
}

// FindAppointmentResources finds appointment.resource records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAppointmentResources(criteria *Criteria, options *Options) (*AppointmentResources, error) {
	ars := &AppointmentResources{}
	if err := c.SearchRead(AppointmentResourceModel, criteria, options, ars); err != nil {
		return nil, err
	}
	return ars, nil
}

// FindAppointmentResourceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAppointmentResourceIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AppointmentResourceModel, criteria, options)
}

// FindAppointmentResourceId finds record id by querying it with criteria.
func (c *Client) FindAppointmentResourceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AppointmentResourceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
