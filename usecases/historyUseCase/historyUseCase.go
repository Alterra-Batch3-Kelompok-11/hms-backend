package historyUseCase

import (
	"fmt"
	"hms-backend/constants"
	"hms-backend/dto"
	"hms-backend/repositories/outpatientSessionRepository"
	"hms-backend/repositories/patientRepository"
	"strconv"
)

type HistoryUseCase interface {
	GetOutpatientSessionHistory(doctorId uint) ([]dto.History, error)
	GetApprovalHistory(doctorId uint) ([]dto.History, error)
}

type historyUseCase struct {
	outptnRep  outpatientSessionRepository.OutpatientSessionRepository
	patientRep patientRepository.PatientRepository
}

func New(
	outptRep outpatientSessionRepository.OutpatientSessionRepository,
	ptRep patientRepository.PatientRepository,
) *historyUseCase {
	return &historyUseCase{outptRep, ptRep}
}

func (uc *historyUseCase) GetOutpatientSessionHistory(doctorId uint) ([]dto.History, error) {
	var res []dto.History

	outpatientSessions, err := uc.outptnRep.GetApprovedAllByDoctorId(doctorId)
	if err != nil {
		return res, err
	}

	for _, outpatientSession := range outpatientSessions {

		patient, err := uc.patientRep.GetById(outpatientSession.PatientId)
		if err != nil {
			continue
		}

		dateString := fmt.Sprintf("%02d", outpatientSession.Schedule.Day()) + "-" + strconv.Itoa(int(outpatientSession.Schedule.Month())) + "-" + strconv.Itoa(outpatientSession.Schedule.Year())
		timeString := fmt.Sprintf("%02d", outpatientSession.Schedule.Hour()) + ":" + fmt.Sprintf("%02d", outpatientSession.Schedule.Minute())

		dateIndoString := fmt.Sprintf("%02d", outpatientSession.Schedule.Day()) + " " +
			constants.Bulan[int(outpatientSession.Schedule.Month())] + " " +
			strconv.Itoa(outpatientSession.Schedule.Year())

		status := "Proses"
		if outpatientSession.IsFinish {
			status = "Selesai"
		}

		res = append(res, dto.History{
			PatientName:      patient.Name,
			Schedule:         outpatientSession.Schedule,
			ScheduleDate:     dateString,
			ScheduleDateIndo: dateIndoString,
			ScheduleTime:     timeString,
			Status:           status,
		})

	}

	return res, nil
}
func (uc *historyUseCase) GetApprovalHistory(doctorId uint) ([]dto.History, error) {
	var res []dto.History

	outpatientSessions, err := uc.outptnRep.GetProcessedAllByDoctorId(doctorId)
	if err != nil {
		return res, err
	}

	for _, outpatientSession := range outpatientSessions {

		patient, err := uc.patientRep.GetById(outpatientSession.PatientId)
		if err != nil {
			continue
		}

		dateString := strconv.Itoa(outpatientSession.Schedule.Year()) + "-" + strconv.Itoa(int(outpatientSession.Schedule.Month())) + "-" + fmt.Sprintf("%02d", outpatientSession.Schedule.Day())
		timeString := fmt.Sprintf("%02d", outpatientSession.Schedule.Hour()) + ":" + fmt.Sprintf("%02d", outpatientSession.Schedule.Minute())

		dateIndoString := fmt.Sprintf("%02d", outpatientSession.Schedule.Day()) + " " +
			constants.Bulan[int(outpatientSession.Schedule.Month())] + " " +
			strconv.Itoa(outpatientSession.Schedule.Year())

		status := "Kunjungan Ditolak"
		if outpatientSession.IsApproved == 1 {
			status = "Disetujui"
		}

		res = append(res, dto.History{
			PatientName:      patient.Name,
			Schedule:         outpatientSession.Schedule,
			ScheduleDate:     dateString,
			ScheduleDateIndo: dateIndoString,
			ScheduleTime:     timeString,
			Status:           status,
		})

	}

	return res, nil
}
