package odoo

// QualitySpreadsheetTemplate represents quality.spreadsheet.template model.
type QualitySpreadsheetTemplate struct {
	CheckCell              *String   `xmlrpc:"check_cell,omitempty"`
	CompanyId              *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate             *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid              *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrentRevisionUuid    *String   `xmlrpc:"current_revision_uuid,omitempty"`
	DisplayName            *String   `xmlrpc:"display_name,omitempty"`
	Id                     *Int      `xmlrpc:"id,omitempty"`
	Name                   *String   `xmlrpc:"name,omitempty"`
	SpreadsheetBinaryData  *String   `xmlrpc:"spreadsheet_binary_data,omitempty"`
	SpreadsheetData        *String   `xmlrpc:"spreadsheet_data,omitempty"`
	SpreadsheetFileName    *String   `xmlrpc:"spreadsheet_file_name,omitempty"`
	SpreadsheetRevisionIds *Relation `xmlrpc:"spreadsheet_revision_ids,omitempty"`
	SpreadsheetSnapshot    *String   `xmlrpc:"spreadsheet_snapshot,omitempty"`
	Thumbnail              *String   `xmlrpc:"thumbnail,omitempty"`
	WriteDate              *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid               *Many2One `xmlrpc:"write_uid,omitempty"`
}

// QualitySpreadsheetTemplates represents array of quality.spreadsheet.template model.
type QualitySpreadsheetTemplates []QualitySpreadsheetTemplate

// QualitySpreadsheetTemplateModel is the odoo model name.
const QualitySpreadsheetTemplateModel = "quality.spreadsheet.template"

// Many2One convert QualitySpreadsheetTemplate to *Many2One.
func (qst *QualitySpreadsheetTemplate) Many2One() *Many2One {
	return NewMany2One(qst.Id.Get(), "")
}

// CreateQualitySpreadsheetTemplate creates a new quality.spreadsheet.template model and returns its id.
func (c *Client) CreateQualitySpreadsheetTemplate(qst *QualitySpreadsheetTemplate) (int64, error) {
	ids, err := c.CreateQualitySpreadsheetTemplates([]*QualitySpreadsheetTemplate{qst})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateQualitySpreadsheetTemplate creates a new quality.spreadsheet.template model and returns its id.
func (c *Client) CreateQualitySpreadsheetTemplates(qsts []*QualitySpreadsheetTemplate) ([]int64, error) {
	var vv []interface{}
	for _, v := range qsts {
		vv = append(vv, v)
	}
	return c.Create(QualitySpreadsheetTemplateModel, vv, nil)
}

// UpdateQualitySpreadsheetTemplate updates an existing quality.spreadsheet.template record.
func (c *Client) UpdateQualitySpreadsheetTemplate(qst *QualitySpreadsheetTemplate) error {
	return c.UpdateQualitySpreadsheetTemplates([]int64{qst.Id.Get()}, qst)
}

// UpdateQualitySpreadsheetTemplates updates existing quality.spreadsheet.template records.
// All records (represented by ids) will be updated by qst values.
func (c *Client) UpdateQualitySpreadsheetTemplates(ids []int64, qst *QualitySpreadsheetTemplate) error {
	return c.Update(QualitySpreadsheetTemplateModel, ids, qst, nil)
}

// DeleteQualitySpreadsheetTemplate deletes an existing quality.spreadsheet.template record.
func (c *Client) DeleteQualitySpreadsheetTemplate(id int64) error {
	return c.DeleteQualitySpreadsheetTemplates([]int64{id})
}

// DeleteQualitySpreadsheetTemplates deletes existing quality.spreadsheet.template records.
func (c *Client) DeleteQualitySpreadsheetTemplates(ids []int64) error {
	return c.Delete(QualitySpreadsheetTemplateModel, ids)
}

// GetQualitySpreadsheetTemplate gets quality.spreadsheet.template existing record.
func (c *Client) GetQualitySpreadsheetTemplate(id int64) (*QualitySpreadsheetTemplate, error) {
	qsts, err := c.GetQualitySpreadsheetTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*qsts)[0]), nil
}

// GetQualitySpreadsheetTemplates gets quality.spreadsheet.template existing records.
func (c *Client) GetQualitySpreadsheetTemplates(ids []int64) (*QualitySpreadsheetTemplates, error) {
	qsts := &QualitySpreadsheetTemplates{}
	if err := c.Read(QualitySpreadsheetTemplateModel, ids, nil, qsts); err != nil {
		return nil, err
	}
	return qsts, nil
}

// FindQualitySpreadsheetTemplate finds quality.spreadsheet.template record by querying it with criteria.
func (c *Client) FindQualitySpreadsheetTemplate(criteria *Criteria) (*QualitySpreadsheetTemplate, error) {
	qsts := &QualitySpreadsheetTemplates{}
	if err := c.SearchRead(QualitySpreadsheetTemplateModel, criteria, NewOptions().Limit(1), qsts); err != nil {
		return nil, err
	}
	return &((*qsts)[0]), nil
}

// FindQualitySpreadsheetTemplates finds quality.spreadsheet.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindQualitySpreadsheetTemplates(criteria *Criteria, options *Options) (*QualitySpreadsheetTemplates, error) {
	qsts := &QualitySpreadsheetTemplates{}
	if err := c.SearchRead(QualitySpreadsheetTemplateModel, criteria, options, qsts); err != nil {
		return nil, err
	}
	return qsts, nil
}

// FindQualitySpreadsheetTemplateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindQualitySpreadsheetTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(QualitySpreadsheetTemplateModel, criteria, options)
}

// FindQualitySpreadsheetTemplateId finds record id by querying it with criteria.
func (c *Client) FindQualitySpreadsheetTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(QualitySpreadsheetTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
