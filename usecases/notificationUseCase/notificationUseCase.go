package notificationUseCase

import (
	"fmt"
	"hms-backend/constants"
	"hms-backend/dto"
	"hms-backend/repositories/doctorRepository"
	"hms-backend/repositories/nurseRepository"
	"hms-backend/repositories/outpatientSessionRepository"
	"hms-backend/repositories/userRepository"
	"strconv"
)

type NotificationUseCase interface {
	GetByUserId(userId uint) ([]dto.Notification, error)
}

type notificationUseCase struct {
	outptnRep outpatientSessionRepository.OutpatientSessionRepository
	dctrRep   doctorRepository.DoctorRepository
	nrsRep    nurseRepository.NurseRepository
	userRep   userRepository.UserRepository
}

func New(
	outptRep outpatientSessionRepository.OutpatientSessionRepository,
	dctrRep doctorRepository.DoctorRepository,
	nrsRep nurseRepository.NurseRepository,
	userRep userRepository.UserRepository,
) *notificationUseCase {
	return &notificationUseCase{outptRep, dctrRep, nrsRep, userRep}
}

func (uc *notificationUseCase) GetByUserId(userId uint) ([]dto.Notification, error) {
	var res []dto.Notification

	user, err := uc.userRep.GetById(userId)
	if err != nil {
		return res, err
	}

	var doctorId uint

	if user.RoleId == 2 {
		doctor, err := uc.dctrRep.GetByUserId(userId)
		if err != nil {
			return res, err
		}

		doctorId = doctor.ID
	} else if user.RoleId == 3 {
		nurse, err := uc.nrsRep.GetByUserId(userId)
		if err != nil {
			return res, err
		}

		doctorId = nurse.DoctorId
	}

	outpatientSessions, err := uc.outptnRep.GetUnprocessedByDoctorId(doctorId)
	if err != nil {
		return res, err
	}

	for _, outpatientSession := range outpatientSessions {

		dateString := strconv.Itoa(outpatientSession.CreatedAt.Year()) + "-" + strconv.Itoa(int(outpatientSession.CreatedAt.Month())) + "-" + fmt.Sprintf("%02d", outpatientSession.CreatedAt.Day())
		timeString := fmt.Sprintf("%02d", outpatientSession.CreatedAt.Hour()) + ":" + fmt.Sprintf("%02d", outpatientSession.CreatedAt.Minute())

		dateIndoString := fmt.Sprintf("%02d", outpatientSession.CreatedAt.Day()) + " " +
			constants.Bulan[int(outpatientSession.CreatedAt.Month())] + " " +
			strconv.Itoa(outpatientSession.CreatedAt.Year())

		description := "Request Kunjungan"

		res = append(res, dto.Notification{
			OutpatientSessionID: outpatientSession.ID,
			Description:         description,
			DateString:          dateString,
			DateStringIndo:      dateIndoString,
			TimeString:          timeString,
		})
	}

	return res, nil
}
