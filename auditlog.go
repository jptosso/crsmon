package crsmon

type AuditLogParts struct {
	RequestHeaders  bool // B
	RequestBody     bool // C
	Trailer         bool // H
	RequestMetadata bool // I
	FileUploadInfo  bool // J
	MatchedRules    bool // K
}

func (alp *AuditLogParts) String() string {
	str := "A"
	if alp.RequestHeaders {
		str += "B"
	}
	if alp.RequestBody {
		str += "C"
	}
	if alp.Trailer {
		str += "H"
	}
	if alp.RequestMetadata {
		str += "I"
	}
	if alp.FileUploadInfo {
		str += "J"
	}
	if alp.MatchedRules {
		str += "K"
	}
	return str
}

func (alp *AuditLogParts) EnableAll() {
	alp.RequestHeaders = true
	alp.RequestBody = true
	alp.Trailer = true
	alp.RequestMetadata = true
	alp.FileUploadInfo = true
	alp.MatchedRules = true
}

func NewAuditLogParts() *AuditLogParts {
	return &AuditLogParts{}
}
