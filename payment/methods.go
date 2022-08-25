package payment

type CardNetwork string

const (
	CardNetworkVisa               CardNetwork = "visa"
	CardNetworkMasterCard         CardNetwork = "mastercard"
	CardNetworkAmericanExpress    CardNetwork = "americanexpress"
	CardNetworkDiscover           CardNetwork = "discover"
	CardNetworkJCB                CardNetwork = "jcb"
	CardNetworkDiners             CardNetwork = "diners"
	CardNetworkUnionPay           CardNetwork = "unionpay"
	CardNetworkAirPlus            CardNetwork = "airplus"
	CardNetworkAurore             CardNetwork = "aurore"
	CardNetworkCarteBancaire      CardNetwork = "cartebancaire"
	CardNetworkCarteBleue         CardNetwork = "cartebleue"
	CardNetworkDankort            CardNetwork = "dankort"
	CardNetworkGECapital          CardNetwork = "gecapital"
	CardNetworkUATP               CardNetwork = "uatp"
	CardNetworkBancomat           CardNetwork = "bancomat"
	CardNetworkBCCard             CardNetwork = "bccard"
	CardNetworkBCACard            CardNetwork = "bcacard"
	CardNetworkCabcharge          CardNetwork = "cabcharge"
	CardNetworkEftpos             CardNetwork = "eftpos"
	CardNetworkEps                CardNetwork = "eps"
	CardNetworkElo                CardNetwork = "elo"
	CardNetworkForbrugsforeningen CardNetwork = "forbrugsforeningen"
	CardNetworkGirocard           CardNetwork = "girocard"
	CardNetworkInterac            CardNetwork = "interac"
	CardNetworkIsracard           CardNetwork = "isracard"
	CardNetworkMir                CardNetwork = "mir"
	CardNetworkMEPS               CardNetwork = "meps"
	CardNetworkNETS               CardNetwork = "nets"
	CardNetworkPayPak             CardNetwork = "paypak"
	CardNetworkRuPay              CardNetwork = "rupay"
	CardNetworkTroy               CardNetwork = "troy"
	CardNetworkVpay               CardNetwork = "vpay"
	CardNetworkVerve              CardNetwork = "verve"
)
