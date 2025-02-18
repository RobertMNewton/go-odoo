package odoo

// GamificationChallengeLine represents gamification.challenge.line model.
type GamificationChallengeLine struct {
	ChallengeId          *Many2One  `xmlrpc:"challenge_id,omitempty"`
	Condition            *Selection `xmlrpc:"condition,omitempty"`
	CreateDate           *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid            *Many2One  `xmlrpc:"create_uid,omitempty"`
	DefinitionFullSuffix *String    `xmlrpc:"definition_full_suffix,omitempty"`
	DefinitionId         *Many2One  `xmlrpc:"definition_id,omitempty"`
	DefinitionMonetary   *Bool      `xmlrpc:"definition_monetary,omitempty"`
	DefinitionSuffix     *String    `xmlrpc:"definition_suffix,omitempty"`
	DisplayName          *String    `xmlrpc:"display_name,omitempty"`
	Id                   *Int       `xmlrpc:"id,omitempty"`
	Name                 *String    `xmlrpc:"name,omitempty"`
	Sequence             *Int       `xmlrpc:"sequence,omitempty"`
	TargetGoal           *Float     `xmlrpc:"target_goal,omitempty"`
	WriteDate            *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid             *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// GamificationChallengeLines represents array of gamification.challenge.line model.
type GamificationChallengeLines []GamificationChallengeLine

// GamificationChallengeLineModel is the odoo model name.
const GamificationChallengeLineModel = "gamification.challenge.line"

// Many2One convert GamificationChallengeLine to *Many2One.
func (gcl *GamificationChallengeLine) Many2One() *Many2One {
	return NewMany2One(gcl.Id.Get(), "")
}

// CreateGamificationChallengeLine creates a new gamification.challenge.line model and returns its id.
func (c *Client) CreateGamificationChallengeLine(gcl *GamificationChallengeLine) (int64, error) {
	ids, err := c.CreateGamificationChallengeLines([]*GamificationChallengeLine{gcl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateGamificationChallengeLine creates a new gamification.challenge.line model and returns its id.
func (c *Client) CreateGamificationChallengeLines(gcls []*GamificationChallengeLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range gcls {
		vv = append(vv, v)
	}
	return c.Create(GamificationChallengeLineModel, vv, nil)
}

// UpdateGamificationChallengeLine updates an existing gamification.challenge.line record.
func (c *Client) UpdateGamificationChallengeLine(gcl *GamificationChallengeLine) error {
	return c.UpdateGamificationChallengeLines([]int64{gcl.Id.Get()}, gcl)
}

// UpdateGamificationChallengeLines updates existing gamification.challenge.line records.
// All records (represented by ids) will be updated by gcl values.
func (c *Client) UpdateGamificationChallengeLines(ids []int64, gcl *GamificationChallengeLine) error {
	return c.Update(GamificationChallengeLineModel, ids, gcl, nil)
}

// DeleteGamificationChallengeLine deletes an existing gamification.challenge.line record.
func (c *Client) DeleteGamificationChallengeLine(id int64) error {
	return c.DeleteGamificationChallengeLines([]int64{id})
}

// DeleteGamificationChallengeLines deletes existing gamification.challenge.line records.
func (c *Client) DeleteGamificationChallengeLines(ids []int64) error {
	return c.Delete(GamificationChallengeLineModel, ids)
}

// GetGamificationChallengeLine gets gamification.challenge.line existing record.
func (c *Client) GetGamificationChallengeLine(id int64) (*GamificationChallengeLine, error) {
	gcls, err := c.GetGamificationChallengeLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*gcls)[0]), nil
}

// GetGamificationChallengeLines gets gamification.challenge.line existing records.
func (c *Client) GetGamificationChallengeLines(ids []int64) (*GamificationChallengeLines, error) {
	gcls := &GamificationChallengeLines{}
	if err := c.Read(GamificationChallengeLineModel, ids, nil, gcls); err != nil {
		return nil, err
	}
	return gcls, nil
}

// FindGamificationChallengeLine finds gamification.challenge.line record by querying it with criteria.
func (c *Client) FindGamificationChallengeLine(criteria *Criteria) (*GamificationChallengeLine, error) {
	gcls := &GamificationChallengeLines{}
	if err := c.SearchRead(GamificationChallengeLineModel, criteria, NewOptions().Limit(1), gcls); err != nil {
		return nil, err
	}
	return &((*gcls)[0]), nil
}

// FindGamificationChallengeLines finds gamification.challenge.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindGamificationChallengeLines(criteria *Criteria, options *Options) (*GamificationChallengeLines, error) {
	gcls := &GamificationChallengeLines{}
	if err := c.SearchRead(GamificationChallengeLineModel, criteria, options, gcls); err != nil {
		return nil, err
	}
	return gcls, nil
}

// FindGamificationChallengeLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindGamificationChallengeLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(GamificationChallengeLineModel, criteria, options)
}

// FindGamificationChallengeLineId finds record id by querying it with criteria.
func (c *Client) FindGamificationChallengeLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(GamificationChallengeLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
