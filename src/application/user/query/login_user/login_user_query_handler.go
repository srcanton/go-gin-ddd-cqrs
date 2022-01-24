package login_user

func (bus *Bus) handleLoginUserQuery(
	query *LoginUserQuery,
) (map[string]interface{}, error) {
	return bus.loginService.Login(
		query.Email,
		query.Password,
	)
}
