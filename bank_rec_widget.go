package odoo

// BankRecWidget represents bank.rec.widget model.
type BankRecWidget struct {
	AvailableRecoModelIds           *Relation   `xmlrpc:"available_reco_model_ids,omitempty"`
	CompanyCurrencyId               *Many2One   `xmlrpc:"company_currency_id,omitempty"`
	CompanyId                       *Many2One   `xmlrpc:"company_id,omitempty"`
	CountryCode                     *String     `xmlrpc:"country_code,omitempty"`
	DisplayName                     *String     `xmlrpc:"display_name,omitempty"`
	FormIndex                       *String     `xmlrpc:"form_index,omitempty"`
	Id                              *Int        `xmlrpc:"id,omitempty"`
	IsMultiCurrency                 *Bool       `xmlrpc:"is_multi_currency,omitempty"`
	JournalCurrencyId               *Many2One   `xmlrpc:"journal_currency_id,omitempty"`
	LineIds                         *Relation   `xmlrpc:"line_ids,omitempty"`
	MatchedSaleOrderIds             *Relation   `xmlrpc:"matched_sale_order_ids,omitempty"`
	MatchingRulesAllowAutoReconcile *Bool       `xmlrpc:"matching_rules_allow_auto_reconcile,omitempty"`
	MoveId                          *Many2One   `xmlrpc:"move_id,omitempty"`
	PartnerId                       *Many2One   `xmlrpc:"partner_id,omitempty"`
	PartnerName                     *String     `xmlrpc:"partner_name,omitempty"`
	ReturnTodoCommand               interface{} `xmlrpc:"return_todo_command,omitempty"`
	SelectedAmlIds                  *Relation   `xmlrpc:"selected_aml_ids,omitempty"`
	SelectedBatchPaymentIds         *Relation   `xmlrpc:"selected_batch_payment_ids,omitempty"`
	SelectedRecoModelId             *Many2One   `xmlrpc:"selected_reco_model_id,omitempty"`
	StLineChecked                   *Bool       `xmlrpc:"st_line_checked,omitempty"`
	StLineId                        *Many2One   `xmlrpc:"st_line_id,omitempty"`
	StLineIsReconciled              *Bool       `xmlrpc:"st_line_is_reconciled,omitempty"`
	StLineJournalId                 *Many2One   `xmlrpc:"st_line_journal_id,omitempty"`
	StLineTransactionDetails        *String     `xmlrpc:"st_line_transaction_details,omitempty"`
	State                           *Selection  `xmlrpc:"state,omitempty"`
	TodoCommand                     interface{} `xmlrpc:"todo_command,omitempty"`
	TransactionCurrencyId           *Many2One   `xmlrpc:"transaction_currency_id,omitempty"`
}

// BankRecWidgets represents array of bank.rec.widget model.
type BankRecWidgets []BankRecWidget

// BankRecWidgetModel is the odoo model name.
const BankRecWidgetModel = "bank.rec.widget"

// Many2One convert BankRecWidget to *Many2One.
func (brw *BankRecWidget) Many2One() *Many2One {
	return NewMany2One(brw.Id.Get(), "")
}

// CreateBankRecWidget creates a new bank.rec.widget model and returns its id.
func (c *Client) CreateBankRecWidget(brw *BankRecWidget) (int64, error) {
	ids, err := c.CreateBankRecWidgets([]*BankRecWidget{brw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBankRecWidget creates a new bank.rec.widget model and returns its id.
func (c *Client) CreateBankRecWidgets(brws []*BankRecWidget) ([]int64, error) {
	var vv []interface{}
	for _, v := range brws {
		vv = append(vv, v)
	}
	return c.Create(BankRecWidgetModel, vv, nil)
}

// UpdateBankRecWidget updates an existing bank.rec.widget record.
func (c *Client) UpdateBankRecWidget(brw *BankRecWidget) error {
	return c.UpdateBankRecWidgets([]int64{brw.Id.Get()}, brw)
}

// UpdateBankRecWidgets updates existing bank.rec.widget records.
// All records (represented by ids) will be updated by brw values.
func (c *Client) UpdateBankRecWidgets(ids []int64, brw *BankRecWidget) error {
	return c.Update(BankRecWidgetModel, ids, brw, nil)
}

// DeleteBankRecWidget deletes an existing bank.rec.widget record.
func (c *Client) DeleteBankRecWidget(id int64) error {
	return c.DeleteBankRecWidgets([]int64{id})
}

// DeleteBankRecWidgets deletes existing bank.rec.widget records.
func (c *Client) DeleteBankRecWidgets(ids []int64) error {
	return c.Delete(BankRecWidgetModel, ids)
}

// GetBankRecWidget gets bank.rec.widget existing record.
func (c *Client) GetBankRecWidget(id int64) (*BankRecWidget, error) {
	brws, err := c.GetBankRecWidgets([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*brws)[0]), nil
}

// GetBankRecWidgets gets bank.rec.widget existing records.
func (c *Client) GetBankRecWidgets(ids []int64) (*BankRecWidgets, error) {
	brws := &BankRecWidgets{}
	if err := c.Read(BankRecWidgetModel, ids, nil, brws); err != nil {
		return nil, err
	}
	return brws, nil
}

// FindBankRecWidget finds bank.rec.widget record by querying it with criteria.
func (c *Client) FindBankRecWidget(criteria *Criteria) (*BankRecWidget, error) {
	brws := &BankRecWidgets{}
	if err := c.SearchRead(BankRecWidgetModel, criteria, NewOptions().Limit(1), brws); err != nil {
		return nil, err
	}
	return &((*brws)[0]), nil
}

// FindBankRecWidgets finds bank.rec.widget records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBankRecWidgets(criteria *Criteria, options *Options) (*BankRecWidgets, error) {
	brws := &BankRecWidgets{}
	if err := c.SearchRead(BankRecWidgetModel, criteria, options, brws); err != nil {
		return nil, err
	}
	return brws, nil
}

// FindBankRecWidgetIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBankRecWidgetIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(BankRecWidgetModel, criteria, options)
}

// FindBankRecWidgetId finds record id by querying it with criteria.
func (c *Client) FindBankRecWidgetId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BankRecWidgetModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
