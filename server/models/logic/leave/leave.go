package leave

import (
	"errors"
	"server/helpers"
	"server/models/logic/user"
	structAPI "server/structs/api"
	structLogic "server/structs/logic"
	"time"
)

// CreateLeaveRequestEmployee ...
func CreateLeaveRequestEmployee(
	employeeNumber int64,
	typeLeaveID int64,
	reason string,
	dateFrom string,
	dateTo string,
	halfDates []string,
	backOn string,
	total float64,
	address string,
	contactLeave string,
	status string,
	notes string,
) error {

	getEmployee, errGetEmployee := DBUser.GetEmployee(employeeNumber)
	if errGetEmployee != nil {
		helpers.CheckErr("Error get employee @CreateLeaveRequestEmployee", errGetEmployee)
		return errGetEmployee
	}

	// Check Working date must < 1 year for annual leave = 11
	if typeLeaveID == 11 {
		startWorkingDate, err := time.Parse("02-01-2006", getEmployee.StartWorkingDate)
		helpers.CheckErr("Error parse startWorkingDate @CreateLeaveRequestEmployee", err)
		after1YearWorkingDate := startWorkingDate.AddDate(1, 0, 0)

		dateFromCheck, err := time.Parse("02-01-2006", dateFrom)
		helpers.CheckErr("Error parse dateFrom @CreateLeaveRequestEmployee", err)

		dateToCheck, err := time.Parse("02-01-2006", dateTo)
		helpers.CheckErr("Error parse dateTo @CreateLeaveRequestEmployee", err)

		if ((after1YearWorkingDate.Equal(dateFromCheck) || after1YearWorkingDate.After(dateFromCheck)) &&
			(after1YearWorkingDate.Equal(dateToCheck) || after1YearWorkingDate.Before(dateToCheck))) ||
			after1YearWorkingDate.After(dateToCheck) {
			return errors.New("Employee not working > 1 year")
		}
	}

	//inquiry leave on the date
	errInquiry := DBLeave.InquiryLeaveRequest(employeeNumber, dateFrom)
	if errInquiry != nil {
		helpers.CheckErr("Error delete leave request @CreateLeaveRequestEmployee - logicLeave", errInquiry)
		return errInquiry
	}

	if typeLeaveID == 22 && total > 1 {
		err := errors.New("Cannot do errand leave on 2 or more consecutive days")
		return err
	}

	if typeLeaveID == 44 || typeLeaveID == 55 || typeLeaveID == 66 {
		errSpecial := DBLeave.InquiryLeaveRequestSpecial(employeeNumber)
		if errSpecial != nil {
			helpers.CheckErr("Special Leave - logicLeave", errSpecial)
			return errSpecial
		}
	}

	errInsert := DBLeave.CreateLeaveRequestEmployee(employeeNumber, typeLeaveID, reason, dateFrom, dateTo, halfDates, backOn, total, address, contactLeave, status, notes)
	if errInsert != nil {
		helpers.CheckErr("Error delete leave request @CreateLeaveRequestEmployee - logicLeave", errInsert)
		return errInsert
	}

	getSupervisorID, errGetSupervisorID := DBUser.GetSupervisor(employeeNumber)
	helpers.CheckErr("Error get supervisor id @CreateLeaveRequestEmployee", errGetSupervisorID)

	getSupervisor, errGetSupervisor := DBUser.GetEmployee(getSupervisorID.SupervisorID)
	helpers.CheckErr("Error get supervisor @CreateLeaveRequestEmployee", errGetSupervisor)

	go func() {
		helpers.GoMailSupervisor(getSupervisor.Email, getEmployee.Name, getSupervisor.Name)
	}()

	return errInsert
}

// CreateLeaveRequestSupervisor ...
func CreateLeaveRequestSupervisor(
	employeeNumber int64,
	typeLeaveID int64,
	reason string,
	dateFrom string,
	dateTo string,
	halfDates []string,
	backOn string,
	total float64,
	address string,
	contactLeave string,
	status string,
	notes string,
) error {

	getEmployee, errGetEmployee := DBUser.GetEmployee(employeeNumber)
	if errGetEmployee != nil {
		helpers.CheckErr("Error get employee @CreateLeaveRequestSupervisor - logicLeave", errGetEmployee)
		return errGetEmployee
	}

	//inquiry leave on the date
	errInquiry := DBLeave.InquiryLeaveRequest(employeeNumber, dateFrom)
	if errInquiry != nil {
		helpers.CheckErr("Error delete leave request @CreateLeaveRequestEmployee - logicLeave", errInquiry)
		return errInquiry
	}

	if typeLeaveID == 22 && total > 1 {
		err := errors.New("Cannot do errand leave on 2 or more consecutive days")
		return err
	}

	if typeLeaveID == 44 || typeLeaveID == 55 || typeLeaveID == 66 {
		errSpecial := DBLeave.InquiryLeaveRequestSpecial(employeeNumber)
		if errSpecial != nil {
			helpers.CheckErr("Special Leave - logicLeave", errSpecial)
			return errSpecial
		}
	}

	// Check Working date must < 1 year for annual leave = 11
	if typeLeaveID == 11 {
		startWorkingDate, err := time.Parse("02-01-2006", getEmployee.StartWorkingDate)
		helpers.CheckErr("Error parse startWorkingDate @CreateLeaveRequestSupervisor", err)
		after1YearWorkingDate := startWorkingDate.AddDate(1, 0, 0)

		dateFromCheck, err := time.Parse("02-01-2006", dateFrom)
		helpers.CheckErr("Error parse dateFrom @CreateLeaveRequestSupervisor", err)

		dateToCheck, err := time.Parse("02-01-2006", dateTo)
		helpers.CheckErr("Error parse dateTo @CreateLeaveRequestSupervisor", err)

		if ((after1YearWorkingDate.Equal(dateFromCheck) || after1YearWorkingDate.After(dateFromCheck)) &&
			(after1YearWorkingDate.Equal(dateToCheck) || after1YearWorkingDate.Before(dateToCheck))) ||
			after1YearWorkingDate.After(dateToCheck) {
			return errors.New("Employee not working > 1 year")
		}
	}

	getDirector, errGetDirector := user.GetDirector()
	helpers.CheckErr("Error get employee @CreateLeaveRequestSupervisor", errGetDirector)

	errInsert := DBLeave.CreateLeaveRequestSupervisor(employeeNumber, typeLeaveID, reason, dateFrom, dateTo, halfDates, backOn, total, address, contactLeave, status, notes)
	if errInsert != nil {
		helpers.CheckErr("Error delete leave request @CreateLeaveRequestSupervisor - logicLeave", errInsert)
		return errInsert
	}

	go func() {
		helpers.GoMailDirectorFromSupervisor(getDirector.Email, getEmployee.Name, getDirector.Name)
	}()

	return errInsert
}

