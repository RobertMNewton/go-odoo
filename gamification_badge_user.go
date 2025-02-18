package odoo

// GamificationBadgeUser represents gamification.badge.user model.
type GamificationBadgeUser struct {
	BadgeId       *Many2One  `xmlrpc:"badge_id,omitempty"`
	BadgeName     *String    `xmlrpc:"badge_name,omitempty"`
	ChallengeId   *Many2One  `xmlrpc:"challenge_id,omitempty"`
	Comment       *String    `xmlrpc:"comment,omitempty"`
	CreateDate    *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName   *String    `xmlrpc:"display_name,omitempty"`
	EmployeeId    *Many2One  `xmlrpc:"employee_id,omitempty"`
	Id            *Int       `xmlrpc:"id,omitempty"`
	Level         *Selection `xmlrpc:"level,omitempty"`
	SenderId      *Many2One  `xmlrpc:"sender_id,omitempty"`
	UserId        *Many2One  `xmlrpc:"user_id,omitempty"`
	UserPartnerId *Many2One  `xmlrpc:"user_partner_id,omitempty"`
	WriteDate     *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// GamificationBadgeUsers represents array of gamification.badge.user model.
type GamificationBadgeUsers []GamificationBadgeUser

// GamificationBadgeUserModel is the odoo model name.
const GamificationBadgeUserModel = "gamification.badge.user"

// Many2One convert GamificationBadgeUser to *Many2One.
func (gbu *GamificationBadgeUser) Many2One() *Many2One {
	return NewMany2One(gbu.Id.Get(), "")
}

// CreateGamificationBadgeUser creates a new gamification.badge.user model and returns its id.
func (c *Client) CreateGamificationBadgeUser(gbu *GamificationBadgeUser) (int64, error) {
	ids, err := c.CreateGamificationBadgeUsers([]*GamificationBadgeUser{gbu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateGamificationBadgeUser creates a new gamification.badge.user model and returns its id.
func (c *Client) CreateGamificationBadgeUsers(gbus []*GamificationBadgeUser) ([]int64, error) {
	var vv []interface{}
	for _, v := range gbus {
		vv = append(vv, v)
	}
	return c.Create(GamificationBadgeUserModel, vv, nil)
}

// UpdateGamificationBadgeUser updates an existing gamification.badge.user record.
func (c *Client) UpdateGamificationBadgeUser(gbu *GamificationBadgeUser) error {
	return c.UpdateGamificationBadgeUsers([]int64{gbu.Id.Get()}, gbu)
}

// UpdateGamificationBadgeUsers updates existing gamification.badge.user records.
// All records (represented by ids) will be updated by gbu values.
func (c *Client) UpdateGamificationBadgeUsers(ids []int64, gbu *GamificationBadgeUser) error {
	return c.Update(GamificationBadgeUserModel, ids, gbu, nil)
}

// DeleteGamificationBadgeUser deletes an existing gamification.badge.user record.
func (c *Client) DeleteGamificationBadgeUser(id int64) error {
	return c.DeleteGamificationBadgeUsers([]int64{id})
}

// DeleteGamificationBadgeUsers deletes existing gamification.badge.user records.
func (c *Client) DeleteGamificationBadgeUsers(ids []int64) error {
	return c.Delete(GamificationBadgeUserModel, ids)
}

// GetGamificationBadgeUser gets gamification.badge.user existing record.
func (c *Client) GetGamificationBadgeUser(id int64) (*GamificationBadgeUser, error) {
	gbus, err := c.GetGamificationBadgeUsers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*gbus)[0]), nil
}

// GetGamificationBadgeUsers gets gamification.badge.user existing records.
func (c *Client) GetGamificationBadgeUsers(ids []int64) (*GamificationBadgeUsers, error) {
	gbus := &GamificationBadgeUsers{}
	if err := c.Read(GamificationBadgeUserModel, ids, nil, gbus); err != nil {
		return nil, err
	}
	return gbus, nil
}

// FindGamificationBadgeUser finds gamification.badge.user record by querying it with criteria.
func (c *Client) FindGamificationBadgeUser(criteria *Criteria) (*GamificationBadgeUser, error) {
	gbus := &GamificationBadgeUsers{}
	if err := c.SearchRead(GamificationBadgeUserModel, criteria, NewOptions().Limit(1), gbus); err != nil {
		return nil, err
	}
	return &((*gbus)[0]), nil
}

// FindGamificationBadgeUsers finds gamification.badge.user records by querying it
// and filtering it with criteria and options.
func (c *Client) FindGamificationBadgeUsers(criteria *Criteria, options *Options) (*GamificationBadgeUsers, error) {
	gbus := &GamificationBadgeUsers{}
	if err := c.SearchRead(GamificationBadgeUserModel, criteria, options, gbus); err != nil {
		return nil, err
	}
	return gbus, nil
}

// FindGamificationBadgeUserIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindGamificationBadgeUserIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(GamificationBadgeUserModel, criteria, options)
}

// FindGamificationBadgeUserId finds record id by querying it with criteria.
func (c *Client) FindGamificationBadgeUserId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(GamificationBadgeUserModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
