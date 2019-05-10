package dataobj

type Api104Record struct {
	ADDRESS           string `json:"ADDRESS"`
	ADDR_INDZONE      string `json:"ADDR_INDZONE"`
	ADDR_NO_DESCRIPT  string `json:"ADDR_NO_DESCRIPT"`
	APPEAR_DATE       string `json:"APPEAR_DATE"`
	APPEAR_TIME       string `json:"APPEAR_TIME"`
	C                 string `json:"C"`
	CERT_ALL_DESCRIPT string `json:"CERT_ALL_DESCRIPT"`
	DESCRIPTION       string `json:"DESCRIPTION"`
	INDCAT            string `json:"INDCAT"`
	J                 string `json:"J"`
	JOB               string `json:"JOB"`
	JOBCAT_DESCRIPT   string `json:"JOBCAT_DESCRIPT"`
	LAT               string `json:"LAT"`
	LON               string `json:"LON"`
	NAME              string `json:"NAME"`
	PRODUCT           string `json:"PRODUCT"`
	PROFILE           string `json:"PROFILE"`
	SAL_MONTH_HIGH    string `json:"SAL_MONTH_HIGH"`
	SAL_MONTH_LOW     string `json:"SAL_MONTH_LOW"`
	WELFARE           string `json:"WELFARE"`
}

type Api104Response struct {
	RECORDCOUNT string         `json:"RECORDCOUNT"`
	PAGECOUNT   string         `json:"PAGECOUNT"`
	PAGE        string         `json:"PAGE"`
	TOTALPAGE   string         `json:"TOTALPAGE"`
	DATA        []Api104Record `json:"data"`
}
