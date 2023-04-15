package serializers

type LoginSerializer struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
