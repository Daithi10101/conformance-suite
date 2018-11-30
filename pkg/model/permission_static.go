package model

var (
	// Get /accounts example json response from ozone
	permissionStaticData = []byte(
		`[
			{
			  "permission": "ReadAccountsBasic",
			  "endpoints": [
				"/accounts",
				"/accounts/{AccountId}"
			  ],
			  "default": true
			},
			{
			  "permission": "ReadAccountsDetail",
			  "endpoints": [
				"/accounts",
				"/accounts/{AccountId}"
			  ]
			},
			{
			  "permission": "ReadBalances",
			  "endpoints": [
				"/balances",
				"/accounts/{AccountId}/balances"
			  ]
			},
			{
			  "permission": "ReadBeneficiariesBasic",
			  "endpoints": [
				"/beneficiaries",
				"/accounts/{AccountId}/beneficiaries"
			  ],
			  "default": true
			},
			{
			  "permission": "ReadBeneficiariesDetail",
			  "endpoints": [
				"/beneficiaries",
				"/accounts/{AccountId}/beneficiaries"
			  ]
			},
			{
			  "permission": "ReadDirectDebits",
			  "endpoints": [
				"/direct-debits",
				"/accounts/{AccountId}/direct-debits"
			  ]
			},
			{
			  "permission": "ReadStandingOrdersBasic",
			  "endpoints": [
				"/standing-orders",
				"/accounts/{AccountId}/standing-orders"
			  ],
			  "default": true
			},
			{
			  "permission": "ReadStandingOrdersDetail",
			  "endpoints": [
				"/standing-orders",
				"/accounts/{AccountId}/standing-orders"
			  ]
			},
			{
			  "permission": "ReadTransactionsBasic",
			  "endpoints": [
				"/transactions",
				"/accounts/{AccountId}/transactions",
				"/accounts/{AccountId}/statements/{StatementId}/transactions"
			  ],
			  "required-one-or-more": [
				"ReadTransactionsCredits",
				"ReadTransactionsDebits"
			  ],
			  "optional": [
				"ReadPAN"
			  ],
			  "default": true
			},
			{
			  "permission": "ReadTransactionsDetail",
			  "endpoints": [
				"/transactions",
				"/accounts/{AccountId}/transactions",
				"/accounts/{AccountId}/statements/{StatementId}/transactions"
			  ],
			  "required-one-or-more": [
				"ReadTransactionsCredits",
				"ReadTransactionsDebits"
			  ],
			  "optional": [
				"ReadPAN"
			  ]
			},
			{
			  "permission": "ReadTransactionsCredits",
			  "endpoints": [
				"/transactions",
				"/accounts/{AccountId}/transactions",
				"/accounts/{AccountId}/statements/{StatementId}/transactions"
			  ],
			  "required-one-or-more": [
				"ReadTransactionsBasic",
				"ReadTransactionsDetail"
			  ],
			  "optional": [
				"ReadPAN"
			  ]
			},
			{
			  "permission": "ReadTransactionsDebits",
			  "endpoints": [
				"/transactions",
				"/accounts/{AccountId}/transactions",
				"/accounts/{AccountId}/statements/{StatementId}/transactions"
			  ],
			  "required-one-or-more": [
				"ReadTransactionsBasic",
				"ReadTransactionsDetail"
			  ],
			  "optional": [
				"ReadPAN"
			  ]
			},
			{
			  "permission": "ReadStatementsBasic",
			  "endpoints": [
				"/statements",
				" /accounts/{AccountId}/statements"
			  ]
			},
			{
			  "permission": "ReadStatementsDetail",
			  "endpoints": [
				"/statements",
				" /accounts/{AccountId}/statements",
				"/accounts/{AccountId}/statements/{StatementId}/file"
			  ]
			},
			{
			  "permission": "ReadProducts",
			  "endpoints": [
				"/products",
				"/accounts/{AccountId}/product"
			  ]
			},
			{
			  "permission": "ReadOffers",
			  "endpoints": [
				"/offers",
				"/accounts/{AccountId}/offers"
			  ]
			},
			{
			  "permission": "ReadParty",
			  "endpoints": [
				"/accounts/{AccountId}/party"
			  ]
			},
			{
			  "permission": "ReadPartyPSU",
			  "endpoints": [
				"/party"
			  ]
			},
			{
			  "permission": "ReadScheduledPaymentsBasic",
			  "endpoints": [
				"/scheduled-payments",
				"/accounts/{AccountId}/scheduled-payments"
			  ],
			  "default": true
			},
			{
			  "permission": "ReadScheduledPaymentsDetail",
			  "endpoints": [
				"/scheduled-payments",
				"/accounts/{AccountId}/scheduled-payments"
			  ]
			},
			{
			  "permission": "ReadPAN",
			  "endpoints": [
				"",
				""
			  ]
			}
		  ]
		  `)
)
