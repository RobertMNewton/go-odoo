package odoo

// StockBarcodeCancelOperation represents stock_barcode.cancel.operation model.
type StockBarcodeCancelOperation struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	PickingId   *Many2One `xmlrpc:"picking_id,omitempty"`
	PickingName *String   `xmlrpc:"picking_name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StockBarcodeCancelOperations represents array of stock_barcode.cancel.operation model.
type StockBarcodeCancelOperations []StockBarcodeCancelOperation

// StockBarcodeCancelOperationModel is the odoo model name.
const StockBarcodeCancelOperationModel = "stock_barcode.cancel.operation"

// Many2One convert StockBarcodeCancelOperation to *Many2One.
func (sco *StockBarcodeCancelOperation) Many2One() *Many2One {
	return NewMany2One(sco.Id.Get(), "")
}

// CreateStockBarcodeCancelOperation creates a new stock_barcode.cancel.operation model and returns its id.
func (c *Client) CreateStockBarcodeCancelOperation(sco *StockBarcodeCancelOperation) (int64, error) {
	ids, err := c.CreateStockBarcodeCancelOperations([]*StockBarcodeCancelOperation{sco})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockBarcodeCancelOperation creates a new stock_barcode.cancel.operation model and returns its id.
func (c *Client) CreateStockBarcodeCancelOperations(scos []*StockBarcodeCancelOperation) ([]int64, error) {
	var vv []interface{}
	for _, v := range scos {
		vv = append(vv, v)
	}
	return c.Create(StockBarcodeCancelOperationModel, vv, nil)
}

// UpdateStockBarcodeCancelOperation updates an existing stock_barcode.cancel.operation record.
func (c *Client) UpdateStockBarcodeCancelOperation(sco *StockBarcodeCancelOperation) error {
	return c.UpdateStockBarcodeCancelOperations([]int64{sco.Id.Get()}, sco)
}

// UpdateStockBarcodeCancelOperations updates existing stock_barcode.cancel.operation records.
// All records (represented by ids) will be updated by sco values.
func (c *Client) UpdateStockBarcodeCancelOperations(ids []int64, sco *StockBarcodeCancelOperation) error {
	return c.Update(StockBarcodeCancelOperationModel, ids, sco, nil)
}

// DeleteStockBarcodeCancelOperation deletes an existing stock_barcode.cancel.operation record.
func (c *Client) DeleteStockBarcodeCancelOperation(id int64) error {
	return c.DeleteStockBarcodeCancelOperations([]int64{id})
}

// DeleteStockBarcodeCancelOperations deletes existing stock_barcode.cancel.operation records.
func (c *Client) DeleteStockBarcodeCancelOperations(ids []int64) error {
	return c.Delete(StockBarcodeCancelOperationModel, ids)
}

// GetStockBarcodeCancelOperation gets stock_barcode.cancel.operation existing record.
func (c *Client) GetStockBarcodeCancelOperation(id int64) (*StockBarcodeCancelOperation, error) {
	scos, err := c.GetStockBarcodeCancelOperations([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*scos)[0]), nil
}

// GetStockBarcodeCancelOperations gets stock_barcode.cancel.operation existing records.
func (c *Client) GetStockBarcodeCancelOperations(ids []int64) (*StockBarcodeCancelOperations, error) {
	scos := &StockBarcodeCancelOperations{}
	if err := c.Read(StockBarcodeCancelOperationModel, ids, nil, scos); err != nil {
		return nil, err
	}
	return scos, nil
}

// FindStockBarcodeCancelOperation finds stock_barcode.cancel.operation record by querying it with criteria.
func (c *Client) FindStockBarcodeCancelOperation(criteria *Criteria) (*StockBarcodeCancelOperation, error) {
	scos := &StockBarcodeCancelOperations{}
	if err := c.SearchRead(StockBarcodeCancelOperationModel, criteria, NewOptions().Limit(1), scos); err != nil {
		return nil, err
	}
	return &((*scos)[0]), nil
}

// FindStockBarcodeCancelOperations finds stock_barcode.cancel.operation records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockBarcodeCancelOperations(criteria *Criteria, options *Options) (*StockBarcodeCancelOperations, error) {
	scos := &StockBarcodeCancelOperations{}
	if err := c.SearchRead(StockBarcodeCancelOperationModel, criteria, options, scos); err != nil {
		return nil, err
	}
	return scos, nil
}

// FindStockBarcodeCancelOperationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockBarcodeCancelOperationIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockBarcodeCancelOperationModel, criteria, options)
}

// FindStockBarcodeCancelOperationId finds record id by querying it with criteria.
func (c *Client) FindStockBarcodeCancelOperationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockBarcodeCancelOperationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
