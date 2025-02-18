package odoo

// StockReport represents stock.report model.
type StockReport struct {
	CategId         *Many2One  `xmlrpc:"categ_id,omitempty"`
	CompanyId       *Many2One  `xmlrpc:"company_id,omitempty"`
	CreationDate    *Time      `xmlrpc:"creation_date,omitempty"`
	CycleTime       *Float     `xmlrpc:"cycle_time,omitempty"`
	DateDone        *Time      `xmlrpc:"date_done,omitempty"`
	Delay           *Float     `xmlrpc:"delay,omitempty"`
	DisplayName     *String    `xmlrpc:"display_name,omitempty"`
	Id              *Int       `xmlrpc:"id,omitempty"`
	IsBackorder     *Bool      `xmlrpc:"is_backorder,omitempty"`
	IsLate          *Bool      `xmlrpc:"is_late,omitempty"`
	OperationTypeId *Many2One  `xmlrpc:"operation_type_id,omitempty"`
	PartnerId       *Many2One  `xmlrpc:"partner_id,omitempty"`
	PickingId       *Many2One  `xmlrpc:"picking_id,omitempty"`
	PickingName     *String    `xmlrpc:"picking_name,omitempty"`
	PickingTypeCode *Selection `xmlrpc:"picking_type_code,omitempty"`
	ProductId       *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductQty      *Float     `xmlrpc:"product_qty,omitempty"`
	Reference       *String    `xmlrpc:"reference,omitempty"`
	ScheduledDate   *Time      `xmlrpc:"scheduled_date,omitempty"`
	State           *Selection `xmlrpc:"state,omitempty"`
}

// StockReports represents array of stock.report model.
type StockReports []StockReport

// StockReportModel is the odoo model name.
const StockReportModel = "stock.report"

// Many2One convert StockReport to *Many2One.
func (sr *StockReport) Many2One() *Many2One {
	return NewMany2One(sr.Id.Get(), "")
}

// CreateStockReport creates a new stock.report model and returns its id.
func (c *Client) CreateStockReport(sr *StockReport) (int64, error) {
	ids, err := c.CreateStockReports([]*StockReport{sr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockReport creates a new stock.report model and returns its id.
func (c *Client) CreateStockReports(srs []*StockReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range srs {
		vv = append(vv, v)
	}
	return c.Create(StockReportModel, vv, nil)
}

// UpdateStockReport updates an existing stock.report record.
func (c *Client) UpdateStockReport(sr *StockReport) error {
	return c.UpdateStockReports([]int64{sr.Id.Get()}, sr)
}

// UpdateStockReports updates existing stock.report records.
// All records (represented by ids) will be updated by sr values.
func (c *Client) UpdateStockReports(ids []int64, sr *StockReport) error {
	return c.Update(StockReportModel, ids, sr, nil)
}

// DeleteStockReport deletes an existing stock.report record.
func (c *Client) DeleteStockReport(id int64) error {
	return c.DeleteStockReports([]int64{id})
}

// DeleteStockReports deletes existing stock.report records.
func (c *Client) DeleteStockReports(ids []int64) error {
	return c.Delete(StockReportModel, ids)
}

// GetStockReport gets stock.report existing record.
func (c *Client) GetStockReport(id int64) (*StockReport, error) {
	srs, err := c.GetStockReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*srs)[0]), nil
}

// GetStockReports gets stock.report existing records.
func (c *Client) GetStockReports(ids []int64) (*StockReports, error) {
	srs := &StockReports{}
	if err := c.Read(StockReportModel, ids, nil, srs); err != nil {
		return nil, err
	}
	return srs, nil
}

// FindStockReport finds stock.report record by querying it with criteria.
func (c *Client) FindStockReport(criteria *Criteria) (*StockReport, error) {
	srs := &StockReports{}
	if err := c.SearchRead(StockReportModel, criteria, NewOptions().Limit(1), srs); err != nil {
		return nil, err
	}
	return &((*srs)[0]), nil
}

// FindStockReports finds stock.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockReports(criteria *Criteria, options *Options) (*StockReports, error) {
	srs := &StockReports{}
	if err := c.SearchRead(StockReportModel, criteria, options, srs); err != nil {
		return nil, err
	}
	return srs, nil
}

// FindStockReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockReportModel, criteria, options)
}

// FindStockReportId finds record id by querying it with criteria.
func (c *Client) FindStockReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
