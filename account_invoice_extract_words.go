package odoo

// AccountInvoiceExtractWords represents account.invoice_extract.words model.
type AccountInvoiceExtractWords struct {
	CreateDate    *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	Field         *String   `xmlrpc:"field,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	InvoiceId     *Many2One `xmlrpc:"invoice_id,omitempty"`
	OcrSelected   *Bool     `xmlrpc:"ocr_selected,omitempty"`
	UserSelected  *Bool     `xmlrpc:"user_selected,omitempty"`
	WordBoxAngle  *Float    `xmlrpc:"word_box_angle,omitempty"`
	WordBoxHeight *Float    `xmlrpc:"word_box_height,omitempty"`
	WordBoxMidX   *Float    `xmlrpc:"word_box_midX,omitempty"`
	WordBoxMidY   *Float    `xmlrpc:"word_box_midY,omitempty"`
	WordBoxWidth  *Float    `xmlrpc:"word_box_width,omitempty"`
	WordPage      *Int      `xmlrpc:"word_page,omitempty"`
	WordText      *String   `xmlrpc:"word_text,omitempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountInvoiceExtractWordss represents array of account.invoice_extract.words model.
type AccountInvoiceExtractWordss []AccountInvoiceExtractWords

// AccountInvoiceExtractWordsModel is the odoo model name.
const AccountInvoiceExtractWordsModel = "account.invoice_extract.words"

// Many2One convert AccountInvoiceExtractWords to *Many2One.
func (aiw *AccountInvoiceExtractWords) Many2One() *Many2One {
	return NewMany2One(aiw.Id.Get(), "")
}

// CreateAccountInvoiceExtractWords creates a new account.invoice_extract.words model and returns its id.
func (c *Client) CreateAccountInvoiceExtractWords(aiw *AccountInvoiceExtractWords) (int64, error) {
	ids, err := c.CreateAccountInvoiceExtractWordss([]*AccountInvoiceExtractWords{aiw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountInvoiceExtractWords creates a new account.invoice_extract.words model and returns its id.
func (c *Client) CreateAccountInvoiceExtractWordss(aiws []*AccountInvoiceExtractWords) ([]int64, error) {
	var vv []interface{}
	for _, v := range aiws {
		vv = append(vv, v)
	}
	return c.Create(AccountInvoiceExtractWordsModel, vv, nil)
}

// UpdateAccountInvoiceExtractWords updates an existing account.invoice_extract.words record.
func (c *Client) UpdateAccountInvoiceExtractWords(aiw *AccountInvoiceExtractWords) error {
	return c.UpdateAccountInvoiceExtractWordss([]int64{aiw.Id.Get()}, aiw)
}

// UpdateAccountInvoiceExtractWordss updates existing account.invoice_extract.words records.
// All records (represented by ids) will be updated by aiw values.
func (c *Client) UpdateAccountInvoiceExtractWordss(ids []int64, aiw *AccountInvoiceExtractWords) error {
	return c.Update(AccountInvoiceExtractWordsModel, ids, aiw, nil)
}

// DeleteAccountInvoiceExtractWords deletes an existing account.invoice_extract.words record.
func (c *Client) DeleteAccountInvoiceExtractWords(id int64) error {
	return c.DeleteAccountInvoiceExtractWordss([]int64{id})
}

// DeleteAccountInvoiceExtractWordss deletes existing account.invoice_extract.words records.
func (c *Client) DeleteAccountInvoiceExtractWordss(ids []int64) error {
	return c.Delete(AccountInvoiceExtractWordsModel, ids)
}

// GetAccountInvoiceExtractWords gets account.invoice_extract.words existing record.
func (c *Client) GetAccountInvoiceExtractWords(id int64) (*AccountInvoiceExtractWords, error) {
	aiws, err := c.GetAccountInvoiceExtractWordss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aiws)[0]), nil
}

// GetAccountInvoiceExtractWordss gets account.invoice_extract.words existing records.
func (c *Client) GetAccountInvoiceExtractWordss(ids []int64) (*AccountInvoiceExtractWordss, error) {
	aiws := &AccountInvoiceExtractWordss{}
	if err := c.Read(AccountInvoiceExtractWordsModel, ids, nil, aiws); err != nil {
		return nil, err
	}
	return aiws, nil
}

// FindAccountInvoiceExtractWords finds account.invoice_extract.words record by querying it with criteria.
func (c *Client) FindAccountInvoiceExtractWords(criteria *Criteria) (*AccountInvoiceExtractWords, error) {
	aiws := &AccountInvoiceExtractWordss{}
	if err := c.SearchRead(AccountInvoiceExtractWordsModel, criteria, NewOptions().Limit(1), aiws); err != nil {
		return nil, err
	}
	return &((*aiws)[0]), nil
}

// FindAccountInvoiceExtractWordss finds account.invoice_extract.words records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountInvoiceExtractWordss(criteria *Criteria, options *Options) (*AccountInvoiceExtractWordss, error) {
	aiws := &AccountInvoiceExtractWordss{}
	if err := c.SearchRead(AccountInvoiceExtractWordsModel, criteria, options, aiws); err != nil {
		return nil, err
	}
	return aiws, nil
}

// FindAccountInvoiceExtractWordsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountInvoiceExtractWordsIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountInvoiceExtractWordsModel, criteria, options)
}

// FindAccountInvoiceExtractWordsId finds record id by querying it with criteria.
func (c *Client) FindAccountInvoiceExtractWordsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountInvoiceExtractWordsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
