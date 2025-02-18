package odoo

// HrEmployeeBase represents hr.employee.base model.
type HrEmployeeBase struct {
	Active             *Bool      `xmlrpc:"active,omitempty"`
	AddressId          *Many2One  `xmlrpc:"address_id,omitempty"`
	BadgeIds           *Relation  `xmlrpc:"badge_ids,omitempty"`
	ChildAllCount      *Int       `xmlrpc:"child_all_count,omitempty"`
	ChildCount         *Int       `xmlrpc:"child_count,omitempty"`
	CoachId            *Many2One  `xmlrpc:"coach_id,omitempty"`
	Color              *Int       `xmlrpc:"color,omitempty"`
	CompanyId          *Many2One  `xmlrpc:"company_id,omitempty"`
	DepartmentColor    *Int       `xmlrpc:"department_color,omitempty"`
	DepartmentId       *Many2One  `xmlrpc:"department_id,omitempty"`
	DirectBadgeIds     *Relation  `xmlrpc:"direct_badge_ids,omitempty"`
	DisplayName        *String    `xmlrpc:"display_name,omitempty"`
	Email              *String    `xmlrpc:"email,omitempty"`
	GoalIds            *Relation  `xmlrpc:"goal_ids,omitempty"`
	HasBadges          *Bool      `xmlrpc:"has_badges,omitempty"`
	HasSlots           *Bool      `xmlrpc:"has_slots,omitempty"`
	HrIconDisplay      *Selection `xmlrpc:"hr_icon_display,omitempty"`
	HrPresenceState    *Selection `xmlrpc:"hr_presence_state,omitempty"`
	Id                 *Int       `xmlrpc:"id,omitempty"`
	ImStatus           *String    `xmlrpc:"im_status,omitempty"`
	IsFlexible         *Bool      `xmlrpc:"is_flexible,omitempty"`
	IsFullyFlexible    *Bool      `xmlrpc:"is_fully_flexible,omitempty"`
	JobId              *Many2One  `xmlrpc:"job_id,omitempty"`
	JobTitle           *String    `xmlrpc:"job_title,omitempty"`
	LastActivity       *Time      `xmlrpc:"last_activity,omitempty"`
	LastActivityTime   *String    `xmlrpc:"last_activity_time,omitempty"`
	MemberOfDepartment *Bool      `xmlrpc:"member_of_department,omitempty"`
	MobilePhone        *String    `xmlrpc:"mobile_phone,omitempty"`
	Name               *String    `xmlrpc:"name,omitempty"`
	NewlyHired         *Bool      `xmlrpc:"newly_hired,omitempty"`
	ParentId           *Many2One  `xmlrpc:"parent_id,omitempty"`
	Phone              *String    `xmlrpc:"phone,omitempty"`
	ResourceCalendarId *Many2One  `xmlrpc:"resource_calendar_id,omitempty"`
	ResourceId         *Many2One  `xmlrpc:"resource_id,omitempty"`
	Share              *Bool      `xmlrpc:"share,omitempty"`
	ShowHrIconDisplay  *Bool      `xmlrpc:"show_hr_icon_display,omitempty"`
	Tz                 *Selection `xmlrpc:"tz,omitempty"`
	UserId             *Many2One  `xmlrpc:"user_id,omitempty"`
	WorkContactId      *Many2One  `xmlrpc:"work_contact_id,omitempty"`
	WorkEmail          *String    `xmlrpc:"work_email,omitempty"`
	WorkLocationId     *Many2One  `xmlrpc:"work_location_id,omitempty"`
	WorkLocationName   *String    `xmlrpc:"work_location_name,omitempty"`
	WorkLocationType   *Selection `xmlrpc:"work_location_type,omitempty"`
	WorkPhone          *String    `xmlrpc:"work_phone,omitempty"`
}

// HrEmployeeBases represents array of hr.employee.base model.
type HrEmployeeBases []HrEmployeeBase

// HrEmployeeBaseModel is the odoo model name.
const HrEmployeeBaseModel = "hr.employee.base"

// Many2One convert HrEmployeeBase to *Many2One.
func (heb *HrEmployeeBase) Many2One() *Many2One {
	return NewMany2One(heb.Id.Get(), "")
}

// CreateHrEmployeeBase creates a new hr.employee.base model and returns its id.
func (c *Client) CreateHrEmployeeBase(heb *HrEmployeeBase) (int64, error) {
	ids, err := c.CreateHrEmployeeBases([]*HrEmployeeBase{heb})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrEmployeeBase creates a new hr.employee.base model and returns its id.
func (c *Client) CreateHrEmployeeBases(hebs []*HrEmployeeBase) ([]int64, error) {
	var vv []interface{}
	for _, v := range hebs {
		vv = append(vv, v)
	}
	return c.Create(HrEmployeeBaseModel, vv, nil)
}

// UpdateHrEmployeeBase updates an existing hr.employee.base record.
func (c *Client) UpdateHrEmployeeBase(heb *HrEmployeeBase) error {
	return c.UpdateHrEmployeeBases([]int64{heb.Id.Get()}, heb)
}

// UpdateHrEmployeeBases updates existing hr.employee.base records.
// All records (represented by ids) will be updated by heb values.
func (c *Client) UpdateHrEmployeeBases(ids []int64, heb *HrEmployeeBase) error {
	return c.Update(HrEmployeeBaseModel, ids, heb, nil)
}

// DeleteHrEmployeeBase deletes an existing hr.employee.base record.
func (c *Client) DeleteHrEmployeeBase(id int64) error {
	return c.DeleteHrEmployeeBases([]int64{id})
}

// DeleteHrEmployeeBases deletes existing hr.employee.base records.
func (c *Client) DeleteHrEmployeeBases(ids []int64) error {
	return c.Delete(HrEmployeeBaseModel, ids)
}

// GetHrEmployeeBase gets hr.employee.base existing record.
func (c *Client) GetHrEmployeeBase(id int64) (*HrEmployeeBase, error) {
	hebs, err := c.GetHrEmployeeBases([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hebs)[0]), nil
}

// GetHrEmployeeBases gets hr.employee.base existing records.
func (c *Client) GetHrEmployeeBases(ids []int64) (*HrEmployeeBases, error) {
	hebs := &HrEmployeeBases{}
	if err := c.Read(HrEmployeeBaseModel, ids, nil, hebs); err != nil {
		return nil, err
	}
	return hebs, nil
}

// FindHrEmployeeBase finds hr.employee.base record by querying it with criteria.
func (c *Client) FindHrEmployeeBase(criteria *Criteria) (*HrEmployeeBase, error) {
	hebs := &HrEmployeeBases{}
	if err := c.SearchRead(HrEmployeeBaseModel, criteria, NewOptions().Limit(1), hebs); err != nil {
		return nil, err
	}
	return &((*hebs)[0]), nil
}

// FindHrEmployeeBases finds hr.employee.base records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployeeBases(criteria *Criteria, options *Options) (*HrEmployeeBases, error) {
	hebs := &HrEmployeeBases{}
	if err := c.SearchRead(HrEmployeeBaseModel, criteria, options, hebs); err != nil {
		return nil, err
	}
	return hebs, nil
}

// FindHrEmployeeBaseIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployeeBaseIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrEmployeeBaseModel, criteria, options)
}

// FindHrEmployeeBaseId finds record id by querying it with criteria.
func (c *Client) FindHrEmployeeBaseId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrEmployeeBaseModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