// CreateLeaveRequestAdmin ...
func CreateLeaveRequestAdmin(
	employeeNumber int64,
	typeLeaveID int64,
	reason string,
	dateFrom string,
	dateTo string,
	halfDates []string,
	backOn string,
	total float64,
	address string,
	contactLeave string,
	status string,
	notes string,
) error {

	getEmployee, errGetEmployee := DBUser.GetEmployeeByEmployeeNumber(employeeNumber)
	if errGetEmployee != nil {
		helpers.CheckErr("Error get employee @CreateLeaveRequestAdmin - logicLeave", errGetEmployee)
		return errGetEmployee
	}

	getDirector, errGetDirector := user.GetDirector()
	helpers.CheckErr("Error get employee @CreateLeaveRequestAdmin", errGetDirector)

	errInsert := DBLeave.CreateLeaveRequestSupervisor(employeeNumber, typeLeaveID, reason, dateFrom, dateTo, halfDates, backOn, total, address, contactLeave, status, notes)
	if errInsert != nil {
		helpers.CheckErr("Error insert leave request @CreateLeaveRequestAdmin - logicLeave", errInsert)
		return errInsert
	}

	go func() {
		helpers.GoMailDirectorFromSupervisor(getDirector.Email, getEmployee.Name, getDirector.Name)
	}()

	return errInsert
}

// // UpdateRequest ...
// func UpdateRequest(e *structAPI.UpdateLeaveRequest) error {
// 	errUpdate := DBLeave.UpdateRequest(e)
// 	if errUpdate != nil {
// 		helpers.CheckErr("Error update leave request @UpdateRequest - logicLeave", errUpdate)
// 	}

// 	return errUpdate
// }

// GetLeave ...
func GetLeave(id int64) (structLogic.GetLeave, error) {
	respGet, errGet := DBLeave.GetLeave(id)
	if errGet != nil {
		helpers.CheckErr("Error get leave request @GetLeave - logicLeave", errGet)
	}

	return respGet, errGet
}

// DeleteRequest ...
func DeleteRequest(id int64) (err error) {
	errDelete := DBLeave.DeleteRequest(id)
	if errDelete != nil {
		helpers.CheckErr("Error delete leave request @DeleteRequest - logicLeave", errDelete)
		return errDelete
	}

	return errDelete
}

// UpdateLeaveRemaningApprove ...
func UpdateLeaveRemaningApprove(total float64, employeeNumber int64, typeID int64) (err error) {
	errUpdate := DBLeave.UpdateLeaveRemaningApprove(total, employeeNumber, typeID)
	if errUpdate != nil {
		helpers.CheckErr("Error update leave balance @UpdateLeaveRemaningApprove - logicLeave", errUpdate)
		return errUpdate
	}

	return errUpdate
}

// UpdateLeaveRemaningCancel ...
func UpdateLeaveRemaningCancel(total float64, employeeNumber int64, typeID int64) (err error) {
	errUpdate := DBLeave.UpdateLeaveRemaningCancel(total, employeeNumber, typeID)
	if errUpdate != nil {
		helpers.CheckErr("Error update leave balance @UpdateLeaveRemaningCancel - logicLeave", errUpdate)
		return errUpdate
	}

	return errUpdate
}

// DownloadReportCSV ...
func DownloadReportCSV(query *structAPI.RequestReport, path string) error {
	errGet := DBLeave.DownloadReportCSV(query.FromDate, query.ToDate, path)
	if errGet != nil {
		helpers.CheckErr("Error get report @DownloadReportCSV - logicLeave", errGet)
	}

	return errGet
}

// ReportLeaveRequest ...
func ReportLeaveRequest(query *structAPI.RequestReport) (report []structLogic.ReportLeaveRequest, err error) {
	respGet, errGet := DBLeave.ReportLeaveRequest(query.FromDate, query.ToDate)
	if errGet != nil {
		helpers.CheckErr("Error get report @ReportLeaveRequest - logicLeave", errGet)
	}

	return respGet, errGet
}

// ReportLeaveRequestTypeLeave ...
func ReportLeaveRequestTypeLeave(query *structAPI.RequestReportTypeLeave) (report []structLogic.ReportLeaveRequest, err error) {
	respGet, errGet := DBLeave.ReportLeaveRequestTypeLeave(query.FromDate, query.ToDate, query.TypeLeaveID)
	if errGet != nil {
		helpers.CheckErr("Error get report type leave @ReportLeaveRequestTypeLeave - logicLeave", errGet)
	}

	return respGet, errGet
}
