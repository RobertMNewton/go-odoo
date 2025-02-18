package odoo

// IrActionsAccountReportDownload represents ir_actions_account_report_download model.
type IrActionsAccountReportDownload struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// IrActionsAccountReportDownloads represents array of ir_actions_account_report_download model.
type IrActionsAccountReportDownloads []IrActionsAccountReportDownload

// IrActionsAccountReportDownloadModel is the odoo model name.
const IrActionsAccountReportDownloadModel = "ir_actions_account_report_download"

// Many2One convert IrActionsAccountReportDownload to *Many2One.
func (i *IrActionsAccountReportDownload) Many2One() *Many2One {
	return NewMany2One(i.Id.Get(), "")
}

// CreateIrActionsAccountReportDownload creates a new ir_actions_account_report_download model and returns its id.
func (c *Client) CreateIrActionsAccountReportDownload(i *IrActionsAccountReportDownload) (int64, error) {
	ids, err := c.CreateIrActionsAccountReportDownloads([]*IrActionsAccountReportDownload{i})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrActionsAccountReportDownload creates a new ir_actions_account_report_download model and returns its id.
func (c *Client) CreateIrActionsAccountReportDownloads(is []*IrActionsAccountReportDownload) ([]int64, error) {
	var vv []interface{}
	for _, v := range is {
		vv = append(vv, v)
	}
	return c.Create(IrActionsAccountReportDownloadModel, vv, nil)
}

// UpdateIrActionsAccountReportDownload updates an existing ir_actions_account_report_download record.
func (c *Client) UpdateIrActionsAccountReportDownload(i *IrActionsAccountReportDownload) error {
	return c.UpdateIrActionsAccountReportDownloads([]int64{i.Id.Get()}, i)
}

// UpdateIrActionsAccountReportDownloads updates existing ir_actions_account_report_download records.
// All records (represented by ids) will be updated by i values.
func (c *Client) UpdateIrActionsAccountReportDownloads(ids []int64, i *IrActionsAccountReportDownload) error {
	return c.Update(IrActionsAccountReportDownloadModel, ids, i, nil)
}

// DeleteIrActionsAccountReportDownload deletes an existing ir_actions_account_report_download record.
func (c *Client) DeleteIrActionsAccountReportDownload(id int64) error {
	return c.DeleteIrActionsAccountReportDownloads([]int64{id})
}

// DeleteIrActionsAccountReportDownloads deletes existing ir_actions_account_report_download records.
func (c *Client) DeleteIrActionsAccountReportDownloads(ids []int64) error {
	return c.Delete(IrActionsAccountReportDownloadModel, ids)
}

// GetIrActionsAccountReportDownload gets ir_actions_account_report_download existing record.
func (c *Client) GetIrActionsAccountReportDownload(id int64) (*IrActionsAccountReportDownload, error) {
	is, err := c.GetIrActionsAccountReportDownloads([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*is)[0]), nil
}

// GetIrActionsAccountReportDownloads gets ir_actions_account_report_download existing records.
func (c *Client) GetIrActionsAccountReportDownloads(ids []int64) (*IrActionsAccountReportDownloads, error) {
	is := &IrActionsAccountReportDownloads{}
	if err := c.Read(IrActionsAccountReportDownloadModel, ids, nil, is); err != nil {
		return nil, err
	}
	return is, nil
}

// FindIrActionsAccountReportDownload finds ir_actions_account_report_download record by querying it with criteria.
func (c *Client) FindIrActionsAccountReportDownload(criteria *Criteria) (*IrActionsAccountReportDownload, error) {
	is := &IrActionsAccountReportDownloads{}
	if err := c.SearchRead(IrActionsAccountReportDownloadModel, criteria, NewOptions().Limit(1), is); err != nil {
		return nil, err
	}
	return &((*is)[0]), nil
}

// FindIrActionsAccountReportDownloads finds ir_actions_account_report_download records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrActionsAccountReportDownloads(criteria *Criteria, options *Options) (*IrActionsAccountReportDownloads, error) {
	is := &IrActionsAccountReportDownloads{}
	if err := c.SearchRead(IrActionsAccountReportDownloadModel, criteria, options, is); err != nil {
		return nil, err
	}
	return is, nil
}

// FindIrActionsAccountReportDownloadIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrActionsAccountReportDownloadIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrActionsAccountReportDownloadModel, criteria, options)
}

// FindIrActionsAccountReportDownloadId finds record id by querying it with criteria.
func (c *Client) FindIrActionsAccountReportDownloadId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrActionsAccountReportDownloadModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
