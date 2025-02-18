package odoo

// GamificationGoalDefinition represents gamification.goal.definition model.
type GamificationGoalDefinition struct {
	ActionId              *Many2One  `xmlrpc:"action_id,omitempty"`
	BatchDistinctiveField *Many2One  `xmlrpc:"batch_distinctive_field,omitempty"`
	BatchMode             *Bool      `xmlrpc:"batch_mode,omitempty"`
	BatchUserExpression   *String    `xmlrpc:"batch_user_expression,omitempty"`
	ComputationMode       *Selection `xmlrpc:"computation_mode,omitempty"`
	ComputeCode           *String    `xmlrpc:"compute_code,omitempty"`
	Condition             *Selection `xmlrpc:"condition,omitempty"`
	CreateDate            *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid             *Many2One  `xmlrpc:"create_uid,omitempty"`
	Description           *String    `xmlrpc:"description,omitempty"`
	DisplayMode           *Selection `xmlrpc:"display_mode,omitempty"`
	DisplayName           *String    `xmlrpc:"display_name,omitempty"`
	Domain                *String    `xmlrpc:"domain,omitempty"`
	FieldDateId           *Many2One  `xmlrpc:"field_date_id,omitempty"`
	FieldId               *Many2One  `xmlrpc:"field_id,omitempty"`
	FullSuffix            *String    `xmlrpc:"full_suffix,omitempty"`
	Id                    *Int       `xmlrpc:"id,omitempty"`
	ModelId               *Many2One  `xmlrpc:"model_id,omitempty"`
	ModelInheritedIds     *Relation  `xmlrpc:"model_inherited_ids,omitempty"`
	Monetary              *Bool      `xmlrpc:"monetary,omitempty"`
	Name                  *String    `xmlrpc:"name,omitempty"`
	ResIdField            *String    `xmlrpc:"res_id_field,omitempty"`
	Suffix                *String    `xmlrpc:"suffix,omitempty"`
	WriteDate             *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid              *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// GamificationGoalDefinitions represents array of gamification.goal.definition model.
type GamificationGoalDefinitions []GamificationGoalDefinition

// GamificationGoalDefinitionModel is the odoo model name.
const GamificationGoalDefinitionModel = "gamification.goal.definition"

// Many2One convert GamificationGoalDefinition to *Many2One.
func (ggd *GamificationGoalDefinition) Many2One() *Many2One {
	return NewMany2One(ggd.Id.Get(), "")
}

// CreateGamificationGoalDefinition creates a new gamification.goal.definition model and returns its id.
func (c *Client) CreateGamificationGoalDefinition(ggd *GamificationGoalDefinition) (int64, error) {
	ids, err := c.CreateGamificationGoalDefinitions([]*GamificationGoalDefinition{ggd})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateGamificationGoalDefinition creates a new gamification.goal.definition model and returns its id.
func (c *Client) CreateGamificationGoalDefinitions(ggds []*GamificationGoalDefinition) ([]int64, error) {
	var vv []interface{}
	for _, v := range ggds {
		vv = append(vv, v)
	}
	return c.Create(GamificationGoalDefinitionModel, vv, nil)
}

// UpdateGamificationGoalDefinition updates an existing gamification.goal.definition record.
func (c *Client) UpdateGamificationGoalDefinition(ggd *GamificationGoalDefinition) error {
	return c.UpdateGamificationGoalDefinitions([]int64{ggd.Id.Get()}, ggd)
}

// UpdateGamificationGoalDefinitions updates existing gamification.goal.definition records.
// All records (represented by ids) will be updated by ggd values.
func (c *Client) UpdateGamificationGoalDefinitions(ids []int64, ggd *GamificationGoalDefinition) error {
	return c.Update(GamificationGoalDefinitionModel, ids, ggd, nil)
}

// DeleteGamificationGoalDefinition deletes an existing gamification.goal.definition record.
func (c *Client) DeleteGamificationGoalDefinition(id int64) error {
	return c.DeleteGamificationGoalDefinitions([]int64{id})
}

// DeleteGamificationGoalDefinitions deletes existing gamification.goal.definition records.
func (c *Client) DeleteGamificationGoalDefinitions(ids []int64) error {
	return c.Delete(GamificationGoalDefinitionModel, ids)
}

// GetGamificationGoalDefinition gets gamification.goal.definition existing record.
func (c *Client) GetGamificationGoalDefinition(id int64) (*GamificationGoalDefinition, error) {
	ggds, err := c.GetGamificationGoalDefinitions([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ggds)[0]), nil
}

// GetGamificationGoalDefinitions gets gamification.goal.definition existing records.
func (c *Client) GetGamificationGoalDefinitions(ids []int64) (*GamificationGoalDefinitions, error) {
	ggds := &GamificationGoalDefinitions{}
	if err := c.Read(GamificationGoalDefinitionModel, ids, nil, ggds); err != nil {
		return nil, err
	}
	return ggds, nil
}

// FindGamificationGoalDefinition finds gamification.goal.definition record by querying it with criteria.
func (c *Client) FindGamificationGoalDefinition(criteria *Criteria) (*GamificationGoalDefinition, error) {
	ggds := &GamificationGoalDefinitions{}
	if err := c.SearchRead(GamificationGoalDefinitionModel, criteria, NewOptions().Limit(1), ggds); err != nil {
		return nil, err
	}
	return &((*ggds)[0]), nil
}

// FindGamificationGoalDefinitions finds gamification.goal.definition records by querying it
// and filtering it with criteria and options.
func (c *Client) FindGamificationGoalDefinitions(criteria *Criteria, options *Options) (*GamificationGoalDefinitions, error) {
	ggds := &GamificationGoalDefinitions{}
	if err := c.SearchRead(GamificationGoalDefinitionModel, criteria, options, ggds); err != nil {
		return nil, err
	}
	return ggds, nil
}

// FindGamificationGoalDefinitionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindGamificationGoalDefinitionIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(GamificationGoalDefinitionModel, criteria, options)
}

// FindGamificationGoalDefinitionId finds record id by querying it with criteria.
func (c *Client) FindGamificationGoalDefinitionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(GamificationGoalDefinitionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
