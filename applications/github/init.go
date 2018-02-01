/*
 * MIT License
 *
 * Copyright (c) 2017 SmartestEE Co., Ltd..
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
 *     Initial: 2017/12/28        Jia Chenhui
 */

package main

import (
	"github.com/fengyfei/gu/applications/core"
	"github.com/fengyfei/gu/applications/github/conf"
	"github.com/fengyfei/gu/applications/github/router"
	"github.com/fengyfei/gu/libs/http/server"
	"github.com/fengyfei/gu/libs/http/server/middleware"
	"github.com/fengyfei/gu/libs/logger"
)

var (
	ep *server.Entrypoint
)

// startServer starts a HTTP server.
func startServer() {
	serverConfig := &server.Configuration{
		Address: conf.GithubConfig.Address,
	}

	ep = server.NewEntrypoint(serverConfig, nil)

	// add middlewares
	ep.AttachMiddleware(middleware.NegroniRecoverHandler())
	ep.AttachMiddleware(middleware.NegroniLoggerHandler())
	ep.AttachMiddleware(middleware.JWTWithConfig(middleware.JWTConfig{
		Skipper:    core.CustomSkipper,
		SigningKey: []byte(core.TokenHMACKey),
		ContextKey: core.ClaimsKey,
	}))
	ep.AttachMiddleware(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowedOrigins: conf.GithubConfig.CorsHosts,
		AllowedMethods: []string{server.GET, server.POST},
	}))

	if err := ep.Start(router.Router.Handler()); err != nil {
		logger.Error(err)
		return
	}

	ep.Wait()
}
