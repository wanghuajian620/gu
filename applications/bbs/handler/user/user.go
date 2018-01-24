/*
 * MIT License
 *
 * Copyright (c) 2018 SmartestEE Co., Ltd..
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
 *     Initial: 2018/01/21        Chen Yanchen
 */

package user

import (
	"github.com/fengyfei/gu/applications/core"
	"github.com/fengyfei/gu/libs/constants"
	"github.com/fengyfei/gu/libs/http/server"
	"github.com/fengyfei/gu/libs/logger"

	"github.com/fengyfei/gu/models/bbs/user"
)

//
func Login(u *server.Context) error {
	var req struct {
		Name     string
		Password string
	}

	if err := u.JSONBody(&req); err != nil {
		logger.Debug(err)
		return core.WriteStatusAndDataJSON(u, constants.ErrInvalidParam, nil)
	}

	err := user.UserServer.Login(req.Name, req.Password)
	if err != nil {
		logger.Debug(err)
		return core.WriteStatusAndDataJSON(u, constants.ErrMysql, nil)
	}

	return core.WriteStatusAndDataJSON(u, constants.ErrSucceed, nil)
	u.ServeJSON()
}

func Register(u *server.Context)  {

}