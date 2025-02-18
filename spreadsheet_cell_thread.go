package odoo

// SpreadsheetCellThread represents spreadsheet.cell.thread model.
type SpreadsheetCellThread struct {
	CreateDate               *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omitempty"`
	DashboardId              *Many2One `xmlrpc:"dashboard_id,omitempty"`
	DisplayName              *String   `xmlrpc:"display_name,omitempty"`
	DocumentId               *Many2One `xmlrpc:"document_id,omitempty"`
	HasMessage               *Bool     `xmlrpc:"has_message,omitempty"`
	Id                       *Int      `xmlrpc:"id,omitempty"`
	MessageAttachmentCount   *Int      `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds       *Relation `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError          *Bool     `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter   *Int      `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError       *Bool     `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds               *Relation `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower        *Bool     `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction        *Bool     `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int      `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation `xmlrpc:"message_partner_ids,omitempty"`
	RatingIds                *Relation `xmlrpc:"rating_ids,omitempty"`
	SaleOrderSpreadsheetId   *Many2One `xmlrpc:"sale_order_spreadsheet_id,omitempty"`
	TemplateId               *Many2One `xmlrpc:"template_id,omitempty"`
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SpreadsheetCellThreads represents array of spreadsheet.cell.thread model.
type SpreadsheetCellThreads []SpreadsheetCellThread

// SpreadsheetCellThreadModel is the odoo model name.
const SpreadsheetCellThreadModel = "spreadsheet.cell.thread"

// Many2One convert SpreadsheetCellThread to *Many2One.
func (sct *SpreadsheetCellThread) Many2One() *Many2One {
	return NewMany2One(sct.Id.Get(), "")
}

// CreateSpreadsheetCellThread creates a new spreadsheet.cell.thread model and returns its id.
func (c *Client) CreateSpreadsheetCellThread(sct *SpreadsheetCellThread) (int64, error) {
	ids, err := c.CreateSpreadsheetCellThreads([]*SpreadsheetCellThread{sct})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSpreadsheetCellThread creates a new spreadsheet.cell.thread model and returns its id.
func (c *Client) CreateSpreadsheetCellThreads(scts []*SpreadsheetCellThread) ([]int64, error) {
	var vv []interface{}
	for _, v := range scts {
		vv = append(vv, v)
	}
	return c.Create(SpreadsheetCellThreadModel, vv, nil)
}

// UpdateSpreadsheetCellThread updates an existing spreadsheet.cell.thread record.
func (c *Client) UpdateSpreadsheetCellThread(sct *SpreadsheetCellThread) error {
	return c.UpdateSpreadsheetCellThreads([]int64{sct.Id.Get()}, sct)
}

// UpdateSpreadsheetCellThreads updates existing spreadsheet.cell.thread records.
// All records (represented by ids) will be updated by sct values.
func (c *Client) UpdateSpreadsheetCellThreads(ids []int64, sct *SpreadsheetCellThread) error {
	return c.Update(SpreadsheetCellThreadModel, ids, sct, nil)
}

// DeleteSpreadsheetCellThread deletes an existing spreadsheet.cell.thread record.
func (c *Client) DeleteSpreadsheetCellThread(id int64) error {
	return c.DeleteSpreadsheetCellThreads([]int64{id})
}

// DeleteSpreadsheetCellThreads deletes existing spreadsheet.cell.thread records.
func (c *Client) DeleteSpreadsheetCellThreads(ids []int64) error {
	return c.Delete(SpreadsheetCellThreadModel, ids)
}

// GetSpreadsheetCellThread gets spreadsheet.cell.thread existing record.
func (c *Client) GetSpreadsheetCellThread(id int64) (*SpreadsheetCellThread, error) {
	scts, err := c.GetSpreadsheetCellThreads([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*scts)[0]), nil
}

// GetSpreadsheetCellThreads gets spreadsheet.cell.thread existing records.
func (c *Client) GetSpreadsheetCellThreads(ids []int64) (*SpreadsheetCellThreads, error) {
	scts := &SpreadsheetCellThreads{}
	if err := c.Read(SpreadsheetCellThreadModel, ids, nil, scts); err != nil {
		return nil, err
	}
	return scts, nil
}

// FindSpreadsheetCellThread finds spreadsheet.cell.thread record by querying it with criteria.
func (c *Client) FindSpreadsheetCellThread(criteria *Criteria) (*SpreadsheetCellThread, error) {
	scts := &SpreadsheetCellThreads{}
	if err := c.SearchRead(SpreadsheetCellThreadModel, criteria, NewOptions().Limit(1), scts); err != nil {
		return nil, err
	}
	return &((*scts)[0]), nil
}

// FindSpreadsheetCellThreads finds spreadsheet.cell.thread records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSpreadsheetCellThreads(criteria *Criteria, options *Options) (*SpreadsheetCellThreads, error) {
	scts := &SpreadsheetCellThreads{}
	if err := c.SearchRead(SpreadsheetCellThreadModel, criteria, options, scts); err != nil {
		return nil, err
	}
	return scts, nil
}

// FindSpreadsheetCellThreadIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSpreadsheetCellThreadIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SpreadsheetCellThreadModel, criteria, options)
}

// FindSpreadsheetCellThreadId finds record id by querying it with criteria.
func (c *Client) FindSpreadsheetCellThreadId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SpreadsheetCellThreadModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
