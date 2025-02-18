package odoo

// BankRecWidgetLine represents bank.rec.widget.line model.
type BankRecWidgetLine struct {
	AccountId                      *Many2One   `xmlrpc:"account_id,omitempty"`
	AmountCurrency                 *Float      `xmlrpc:"amount_currency,omitempty"`
	AmountTransactionCurrency      *Float      `xmlrpc:"amount_transaction_currency,omitempty"`
	AnalyticDistribution           interface{} `xmlrpc:"analytic_distribution,omitempty"`
	AnalyticPrecision              *Int        `xmlrpc:"analytic_precision,omitempty"`
	Balance                        *Float      `xmlrpc:"balance,omitempty"`
	BankAccount                    *String     `xmlrpc:"bank_account,omitempty"`
	CompanyCurrencyId              *Many2One   `xmlrpc:"company_currency_id,omitempty"`
	CompanyId                      *Many2One   `xmlrpc:"company_id,omitempty"`
	CountryCode                    *String     `xmlrpc:"country_code,omitempty"`
	Credit                         *Float      `xmlrpc:"credit,omitempty"`
	CurrencyId                     *Many2One   `xmlrpc:"currency_id,omitempty"`
	Date                           *Time       `xmlrpc:"date,omitempty"`
	Debit                          *Float      `xmlrpc:"debit,omitempty"`
	DisplayName                    *String     `xmlrpc:"display_name,omitempty"`
	DisplayStrokedAmountCurrency   *Bool       `xmlrpc:"display_stroked_amount_currency,omitempty"`
	DisplayStrokedBalance          *Bool       `xmlrpc:"display_stroked_balance,omitempty"`
	DistributionAnalyticAccountIds *Relation   `xmlrpc:"distribution_analytic_account_ids,omitempty"`
	Flag                           *Selection  `xmlrpc:"flag,omitempty"`
	ForcePriceIncludedTaxes        *Bool       `xmlrpc:"force_price_included_taxes,omitempty"`
	GroupTaxId                     *Many2One   `xmlrpc:"group_tax_id,omitempty"`
	Id                             *Int        `xmlrpc:"id,omitempty"`
	Index                          *String     `xmlrpc:"index,omitempty"`
	JournalDefaultAccountId        *Many2One   `xmlrpc:"journal_default_account_id,omitempty"`
	ManuallyModified               *Bool       `xmlrpc:"manually_modified,omitempty"`
	Name                           *String     `xmlrpc:"name,omitempty"`
	Narration                      *String     `xmlrpc:"narration,omitempty"`
	PartnerCurrencyId              *Many2One   `xmlrpc:"partner_currency_id,omitempty"`
	PartnerId                      *Many2One   `xmlrpc:"partner_id,omitempty"`
	PartnerPayableAccountId        *Many2One   `xmlrpc:"partner_payable_account_id,omitempty"`
	PartnerPayableAmount           *Float      `xmlrpc:"partner_payable_amount,omitempty"`
	PartnerReceivableAccountId     *Many2One   `xmlrpc:"partner_receivable_account_id,omitempty"`
	PartnerReceivableAmount        *Float      `xmlrpc:"partner_receivable_amount,omitempty"`
	ReconcileModelId               *Many2One   `xmlrpc:"reconcile_model_id,omitempty"`
	Ref                            *String     `xmlrpc:"ref,omitempty"`
	SourceAmlId                    *Many2One   `xmlrpc:"source_aml_id,omitempty"`
	SourceAmlMoveId                *Many2One   `xmlrpc:"source_aml_move_id,omitempty"`
	SourceAmlMoveName              *String     `xmlrpc:"source_aml_move_name,omitempty"`
	SourceAmountCurrency           *Float      `xmlrpc:"source_amount_currency,omitempty"`
	SourceBalance                  *Float      `xmlrpc:"source_balance,omitempty"`
	SourceBatchPaymentId           *Many2One   `xmlrpc:"source_batch_payment_id,omitempty"`
	SourceBatchPaymentName         *String     `xmlrpc:"source_batch_payment_name,omitempty"`
	SourceCredit                   *Float      `xmlrpc:"source_credit,omitempty"`
	SourceDebit                    *Float      `xmlrpc:"source_debit,omitempty"`
	SourceRate                     *Float      `xmlrpc:"source_rate,omitempty"`
	SuggestionAmountCurrency       *Float      `xmlrpc:"suggestion_amount_currency,omitempty"`
	SuggestionBalance              *Float      `xmlrpc:"suggestion_balance,omitempty"`
	SuggestionHtml                 *String     `xmlrpc:"suggestion_html,omitempty"`
	TaxBaseAmountCurrency          *Float      `xmlrpc:"tax_base_amount_currency,omitempty"`
	TaxIds                         *Relation   `xmlrpc:"tax_ids,omitempty"`
	TaxRepartitionLineId           *Many2One   `xmlrpc:"tax_repartition_line_id,omitempty"`
	TaxTagIds                      *Relation   `xmlrpc:"tax_tag_ids,omitempty"`
	TransactionCurrencyId          *Many2One   `xmlrpc:"transaction_currency_id,omitempty"`
	WizardId                       *Many2One   `xmlrpc:"wizard_id,omitempty"`
}

