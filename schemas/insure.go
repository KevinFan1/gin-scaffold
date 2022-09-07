package schemas

import "encoding/xml"

type Head struct {
	UUID         string `xml:"UUID"`
	CRequestType string `xml:"CRequestType"`
	CBusiChnl    string `xml:"CBusiChnl"`
	TAcctTm      string `xml:"TAcctTm"`
	CheckCode    string `xml:"CheckCode"`
}

type Base struct {
	CProdPlan    string `xml:"CProdPlan"`
	TAppTm       string `xml:"TAppTm"`
	TInsrncBgnTm string `xml:"TInsrncBgnTm"`
	TInsrncEndTm string `xml:"TInsrncEndTm"`
	CTmSysCde    string `xml:"CTmSysCde"`
	CUnfixSpc    string `xml:"CUnfixSpc"`
	OrderNo      string `xml:"OrderNo"`
	NAmt         string `xml:"NAmt"`
	NPrm         string `xml:"NPrm"`
}

type Applicant struct {
	CAppNme   string `xml:"CAppNme"`
	CClntMrk  string `xml:"CClntMrk"`
	CCertfCls string `xml:"CCertfCls"`
	CCertfCde string `xml:"CCertfCde"`
	CMobile   string `xml:"CMobile"`
	CRelCode  string `xml:"CRelCode"`
	CTelPhone string `xml:"CTelPhone"`
	CEmail    string `xml:"CEmail"`
	CClntAddr string `xml:"CClntAddr"`
	CZipCde   string `xml:"CZipCde"`
}

type GrpMember struct {
	CNme       string `xml:"CNme"`
	CClntMrk   string `xml:"CClntMrk"`
	CCertTyp   string `xml:"CCertTyp"`
	CCertNo    string `xml:"CCertNo"`
	CSex       string `xml:"CSex"`
	CMobile    string `xml:"CMobile"`
	CClntAddr  string `xml:"CClntAddr"`
	CTBirthDay string `xml:"CTBirthDay"`
	CZipCde    string `xml:"CZipCde"`
}

type Tgt struct {
	CPayerName   string `xml:"CPayerName"`
	CTranNo      string `xml:"CTranNo"`
	CArrivalTime string `xml:"CArrivalTime"`
	NTranAmt     string `xml:"NTranAmt"`
}

type Body struct {
	Base      Base      `xml:"Base"`
	Applicant Applicant `xml:"Applicant"`
	GrpMember GrpMember `xml:"GrpMember"`
	Tgt       GrpMember `xml:"Tgt"`
}

type Packet struct {
	Packet xml.Name `xml:"Packet"`
	Head   Head     `xml:"Head"`
	Body   Body     `xml:"Body"`
}

type InsureResponse struct {
	No   string      `json:"no"`
	Data interface{} `json:"data"`
}
