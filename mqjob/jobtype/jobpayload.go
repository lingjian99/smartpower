package jobtype

type WorkOrderPushPayload struct {
	Title   string      `json:"title"`
	Content string      `json:"content"`
	Image   string      `json:"image"`
	Extra   interface{} `json:"extra"`
	Users   []string    `json:"users"` // REGISTRATION_ID
}

type WorkOrder struct {
	Id           int64 `json:"id"`
	Type         int   `json:"type"`
	SiteId       int64 `json:"site_id"`
	DepartmentId int64 `json:"department_id"`
	WorkerId     int64 `json:"worker_id"`
	FactoryId    int64 `json:"factory_id"`
	Status       int   `json:"status"`
	// 工单标题
	Title string `json:"title"`
}

type NotiUser struct {
	App    string
	UserId int64
}
