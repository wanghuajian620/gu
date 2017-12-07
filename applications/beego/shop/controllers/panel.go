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
 *     Initial: 2017/11/30        Wang RiYu
 */

package controllers

import (
  "github.com/fengyfei/gu/applications/beego/shop/mysql"
  "github.com/fengyfei/gu/applications/beego/base"
  "github.com/fengyfei/gu/models/shop/panel"
  "github.com/fengyfei/gu/models/shop/ware"
  "github.com/fengyfei/gu/libs/logger"
  "github.com/fengyfei/gu/libs/orm"
  "github.com/fengyfei/gu/libs/constants"
  "encoding/json"
  "errors"
  "strings"
  "github.com/fengyfei/gu/applications/beego/shop/util"
)

type (
  PanelController struct {
    base.Controller
  }
)

// add promotion panel
func (this *PanelController) AddPanel() {
  var (
    err error
    addReq panel.PanelReq
    conn orm.Connection
  )

  conn, err = mysql.Pool.Get()
  defer mysql.Pool.Release(conn)
  if err != nil {
    logger.Error(err)
    this.Data["json"] = map[string]interface{}{constants.RespKeyStatus: constants.ErrMysql}

    goto finish
  }

  err = json.Unmarshal(this.Ctx.Input.RequestBody, &addReq)
  if err != nil {
    logger.Error(err)
    this.Data["json"] = map[string]interface{}{constants.RespKeyStatus: constants.ErrInvalidParam}

    goto finish
  }

  err = this.Validate(addReq)
  if err != nil {
    logger.Error(err)
    this.Data["json"] = map[string]interface{}{constants.RespKeyStatus: constants.ErrInvalidParam}

    goto finish
  }

  err = panel.Service.CreatePanel(conn, addReq)
  if err != nil {
    logger.Error(err)
    this.Data["json"] = map[string]interface{}{constants.RespKeyStatus: constants.ErrMysql}

    goto finish
  }
  this.Data["json"] = map[string]interface{}{constants.RespKeyStatus: constants.ErrSucceed}
  logger.Info("create panel", addReq.Title, "success")

finish:
  this.ServeJSON(true)
}

// add promotion list
func (this *PanelController) AddPromotion() {
  var (
    err error
    addReq panel.PromotionReq
    conn orm.Connection
  )

  conn, err = mysql.Pool.Get()
  defer mysql.Pool.Release(conn)
  if err != nil {
    logger.Error(err)
    this.Data["json"] = map[string]interface{}{constants.RespKeyStatus: constants.ErrMysql}

    goto finish
  }

  err = json.Unmarshal(this.Ctx.Input.RequestBody, &addReq)
  if err != nil {
    logger.Error(err)
    this.Data["json"] = map[string]interface{}{constants.RespKeyStatus: constants.ErrInvalidParam}

    goto finish
  }

  err = this.Validate(addReq)
  if err != nil {
    logger.Error(err)
    this.Data["json"] = map[string]interface{}{constants.RespKeyStatus: constants.ErrInvalidParam}

    goto finish
  }

  err = panel.Service.AddPromotionList(conn, addReq)
  if err != nil {
    logger.Error(err)
    this.Data["json"] = map[string]interface{}{constants.RespKeyStatus: constants.ErrMysql}

    goto finish
  }
  this.Data["json"] = map[string]interface{}{constants.RespKeyStatus: constants.ErrSucceed}
  logger.Info("add promotion list success")

finish:
  this.ServeJSON(true)
}

// add recommend
func (this *PanelController) AddRecommend() {
  var (
    err error
    addReq panel.RecommendReq
    conn orm.Connection
  )

  conn, err = mysql.Pool.Get()
  defer mysql.Pool.Release(conn)
  if err != nil {
    logger.Error(err)
    this.Data["json"] = map[string]interface{}{constants.RespKeyStatus: constants.ErrMysql}

    goto finish
  }

  err = json.Unmarshal(this.Ctx.Input.RequestBody, &addReq)
  if err != nil {
    logger.Error(err)
    this.Data["json"] = map[string]interface{}{constants.RespKeyStatus: constants.ErrInvalidParam}

    goto finish
  }

  err = this.Validate(addReq)
  if err != nil {
    logger.Error(err)
    this.Data["json"] = map[string]interface{}{constants.RespKeyStatus: constants.ErrInvalidParam}

    goto finish
  }

  addReq.Picture, err = util.SavePicture(addReq.Picture, "recommend/")
  if err != nil {
    logger.Error(err)
    this.Data["json"] = map[string]interface{}{constants.RespKeyStatus: constants.ErrInternalServerError}

    goto finish
  }

  err = panel.Service.AddRecommend(conn, addReq)
  if err != nil {
    logger.Error(err)
    if !util.DeletePicture(addReq.Picture) {
      logger.Error(errors.New("add recommend failed and delete it's pictures go wrong, please delete picture manually"))
    }
    this.Data["json"] = map[string]interface{}{constants.RespKeyStatus: constants.ErrMysql}

    goto finish
  }
  this.Data["json"] = map[string]interface{}{constants.RespKeyStatus: constants.ErrSucceed}
  logger.Info("add recommend of panel success")

finish:
  this.ServeJSON(true)
}

// TODO: add second-hand
func (this *PanelController) AddSecondHand() {}

// get panel page
func (this *PanelController) GetPanelPage() {
  var (
    err error
    res []panel.PanelsPage
  )

  conn, err := mysql.Pool.Get()
  defer mysql.Pool.Release(conn)
  if err != nil {
    logger.Error(err)
    this.Data["json"] = map[string]interface{}{constants.RespKeyStatus: constants.ErrMysql}

    goto finish
  }

  res, err = panel.Service.GetPanels(conn)
  if err != nil {
    logger.Error(err)
    this.Data["json"] = map[string]interface{}{constants.RespKeyStatus: constants.ErrMysql}

    goto finish
  }

  for i := range res {
    if res[i].Type == 1 {
      detail, err := panel.Service.GetDetail(conn, res[i].ID)
      if err != nil {
        logger.Error(err)

        res[i].Content = []interface{}{}
      } else {
        ids := strings.Split(detail.Content, "#")

        wares, err := ware.Service.GetByIDs(conn, ids)
        if err != nil {
          logger.Error(err)

          res[i].Content = []interface{}{}
        } else {
          for k := range wares {
            res[i].Content = append(res[i].Content, wares[k])
          }
        }
      }
    }
    if res[i].Type == 2 {
      detail, err := panel.Service.GetDetail(conn, res[i].ID)
      if err != nil {
        logger.Error(err)

        res[i].Content = []interface{}{}
      } else {
        res[i].Content = append(res[i].Content, detail.Picture)
      }
    }
  }
  this.Data["json"] = res

finish:
  this.ServeJSON(true)
}
