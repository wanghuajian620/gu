/*
 * MIT License
 *
 * Copyright (c) 2017 SmartestEE Co., Ltd.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*
 * Revision History:
 *     Initial: 2017/11/09        Jia Chenhui
 */

package role

import (
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	"github.com/fengyfei/gu/applications/echo/admin/mysql"
	"github.com/fengyfei/gu/applications/echo/core"
	"github.com/fengyfei/gu/libs/constants"
	"github.com/fengyfei/gu/models/staff"
)

type (
	// createReq - The request struct that create role information.
	createReq struct {
		Name  *string `json:"name" validate:"required,alpha,min=4,max=64"`
		Intro *string `json:"intro" validate:"required,alphanumunicode,min=4,max=256"`
	}

	// modifyReq - The request struct that modify role information.
	modifyReq struct {
		Id    int16   `json:"id" validate:"required"`
		Name  *string `json:"name" validate:"required,alpha,min=4,max=64"`
		Intro *string `json:"intro" validate:"required,alphanumunicode,min=4,max=256"`
	}

	// activateReq - The request struct that modify role status.
	activateReq struct {
		Id     int16 `json:"id" validate:"required"`
		Active bool  `json:"active"`
	}

	// infoReq - The request struct for get detail of specified role.
	infoReq struct {
		Id int16 `json:"id" validate:"required"`
	}

	// infoResp - The detail information for role.
	infoResp struct {
		Id      int16     `json:"id"`
		Name    string    `json:"name"`
		Intro   string    `json:"intro"`
		Active  bool      `json:"active"`
		Created time.Time `json:"created"`
	}
)

// Create - Create role information.
func Create(c echo.Context) error {
	var (
		err error
		req createReq
	)

	if err = c.Bind(&req); err != nil {
		return core.NewErrorWithMsg(constants.ErrInvalidParam, err.Error())
	}

	if err = c.Validate(&req); err != nil {
		return core.NewErrorWithMsg(constants.ErrInvalidParam, err.Error())
	}

	conn, err := mysql.Pool.Get()
	if err != nil {
		return core.NewErrorWithMsg(constants.ErrMysql, err.Error())
	}
	defer mysql.Pool.Release(conn)

	if err = staff.Service.CreateRole(conn, req.Name, req.Intro); err != nil {
		return core.NewErrorWithMsg(constants.ErrMysql, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		constants.RespKeyStatus: constants.ErrSucceed,
	})
}

// Modify - Modify role information.
func Modify(c echo.Context) error {
	var (
		err error
		req modifyReq
	)

	if err = c.Bind(&req); err != nil {
		return core.NewErrorWithMsg(constants.ErrInvalidParam, err.Error())
	}

	if err = c.Validate(&req); err != nil {
		return core.NewErrorWithMsg(constants.ErrInvalidParam, err.Error())
	}

	conn, err := mysql.Pool.Get()
	if err != nil {
		return core.NewErrorWithMsg(constants.ErrMysql, err.Error())
	}
	defer mysql.Pool.Release(conn)

	if err = staff.Service.ModifyRole(conn, req.Id, req.Name, req.Intro); err != nil {
		return core.NewErrorWithMsg(constants.ErrMysql, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		constants.RespKeyStatus: constants.ErrSucceed,
	})
}

// ModifyActive - Modify role status.
func ModifyActive(c echo.Context) error {
	var (
		err error
		req activateReq
	)

	if err = c.Bind(&req); err != nil {
		return core.NewErrorWithMsg(constants.ErrInvalidParam, err.Error())
	}

	if err = c.Validate(&req); err != nil {
		return core.NewErrorWithMsg(constants.ErrInvalidParam, err.Error())
	}

	conn, err := mysql.Pool.Get()
	if err != nil {
		return core.NewErrorWithMsg(constants.ErrMysql, err.Error())
	}
	defer mysql.Pool.Release(conn)

	if err = staff.Service.ModifyRoleActive(conn, req.Id, req.Active); err != nil {
		return core.NewErrorWithMsg(constants.ErrMysql, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		constants.RespKeyStatus: constants.ErrSucceed,
	})
}

// List - Get a list of active role details.
func List(c echo.Context) error {
	var resp []infoResp = make([]infoResp, 0)

	conn, err := mysql.Pool.Get()
	if err != nil {
		return core.NewErrorWithMsg(constants.ErrMysql, err.Error())
	}
	defer mysql.Pool.Release(conn)

	rlist, err := staff.Service.RoleList(conn)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return core.NewErrorWithMsg(constants.ErrMysql, err.Error())
		}

		return core.NewErrorWithMsg(constants.ErrMysql, err.Error())
	}

	for _, r := range rlist {
		info := infoResp{
			Id:      r.Id,
			Name:    r.Name,
			Intro:   r.Intro,
			Active:  r.Active,
			Created: *r.Created,
		}

		resp = append(resp, info)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		constants.RespKeyStatus: constants.ErrSucceed,
		constants.RespKeyData:   resp,
	})
}

// Info - Get detail information for specified role.
func Info(c echo.Context) error {
	var (
		err error
		req infoReq
	)

	if err = c.Bind(&req); err != nil {
		return core.NewErrorWithMsg(constants.ErrInvalidParam, err.Error())
	}

	if err = c.Validate(&req); err != nil {
		return core.NewErrorWithMsg(constants.ErrInvalidParam, err.Error())
	}

	conn, err := mysql.Pool.Get()
	if err != nil {
		return core.NewErrorWithMsg(constants.ErrMysql, err.Error())
	}
	defer mysql.Pool.Release(conn)

	info, err := staff.Service.GetRoleByID(conn, req.Id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return core.NewErrorWithMsg(constants.ErrMysql, err.Error())
		}

		return core.NewErrorWithMsg(constants.ErrMysql, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		constants.RespKeyStatus: constants.ErrSucceed,
		constants.RespKeyData:   *info,
	})
}
