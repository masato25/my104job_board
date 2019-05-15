package dataobj

type Api104Record struct {
	ADDRESS         string `json:"ADDRESS"`
	AddrINDZONE     string `json:"ADDR_INDZONE"`
	AddrNoDESCRIPT  string `json:"ADDR_NO_DESCRIPT"`
	AppearDATE      string `json:"APPEAR_DATE"`
	AppearTIME      string `json:"APPEAR_TIME"`
	C               string `json:"C"`
	CertAllDESCRIPT string `json:"CERT_ALL_DESCRIPT"`
	DESCRIPTION     string `json:"DESCRIPTION"`
	INDCAT          string `json:"INDCAT"`
	J               string `json:"J"`
	S2              string `json:"S2"`
	S3              string `json:"S3"`
	LINK            string `json:"LINK"`
	JOB             string `json:"JOB"`
	JobcatDESCRIPT  string `json:"JOBCAT_DESCRIPT"`
	LAT             string `json:"LAT"`
	LON             string `json:"LON"`
	NAME            string `json:"NAME"`
	PRODUCT         string `json:"PRODUCT"`
	PROFILE         string `json:"PROFILE"`
	SalMonthHIGH    string `json:"SAL_MONTH_HIGH"`
	SalMonthLOW     string `json:"SAL_MONTH_LOW"`
	WELFARE         string `json:"WELFARE"`
}

type Api104Response struct {
	RECORDCOUNT string         `json:"RECORDCOUNT"`
	PAGECOUNT   string         `json:"PAGECOUNT"`
	PAGE        string         `json:"PAGE"`
	TOTALPAGE   string         `json:"TOTALPAGE"`
	DATA        []Api104Record `json:"data"`
}
