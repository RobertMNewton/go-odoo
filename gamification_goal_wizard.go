package odoo

// GamificationGoalWizard represents gamification.goal.wizard model.
type GamificationGoalWizard struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	Current     *Float    `xmlrpc:"current,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	GoalId      *Many2One `xmlrpc:"goal_id,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// GamificationGoalWizards represents array of gamification.goal.wizard model.
type GamificationGoalWizards []GamificationGoalWizard

// GamificationGoalWizardModel is the odoo model name.
const GamificationGoalWizardModel = "gamification.goal.wizard"

// Many2One convert GamificationGoalWizard to *Many2One.
func (ggw *GamificationGoalWizard) Many2One() *Many2One {
	return NewMany2One(ggw.Id.Get(), "")
}

// CreateGamificationGoalWizard creates a new gamification.goal.wizard model and returns its id.
func (c *Client) CreateGamificationGoalWizard(ggw *GamificationGoalWizard) (int64, error) {
	ids, err := c.CreateGamificationGoalWizards([]*GamificationGoalWizard{ggw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateGamificationGoalWizard creates a new gamification.goal.wizard model and returns its id.
func (c *Client) CreateGamificationGoalWizards(ggws []*GamificationGoalWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range ggws {
		vv = append(vv, v)
	}
	return c.Create(GamificationGoalWizardModel, vv, nil)
}

// UpdateGamificationGoalWizard updates an existing gamification.goal.wizard record.
func (c *Client) UpdateGamificationGoalWizard(ggw *GamificationGoalWizard) error {
	return c.UpdateGamificationGoalWizards([]int64{ggw.Id.Get()}, ggw)
}

// UpdateGamificationGoalWizards updates existing gamification.goal.wizard records.
// All records (represented by ids) will be updated by ggw values.
func (c *Client) UpdateGamificationGoalWizards(ids []int64, ggw *GamificationGoalWizard) error {
	return c.Update(GamificationGoalWizardModel, ids, ggw, nil)
}

// DeleteGamificationGoalWizard deletes an existing gamification.goal.wizard record.
func (c *Client) DeleteGamificationGoalWizard(id int64) error {
	return c.DeleteGamificationGoalWizards([]int64{id})
}

// DeleteGamificationGoalWizards deletes existing gamification.goal.wizard records.
func (c *Client) DeleteGamificationGoalWizards(ids []int64) error {
	return c.Delete(GamificationGoalWizardModel, ids)
}

// GetGamificationGoalWizard gets gamification.goal.wizard existing record.
func (c *Client) GetGamificationGoalWizard(id int64) (*GamificationGoalWizard, error) {
	ggws, err := c.GetGamificationGoalWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ggws)[0]), nil
}

// GetGamificationGoalWizards gets gamification.goal.wizard existing records.
func (c *Client) GetGamificationGoalWizards(ids []int64) (*GamificationGoalWizards, error) {
	ggws := &GamificationGoalWizards{}
	if err := c.Read(GamificationGoalWizardModel, ids, nil, ggws); err != nil {
		return nil, err
	}
	return ggws, nil
}

// FindGamificationGoalWizard finds gamification.goal.wizard record by querying it with criteria.
func (c *Client) FindGamificationGoalWizard(criteria *Criteria) (*GamificationGoalWizard, error) {
	ggws := &GamificationGoalWizards{}
	if err := c.SearchRead(GamificationGoalWizardModel, criteria, NewOptions().Limit(1), ggws); err != nil {
		return nil, err
	}
	return &((*ggws)[0]), nil
}

// FindGamificationGoalWizards finds gamification.goal.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindGamificationGoalWizards(criteria *Criteria, options *Options) (*GamificationGoalWizards, error) {
	ggws := &GamificationGoalWizards{}
	if err := c.SearchRead(GamificationGoalWizardModel, criteria, options, ggws); err != nil {
		return nil, err
	}
	return ggws, nil
}

// FindGamificationGoalWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindGamificationGoalWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(GamificationGoalWizardModel, criteria, options)
}

// FindGamificationGoalWizardId finds record id by querying it with criteria.
func (c *Client) FindGamificationGoalWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(GamificationGoalWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
