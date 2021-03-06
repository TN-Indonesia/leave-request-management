package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"server/helpers"
	logicAdmin "server/models/logic/admin"
	logic "server/models/logic/user"
	logicUser "server/models/logic/user"
	structAPI "server/structs/api"
	structDB "server/structs/db"
	"strconv"

	"github.com/astaxie/beego"
)

// AdminController ...
type AdminController struct {
	beego.Controller
}

// CreateUser ...
func (c *AdminController) CreateUser() {
	var (
		resp    structAPI.RespData
		reqUser structAPI.ReqRegister
	)

	body := c.Ctx.Input.RequestBody
	fmt.Println("CREATE-USER=======>", string(body))

	errMarshal := json.Unmarshal(body, &reqUser)
	if errMarshal != nil {
		helpers.CheckErr("Failed unmarshall req body failed @CreateUser - controller", errMarshal)
		resp.Error = errors.New("Type request malform").Error()
		c.Ctx.Output.SetStatus(400)
		c.Ctx.Output.JSON(resp, false, false)
		return
	}

	user := structDB.User{
		EmployeeNumber:   reqUser.EmployeeNumber,
		Name:             reqUser.Name,
		Gender:           reqUser.Gender,
		Position:         reqUser.Position,
		StartWorkingDate: reqUser.StartWorkingDate,
		MobilePhone:      reqUser.MobilePhone,
		Email:            reqUser.Email,
		Password:         reqUser.Password,
		Role:             reqUser.Role,
		SupervisorID:     reqUser.SupervisorID,
	}

	lastRowID, errAddUser := logicAdmin.CreateUser(user)
	if errAddUser != nil {
		resp.Error = errAddUser.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = "Create user success"
	}

	go func() {
		if reqUser.Gender == "Male" && reqUser.Role == "employee" {
			logicUser.CreateUserTypeLeave(lastRowID, 11, 12.0)
			// logicUser.CreateUserTypeLeave(lastRowID, 22, 3.0)
			logicUser.CreateUserTypeLeave(lastRowID, 33, 30.0)
			logicUser.CreateUserTypeLeave(lastRowID, 44, 2.0)
			logicUser.CreateUserTypeLeave(lastRowID, 66, 99.0)
		} else if reqUser.Gender == "Male" && reqUser.Role == "supervisor" {
			logicUser.CreateUserTypeLeave(lastRowID, 11, 12.0)
			// logicUser.CreateUserTypeLeave(lastRowID, 22, 3.0)
			logicUser.CreateUserTypeLeave(lastRowID, 33, 30.0)
			logicUser.CreateUserTypeLeave(lastRowID, 44, 2.0)
			logicUser.CreateUserTypeLeave(lastRowID, 66, 99.0)
		} else if reqUser.Gender == "Female" && reqUser.Role == "employee" {
			logicUser.CreateUserTypeLeave(lastRowID, 11, 12.0)
			// logicUser.CreateUserTypeLeave(lastRowID, 22, 3.0)
			logicUser.CreateUserTypeLeave(lastRowID, 33, 30.0)
			logicUser.CreateUserTypeLeave(lastRowID, 44, 2.0)
			logicUser.CreateUserTypeLeave(lastRowID, 55, 90.0)
			logicUser.CreateUserTypeLeave(lastRowID, 66, 99.0)
		} else if reqUser.Gender == "Female" && reqUser.Role == "supervisor" {
			logicUser.CreateUserTypeLeave(lastRowID, 11, 12.0)
			// logicUser.CreateUserTypeLeave(lastRowID, 22, 3.0)
			logicUser.CreateUserTypeLeave(lastRowID, 33, 30.0)
			logicUser.CreateUserTypeLeave(lastRowID, 44, 2.0)
			logicUser.CreateUserTypeLeave(lastRowID, 55, 90.0)
			logicUser.CreateUserTypeLeave(lastRowID, 66, 99.0)
		}
	}()

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("Failed giving output @CreateUser - controller", err)
	}
}

// GetUsers ...
func (c *AdminController) GetUsers() {
	var resp structAPI.RespData

	res, errGet := logicAdmin.GetUsers()
	if errGet != nil {
		resp.Error = errGet.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = res
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("Failed giving output @GetUsers - controller", err)
	}
}

// GetUser ...
func (c *AdminController) GetUser() {
	var resp structAPI.RespData

	idStr := c.Ctx.Input.Param(":id")
	employeeNumber, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("Convert id failed @GetRequestAccept - controller", errCon)
		resp.Error = errors.New("Convert id failed").Error()
		return
	}

	res, errGet := logicAdmin.GetUser(employeeNumber)
	if errGet != nil {
		resp.Error = errGet.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = res
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("Failed giving output @GetUsers - controller", err)
	}
}

