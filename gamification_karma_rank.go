package odoo

// GamificationKarmaRank represents gamification.karma.rank model.
type GamificationKarmaRank struct {
	CreateDate              *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid               *Many2One `xmlrpc:"create_uid,omitempty"`
	Description             *String   `xmlrpc:"description,omitempty"`
	DescriptionMotivational *String   `xmlrpc:"description_motivational,omitempty"`
	DisplayName             *String   `xmlrpc:"display_name,omitempty"`
	Id                      *Int      `xmlrpc:"id,omitempty"`
	Image1024               *String   `xmlrpc:"image_1024,omitempty"`
	Image128                *String   `xmlrpc:"image_128,omitempty"`
	Image1920               *String   `xmlrpc:"image_1920,omitempty"`
	Image256                *String   `xmlrpc:"image_256,omitempty"`
	Image512                *String   `xmlrpc:"image_512,omitempty"`
	KarmaMin                *Int      `xmlrpc:"karma_min,omitempty"`
	Name                    *String   `xmlrpc:"name,omitempty"`
	RankUsersCount          *Int      `xmlrpc:"rank_users_count,omitempty"`
	UserIds                 *Relation `xmlrpc:"user_ids,omitempty"`
	WriteDate               *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                *Many2One `xmlrpc:"write_uid,omitempty"`
}

// GamificationKarmaRanks represents array of gamification.karma.rank model.
type GamificationKarmaRanks []GamificationKarmaRank

// GamificationKarmaRankModel is the odoo model name.
const GamificationKarmaRankModel = "gamification.karma.rank"

// Many2One convert GamificationKarmaRank to *Many2One.
func (gkr *GamificationKarmaRank) Many2One() *Many2One {
	return NewMany2One(gkr.Id.Get(), "")
}

// CreateGamificationKarmaRank creates a new gamification.karma.rank model and returns its id.
func (c *Client) CreateGamificationKarmaRank(gkr *GamificationKarmaRank) (int64, error) {
	ids, err := c.CreateGamificationKarmaRanks([]*GamificationKarmaRank{gkr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateGamificationKarmaRank creates a new gamification.karma.rank model and returns its id.
func (c *Client) CreateGamificationKarmaRanks(gkrs []*GamificationKarmaRank) ([]int64, error) {
	var vv []interface{}
	for _, v := range gkrs {
		vv = append(vv, v)
	}
	return c.Create(GamificationKarmaRankModel, vv, nil)
}

// UpdateGamificationKarmaRank updates an existing gamification.karma.rank record.
func (c *Client) UpdateGamificationKarmaRank(gkr *GamificationKarmaRank) error {
	return c.UpdateGamificationKarmaRanks([]int64{gkr.Id.Get()}, gkr)
}

// UpdateGamificationKarmaRanks updates existing gamification.karma.rank records.
// All records (represented by ids) will be updated by gkr values.
func (c *Client) UpdateGamificationKarmaRanks(ids []int64, gkr *GamificationKarmaRank) error {
	return c.Update(GamificationKarmaRankModel, ids, gkr, nil)
}

// DeleteGamificationKarmaRank deletes an existing gamification.karma.rank record.
func (c *Client) DeleteGamificationKarmaRank(id int64) error {
	return c.DeleteGamificationKarmaRanks([]int64{id})
}

// DeleteGamificationKarmaRanks deletes existing gamification.karma.rank records.
func (c *Client) DeleteGamificationKarmaRanks(ids []int64) error {
	return c.Delete(GamificationKarmaRankModel, ids)
}

// GetGamificationKarmaRank gets gamification.karma.rank existing record.
func (c *Client) GetGamificationKarmaRank(id int64) (*GamificationKarmaRank, error) {
	gkrs, err := c.GetGamificationKarmaRanks([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*gkrs)[0]), nil
}

// GetGamificationKarmaRanks gets gamification.karma.rank existing records.
func (c *Client) GetGamificationKarmaRanks(ids []int64) (*GamificationKarmaRanks, error) {
	gkrs := &GamificationKarmaRanks{}
	if err := c.Read(GamificationKarmaRankModel, ids, nil, gkrs); err != nil {
		return nil, err
	}
	return gkrs, nil
}

// FindGamificationKarmaRank finds gamification.karma.rank record by querying it with criteria.
func (c *Client) FindGamificationKarmaRank(criteria *Criteria) (*GamificationKarmaRank, error) {
	gkrs := &GamificationKarmaRanks{}
	if err := c.SearchRead(GamificationKarmaRankModel, criteria, NewOptions().Limit(1), gkrs); err != nil {
		return nil, err
	}
	return &((*gkrs)[0]), nil
}

// FindGamificationKarmaRanks finds gamification.karma.rank records by querying it
// and filtering it with criteria and options.
func (c *Client) FindGamificationKarmaRanks(criteria *Criteria, options *Options) (*GamificationKarmaRanks, error) {
	gkrs := &GamificationKarmaRanks{}
	if err := c.SearchRead(GamificationKarmaRankModel, criteria, options, gkrs); err != nil {
		return nil, err
	}
	return gkrs, nil
}

// FindGamificationKarmaRankIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindGamificationKarmaRankIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(GamificationKarmaRankModel, criteria, options)
}

// FindGamificationKarmaRankId finds record id by querying it with criteria.
func (c *Client) FindGamificationKarmaRankId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(GamificationKarmaRankModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
