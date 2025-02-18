package odoo

// SignLog represents sign.log model.
type SignLog struct {
	Action            *Selection `xmlrpc:"action,omitempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName       *String    `xmlrpc:"display_name,omitempty"`
	Id                *Int       `xmlrpc:"id,omitempty"`
	Ip                *String    `xmlrpc:"ip,omitempty"`
	Latitude          *Float     `xmlrpc:"latitude,omitempty"`
	LogDate           *Time      `xmlrpc:"log_date,omitempty"`
	LogHash           *String    `xmlrpc:"log_hash,omitempty"`
	Longitude         *Float     `xmlrpc:"longitude,omitempty"`
	PartnerId         *Many2One  `xmlrpc:"partner_id,omitempty"`
	RequestState      *Selection `xmlrpc:"request_state,omitempty"`
	SignRequestId     *Many2One  `xmlrpc:"sign_request_id,omitempty"`
	SignRequestItemId *Many2One  `xmlrpc:"sign_request_item_id,omitempty"`
	Token             *String    `xmlrpc:"token,omitempty"`
	UserId            *Many2One  `xmlrpc:"user_id,omitempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// SignLogs represents array of sign.log model.
type SignLogs []SignLog

// SignLogModel is the odoo model name.
const SignLogModel = "sign.log"

// Many2One convert SignLog to *Many2One.
func (sl *SignLog) Many2One() *Many2One {
	return NewMany2One(sl.Id.Get(), "")
}

// CreateSignLog creates a new sign.log model and returns its id.
func (c *Client) CreateSignLog(sl *SignLog) (int64, error) {
	ids, err := c.CreateSignLogs([]*SignLog{sl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSignLog creates a new sign.log model and returns its id.
func (c *Client) CreateSignLogs(sls []*SignLog) ([]int64, error) {
	var vv []interface{}
	for _, v := range sls {
		vv = append(vv, v)
	}
	return c.Create(SignLogModel, vv, nil)
}

// UpdateSignLog updates an existing sign.log record.
func (c *Client) UpdateSignLog(sl *SignLog) error {
	return c.UpdateSignLogs([]int64{sl.Id.Get()}, sl)
}

// UpdateSignLogs updates existing sign.log records.
// All records (represented by ids) will be updated by sl values.
func (c *Client) UpdateSignLogs(ids []int64, sl *SignLog) error {
	return c.Update(SignLogModel, ids, sl, nil)
}

// DeleteSignLog deletes an existing sign.log record.
func (c *Client) DeleteSignLog(id int64) error {
	return c.DeleteSignLogs([]int64{id})
}

// DeleteSignLogs deletes existing sign.log records.
func (c *Client) DeleteSignLogs(ids []int64) error {
	return c.Delete(SignLogModel, ids)
}

// GetSignLog gets sign.log existing record.
func (c *Client) GetSignLog(id int64) (*SignLog, error) {
	sls, err := c.GetSignLogs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sls)[0]), nil
}

// GetSignLogs gets sign.log existing records.
func (c *Client) GetSignLogs(ids []int64) (*SignLogs, error) {
	sls := &SignLogs{}
	if err := c.Read(SignLogModel, ids, nil, sls); err != nil {
		return nil, err
	}
	return sls, nil
}

// FindSignLog finds sign.log record by querying it with criteria.
func (c *Client) FindSignLog(criteria *Criteria) (*SignLog, error) {
	sls := &SignLogs{}
	if err := c.SearchRead(SignLogModel, criteria, NewOptions().Limit(1), sls); err != nil {
		return nil, err
	}
	return &((*sls)[0]), nil
}

// FindSignLogs finds sign.log records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSignLogs(criteria *Criteria, options *Options) (*SignLogs, error) {
	sls := &SignLogs{}
	if err := c.SearchRead(SignLogModel, criteria, options, sls); err != nil {
		return nil, err
	}
	return sls, nil
}

// FindSignLogIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSignLogIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SignLogModel, criteria, options)
}

// FindSignLogId finds record id by querying it with criteria.
func (c *Client) FindSignLogId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SignLogModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
