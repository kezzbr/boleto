package goboleto

import "encoding/base64"

/*
Bank constants
@ID FEBRABAN bank identifier
@Aceite if the payer accept the billet
@Currency the currency identifier
@CurrencyName the currency name
 */
var configBradesco = map[string]interface{}{
	"ID": 237,
	"Aceite": "N",
	"Currency": 9,
	"CurrencyName": "R$",
}

/*
Bradesco
Source: (https://banco.bradesco/assets/pessoajuridica/pdf/4008-524-0121-08-layout-cobranca-versao-portuguesSS28785.pdf)
 */
type Bradesco struct {
	Agency 			int
	Account			int
	Convenio		int
	Contrato		int
	Carteira		int
	VariacaoCarteira	int
	FormatacaoConvenio	int
	FormatacaoNossoNumero	int
	Company			Company
}

func (b Bradesco) Barcode(d Document) string {
	return "1001.011011.1 123002 2"
}

func (b Bradesco) BarcodeImage(d Document) base64.Encoding {
	// TODO
	return base64.Encoding{}
}

func (b Bradesco) Layout(d Document) {
	// TODO
}