package apis

import (
	"github.com/Zhanat87/go/daos"
	"github.com/go-ozzo/ozzo-routing"
	"github.com/Zhanat87/go/app"
	"github.com/Zhanat87/go/util"
	"fmt"
)

func UserEmail(userDAO *daos.UserDAO) routing.Handler {
	return func(c *routing.Context) error {
		var response error

		rs := app.GetRequestScope(c)
		model, err := userDAO.FindByUsername(rs, c.Param("username"))
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				response = c.Write(fmt.Sprintf("user with username <<%s>> not found", c.Param("username")))
			} else {
				return err
			}
		} else {
			response = c.Write(util.H{
				"email": model.Email,
			})
		}

		return response
	}
}
