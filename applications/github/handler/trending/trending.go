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
 *     Initial: 2017/12/28        Jia Chenhui
 */

package trending

import (
	"github.com/fengyfei/gu/applications/core"
	"github.com/fengyfei/gu/applications/github/crawler"
	"github.com/fengyfei/gu/applications/nats"
	"github.com/fengyfei/gu/libs/constants"
	"github.com/fengyfei/gu/libs/crawler/github"
	"github.com/fengyfei/gu/libs/http/server"
	"github.com/fengyfei/gu/libs/logger"
)

type (
	// langReq - The request struct that get the trending of the day of a language.
	langReq struct {
		Lang *string `json:"lang" validate:"required"`
	}

	// infoResp - The response struct that represents the trending of the day of a language.
	infoResp struct {
		Title    string `json:"title"`
		Abstract string `json:"abstract"`
		Lang     string `json:"lang"`
		Date     string `json:"date"`
		Stars    int    `json:"stars"`
		Today    int    `json:"today"`
	}
)

// LangInfo - Get library trending based on the language.
// If there is no data in cache, get data from GitHub.
func LangInfo(c *server.Context) error {
	var (
		err   error
		ok    bool
		req   langReq
		resp  []infoResp = make([]infoResp, 0)
		t     *github.Trending
		tList []*github.Trending
		info  infoResp
	)

	if err = c.JSONBody(&req); err != nil {
		logger.Error(err)
		return core.WriteStatusAndDataJSON(c, constants.ErrInvalidParam, nil)
	}

	if err = c.Validate(&req); err != nil {
		logger.Error(err)
		return core.WriteStatusAndDataJSON(c, constants.ErrInvalidParam, nil)
	}

readFromCache:
	tList, ok = crawler.TrendingCache.Load(*req.Lang)
	if !ok {
		nats.StartLangCrawler(req.Lang)
		goto readFromCache
	}

	for _, t = range tList {
		info = infoResp{
			Title:    t.Title,
			Abstract: t.Abstract,
			Lang:     t.Lang,
			Date:     t.Date,
			Stars:    t.Stars,
			Today:    t.Today,
		}

		resp = append(resp, info)
	}

	return core.WriteStatusAndDataJSON(c, constants.ErrSucceed, resp)
}
