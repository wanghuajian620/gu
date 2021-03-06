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
 *     Initial: 2018/01/27        Chen Yanchen
 */

package initialize

import (
	"fmt"

	"github.com/fengyfei/gu/applications/bbs/conf"
	"github.com/fengyfei/gu/libs/orm/mysql"
)

const (
	poolSize = 20
)

var (
	Pool *mysql.Pool
)

func init() {
	dataSource := fmt.Sprintf(conf.BBSConfig.MysqlUser + ":" + conf.BBSConfig.MysqlPass + "@" + "tcp(" + conf.BBSConfig.MysqlHost + ":" + conf.BBSConfig.MysqlPort + ")/" + conf.BBSConfig.MysqlDb + "?charset=utf8&parseTime=True&loc=Local")
	InitPool(dataSource)
}

// InitPool initialize the connection pool.
func InitPool(db string) {
	Pool = mysql.NewPool(db, poolSize)

	if Pool == nil {
		panic("MySQL DB connection error.")
	}
}
