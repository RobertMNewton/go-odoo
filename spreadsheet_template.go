package odoo

// SpreadsheetTemplate represents spreadsheet.template model.
type SpreadsheetTemplate struct {
	CreateDate             *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid              *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrentRevisionUuid    *String   `xmlrpc:"current_revision_uuid,omitempty"`
	DisplayName            *String   `xmlrpc:"display_name,omitempty"`
	FileName               *String   `xmlrpc:"file_name,omitempty"`
	Id                     *Int      `xmlrpc:"id,omitempty"`
	Name                   *String   `xmlrpc:"name,omitempty"`
	Sequence               *Int      `xmlrpc:"sequence,omitempty"`
	SpreadsheetBinaryData  *String   `xmlrpc:"spreadsheet_binary_data,omitempty"`
	SpreadsheetData        *String   `xmlrpc:"spreadsheet_data,omitempty"`
	SpreadsheetFileName    *String   `xmlrpc:"spreadsheet_file_name,omitempty"`
	SpreadsheetRevisionIds *Relation `xmlrpc:"spreadsheet_revision_ids,omitempty"`
	SpreadsheetSnapshot    *String   `xmlrpc:"spreadsheet_snapshot,omitempty"`
	Thumbnail              *String   `xmlrpc:"thumbnail,omitempty"`
	WriteDate              *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid               *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SpreadsheetTemplates represents array of spreadsheet.template model.
type SpreadsheetTemplates []SpreadsheetTemplate

// SpreadsheetTemplateModel is the odoo model name.
const SpreadsheetTemplateModel = "spreadsheet.template"

// Many2One convert SpreadsheetTemplate to *Many2One.
func (st *SpreadsheetTemplate) Many2One() *Many2One {
	return NewMany2One(st.Id.Get(), "")
}

// CreateSpreadsheetTemplate creates a new spreadsheet.template model and returns its id.
func (c *Client) CreateSpreadsheetTemplate(st *SpreadsheetTemplate) (int64, error) {
	ids, err := c.CreateSpreadsheetTemplates([]*SpreadsheetTemplate{st})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSpreadsheetTemplate creates a new spreadsheet.template model and returns its id.
func (c *Client) CreateSpreadsheetTemplates(sts []*SpreadsheetTemplate) ([]int64, error) {
	var vv []interface{}
	for _, v := range sts {
		vv = append(vv, v)
	}
	return c.Create(SpreadsheetTemplateModel, vv, nil)
}

// UpdateSpreadsheetTemplate updates an existing spreadsheet.template record.
func (c *Client) UpdateSpreadsheetTemplate(st *SpreadsheetTemplate) error {
	return c.UpdateSpreadsheetTemplates([]int64{st.Id.Get()}, st)
}

// UpdateSpreadsheetTemplates updates existing spreadsheet.template records.
// All records (represented by ids) will be updated by st values.
func (c *Client) UpdateSpreadsheetTemplates(ids []int64, st *SpreadsheetTemplate) error {
	return c.Update(SpreadsheetTemplateModel, ids, st, nil)
}

// DeleteSpreadsheetTemplate deletes an existing spreadsheet.template record.
func (c *Client) DeleteSpreadsheetTemplate(id int64) error {
	return c.DeleteSpreadsheetTemplates([]int64{id})
}

// DeleteSpreadsheetTemplates deletes existing spreadsheet.template records.
func (c *Client) DeleteSpreadsheetTemplates(ids []int64) error {
	return c.Delete(SpreadsheetTemplateModel, ids)
}

// GetSpreadsheetTemplate gets spreadsheet.template existing record.
func (c *Client) GetSpreadsheetTemplate(id int64) (*SpreadsheetTemplate, error) {
	sts, err := c.GetSpreadsheetTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sts)[0]), nil
}

// GetSpreadsheetTemplates gets spreadsheet.template existing records.
func (c *Client) GetSpreadsheetTemplates(ids []int64) (*SpreadsheetTemplates, error) {
	sts := &SpreadsheetTemplates{}
	if err := c.Read(SpreadsheetTemplateModel, ids, nil, sts); err != nil {
		return nil, err
	}
	return sts, nil
}

// FindSpreadsheetTemplate finds spreadsheet.template record by querying it with criteria.
func (c *Client) FindSpreadsheetTemplate(criteria *Criteria) (*SpreadsheetTemplate, error) {
	sts := &SpreadsheetTemplates{}
	if err := c.SearchRead(SpreadsheetTemplateModel, criteria, NewOptions().Limit(1), sts); err != nil {
		return nil, err
	}
	return &((*sts)[0]), nil
}

// FindSpreadsheetTemplates finds spreadsheet.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSpreadsheetTemplates(criteria *Criteria, options *Options) (*SpreadsheetTemplates, error) {
	sts := &SpreadsheetTemplates{}
	if err := c.SearchRead(SpreadsheetTemplateModel, criteria, options, sts); err != nil {
		return nil, err
	}
	return sts, nil
}

// FindSpreadsheetTemplateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSpreadsheetTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SpreadsheetTemplateModel, criteria, options)
}

// FindSpreadsheetTemplateId finds record id by querying it with criteria.
func (c *Client) FindSpreadsheetTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SpreadsheetTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
