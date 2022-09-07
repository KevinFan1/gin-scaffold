package schemas

import "encoding/xml"

type XMLHeadData struct {
	UUID         string `xml:"UUID" json:"uuid"`
	CRequestType string `xml:"CRequestType" json:"c_request_type"`
	CBusiChnl    string `xml:"CBusiChnl" json:"c_busi_chnl"`
	TAcctTm      string `xml:"TAcctTm" json:"t_acct_tm"`
	CheckCode    string `xml:"CheckCode" json:"check_code"`
}

type XMLData struct {
	XMLName xml.Name `xml:"Packet"`
	Type    string   `xml:"type,attr"`
	Version string   `xml:"version,attr"`
	Head    XMLHeadData
}