// DeleteUser ...
func (c *AdminController) DeleteUser() {
	var resp structAPI.RespData

	idStr := c.Ctx.Input.Param(":id")
	EmployeeNumber, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("Convert id failed @DeleteUser - controller", errCon)
		resp.Error = errors.New("Convert id failed").Error()
		return
	}

	if err := logicAdmin.DeleteUser(EmployeeNumber); err == nil {
		resp.Body = "Deleted success"
	} else {
		resp.Error = err.Error()
		c.Ctx.Output.SetStatus(400)
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("Failed giving output @DeleteUser - controller", err)
	}
}

// UpdateUser ...
func (c *AdminController) UpdateUser() {
	var (
		resp    structAPI.RespData
		reqUser structAPI.ReqRegister
	)

	body := c.Ctx.Input.RequestBody
	fmt.Println("UPDATE-USER=======>", string(body))

	err := json.Unmarshal(body, &reqUser)
	if err != nil {
		helpers.CheckErr("Failed unmarshall req body @UpdateUser - controller", err)
		resp.Error = errors.New("Type request malform").Error()
		c.Ctx.Output.JSON(resp, false, false)
		return
	}

	idStr := c.Ctx.Input.Param(":id")
	EmployeeNumber, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("Convert id failed @UpdateUser - controller", errCon)
		resp.Error = errors.New("Convert id failed").Error()
		return
	}

	resTime, errTime := helpers.NowLoc("Asia/Jakarta")
	helpers.CheckErr("Error get location time time @UpdateUser - controller", errTime)

	user := structDB.User{
		EmployeeNumber:   reqUser.EmployeeNumber,
		Name:             reqUser.Name,
		Gender:           reqUser.Gender,
		Position:         reqUser.Position,
		StartWorkingDate: reqUser.StartWorkingDate,
		MobilePhone:      reqUser.MobilePhone,
		Email:            reqUser.Email,
		Password:         reqUser.Password,
		Role:             reqUser.Role,
		SupervisorID:     reqUser.SupervisorID,
		UpdatedAt:        resTime,
	}

	errUpdate := logic.DBPostAdmin.UpdateUser(&user, EmployeeNumber)
	if errUpdate != nil {
		resp.Error = errUpdate.Error()
	} else {
		resp.Body = "Update user success"
	}

	err = c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("Failed giving output @UpdateUser - controller", err)
	}
}

// GetRequestPending ...
func (c *AdminController) GetRequestPending() {
	var resp structAPI.RespData

	resGet, errGetPending := logicAdmin.GetLeaveRequestPending()
	if errGetPending != nil {
		resp.Error = errGetPending.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = resGet
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @GetRequestPending", err)
	}
}

// GetRequestAccept ...
func (c *AdminController) GetRequestAccept() {
	var resp structAPI.RespData

	resGet, errGetAccept := logicAdmin.GetLeaveRequestApproved()
	if errGetAccept != nil {
		resp.Error = errGetAccept.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = resGet
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @GetRequestAccept", err)
	}
}

// GetRequestReject ...
func (c *AdminController) GetRequestReject() {
	var resp structAPI.RespData

	resGet, errGetReject := logicAdmin.GetLeaveRequestRejected()
	if errGetReject != nil {
		resp.Error = errGetReject.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = resGet
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @GetRequestReject", err)
	}
}

// CancelRequestLeave ...
func (c *AdminController) CancelRequestLeave() {
	var (
		resp structAPI.RespData
	)

	idStr := c.Ctx.Input.Param(":id")
	id, errCon := strconv.ParseInt(idStr, 0, 64)
	if errCon != nil {
		helpers.CheckErr("Convert id failed @CancelRequestLeave", errCon)
		resp.Error = errors.New("Convert id failed").Error()
		return
	}

	errUpStat := logicAdmin.CancelRequestLeave(id)
	if errUpStat != nil {
		resp.Error = errUpStat.Error()
	} else {
		resp.Body = "Leave request has been canceled and deleted"
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("Failed giving output @CancelRequestLeave", err)
	}
}

// ResetLeaveBalance ...
func (c *AdminController) ResetLeaveBalance() {
	var resp structAPI.RespData

	errReset := logicAdmin.ResetUserTypeLeave(12.0, 11)
	// errReset = logicAdmin.ResetUserTypeLeave(3, 22)
	errReset = logicAdmin.ResetUserTypeLeave(30.0, 33)
	errReset = logicAdmin.ResetUserTypeLeave(2.0, 44)
	errReset = logicAdmin.ResetUserTypeLeave(90.0, 55)
	errReset = logicAdmin.ResetUserTypeLeave(99.0, 66)

	if errReset != nil {
		resp.Error = errReset.Error()
		c.Ctx.Output.SetStatus(400)
	} else {
		resp.Body = "reset leave balance success"
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		helpers.CheckErr("failed giving output @ResetLeaveBalance", err)
	}
}
