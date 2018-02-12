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
 *     Initial: 2018/01/21        Li Zebang
 */

package segment

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/fengyfei/gu/libs/logger"
)

type Blog struct {
	Title string
	URL   string
	Date  string
	Blog  string
}

var (
	topMap  = map[string]string{"&amp;": "&", "&#x27;": "'", "\n": ""}
	bodyMap = map[string]string{"&#x27;": "'", "&amp;": "&", "&quot;": "\"", "&lt;": "<", "&gt;": ">"}
)

func (c *segmentCrawler) startBlog() {
	for {
		if url, ready := <-c.urlCh; ready {
			err := c.getBlog(url)
			if err != nil {
				c.errCh <- err
			}
			c.blogFinish <- ok{}
		}
	}
}

func (c *segmentCrawler) getBlog(url *string) error {
	cli := &http.Client{}

	req, err := http.NewRequest("GET", *url, nil)
	if err != nil {
		logger.Error("error in getting blog", err)
		return err
	}

	resp, err := cli.Do(req)
	if err != nil {
		logger.Error("error in getting blog", err)
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("error in getting blog", err)
		return err
	}

	b := &Blog{URL: *url}
	b.parseBlog(string(data))
	c.blogCh <- b
	return nil
}

func (b *Blog) parseBlog(s string) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(err, b.URL)
		}
	}()
	s = strings.SplitN(s, "<h1 class=\"Article-title\" data-reactid=\"39\">", 2)[1]
	s = strings.SplitN(s, "<footer class=\"Article-footer\" data-reactid=\"", 2)[0]
	data := strings.SplitN(s, "<div class=\"Article-body Content\" data-swiftype-name=\"body\" data-swiftype-type=\"text\" data-reactid=", 2)

	b.parseTop(data[0])
	b.parseBody(data[1])
}

func (b *Blog) parseTop(s string) {
	text := strings.SplitN(s, "</h1>", 2)
	b.Title = text[0]
	for k, v := range topMap {
		b.Title = strings.Replace(b.Title, k, v, -1)
	}
	b.Blog += "# " + b.Title + "\n\n"

	count := strings.Count(text[1], "<a class=\"Author-name\" href=\"")
	for i := 0; i < count; i++ {
		if i != 0 {
			b.Blog += " and "
		}
		text = strings.SplitN(text[1], "<a class=\"Author-name\" href=\"", 2)
		text = strings.SplitN(text[1], "\"", 2)
		url := site + text[0]
		text = strings.SplitN(text[1], ">", 2)
		text = strings.SplitN(text[1], "<", 2)
		name := text[0]
		b.Blog = fmt.Sprintf("%s[%s](%s)", b.Blog, name, url)
	}

	text = strings.SplitN(text[1], "<!-- /react-text -->", 3)
	text = strings.SplitN(text[1], "-->", 2)
	b.Date = text[1]
	b.Blog = fmt.Sprintf("%s on %s\n", b.Blog, b.Date)
}

func (b *Blog) parseBody(s string) {
	s = s[5 : len(s)-6]
	s = strings.Replace(s, "<hr/>", "\n---\n\n", -1)
	s = strings.Replace(s, "<br/>", "\n", -1)
	s = parseMD(s)
	for k, v := range bodyMap {
		s = strings.Replace(s, k, v, -1)
	}
	b.Blog += s
}