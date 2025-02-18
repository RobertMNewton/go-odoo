package odoo

// SaveSpreadsheetTemplate represents save.spreadsheet.template model.
type SaveSpreadsheetTemplate struct {
	CreateDate             *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid              *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrentRevisionUuid    *String   `xmlrpc:"current_revision_uuid,omitempty"`
	DisplayName            *String   `xmlrpc:"display_name,omitempty"`
	Id                     *Int      `xmlrpc:"id,omitempty"`
	SpreadsheetBinaryData  *String   `xmlrpc:"spreadsheet_binary_data,omitempty"`
	SpreadsheetData        *String   `xmlrpc:"spreadsheet_data,omitempty"`
	SpreadsheetFileName    *String   `xmlrpc:"spreadsheet_file_name,omitempty"`
	SpreadsheetRevisionIds *Relation `xmlrpc:"spreadsheet_revision_ids,omitempty"`
	SpreadsheetSnapshot    *String   `xmlrpc:"spreadsheet_snapshot,omitempty"`
	TemplateName           *String   `xmlrpc:"template_name,omitempty"`
	Thumbnail              *String   `xmlrpc:"thumbnail,omitempty"`
	WriteDate              *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid               *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SaveSpreadsheetTemplates represents array of save.spreadsheet.template model.
type SaveSpreadsheetTemplates []SaveSpreadsheetTemplate

// SaveSpreadsheetTemplateModel is the odoo model name.
const SaveSpreadsheetTemplateModel = "save.spreadsheet.template"

// Many2One convert SaveSpreadsheetTemplate to *Many2One.
func (sst *SaveSpreadsheetTemplate) Many2One() *Many2One {
	return NewMany2One(sst.Id.Get(), "")
}

// CreateSaveSpreadsheetTemplate creates a new save.spreadsheet.template model and returns its id.
func (c *Client) CreateSaveSpreadsheetTemplate(sst *SaveSpreadsheetTemplate) (int64, error) {
	ids, err := c.CreateSaveSpreadsheetTemplates([]*SaveSpreadsheetTemplate{sst})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaveSpreadsheetTemplate creates a new save.spreadsheet.template model and returns its id.
func (c *Client) CreateSaveSpreadsheetTemplates(ssts []*SaveSpreadsheetTemplate) ([]int64, error) {
	var vv []interface{}
	for _, v := range ssts {
		vv = append(vv, v)
	}
	return c.Create(SaveSpreadsheetTemplateModel, vv, nil)
}

// UpdateSaveSpreadsheetTemplate updates an existing save.spreadsheet.template record.
func (c *Client) UpdateSaveSpreadsheetTemplate(sst *SaveSpreadsheetTemplate) error {
	return c.UpdateSaveSpreadsheetTemplates([]int64{sst.Id.Get()}, sst)
}

// UpdateSaveSpreadsheetTemplates updates existing save.spreadsheet.template records.
// All records (represented by ids) will be updated by sst values.
func (c *Client) UpdateSaveSpreadsheetTemplates(ids []int64, sst *SaveSpreadsheetTemplate) error {
	return c.Update(SaveSpreadsheetTemplateModel, ids, sst, nil)
}

// DeleteSaveSpreadsheetTemplate deletes an existing save.spreadsheet.template record.
func (c *Client) DeleteSaveSpreadsheetTemplate(id int64) error {
	return c.DeleteSaveSpreadsheetTemplates([]int64{id})
}

// DeleteSaveSpreadsheetTemplates deletes existing save.spreadsheet.template records.
func (c *Client) DeleteSaveSpreadsheetTemplates(ids []int64) error {
	return c.Delete(SaveSpreadsheetTemplateModel, ids)
}

// GetSaveSpreadsheetTemplate gets save.spreadsheet.template existing record.
func (c *Client) GetSaveSpreadsheetTemplate(id int64) (*SaveSpreadsheetTemplate, error) {
	ssts, err := c.GetSaveSpreadsheetTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ssts)[0]), nil
}

// GetSaveSpreadsheetTemplates gets save.spreadsheet.template existing records.
func (c *Client) GetSaveSpreadsheetTemplates(ids []int64) (*SaveSpreadsheetTemplates, error) {
	ssts := &SaveSpreadsheetTemplates{}
	if err := c.Read(SaveSpreadsheetTemplateModel, ids, nil, ssts); err != nil {
		return nil, err
	}
	return ssts, nil
}

// FindSaveSpreadsheetTemplate finds save.spreadsheet.template record by querying it with criteria.
func (c *Client) FindSaveSpreadsheetTemplate(criteria *Criteria) (*SaveSpreadsheetTemplate, error) {
	ssts := &SaveSpreadsheetTemplates{}
	if err := c.SearchRead(SaveSpreadsheetTemplateModel, criteria, NewOptions().Limit(1), ssts); err != nil {
		return nil, err
	}
	return &((*ssts)[0]), nil
}

// FindSaveSpreadsheetTemplates finds save.spreadsheet.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaveSpreadsheetTemplates(criteria *Criteria, options *Options) (*SaveSpreadsheetTemplates, error) {
	ssts := &SaveSpreadsheetTemplates{}
	if err := c.SearchRead(SaveSpreadsheetTemplateModel, criteria, options, ssts); err != nil {
		return nil, err
	}
	return ssts, nil
}

// FindSaveSpreadsheetTemplateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaveSpreadsheetTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SaveSpreadsheetTemplateModel, criteria, options)
}

// FindSaveSpreadsheetTemplateId finds record id by querying it with criteria.
func (c *Client) FindSaveSpreadsheetTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaveSpreadsheetTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
