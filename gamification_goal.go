package odoo

// GamificationGoal represents gamification.goal model.
type GamificationGoal struct {
	ChallengeId           *Many2One  `xmlrpc:"challenge_id,omitempty"`
	Closed                *Bool      `xmlrpc:"closed,omitempty"`
	Color                 *Int       `xmlrpc:"color,omitempty"`
	Completeness          *Float     `xmlrpc:"completeness,omitempty"`
	ComputationMode       *Selection `xmlrpc:"computation_mode,omitempty"`
	CreateDate            *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid             *Many2One  `xmlrpc:"create_uid,omitempty"`
	Current               *Float     `xmlrpc:"current,omitempty"`
	DefinitionCondition   *Selection `xmlrpc:"definition_condition,omitempty"`
	DefinitionDescription *String    `xmlrpc:"definition_description,omitempty"`
	DefinitionDisplay     *Selection `xmlrpc:"definition_display,omitempty"`
	DefinitionId          *Many2One  `xmlrpc:"definition_id,omitempty"`
	DefinitionSuffix      *String    `xmlrpc:"definition_suffix,omitempty"`
	DisplayName           *String    `xmlrpc:"display_name,omitempty"`
	EndDate               *Time      `xmlrpc:"end_date,omitempty"`
	Id                    *Int       `xmlrpc:"id,omitempty"`
	LastUpdate            *Time      `xmlrpc:"last_update,omitempty"`
	LineId                *Many2One  `xmlrpc:"line_id,omitempty"`
	RemindUpdateDelay     *Int       `xmlrpc:"remind_update_delay,omitempty"`
	StartDate             *Time      `xmlrpc:"start_date,omitempty"`
	State                 *Selection `xmlrpc:"state,omitempty"`
	TargetGoal            *Float     `xmlrpc:"target_goal,omitempty"`
	ToUpdate              *Bool      `xmlrpc:"to_update,omitempty"`
	UserId                *Many2One  `xmlrpc:"user_id,omitempty"`
	UserPartnerId         *Many2One  `xmlrpc:"user_partner_id,omitempty"`
	WriteDate             *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid              *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// GamificationGoals represents array of gamification.goal model.
type GamificationGoals []GamificationGoal

// GamificationGoalModel is the odoo model name.
const GamificationGoalModel = "gamification.goal"

// Many2One convert GamificationGoal to *Many2One.
func (gg *GamificationGoal) Many2One() *Many2One {
	return NewMany2One(gg.Id.Get(), "")
}

// CreateGamificationGoal creates a new gamification.goal model and returns its id.
func (c *Client) CreateGamificationGoal(gg *GamificationGoal) (int64, error) {
	ids, err := c.CreateGamificationGoals([]*GamificationGoal{gg})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateGamificationGoal creates a new gamification.goal model and returns its id.
func (c *Client) CreateGamificationGoals(ggs []*GamificationGoal) ([]int64, error) {
	var vv []interface{}
	for _, v := range ggs {
		vv = append(vv, v)
	}
	return c.Create(GamificationGoalModel, vv, nil)
}

// UpdateGamificationGoal updates an existing gamification.goal record.
func (c *Client) UpdateGamificationGoal(gg *GamificationGoal) error {
	return c.UpdateGamificationGoals([]int64{gg.Id.Get()}, gg)
}

// UpdateGamificationGoals updates existing gamification.goal records.
// All records (represented by ids) will be updated by gg values.
func (c *Client) UpdateGamificationGoals(ids []int64, gg *GamificationGoal) error {
	return c.Update(GamificationGoalModel, ids, gg, nil)
}

// DeleteGamificationGoal deletes an existing gamification.goal record.
func (c *Client) DeleteGamificationGoal(id int64) error {
	return c.DeleteGamificationGoals([]int64{id})
}

// DeleteGamificationGoals deletes existing gamification.goal records.
func (c *Client) DeleteGamificationGoals(ids []int64) error {
	return c.Delete(GamificationGoalModel, ids)
}

// GetGamificationGoal gets gamification.goal existing record.
func (c *Client) GetGamificationGoal(id int64) (*GamificationGoal, error) {
	ggs, err := c.GetGamificationGoals([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ggs)[0]), nil
}

// GetGamificationGoals gets gamification.goal existing records.
func (c *Client) GetGamificationGoals(ids []int64) (*GamificationGoals, error) {
	ggs := &GamificationGoals{}
	if err := c.Read(GamificationGoalModel, ids, nil, ggs); err != nil {
		return nil, err
	}
	return ggs, nil
}

// FindGamificationGoal finds gamification.goal record by querying it with criteria.
func (c *Client) FindGamificationGoal(criteria *Criteria) (*GamificationGoal, error) {
	ggs := &GamificationGoals{}
	if err := c.SearchRead(GamificationGoalModel, criteria, NewOptions().Limit(1), ggs); err != nil {
		return nil, err
	}
	return &((*ggs)[0]), nil
}

// FindGamificationGoals finds gamification.goal records by querying it
// and filtering it with criteria and options.
func (c *Client) FindGamificationGoals(criteria *Criteria, options *Options) (*GamificationGoals, error) {
	ggs := &GamificationGoals{}
	if err := c.SearchRead(GamificationGoalModel, criteria, options, ggs); err != nil {
		return nil, err
	}
	return ggs, nil
}

// FindGamificationGoalIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindGamificationGoalIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(GamificationGoalModel, criteria, options)
}

// FindGamificationGoalId finds record id by querying it with criteria.
func (c *Client) FindGamificationGoalId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(GamificationGoalModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
