package odoo

// QualityCheckSpreadsheet represents quality.check.spreadsheet model.
type QualityCheckSpreadsheet struct {
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

// QualityCheckSpreadsheets represents array of quality.check.spreadsheet model.
type QualityCheckSpreadsheets []QualityCheckSpreadsheet

// QualityCheckSpreadsheetModel is the odoo model name.
const QualityCheckSpreadsheetModel = "quality.check.spreadsheet"

// Many2One convert QualityCheckSpreadsheet to *Many2One.
func (qcs *QualityCheckSpreadsheet) Many2One() *Many2One {
	return NewMany2One(qcs.Id.Get(), "")
}

// CreateQualityCheckSpreadsheet creates a new quality.check.spreadsheet model and returns its id.
func (c *Client) CreateQualityCheckSpreadsheet(qcs *QualityCheckSpreadsheet) (int64, error) {
	ids, err := c.CreateQualityCheckSpreadsheets([]*QualityCheckSpreadsheet{qcs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateQualityCheckSpreadsheet creates a new quality.check.spreadsheet model and returns its id.
func (c *Client) CreateQualityCheckSpreadsheets(qcss []*QualityCheckSpreadsheet) ([]int64, error) {
	var vv []interface{}
	for _, v := range qcss {
		vv = append(vv, v)
	}
	return c.Create(QualityCheckSpreadsheetModel, vv, nil)
}

// UpdateQualityCheckSpreadsheet updates an existing quality.check.spreadsheet record.
func (c *Client) UpdateQualityCheckSpreadsheet(qcs *QualityCheckSpreadsheet) error {
	return c.UpdateQualityCheckSpreadsheets([]int64{qcs.Id.Get()}, qcs)
}

// UpdateQualityCheckSpreadsheets updates existing quality.check.spreadsheet records.
// All records (represented by ids) will be updated by qcs values.
func (c *Client) UpdateQualityCheckSpreadsheets(ids []int64, qcs *QualityCheckSpreadsheet) error {
	return c.Update(QualityCheckSpreadsheetModel, ids, qcs, nil)
}

// DeleteQualityCheckSpreadsheet deletes an existing quality.check.spreadsheet record.
func (c *Client) DeleteQualityCheckSpreadsheet(id int64) error {
	return c.DeleteQualityCheckSpreadsheets([]int64{id})
}

// DeleteQualityCheckSpreadsheets deletes existing quality.check.spreadsheet records.
func (c *Client) DeleteQualityCheckSpreadsheets(ids []int64) error {
	return c.Delete(QualityCheckSpreadsheetModel, ids)
}

// GetQualityCheckSpreadsheet gets quality.check.spreadsheet existing record.
func (c *Client) GetQualityCheckSpreadsheet(id int64) (*QualityCheckSpreadsheet, error) {
	qcss, err := c.GetQualityCheckSpreadsheets([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*qcss)[0]), nil
}

// GetQualityCheckSpreadsheets gets quality.check.spreadsheet existing records.
func (c *Client) GetQualityCheckSpreadsheets(ids []int64) (*QualityCheckSpreadsheets, error) {
	qcss := &QualityCheckSpreadsheets{}
	if err := c.Read(QualityCheckSpreadsheetModel, ids, nil, qcss); err != nil {
		return nil, err
	}
	return qcss, nil
}

// FindQualityCheckSpreadsheet finds quality.check.spreadsheet record by querying it with criteria.
func (c *Client) FindQualityCheckSpreadsheet(criteria *Criteria) (*QualityCheckSpreadsheet, error) {
	qcss := &QualityCheckSpreadsheets{}
	if err := c.SearchRead(QualityCheckSpreadsheetModel, criteria, NewOptions().Limit(1), qcss); err != nil {
		return nil, err
	}
	return &((*qcss)[0]), nil
}

// FindQualityCheckSpreadsheets finds quality.check.spreadsheet records by querying it
// and filtering it with criteria and options.
func (c *Client) FindQualityCheckSpreadsheets(criteria *Criteria, options *Options) (*QualityCheckSpreadsheets, error) {
	qcss := &QualityCheckSpreadsheets{}
	if err := c.SearchRead(QualityCheckSpreadsheetModel, criteria, options, qcss); err != nil {
		return nil, err
	}
	return qcss, nil
}

// FindQualityCheckSpreadsheetIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindQualityCheckSpreadsheetIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(QualityCheckSpreadsheetModel, criteria, options)
}

// FindQualityCheckSpreadsheetId finds record id by querying it with criteria.
func (c *Client) FindQualityCheckSpreadsheetId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(QualityCheckSpreadsheetModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
