package odoo

// PlanningSlot represents planning.slot model.
type PlanningSlot struct {
	AccessToken              *String     `xmlrpc:"access_token,omitempty"`
	AllocatedHours           *Float      `xmlrpc:"allocated_hours,omitempty"`
	AllocatedPercentage      *Float      `xmlrpc:"allocated_percentage,omitempty"`
	AllocationType           *Selection  `xmlrpc:"allocation_type,omitempty"`
	AllowSelfUnassign        *Bool       `xmlrpc:"allow_self_unassign,omitempty"`
	AllowTemplateCreation    *Bool       `xmlrpc:"allow_template_creation,omitempty"`
	Color                    *Int        `xmlrpc:"color,omitempty"`
	CompanyId                *Many2One   `xmlrpc:"company_id,omitempty"`
	ConfirmDelete            *Bool       `xmlrpc:"confirm_delete,omitempty"`
	ConflictingSlotIds       *Relation   `xmlrpc:"conflicting_slot_ids,omitempty"`
	CreateDate               *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One   `xmlrpc:"create_uid,omitempty"`
	DepartmentId             *Many2One   `xmlrpc:"department_id,omitempty"`
	DisplayName              *String     `xmlrpc:"display_name,omitempty"`
	Duration                 *Float      `xmlrpc:"duration,omitempty"`
	EmployeeId               *Many2One   `xmlrpc:"employee_id,omitempty"`
	EmployeeSkillIds         *Relation   `xmlrpc:"employee_skill_ids,omitempty"`
	EndDatetime              *Time       `xmlrpc:"end_datetime,omitempty"`
	Id                       *Int        `xmlrpc:"id,omitempty"`
	IsHatched                *Bool       `xmlrpc:"is_hatched,omitempty"`
	IsPast                   *Bool       `xmlrpc:"is_past,omitempty"`
	IsUnassignDeadlinePassed *Bool       `xmlrpc:"is_unassign_deadline_passed,omitempty"`
	IsUsersRole              *Bool       `xmlrpc:"is_users_role,omitempty"`
	JobTitle                 *String     `xmlrpc:"job_title,omitempty"`
	ManagerId                *Many2One   `xmlrpc:"manager_id,omitempty"`
	Name                     *String     `xmlrpc:"name,omitempty"`
	OverlapSlotCount         *Int        `xmlrpc:"overlap_slot_count,omitempty"`
	PartnerId                *Many2One   `xmlrpc:"partner_id,omitempty"`
	PreviousTemplateId       *Many2One   `xmlrpc:"previous_template_id,omitempty"`
	PublicationWarning       *Bool       `xmlrpc:"publication_warning,omitempty"`
	RecurrenceUpdate         *Selection  `xmlrpc:"recurrence_update,omitempty"`
	RecurrencyId             *Many2One   `xmlrpc:"recurrency_id,omitempty"`
	Repeat                   *Bool       `xmlrpc:"repeat,omitempty"`
	RepeatInterval           *Int        `xmlrpc:"repeat_interval,omitempty"`
	RepeatNumber             *Int        `xmlrpc:"repeat_number,omitempty"`
	RepeatType               *Selection  `xmlrpc:"repeat_type,omitempty"`
	RepeatUnit               *Selection  `xmlrpc:"repeat_unit,omitempty"`
	RepeatUntil              *Time       `xmlrpc:"repeat_until,omitempty"`
	RequestToSwitch          *Bool       `xmlrpc:"request_to_switch,omitempty"`
	ResourceColor            *Int        `xmlrpc:"resource_color,omitempty"`
	ResourceId               *Many2One   `xmlrpc:"resource_id,omitempty"`
	ResourceRoles            *Relation   `xmlrpc:"resource_roles,omitempty"`
	ResourceType             *Selection  `xmlrpc:"resource_type,omitempty"`
	RoleId                   *Many2One   `xmlrpc:"role_id,omitempty"`
	RoleProductIds           *Relation   `xmlrpc:"role_product_ids,omitempty"`
	SaleLineId               *Many2One   `xmlrpc:"sale_line_id,omitempty"`
	SaleLinePlannable        *Bool       `xmlrpc:"sale_line_plannable,omitempty"`
	SaleOrderId              *Many2One   `xmlrpc:"sale_order_id,omitempty"`
	SaleOrderState           *Selection  `xmlrpc:"sale_order_state,omitempty"`
	SelfUnassignDaysBefore   *Int        `xmlrpc:"self_unassign_days_before,omitempty"`
	SlotProperties           interface{} `xmlrpc:"slot_properties,omitempty"`
	StartDatetime            *Time       `xmlrpc:"start_datetime,omitempty"`
	State                    *Selection  `xmlrpc:"state,omitempty"`
	TemplateAutocompleteIds  *Relation   `xmlrpc:"template_autocomplete_ids,omitempty"`
	TemplateCreation         *Bool       `xmlrpc:"template_creation,omitempty"`
	TemplateId               *Many2One   `xmlrpc:"template_id,omitempty"`
	TemplateReset            *Bool       `xmlrpc:"template_reset,omitempty"`
	UnassignDeadline         *Time       `xmlrpc:"unassign_deadline,omitempty"`
	UserId                   *Many2One   `xmlrpc:"user_id,omitempty"`
	WasCopied                *Bool       `xmlrpc:"was_copied,omitempty"`
	WorkEmail                *String     `xmlrpc:"work_email,omitempty"`
	WorkLocationId           *Many2One   `xmlrpc:"work_location_id,omitempty"`
	WriteDate                *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// PlanningSlots represents array of planning.slot model.
type PlanningSlots []PlanningSlot

// PlanningSlotModel is the odoo model name.
const PlanningSlotModel = "planning.slot"

// Many2One convert PlanningSlot to *Many2One.
func (ps *PlanningSlot) Many2One() *Many2One {
	return NewMany2One(ps.Id.Get(), "")
}

// CreatePlanningSlot creates a new planning.slot model and returns its id.
func (c *Client) CreatePlanningSlot(ps *PlanningSlot) (int64, error) {
	ids, err := c.CreatePlanningSlots([]*PlanningSlot{ps})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePlanningSlot creates a new planning.slot model and returns its id.
func (c *Client) CreatePlanningSlots(pss []*PlanningSlot) ([]int64, error) {
	var vv []interface{}
	for _, v := range pss {
		vv = append(vv, v)
	}
	return c.Create(PlanningSlotModel, vv, nil)
}

// UpdatePlanningSlot updates an existing planning.slot record.
func (c *Client) UpdatePlanningSlot(ps *PlanningSlot) error {
	return c.UpdatePlanningSlots([]int64{ps.Id.Get()}, ps)
}

// UpdatePlanningSlots updates existing planning.slot records.
// All records (represented by ids) will be updated by ps values.
func (c *Client) UpdatePlanningSlots(ids []int64, ps *PlanningSlot) error {
	return c.Update(PlanningSlotModel, ids, ps, nil)
}

// DeletePlanningSlot deletes an existing planning.slot record.
func (c *Client) DeletePlanningSlot(id int64) error {
	return c.DeletePlanningSlots([]int64{id})
}

// DeletePlanningSlots deletes existing planning.slot records.
func (c *Client) DeletePlanningSlots(ids []int64) error {
	return c.Delete(PlanningSlotModel, ids)
}

// GetPlanningSlot gets planning.slot existing record.
func (c *Client) GetPlanningSlot(id int64) (*PlanningSlot, error) {
	pss, err := c.GetPlanningSlots([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pss)[0]), nil
}

// GetPlanningSlots gets planning.slot existing records.
func (c *Client) GetPlanningSlots(ids []int64) (*PlanningSlots, error) {
	pss := &PlanningSlots{}
	if err := c.Read(PlanningSlotModel, ids, nil, pss); err != nil {
		return nil, err
	}
	return pss, nil
}

// FindPlanningSlot finds planning.slot record by querying it with criteria.
func (c *Client) FindPlanningSlot(criteria *Criteria) (*PlanningSlot, error) {
	pss := &PlanningSlots{}
	if err := c.SearchRead(PlanningSlotModel, criteria, NewOptions().Limit(1), pss); err != nil {
		return nil, err
	}
	return &((*pss)[0]), nil
}

// FindPlanningSlots finds planning.slot records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPlanningSlots(criteria *Criteria, options *Options) (*PlanningSlots, error) {
	pss := &PlanningSlots{}
	if err := c.SearchRead(PlanningSlotModel, criteria, options, pss); err != nil {
		return nil, err
	}
	return pss, nil
}

// FindPlanningSlotIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPlanningSlotIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PlanningSlotModel, criteria, options)
}

// FindPlanningSlotId finds record id by querying it with criteria.
func (c *Client) FindPlanningSlotId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PlanningSlotModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
