package controllers

import (
	"encoding/json"
	"manageFile/models"

	"github.com/astaxie/beego"
)

// Operations about StoreHouse
type StoreHouseController struct {
	beego.Controller
}

// @Title StoreBook
// @Description store books
// @Param	body		body 	models.StoreHouse	true		"body for StoreHouseController content"
// @Success 200 {int} models.StoreHouse.Id
// @Failure 403 body is empty
// @router / [post]
func (s *StoreHouseController) Post() {
	var store models.StoreHouse
	json.Unmarshal(s.Ctx.Input.RequestBody, &store)
	storeId := models.AddStore(store)
	//todo 上传文件
	
	s.Data["json"] = res.Success("", map[string]interface{}{"storeId": storeId})
	s.ServeJSON()
}

// @Title GetAll
// @Description get all Stores
// @Success 200 {storeHouse} models.StoreHouse
// @router / [get]
func (s *StoreHouseController) GetAll() {
	stores := models.GetAllStore()
	storeList := map[string]interface{}{
		"storeList": stores,
	}
	s.Data["json"] = res.Success("", storeList)
	s.ServeJSON()
}

// @Title Get
// @Description get store by id
// @Param	uid		path 	int64	true		"The key for staticblock"
// @Success 200 {storeHouse} models.StoreHouse
// @Failure 403 :storeId is empty
// @router /:uid [get]
func (s *StoreHouseController) Get() {
	storeId, _ := s.GetInt64(":storeId")
	if storeId != 0 {
		store, err := models.GetStoreById(storeId)
		if err != nil {
			s.Data["json"] = err.Error()
		} else {
			storeInfo := map[string]interface{}{
				"storeInfo": store,
			}
			s.Data["json"] = res.Success("", storeInfo)
		}
	}
	s.ServeJSON()
}

// @Title Update
// @Description update the store
// @Param	uid		path 	int64	true		"The uid you want to update"
// @Param	body		body 	models.StoreHouse	true		"body for store content"
// @Success 200 {storeHouse} models.StoreHouse
// @Failure 403 :storeId is not int
// @router /:storeId [put]
func (s *StoreHouseController) Put() {
	storeId, _ := s.GetInt64(":storeId")
	if storeId != 0 {
		var store models.StoreHouse
		json.Unmarshal(s.Ctx.Input.RequestBody, &store)
		_, err := models.UpdateStore(storeId, &store)
		if err != nil {
			s.Data["json"] = res.Fail(err.Error())
		} else {
			s.Data["json"] = res.Success("update success", make(map[string]interface{}))
		}
	}
	s.ServeJSON()
}

// @Title Delete
// @Description delete the store
// @Param	uid		path 	int64	true		"The store you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 storeId is empty
// @router /:storeId [delete]
func (s *StoreHouseController) Delete() {
	storeId, _ := s.GetInt64(":storeId")
	models.DeleteStore(storeId)
	s.Data["json"] = res.Success("delete success", make(map[string]interface{}))
	s.ServeJSON()
}
