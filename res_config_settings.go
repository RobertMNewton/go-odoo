package odoo

// ResConfigSettings represents res.config.settings model.
type ResConfigSettings struct {
	AccountCashBasisBaseAccountId               *Many2One  `xmlrpc:"account_cash_basis_base_account_id,omitempty"`
	AccountDefaultCreditLimit                   *Float     `xmlrpc:"account_default_credit_limit,omitempty"`
	AccountDiscountExpenseAllocationId          *Many2One  `xmlrpc:"account_discount_expense_allocation_id,omitempty"`
	AccountDiscountIncomeAllocationId           *Many2One  `xmlrpc:"account_discount_income_allocation_id,omitempty"`
	AccountFiscalCountryId                      *Many2One  `xmlrpc:"account_fiscal_country_id,omitempty"`
	AccountFolderId                             *Many2One  `xmlrpc:"account_folder_id,omitempty"`
	AccountJournalEarlyPayDiscountGainAccountId *Many2One  `xmlrpc:"account_journal_early_pay_discount_gain_account_id,omitempty"`
	AccountJournalEarlyPayDiscountLossAccountId *Many2One  `xmlrpc:"account_journal_early_pay_discount_loss_account_id,omitempty"`
	AccountJournalSuspenseAccountId             *Many2One  `xmlrpc:"account_journal_suspense_account_id,omitempty"`
	AccountPriceInclude                         *Selection `xmlrpc:"account_price_include,omitempty"`
	AccountProductionWipAccountId               *Many2One  `xmlrpc:"account_production_wip_account_id,omitempty"`
	AccountProductionWipOverheadAccountId       *Many2One  `xmlrpc:"account_production_wip_overhead_account_id,omitempty"`
	AccountReportsShowPerCompanySetting         *Bool      `xmlrpc:"account_reports_show_per_company_setting,omitempty"`
	AccountStorno                               *Bool      `xmlrpc:"account_storno,omitempty"`
	AccountTaxPeriodicity                       *Selection `xmlrpc:"account_tax_periodicity,omitempty"`
	AccountTaxPeriodicityJournalId              *Many2One  `xmlrpc:"account_tax_periodicity_journal_id,omitempty"`
	AccountTaxPeriodicityReminderDay            *Int       `xmlrpc:"account_tax_periodicity_reminder_day,omitempty"`
	AccountUseCreditLimit                       *Bool      `xmlrpc:"account_use_credit_limit,omitempty"`
	ActiveUserCount                             *Int       `xmlrpc:"active_user_count,omitempty"`
	AliasDomainId                               *Many2One  `xmlrpc:"alias_domain_id,omitempty"`
	AnnualInventoryDay                          *Int       `xmlrpc:"annual_inventory_day,omitempty"`
	AnnualInventoryMonth                        *Selection `xmlrpc:"annual_inventory_month,omitempty"`
	AuthSignupResetPassword                     *Bool      `xmlrpc:"auth_signup_reset_password,omitempty"`
	AuthSignupTemplateUserId                    *Many2One  `xmlrpc:"auth_signup_template_user_id,omitempty"`
	AuthSignupUninvited                         *Selection `xmlrpc:"auth_signup_uninvited,omitempty"`
	AuthTotpEnforce                             *Bool      `xmlrpc:"auth_totp_enforce,omitempty"`
	AuthTotpPolicy                              *Selection `xmlrpc:"auth_totp_policy,omitempty"`
	AutomaticInvoice                            *Bool      `xmlrpc:"automatic_invoice,omitempty"`
	AutopostBills                               *Bool      `xmlrpc:"autopost_bills,omitempty"`
	BarcodeMaxTimeBetweenKeysInMs               *Int       `xmlrpc:"barcode_max_time_between_keys_in_ms,omitempty"`
	BarcodeNomenclatureId                       *Many2One  `xmlrpc:"barcode_nomenclature_id,omitempty"`
	BarcodeRfidBatchTime                        *Int       `xmlrpc:"barcode_rfid_batch_time,omitempty"`
	BarcodeSeparator                            *String    `xmlrpc:"barcode_separator,omitempty"`
	BarcodeSeparatorRegex                       *String    `xmlrpc:"barcode_separator_regex,omitempty"`
	BarcodelookupApiKey                         *String    `xmlrpc:"barcodelookup_api_key,omitempty"`
	CdnActivated                                *Bool      `xmlrpc:"cdn_activated,omitempty"`
	CdnFilters                                  *String    `xmlrpc:"cdn_filters,omitempty"`
	CdnUrl                                      *String    `xmlrpc:"cdn_url,omitempty"`
	ChartTemplate                               *Selection `xmlrpc:"chart_template,omitempty"`
	CheckAccountAuditTrail                      *Bool      `xmlrpc:"check_account_audit_trail,omitempty"`
	CompanyCount                                *Int       `xmlrpc:"company_count,omitempty"`
	CompanyCountryCode                          *String    `xmlrpc:"company_country_code,omitempty"`
	CompanyCurrencyId                           *Many2One  `xmlrpc:"company_currency_id,omitempty"`
	CompanyId                                   *Many2One  `xmlrpc:"company_id,omitempty"`
	CompanyInformations                         *String    `xmlrpc:"company_informations,omitempty"`
	CompanyName                                 *String    `xmlrpc:"company_name,omitempty"`
	CompanySoTemplateId                         *Many2One  `xmlrpc:"company_so_template_id,omitempty"`
	CountryCode                                 *String    `xmlrpc:"country_code,omitempty"`
	CreateDate                                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	CrmAutoAssignmentAction                     *Selection `xmlrpc:"crm_auto_assignment_action,omitempty"`
	CrmAutoAssignmentIntervalNumber             *Int       `xmlrpc:"crm_auto_assignment_interval_number,omitempty"`
	CrmAutoAssignmentIntervalType               *Selection `xmlrpc:"crm_auto_assignment_interval_type,omitempty"`
	CrmAutoAssignmentRunDatetime                *Time      `xmlrpc:"crm_auto_assignment_run_datetime,omitempty"`
	CrmUseAutoAssignment                        *Bool      `xmlrpc:"crm_use_auto_assignment,omitempty"`
	CurrencyExchangeJournalId                   *Many2One  `xmlrpc:"currency_exchange_journal_id,omitempty"`
	CurrencyId                                  *Many2One  `xmlrpc:"currency_id,omitempty"`
	CurrencyIntervalUnit                        *Selection `xmlrpc:"currency_interval_unit,omitempty"`
	CurrencyNextExecutionDate                   *Time      `xmlrpc:"currency_next_execution_date,omitempty"`
	CurrencyProvider                            *Selection `xmlrpc:"currency_provider,omitempty"`
	DaysToPurchase                              *Float     `xmlrpc:"days_to_purchase,omitempty"`
	DefaultInvoicePolicy                        *Selection `xmlrpc:"default_invoice_policy,omitempty"`
	DefaultPickingPolicy                        *Selection `xmlrpc:"default_picking_policy,omitempty"`
	DefaultPurchaseMethod                       *Selection `xmlrpc:"default_purchase_method,omitempty"`
	DeferredExpenseAccountId                    *Many2One  `xmlrpc:"deferred_expense_account_id,omitempty"`
	DeferredExpenseAmountComputationMethod      *Selection `xmlrpc:"deferred_expense_amount_computation_method,omitempty"`
	DeferredExpenseJournalId                    *Many2One  `xmlrpc:"deferred_expense_journal_id,omitempty"`
	DeferredRevenueAccountId                    *Many2One  `xmlrpc:"deferred_revenue_account_id,omitempty"`
	DeferredRevenueAmountComputationMethod      *Selection `xmlrpc:"deferred_revenue_amount_computation_method,omitempty"`
	DeferredRevenueJournalId                    *Many2One  `xmlrpc:"deferred_revenue_journal_id,omitempty"`
	DeletionDelay                               *Int       `xmlrpc:"deletion_delay,omitempty"`
	DigestEmails                                *Bool      `xmlrpc:"digest_emails,omitempty"`
	DigestId                                    *Many2One  `xmlrpc:"digest_id,omitempty"`
	DisableRedirectFirebaseDynamicLink          *Bool      `xmlrpc:"disable_redirect_firebase_dynamic_link,omitempty"`
	DisplayInvoiceAmountTotalWords              *Bool      `xmlrpc:"display_invoice_amount_total_words,omitempty"`
	DisplayInvoiceTaxCompanyCurrency            *Bool      `xmlrpc:"display_invoice_tax_company_currency,omitempty"`
	DisplayName                                 *String    `xmlrpc:"display_name,omitempty"`
	DocumentSpreadsheetFolderId                 *Many2One  `xmlrpc:"document_spreadsheet_folder_id,omitempty"`
	DocumentsAccountSettings                    *Bool      `xmlrpc:"documents_account_settings,omitempty"`
	DocumentsHrFolder                           *Many2One  `xmlrpc:"documents_hr_folder,omitempty"`
	DocumentsHrSettings                         *Bool      `xmlrpc:"documents_hr_settings,omitempty"`
	DocumentsProductSettings                    *Bool      `xmlrpc:"documents_product_settings,omitempty"`
	DocumentsSignFolderId                       *Many2One  `xmlrpc:"documents_sign_folder_id,omitempty"`
	EmailPrimaryColor                           *String    `xmlrpc:"email_primary_color,omitempty"`
	EmailSecondaryColor                         *String    `xmlrpc:"email_secondary_color,omitempty"`
	EnableOcn                                   *Bool      `xmlrpc:"enable_ocn,omitempty"`
	ExpenseCurrencyExchangeAccountId            *Many2One  `xmlrpc:"expense_currency_exchange_account_id,omitempty"`
	ExternalEmailServerDefault                  *Bool      `xmlrpc:"external_email_server_default,omitempty"`
	ExternalReportLayoutId                      *Many2One  `xmlrpc:"external_report_layout_id,omitempty"`
	ExtractBankStatementDigitalizationMode      *Selection `xmlrpc:"extract_bank_statement_digitalization_mode,omitempty"`
	ExtractInInvoiceDigitalizationMode          *Selection `xmlrpc:"extract_in_invoice_digitalization_mode,omitempty"`
	ExtractOutInvoiceDigitalizationMode         *Selection `xmlrpc:"extract_out_invoice_digitalization_mode,omitempty"`
	ExtractSingleLinePerTax                     *Bool      `xmlrpc:"extract_single_line_per_tax,omitempty"`
	FailCounter                                 *Int       `xmlrpc:"fail_counter,omitempty"`
	Favicon                                     *String    `xmlrpc:"favicon,omitempty"`
	FirstProviderLabel                          *String    `xmlrpc:"first_provider_label,omitempty"`
	FiscalyearLastDay                           *Int       `xmlrpc:"fiscalyear_last_day,omitempty"`
	FiscalyearLastMonth                         *Selection `xmlrpc:"fiscalyear_last_month,omitempty"`
	GenerateDeferredExpenseEntriesMethod        *Selection `xmlrpc:"generate_deferred_expense_entries_method,omitempty"`
	GenerateDeferredRevenueEntriesMethod        *Selection `xmlrpc:"generate_deferred_revenue_entries_method,omitempty"`
	GoogleAnalyticsKey                          *String    `xmlrpc:"google_analytics_key,omitempty"`
	GoogleSearchConsole                         *String    `xmlrpc:"google_search_console,omitempty"`
	GoogleTranslateApiKey                       *String    `xmlrpc:"google_translate_api_key,omitempty"`
	GroupAnalyticAccounting                     *Bool      `xmlrpc:"group_analytic_accounting,omitempty"`
	GroupAutoDoneSetting                        *Bool      `xmlrpc:"group_auto_done_setting,omitempty"`
	GroupBarcodeCountEntireLocation             *Bool      `xmlrpc:"group_barcode_count_entire_location,omitempty"`
	GroupBarcodeShowQuantityCount               *Bool      `xmlrpc:"group_barcode_show_quantity_count,omitempty"`
	GroupCashRounding                           *Bool      `xmlrpc:"group_cash_rounding,omitempty"`
	GroupDiscountPerSoLine                      *Bool      `xmlrpc:"group_discount_per_so_line,omitempty"`
	GroupFiscalYear                             *Bool      `xmlrpc:"group_fiscal_year,omitempty"`
	GroupLotOnDeliverySlip                      *Bool      `xmlrpc:"group_lot_on_delivery_slip,omitempty"`
	GroupLotOnInvoice                           *Bool      `xmlrpc:"group_lot_on_invoice,omitempty"`
	GroupManageTemplateAccess                   *Bool      `xmlrpc:"group_manage_template_access,omitempty"`
	GroupMrpByproducts                          *Bool      `xmlrpc:"group_mrp_byproducts,omitempty"`
	GroupMrpReceptionReport                     *Bool      `xmlrpc:"group_mrp_reception_report,omitempty"`
	GroupMrpRoutings                            *Bool      `xmlrpc:"group_mrp_routings,omitempty"`
	GroupMrpWoShopFloor                         *Bool      `xmlrpc:"group_mrp_wo_shop_floor,omitempty"`
	GroupMrpWoTabletTimer                       *Bool      `xmlrpc:"group_mrp_wo_tablet_timer,omitempty"`
	GroupMrpWorkorderDependencies               *Bool      `xmlrpc:"group_mrp_workorder_dependencies,omitempty"`
	GroupMultiCurrency                          *Bool      `xmlrpc:"group_multi_currency,omitempty"`
	GroupMultiWebsite                           *Bool      `xmlrpc:"group_multi_website,omitempty"`
	GroupProductPricelist                       *Bool      `xmlrpc:"group_product_pricelist,omitempty"`
	GroupProductVariant                         *Bool      `xmlrpc:"group_product_variant,omitempty"`
	GroupProformaSales                          *Bool      `xmlrpc:"group_proforma_sales,omitempty"`
	GroupProjectMilestone                       *Bool      `xmlrpc:"group_project_milestone,omitempty"`
	GroupProjectRating                          *Bool      `xmlrpc:"group_project_rating,omitempty"`
	GroupProjectRecurringTasks                  *Bool      `xmlrpc:"group_project_recurring_tasks,omitempty"`
	GroupProjectStages                          *Bool      `xmlrpc:"group_project_stages,omitempty"`
	GroupProjectTaskDependencies                *Bool      `xmlrpc:"group_project_task_dependencies,omitempty"`
	GroupSaleDeliveryAddress                    *Bool      `xmlrpc:"group_sale_delivery_address,omitempty"`
	GroupSaleOrderTemplate                      *Bool      `xmlrpc:"group_sale_order_template,omitempty"`
	GroupSendReminder                           *Bool      `xmlrpc:"group_send_reminder,omitempty"`
	GroupShowPurchaseReceipts                   *Bool      `xmlrpc:"group_show_purchase_receipts,omitempty"`
	GroupShowSaleReceipts                       *Bool      `xmlrpc:"group_show_sale_receipts,omitempty"`
	GroupStockAccountingAutomatic               *Bool      `xmlrpc:"group_stock_accounting_automatic,omitempty"`
	GroupStockAdvLocation                       *Bool      `xmlrpc:"group_stock_adv_location,omitempty"`
	GroupStockLotPrintGs1                       *Bool      `xmlrpc:"group_stock_lot_print_gs1,omitempty"`
	GroupStockMultiLocations                    *Bool      `xmlrpc:"group_stock_multi_locations,omitempty"`
	GroupStockProductionLot                     *Bool      `xmlrpc:"group_stock_production_lot,omitempty"`
	GroupStockReceptionReport                   *Bool      `xmlrpc:"group_stock_reception_report,omitempty"`
	GroupStockSignDelivery                      *Bool      `xmlrpc:"group_stock_sign_delivery,omitempty"`
	GroupStockTrackingLot                       *Bool      `xmlrpc:"group_stock_tracking_lot,omitempty"`
	GroupStockTrackingOwner                     *Bool      `xmlrpc:"group_stock_tracking_owner,omitempty"`
	GroupUnlockedByDefault                      *Bool      `xmlrpc:"group_unlocked_by_default,omitempty"`
	GroupUom                                    *Bool      `xmlrpc:"group_uom,omitempty"`
	GroupUseLead                                *Bool      `xmlrpc:"group_use_lead,omitempty"`
	GroupUseRecurringRevenues                   *Bool      `xmlrpc:"group_use_recurring_revenues,omitempty"`
	GroupWarningAccount                         *Bool      `xmlrpc:"group_warning_account,omitempty"`
	GroupWarningPurchase                        *Bool      `xmlrpc:"group_warning_purchase,omitempty"`
	GroupWarningSale                            *Bool      `xmlrpc:"group_warning_sale,omitempty"`
	GroupWarningStock                           *Bool      `xmlrpc:"group_warning_stock,omitempty"`
	HasAccountingEntries                        *Bool      `xmlrpc:"has_accounting_entries,omitempty"`
	HasChartOfAccounts                          *Bool      `xmlrpc:"has_chart_of_accounts,omitempty"`
	HasDefaultShareImage                        *Bool      `xmlrpc:"has_default_share_image,omitempty"`
	HasGoogleAnalytics                          *Bool      `xmlrpc:"has_google_analytics,omitempty"`
	HasGoogleSearchConsole                      *Bool      `xmlrpc:"has_google_search_console,omitempty"`
	HasPlausibleSharedKey                       *Bool      `xmlrpc:"has_plausible_shared_key,omitempty"`
	HrEmployeeSelfEdit                          *Bool      `xmlrpc:"hr_employee_self_edit,omitempty"`
	HrPresenceControlEmail                      *Bool      `xmlrpc:"hr_presence_control_email,omitempty"`
	HrPresenceControlEmailAmount                *Int       `xmlrpc:"hr_presence_control_email_amount,omitempty"`
	HrPresenceControlIp                         *Bool      `xmlrpc:"hr_presence_control_ip,omitempty"`
	HrPresenceControlIpList                     *String    `xmlrpc:"hr_presence_control_ip_list,omitempty"`
	HrPresenceControlLogin                      *Bool      `xmlrpc:"hr_presence_control_login,omitempty"`
	Id                                          *Int       `xmlrpc:"id,omitempty"`
	IncomeCurrencyExchangeAccountId             *Many2One  `xmlrpc:"income_currency_exchange_account_id,omitempty"`
	IncotermId                                  *Many2One  `xmlrpc:"incoterm_id,omitempty"`
	InvoiceMailTemplateId                       *Many2One  `xmlrpc:"invoice_mail_template_id,omitempty"`
	InvoiceTerms                                *String    `xmlrpc:"invoice_terms,omitempty"`
	InvoiceTermsHtml                            *String    `xmlrpc:"invoice_terms_html,omitempty"`
	InvoicingSwitchThreshold                    *Time      `xmlrpc:"invoicing_switch_threshold,omitempty"`
	IsAccountPeppolEligible                     *Bool      `xmlrpc:"is_account_peppol_eligible,omitempty"`
	IsInstalledSale                             *Bool      `xmlrpc:"is_installed_sale,omitempty"`
	IsMembershipMulti                           *Bool      `xmlrpc:"is_membership_multi,omitempty"`
	IsRootCompany                               *Bool      `xmlrpc:"is_root_company,omitempty"`
	IsStripeSupportedCountry                    *Bool      `xmlrpc:"is_stripe_supported_country,omitempty"`
	LanguageCount                               *Int       `xmlrpc:"language_count,omitempty"`
	LanguageIds                                 *Relation  `xmlrpc:"language_ids,omitempty"`
	LeadEnrichAuto                              *Selection `xmlrpc:"lead_enrich_auto,omitempty"`
	LeadMiningInPipeline                        *Bool      `xmlrpc:"lead_mining_in_pipeline,omitempty"`
	LockConfirmedPo                             *Bool      `xmlrpc:"lock_confirmed_po,omitempty"`
	ManufacturingLead                           *Float     `xmlrpc:"manufacturing_lead,omitempty"`
	MapBoxToken                                 *String    `xmlrpc:"map_box_token,omitempty"`
	MicrosoftOutlookClientIdentifier            *String    `xmlrpc:"microsoft_outlook_client_identifier,omitempty"`
	MicrosoftOutlookClientSecret                *String    `xmlrpc:"microsoft_outlook_client_secret,omitempty"`
	ModuleAccount3WayMatch                      *Bool      `xmlrpc:"module_account_3way_match,omitempty"`
	ModuleAccountAccountant                     *Bool      `xmlrpc:"module_account_accountant,omitempty"`
	ModuleAccountBankStatementExtract           *Bool      `xmlrpc:"module_account_bank_statement_extract,omitempty"`
	ModuleAccountBankStatementImportCamt        *Bool      `xmlrpc:"module_account_bank_statement_import_camt,omitempty"`
	ModuleAccountBankStatementImportCsv         *Bool      `xmlrpc:"module_account_bank_statement_import_csv,omitempty"`
	ModuleAccountBankStatementImportOfx         *Bool      `xmlrpc:"module_account_bank_statement_import_ofx,omitempty"`
	ModuleAccountBankStatementImportQif         *Bool      `xmlrpc:"module_account_bank_statement_import_qif,omitempty"`
	ModuleAccountBatchPayment                   *Bool      `xmlrpc:"module_account_batch_payment,omitempty"`
	ModuleAccountBudget                         *Bool      `xmlrpc:"module_account_budget,omitempty"`
	ModuleAccountCheckPrinting                  *Bool      `xmlrpc:"module_account_check_printing,omitempty"`
	ModuleAccountExtract                        *Bool      `xmlrpc:"module_account_extract,omitempty"`
	ModuleAccountInterCompanyRules              *Bool      `xmlrpc:"module_account_inter_company_rules,omitempty"`
	ModuleAccountIntrastat                      *Bool      `xmlrpc:"module_account_intrastat,omitempty"`
	ModuleAccountInvoiceExtract                 *Bool      `xmlrpc:"module_account_invoice_extract,omitempty"`
	ModuleAccountIso20022                       *Bool      `xmlrpc:"module_account_iso20022,omitempty"`
	ModuleAccountPayment                        *Bool      `xmlrpc:"module_account_payment,omitempty"`
	ModuleAccountPeppol                         *Bool      `xmlrpc:"module_account_peppol,omitempty"`
	ModuleAccountReports                        *Bool      `xmlrpc:"module_account_reports,omitempty"`
	ModuleAccountSepaDirectDebit                *Bool      `xmlrpc:"module_account_sepa_direct_debit,omitempty"`
	ModuleAuthLdap                              *Bool      `xmlrpc:"module_auth_ldap,omitempty"`
	ModuleAuthOauth                             *Bool      `xmlrpc:"module_auth_oauth,omitempty"`
	ModuleBaseGeolocalize                       *Bool      `xmlrpc:"module_base_geolocalize,omitempty"`
	ModuleBaseImport                            *Bool      `xmlrpc:"module_base_import,omitempty"`
	ModuleCrmIapEnrich                          *Bool      `xmlrpc:"module_crm_iap_enrich,omitempty"`
	ModuleCrmIapMine                            *Bool      `xmlrpc:"module_crm_iap_mine,omitempty"`
	ModuleCurrencyRateLive                      *Bool      `xmlrpc:"module_currency_rate_live,omitempty"`
	ModuleDelivery                              *Bool      `xmlrpc:"module_delivery,omitempty"`
	ModuleDeliveryBpost                         *Bool      `xmlrpc:"module_delivery_bpost,omitempty"`
	ModuleDeliveryDhl                           *Bool      `xmlrpc:"module_delivery_dhl,omitempty"`
	ModuleDeliveryEasypost                      *Bool      `xmlrpc:"module_delivery_easypost,omitempty"`
	ModuleDeliveryFedex                         *Bool      `xmlrpc:"module_delivery_fedex,omitempty"`
	ModuleDeliverySendcloud                     *Bool      `xmlrpc:"module_delivery_sendcloud,omitempty"`
	ModuleDeliveryShiprocket                    *Bool      `xmlrpc:"module_delivery_shiprocket,omitempty"`
	ModuleDeliveryStarshipit                    *Bool      `xmlrpc:"module_delivery_starshipit,omitempty"`
	ModuleDeliveryUps                           *Bool      `xmlrpc:"module_delivery_ups,omitempty"`
	ModuleDeliveryUsps                          *Bool      `xmlrpc:"module_delivery_usps,omitempty"`
	ModuleGoogleAddressAutocomplete             *Bool      `xmlrpc:"module_google_address_autocomplete,omitempty"`
	ModuleGoogleCalendar                        *Bool      `xmlrpc:"module_google_calendar,omitempty"`
	ModuleGoogleGmail                           *Bool      `xmlrpc:"module_google_gmail,omitempty"`
	ModuleGoogleRecaptcha                       *Bool      `xmlrpc:"module_google_recaptcha,omitempty"`
	ModuleHrAttendance                          *Bool      `xmlrpc:"module_hr_attendance,omitempty"`
	ModuleHrHomeworking                         *Bool      `xmlrpc:"module_hr_homeworking,omitempty"`
	ModuleHrPresence                            *Bool      `xmlrpc:"module_hr_presence,omitempty"`
	ModuleHrSkills                              *Bool      `xmlrpc:"module_hr_skills,omitempty"`
	ModuleHrTimesheet                           *Bool      `xmlrpc:"module_hr_timesheet,omitempty"`
	ModuleL10NEuOss                             *Bool      `xmlrpc:"module_l10n_eu_oss,omitempty"`
	ModuleLoyalty                               *Bool      `xmlrpc:"module_loyalty,omitempty"`
	ModuleMailPlugin                            *Bool      `xmlrpc:"module_mail_plugin,omitempty"`
	ModuleMarketingAutomation                   *Bool      `xmlrpc:"module_marketing_automation,omitempty"`
	ModuleMicrosoftCalendar                     *Bool      `xmlrpc:"module_microsoft_calendar,omitempty"`
	ModuleMicrosoftOutlook                      *Bool      `xmlrpc:"module_microsoft_outlook,omitempty"`
	ModuleMrpMps                                *Bool      `xmlrpc:"module_mrp_mps,omitempty"`
	ModuleMrpPlm                                *Bool      `xmlrpc:"module_mrp_plm,omitempty"`
	ModuleMrpSubcontracting                     *Bool      `xmlrpc:"module_mrp_subcontracting,omitempty"`
	ModulePartnerAutocomplete                   *Bool      `xmlrpc:"module_partner_autocomplete,omitempty"`
	ModuleProductEmailTemplate                  *Bool      `xmlrpc:"module_product_email_template,omitempty"`
	ModuleProductExpiry                         *Bool      `xmlrpc:"module_product_expiry,omitempty"`
	ModuleProductImages                         *Bool      `xmlrpc:"module_product_images,omitempty"`
	ModuleProductMargin                         *Bool      `xmlrpc:"module_product_margin,omitempty"`
	ModuleProjectForecast                       *Bool      `xmlrpc:"module_project_forecast,omitempty"`
	ModulePurchaseProductMatrix                 *Bool      `xmlrpc:"module_purchase_product_matrix,omitempty"`
	ModulePurchaseRequisition                   *Bool      `xmlrpc:"module_purchase_requisition,omitempty"`
	ModuleQualityControl                        *Bool      `xmlrpc:"module_quality_control,omitempty"`
	ModuleQualityControlWorksheet               *Bool      `xmlrpc:"module_quality_control_worksheet,omitempty"`
	ModuleSaleAmazon                            *Bool      `xmlrpc:"module_sale_amazon,omitempty"`
	ModuleSaleCommission                        *Bool      `xmlrpc:"module_sale_commission,omitempty"`
	ModuleSaleLoyalty                           *Bool      `xmlrpc:"module_sale_loyalty,omitempty"`
	ModuleSaleMargin                            *Bool      `xmlrpc:"module_sale_margin,omitempty"`
	ModuleSalePdfQuoteBuilder                   *Bool      `xmlrpc:"module_sale_pdf_quote_builder,omitempty"`
	ModuleSaleProductMatrix                     *Bool      `xmlrpc:"module_sale_product_matrix,omitempty"`
	ModuleSign                                  *Bool      `xmlrpc:"module_sign,omitempty"`
	ModuleSignItsme                             *Bool      `xmlrpc:"module_sign_itsme,omitempty"`
	ModuleSms                                   *Bool      `xmlrpc:"module_sms,omitempty"`
	ModuleSnailmailAccount                      *Bool      `xmlrpc:"module_snailmail_account,omitempty"`
	ModuleStockBarcode                          *Bool      `xmlrpc:"module_stock_barcode,omitempty"`
	ModuleStockBarcodeBarcodelookup             *Bool      `xmlrpc:"module_stock_barcode_barcodelookup,omitempty"`
	ModuleStockDropshipping                     *Bool      `xmlrpc:"module_stock_dropshipping,omitempty"`
	ModuleStockFleet                            *Bool      `xmlrpc:"module_stock_fleet,omitempty"`
	ModuleStockLandedCosts                      *Bool      `xmlrpc:"module_stock_landed_costs,omitempty"`
	ModuleStockPickingBatch                     *Bool      `xmlrpc:"module_stock_picking_batch,omitempty"`
	ModuleStockSms                              *Bool      `xmlrpc:"module_stock_sms,omitempty"`
	ModuleVoip                                  *Bool      `xmlrpc:"module_voip,omitempty"`
	ModuleWebUnsplash                           *Bool      `xmlrpc:"module_web_unsplash,omitempty"`
	ModuleWebsiteCfTurnstile                    *Bool      `xmlrpc:"module_website_cf_turnstile,omitempty"`
	ModuleWebsiteCrmIapReveal                   *Bool      `xmlrpc:"module_website_crm_iap_reveal,omitempty"`
	ModuleWebsiteLivechat                       *Bool      `xmlrpc:"module_website_livechat,omitempty"`
	PartnerAutocompleteInsufficientCredit       *Bool      `xmlrpc:"partner_autocomplete_insufficient_credit,omitempty"`
	PayInvoicesOnline                           *Bool      `xmlrpc:"pay_invoices_online,omitempty"`
	PlanningEmployeeUnavailabilities            *Selection `xmlrpc:"planning_employee_unavailabilities,omitempty"`
	PlanningGenerationInterval                  *Int       `xmlrpc:"planning_generation_interval,omitempty"`
	PlanningSelfUnassignDaysBefore              *Int       `xmlrpc:"planning_self_unassign_days_before,omitempty"`
	PlausibleSharedKey                          *String    `xmlrpc:"plausible_shared_key,omitempty"`
	PlausibleSite                               *String    `xmlrpc:"plausible_site,omitempty"`
	PoDoubleValidation                          *Selection `xmlrpc:"po_double_validation,omitempty"`
	PoDoubleValidationAmount                    *Float     `xmlrpc:"po_double_validation_amount,omitempty"`
	PoLead                                      *Float     `xmlrpc:"po_lead,omitempty"`
	PoLock                                      *Selection `xmlrpc:"po_lock,omitempty"`
	PoOrderApproval                             *Bool      `xmlrpc:"po_order_approval,omitempty"`
	PortalAllowApiKeys                          *Bool      `xmlrpc:"portal_allow_api_keys,omitempty"`
	PortalConfirmationPay                       *Bool      `xmlrpc:"portal_confirmation_pay,omitempty"`
	PortalConfirmationSign                      *Bool      `xmlrpc:"portal_confirmation_sign,omitempty"`
	PredictBillProduct                          *Bool      `xmlrpc:"predict_bill_product,omitempty"`
	PredictiveLeadScoringFieldLabels            *String    `xmlrpc:"predictive_lead_scoring_field_labels,omitempty"`
	PredictiveLeadScoringFields                 *Relation  `xmlrpc:"predictive_lead_scoring_fields,omitempty"`
	PredictiveLeadScoringFieldsStr              *String    `xmlrpc:"predictive_lead_scoring_fields_str,omitempty"`
	PredictiveLeadScoringStartDate              *Time      `xmlrpc:"predictive_lead_scoring_start_date,omitempty"`
	PredictiveLeadScoringStartDateStr           *String    `xmlrpc:"predictive_lead_scoring_start_date_str,omitempty"`
	PrepaymentPercent                           *Float     `xmlrpc:"prepayment_percent,omitempty"`
	PreviewReady                                *Bool      `xmlrpc:"preview_ready,omitempty"`
	ProductFolderId                             *Many2One  `xmlrpc:"product_folder_id,omitempty"`
	ProductTagIds                               *Relation  `xmlrpc:"product_tag_ids,omitempty"`
	ProductVolumeVolumeInCubicFeet              *Selection `xmlrpc:"product_volume_volume_in_cubic_feet,omitempty"`
	ProductWeightInLbs                          *Selection `xmlrpc:"product_weight_in_lbs,omitempty"`
	ProfilingEnabledUntil                       *Time      `xmlrpc:"profiling_enabled_until,omitempty"`
	PropertyAccountExpenseCategId               *Many2One  `xmlrpc:"property_account_expense_categ_id,omitempty"`
	PropertyAccountIncomeCategId                *Many2One  `xmlrpc:"property_account_income_categ_id,omitempty"`
	PropertyStockAccountInputCategId            *Many2One  `xmlrpc:"property_stock_account_input_categ_id,omitempty"`
	PropertyStockAccountOutputCategId           *Many2One  `xmlrpc:"property_stock_account_output_categ_id,omitempty"`
	PropertyStockAccountProductionCostId        *Many2One  `xmlrpc:"property_stock_account_production_cost_id,omitempty"`
	PropertyStockJournal                        *Many2One  `xmlrpc:"property_stock_journal,omitempty"`
	PropertyStockValuationAccountId             *Many2One  `xmlrpc:"property_stock_valuation_account_id,omitempty"`
	ProvidersState                              *Selection `xmlrpc:"providers_state,omitempty"`
	PurchaseTaxId                               *Many2One  `xmlrpc:"purchase_tax_id,omitempty"`
	QrCode                                      *Bool      `xmlrpc:"qr_code,omitempty"`
	QuickEditMode                               *Selection `xmlrpc:"quick_edit_mode,omitempty"`
	QuotationValidityDays                       *Int       `xmlrpc:"quotation_validity_days,omitempty"`
	RecaptchaMinScore                           *Float     `xmlrpc:"recaptcha_min_score,omitempty"`
	RecaptchaPrivateKey                         *String    `xmlrpc:"recaptcha_private_key,omitempty"`
	RecaptchaPublicKey                          *String    `xmlrpc:"recaptcha_public_key,omitempty"`
	ReportFooter                                *String    `xmlrpc:"report_footer,omitempty"`
	ResourceCalendarId                          *Many2One  `xmlrpc:"resource_calendar_id,omitempty"`
	RestrictTemplateRendering                   *Bool      `xmlrpc:"restrict_template_rendering,omitempty"`
	SaleTaxId                                   *Many2One  `xmlrpc:"sale_tax_id,omitempty"`
	SecurityLead                                *Float     `xmlrpc:"security_lead,omitempty"`
	SfuServerKey                                *String    `xmlrpc:"sfu_server_key,omitempty"`
	SfuServerUrl                                *String    `xmlrpc:"sfu_server_url,omitempty"`
	SharedUserAccount                           *Bool      `xmlrpc:"shared_user_account,omitempty"`
	ShowBarcodeNomenclature                     *Bool      `xmlrpc:"show_barcode_nomenclature,omitempty"`
	ShowEffect                                  *Bool      `xmlrpc:"show_effect,omitempty"`
	SignInvoice                                 *Bool      `xmlrpc:"sign_invoice,omitempty"`
	SignPreviewReady                            *Bool      `xmlrpc:"sign_preview_ready,omitempty"`
	SignTerms                                   *String    `xmlrpc:"sign_terms,omitempty"`
	SignTermsHtml                               *String    `xmlrpc:"sign_terms_html,omitempty"`
	SignTermsType                               *Selection `xmlrpc:"sign_terms_type,omitempty"`
	SigningUser                                 *Many2One  `xmlrpc:"signing_user,omitempty"`
	SnailmailColor                              *Bool      `xmlrpc:"snailmail_color,omitempty"`
	SnailmailCover                              *Bool      `xmlrpc:"snailmail_cover,omitempty"`
	SnailmailCoverReadonly                      *Bool      `xmlrpc:"snailmail_cover_readonly,omitempty"`
	SnailmailDuplex                             *Bool      `xmlrpc:"snailmail_duplex,omitempty"`
	SocialDefaultImage                          *String    `xmlrpc:"social_default_image,omitempty"`
	StockBarcodeDemoActive                      *Bool      `xmlrpc:"stock_barcode_demo_active,omitempty"`
	StockBarcodeMuteSoundNotifications          *Bool      `xmlrpc:"stock_barcode_mute_sound_notifications,omitempty"`
	StockMoveEmailValidation                    *Bool      `xmlrpc:"stock_move_email_validation,omitempty"`
	StockMoveSmsValidation                      *Bool      `xmlrpc:"stock_move_sms_validation,omitempty"`
	StockSmsConfirmationTemplateId              *Many2One  `xmlrpc:"stock_sms_confirmation_template_id,omitempty"`
	TaxCalculationRoundingMethod                *Selection `xmlrpc:"tax_calculation_rounding_method,omitempty"`
	TaxCashBasisJournalId                       *Many2One  `xmlrpc:"tax_cash_basis_journal_id,omitempty"`
	TaxExigibility                              *Bool      `xmlrpc:"tax_exigibility,omitempty"`
	TenorApiKey                                 *String    `xmlrpc:"tenor_api_key,omitempty"`
	TenorContentFilter                          *Selection `xmlrpc:"tenor_content_filter,omitempty"`
	TenorGifLimit                               *Int       `xmlrpc:"tenor_gif_limit,omitempty"`
	TermsType                                   *Selection `xmlrpc:"terms_type,omitempty"`
	TotalsBelowSections                         *Bool      `xmlrpc:"totals_below_sections,omitempty"`
	TransferAccountId                           *Many2One  `xmlrpc:"transfer_account_id,omitempty"`
	TwilioAccountSid                            *String    `xmlrpc:"twilio_account_sid,omitempty"`
	TwilioAccountToken                          *String    `xmlrpc:"twilio_account_token,omitempty"`
	UnsplashAccessKey                           *String    `xmlrpc:"unsplash_access_key,omitempty"`
	UnsplashAppId                               *String    `xmlrpc:"unsplash_app_id,omitempty"`
	UseAngloSaxon                               *Bool      `xmlrpc:"use_anglo_saxon,omitempty"`
	UseInvoiceTerms                             *Bool      `xmlrpc:"use_invoice_terms,omitempty"`
	UseManufacturingLead                        *Bool      `xmlrpc:"use_manufacturing_lead,omitempty"`
	UsePoLead                                   *Bool      `xmlrpc:"use_po_lead,omitempty"`
	UseSecurityLead                             *Bool      `xmlrpc:"use_security_lead,omitempty"`
	UseSignTerms                                *Bool      `xmlrpc:"use_sign_terms,omitempty"`
	UseTwilioRtcServers                         *Bool      `xmlrpc:"use_twilio_rtc_servers,omitempty"`
	UserDefaultRights                           *Bool      `xmlrpc:"user_default_rights,omitempty"`
	WebAppName                                  *String    `xmlrpc:"web_app_name,omitempty"`
	WebsiteBlockThirdPartyDomains               *Bool      `xmlrpc:"website_block_third_party_domains,omitempty"`
	WebsiteCompanyId                            *Many2One  `xmlrpc:"website_company_id,omitempty"`
	WebsiteCookiesBar                           *Bool      `xmlrpc:"website_cookies_bar,omitempty"`
	WebsiteDefaultLangCode                      *String    `xmlrpc:"website_default_lang_code,omitempty"`
	WebsiteDefaultLangId                        *Many2One  `xmlrpc:"website_default_lang_id,omitempty"`
	WebsiteDomain                               *String    `xmlrpc:"website_domain,omitempty"`
	WebsiteHomepageUrl                          *String    `xmlrpc:"website_homepage_url,omitempty"`
	WebsiteId                                   *Many2One  `xmlrpc:"website_id,omitempty"`
	WebsiteLanguageCount                        *Int       `xmlrpc:"website_language_count,omitempty"`
	WebsiteLogo                                 *String    `xmlrpc:"website_logo,omitempty"`
	WebsiteName                                 *String    `xmlrpc:"website_name,omitempty"`
	WriteDate                                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ResConfigSettingss represents array of res.config.settings model.
type ResConfigSettingss []ResConfigSettings

// ResConfigSettingsModel is the odoo model name.
const ResConfigSettingsModel = "res.config.settings"

// Many2One convert ResConfigSettings to *Many2One.
func (rcs *ResConfigSettings) Many2One() *Many2One {
	return NewMany2One(rcs.Id.Get(), "")
}

// CreateResConfigSettings creates a new res.config.settings model and returns its id.
func (c *Client) CreateResConfigSettings(rcs *ResConfigSettings) (int64, error) {
	ids, err := c.CreateResConfigSettingss([]*ResConfigSettings{rcs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResConfigSettings creates a new res.config.settings model and returns its id.
func (c *Client) CreateResConfigSettingss(rcss []*ResConfigSettings) ([]int64, error) {
	var vv []interface{}
	for _, v := range rcss {
		vv = append(vv, v)
	}
	return c.Create(ResConfigSettingsModel, vv, nil)
}

// UpdateResConfigSettings updates an existing res.config.settings record.
func (c *Client) UpdateResConfigSettings(rcs *ResConfigSettings) error {
	return c.UpdateResConfigSettingss([]int64{rcs.Id.Get()}, rcs)
}

// UpdateResConfigSettingss updates existing res.config.settings records.
// All records (represented by ids) will be updated by rcs values.
func (c *Client) UpdateResConfigSettingss(ids []int64, rcs *ResConfigSettings) error {
	return c.Update(ResConfigSettingsModel, ids, rcs, nil)
}

// DeleteResConfigSettings deletes an existing res.config.settings record.
func (c *Client) DeleteResConfigSettings(id int64) error {
	return c.DeleteResConfigSettingss([]int64{id})
}

// DeleteResConfigSettingss deletes existing res.config.settings records.
func (c *Client) DeleteResConfigSettingss(ids []int64) error {
	return c.Delete(ResConfigSettingsModel, ids)
}

// GetResConfigSettings gets res.config.settings existing record.
func (c *Client) GetResConfigSettings(id int64) (*ResConfigSettings, error) {
	rcss, err := c.GetResConfigSettingss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rcss)[0]), nil
}

// GetResConfigSettingss gets res.config.settings existing records.
func (c *Client) GetResConfigSettingss(ids []int64) (*ResConfigSettingss, error) {
	rcss := &ResConfigSettingss{}
	if err := c.Read(ResConfigSettingsModel, ids, nil, rcss); err != nil {
		return nil, err
	}
	return rcss, nil
}

// FindResConfigSettings finds res.config.settings record by querying it with criteria.
func (c *Client) FindResConfigSettings(criteria *Criteria) (*ResConfigSettings, error) {
	rcss := &ResConfigSettingss{}
	if err := c.SearchRead(ResConfigSettingsModel, criteria, NewOptions().Limit(1), rcss); err != nil {
		return nil, err
	}
	return &((*rcss)[0]), nil
}

// FindResConfigSettingss finds res.config.settings records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResConfigSettingss(criteria *Criteria, options *Options) (*ResConfigSettingss, error) {
	rcss := &ResConfigSettingss{}
	if err := c.SearchRead(ResConfigSettingsModel, criteria, options, rcss); err != nil {
		return nil, err
	}
	return rcss, nil
}

// FindResConfigSettingsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResConfigSettingsIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResConfigSettingsModel, criteria, options)
}

// FindResConfigSettingsId finds record id by querying it with criteria.
func (c *Client) FindResConfigSettingsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResConfigSettingsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