// BankRecWidgetLines represents array of bank.rec.widget.line model.
type BankRecWidgetLines []BankRecWidgetLine

// BankRecWidgetLineModel is the odoo model name.
const BankRecWidgetLineModel = "bank.rec.widget.line"

// Many2One convert BankRecWidgetLine to *Many2One.
func (brwl *BankRecWidgetLine) Many2One() *Many2One {
	return NewMany2One(brwl.Id.Get(), "")
}

// CreateBankRecWidgetLine creates a new bank.rec.widget.line model and returns its id.
func (c *Client) CreateBankRecWidgetLine(brwl *BankRecWidgetLine) (int64, error) {
	ids, err := c.CreateBankRecWidgetLines([]*BankRecWidgetLine{brwl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBankRecWidgetLine creates a new bank.rec.widget.line model and returns its id.
func (c *Client) CreateBankRecWidgetLines(brwls []*BankRecWidgetLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range brwls {
		vv = append(vv, v)
	}
	return c.Create(BankRecWidgetLineModel, vv, nil)
}

// UpdateBankRecWidgetLine updates an existing bank.rec.widget.line record.
func (c *Client) UpdateBankRecWidgetLine(brwl *BankRecWidgetLine) error {
	return c.UpdateBankRecWidgetLines([]int64{brwl.Id.Get()}, brwl)
}

// UpdateBankRecWidgetLines updates existing bank.rec.widget.line records.
// All records (represented by ids) will be updated by brwl values.
func (c *Client) UpdateBankRecWidgetLines(ids []int64, brwl *BankRecWidgetLine) error {
	return c.Update(BankRecWidgetLineModel, ids, brwl, nil)
}

// DeleteBankRecWidgetLine deletes an existing bank.rec.widget.line record.
func (c *Client) DeleteBankRecWidgetLine(id int64) error {
	return c.DeleteBankRecWidgetLines([]int64{id})
}

// DeleteBankRecWidgetLines deletes existing bank.rec.widget.line records.
func (c *Client) DeleteBankRecWidgetLines(ids []int64) error {
	return c.Delete(BankRecWidgetLineModel, ids)
}

// GetBankRecWidgetLine gets bank.rec.widget.line existing record.
func (c *Client) GetBankRecWidgetLine(id int64) (*BankRecWidgetLine, error) {
	brwls, err := c.GetBankRecWidgetLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*brwls)[0]), nil
}

// GetBankRecWidgetLines gets bank.rec.widget.line existing records.
func (c *Client) GetBankRecWidgetLines(ids []int64) (*BankRecWidgetLines, error) {
	brwls := &BankRecWidgetLines{}
	if err := c.Read(BankRecWidgetLineModel, ids, nil, brwls); err != nil {
		return nil, err
	}
	return brwls, nil
}

// FindBankRecWidgetLine finds bank.rec.widget.line record by querying it with criteria.
func (c *Client) FindBankRecWidgetLine(criteria *Criteria) (*BankRecWidgetLine, error) {
	brwls := &BankRecWidgetLines{}
	if err := c.SearchRead(BankRecWidgetLineModel, criteria, NewOptions().Limit(1), brwls); err != nil {
		return nil, err
	}
	return &((*brwls)[0]), nil
}

// FindBankRecWidgetLines finds bank.rec.widget.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBankRecWidgetLines(criteria *Criteria, options *Options) (*BankRecWidgetLines, error) {
	brwls := &BankRecWidgetLines{}
	if err := c.SearchRead(BankRecWidgetLineModel, criteria, options, brwls); err != nil {
		return nil, err
	}
	return brwls, nil
}

// FindBankRecWidgetLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBankRecWidgetLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(BankRecWidgetLineModel, criteria, options)
}

// FindBankRecWidgetLineId finds record id by querying it with criteria.
func (c *Client) FindBankRecWidgetLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BankRecWidgetLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
