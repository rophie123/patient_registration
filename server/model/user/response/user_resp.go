package response

type UserInfoResponse struct {
	Roles []string `json:"roles"`
	Name  string   `json:"name"`
}
