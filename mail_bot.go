package odoo

// MailBot represents mail.bot model.
type MailBot struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// MailBots represents array of mail.bot model.
type MailBots []MailBot

// MailBotModel is the odoo model name.
const MailBotModel = "mail.bot"

// Many2One convert MailBot to *Many2One.
func (mb *MailBot) Many2One() *Many2One {
	return NewMany2One(mb.Id.Get(), "")
}

// CreateMailBot creates a new mail.bot model and returns its id.
func (c *Client) CreateMailBot(mb *MailBot) (int64, error) {
	ids, err := c.CreateMailBots([]*MailBot{mb})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailBot creates a new mail.bot model and returns its id.
func (c *Client) CreateMailBots(mbs []*MailBot) ([]int64, error) {
	var vv []interface{}
	for _, v := range mbs {
		vv = append(vv, v)
	}
	return c.Create(MailBotModel, vv, nil)
}

// UpdateMailBot updates an existing mail.bot record.
func (c *Client) UpdateMailBot(mb *MailBot) error {
	return c.UpdateMailBots([]int64{mb.Id.Get()}, mb)
}

// UpdateMailBots updates existing mail.bot records.
// All records (represented by ids) will be updated by mb values.
func (c *Client) UpdateMailBots(ids []int64, mb *MailBot) error {
	return c.Update(MailBotModel, ids, mb, nil)
}

// DeleteMailBot deletes an existing mail.bot record.
func (c *Client) DeleteMailBot(id int64) error {
	return c.DeleteMailBots([]int64{id})
}

// DeleteMailBots deletes existing mail.bot records.
func (c *Client) DeleteMailBots(ids []int64) error {
	return c.Delete(MailBotModel, ids)
}

// GetMailBot gets mail.bot existing record.
func (c *Client) GetMailBot(id int64) (*MailBot, error) {
	mbs, err := c.GetMailBots([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mbs)[0]), nil
}

// GetMailBots gets mail.bot existing records.
func (c *Client) GetMailBots(ids []int64) (*MailBots, error) {
	mbs := &MailBots{}
	if err := c.Read(MailBotModel, ids, nil, mbs); err != nil {
		return nil, err
	}
	return mbs, nil
}

// FindMailBot finds mail.bot record by querying it with criteria.
func (c *Client) FindMailBot(criteria *Criteria) (*MailBot, error) {
	mbs := &MailBots{}
	if err := c.SearchRead(MailBotModel, criteria, NewOptions().Limit(1), mbs); err != nil {
		return nil, err
	}
	return &((*mbs)[0]), nil
}

// FindMailBots finds mail.bot records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailBots(criteria *Criteria, options *Options) (*MailBots, error) {
	mbs := &MailBots{}
	if err := c.SearchRead(MailBotModel, criteria, options, mbs); err != nil {
		return nil, err
	}
	return mbs, nil
}

// FindMailBotIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailBotIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailBotModel, criteria, options)
}

// FindMailBotId finds record id by querying it with criteria.
func (c *Client) FindMailBotId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailBotModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
