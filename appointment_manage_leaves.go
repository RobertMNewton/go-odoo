package odoo

// AppointmentManageLeaves represents appointment.manage.leaves model.
type AppointmentManageLeaves struct {
	AppointmentResourceIds *Relation `xmlrpc:"appointment_resource_ids,omitempty"`
	CreateDate             *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid              *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName            *String   `xmlrpc:"display_name,omitempty"`
	Id                     *Int      `xmlrpc:"id,omitempty"`
	LeaveEndDt             *Time     `xmlrpc:"leave_end_dt,omitempty"`
	LeaveStartDt           *Time     `xmlrpc:"leave_start_dt,omitempty"`
	Reason                 *String   `xmlrpc:"reason,omitempty"`
	WriteDate              *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid               *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AppointmentManageLeavess represents array of appointment.manage.leaves model.
type AppointmentManageLeavess []AppointmentManageLeaves

// AppointmentManageLeavesModel is the odoo model name.
const AppointmentManageLeavesModel = "appointment.manage.leaves"

// Many2One convert AppointmentManageLeaves to *Many2One.
func (aml *AppointmentManageLeaves) Many2One() *Many2One {
	return NewMany2One(aml.Id.Get(), "")
}

// CreateAppointmentManageLeaves creates a new appointment.manage.leaves model and returns its id.
func (c *Client) CreateAppointmentManageLeaves(aml *AppointmentManageLeaves) (int64, error) {
	ids, err := c.CreateAppointmentManageLeavess([]*AppointmentManageLeaves{aml})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAppointmentManageLeaves creates a new appointment.manage.leaves model and returns its id.
func (c *Client) CreateAppointmentManageLeavess(amls []*AppointmentManageLeaves) ([]int64, error) {
	var vv []interface{}
	for _, v := range amls {
		vv = append(vv, v)
	}
	return c.Create(AppointmentManageLeavesModel, vv, nil)
}

// UpdateAppointmentManageLeaves updates an existing appointment.manage.leaves record.
func (c *Client) UpdateAppointmentManageLeaves(aml *AppointmentManageLeaves) error {
	return c.UpdateAppointmentManageLeavess([]int64{aml.Id.Get()}, aml)
}

// UpdateAppointmentManageLeavess updates existing appointment.manage.leaves records.
// All records (represented by ids) will be updated by aml values.
func (c *Client) UpdateAppointmentManageLeavess(ids []int64, aml *AppointmentManageLeaves) error {
	return c.Update(AppointmentManageLeavesModel, ids, aml, nil)
}

// DeleteAppointmentManageLeaves deletes an existing appointment.manage.leaves record.
func (c *Client) DeleteAppointmentManageLeaves(id int64) error {
	return c.DeleteAppointmentManageLeavess([]int64{id})
}

// DeleteAppointmentManageLeavess deletes existing appointment.manage.leaves records.
func (c *Client) DeleteAppointmentManageLeavess(ids []int64) error {
	return c.Delete(AppointmentManageLeavesModel, ids)
}

// GetAppointmentManageLeaves gets appointment.manage.leaves existing record.
func (c *Client) GetAppointmentManageLeaves(id int64) (*AppointmentManageLeaves, error) {
	amls, err := c.GetAppointmentManageLeavess([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*amls)[0]), nil
}

// GetAppointmentManageLeavess gets appointment.manage.leaves existing records.
func (c *Client) GetAppointmentManageLeavess(ids []int64) (*AppointmentManageLeavess, error) {
	amls := &AppointmentManageLeavess{}
	if err := c.Read(AppointmentManageLeavesModel, ids, nil, amls); err != nil {
		return nil, err
	}
	return amls, nil
}

// FindAppointmentManageLeaves finds appointment.manage.leaves record by querying it with criteria.
func (c *Client) FindAppointmentManageLeaves(criteria *Criteria) (*AppointmentManageLeaves, error) {
	amls := &AppointmentManageLeavess{}
	if err := c.SearchRead(AppointmentManageLeavesModel, criteria, NewOptions().Limit(1), amls); err != nil {
		return nil, err
	}
	return &((*amls)[0]), nil
}

// FindAppointmentManageLeavess finds appointment.manage.leaves records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAppointmentManageLeavess(criteria *Criteria, options *Options) (*AppointmentManageLeavess, error) {
	amls := &AppointmentManageLeavess{}
	if err := c.SearchRead(AppointmentManageLeavesModel, criteria, options, amls); err != nil {
		return nil, err
	}
	return amls, nil
}

// FindAppointmentManageLeavesIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAppointmentManageLeavesIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AppointmentManageLeavesModel, criteria, options)
}

// FindAppointmentManageLeavesId finds record id by querying it with criteria.
func (c *Client) FindAppointmentManageLeavesId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AppointmentManageLeavesModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
