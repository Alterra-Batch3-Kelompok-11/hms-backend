package outpatientSessionUseCase

import (
	"fmt"
	"gorm.io/gorm"
	"hms-backend/constants"
	"hms-backend/dto"
	"hms-backend/models"
	"hms-backend/repositories/doctorRepository"
	"hms-backend/repositories/doctorScheduleRepository"
	"hms-backend/repositories/outpatientSessionRepository"
	"hms-backend/repositories/patientRepository"
	"hms-backend/repositories/specialityRepository"
	"hms-backend/repositories/userRepository"
	"strconv"
	"time"
)

type OutpatientSessionUseCase interface {
	GetAll() ([]dto.OutpatientSessionRes, error)
	GetById(id uint) (dto.OutpatientSessionRes, error)
	GetByDoctorId(doctorId uint) ([]dto.OutpatientSessionRes, error)
	GetByPatientId(patientId uint) ([]dto.OutpatientSessionRes, error)
	GetUnprocessedByDoctorId(doctorId uint) ([]dto.OutpatientSessionRes, error)
	GetProcessedByDoctorId(doctorId uint) ([]dto.OutpatientSessionRes, error)
	GetApprovedByDoctorId(doctorId uint) ([]dto.OutpatientSessionRes, error)
	GetRejectedByDoctorId(doctorId uint) ([]dto.OutpatientSessionRes, error)
	Create(payload dto.OutpatientSessionReq) (dto.OutpatientSessionRes, error)
	Update(id uint, payload dto.OutpatientSessionReq) (dto.OutpatientSessionRes, error)
	Approval(id uint, payload dto.ApprovalReq) (dto.OutpatientSessionRes, error)
	Delete(id uint) error
}

type outpatientSessionUseCase struct {
	outptnRep  outpatientSessionRepository.OutpatientSessionRepository
	userRep    userRepository.UserRepository
	doctorRep  doctorRepository.DoctorRepository
	specRep    specialityRepository.SpecialityRepository
	schedRep   doctorScheduleRepository.DoctorScheduleRepository
	patientRep patientRepository.PatientRepository
}

func New(
	outptRep outpatientSessionRepository.OutpatientSessionRepository,
	usrRep userRepository.UserRepository,
	dctRep doctorRepository.DoctorRepository,
	spcRep specialityRepository.SpecialityRepository,
	scdRep doctorScheduleRepository.DoctorScheduleRepository,
	ptRep patientRepository.PatientRepository,
) *outpatientSessionUseCase {
	return &outpatientSessionUseCase{outptRep, usrRep, dctRep, spcRep, scdRep, ptRep}
}

func (uc *outpatientSessionUseCase) GetAll() ([]dto.OutpatientSessionRes, error) {
	var res []dto.OutpatientSessionRes
	outpatientSessions, err := uc.outptnRep.GetAll()
	if err != nil {
		return res, err
	}

	for _, outpatientSession := range outpatientSessions {
		doctor, err := uc.doctorRep.GetById(outpatientSession.DoctorId)
		if err != nil {
			return res, err
		}

		user, err := uc.userRep.GetById(doctor.UserId)
		if err != nil {
			return res, err
		}

		speciality, err := uc.specRep.GetById(doctor.SpecialityId)
		if err != nil {
			return res, err
		}

		patient, err := uc.patientRep.GetById(outpatientSession.PatientId)
		if err != nil {
			return res, err
		}

		dateString := strconv.Itoa(outpatientSession.Schedule.Year()) + "-" + strconv.Itoa(int(outpatientSession.Schedule.Month())) + "-" + fmt.Sprintf("%02d", outpatientSession.Schedule.Day())
		timeString := fmt.Sprintf("%02d", outpatientSession.Schedule.Hour()) + ":" + fmt.Sprintf("%02d", outpatientSession.Schedule.Minute())

		dateIndoString := fmt.Sprintf("%02d", outpatientSession.Schedule.Day()) + " " +
			constants.Bulan[int(outpatientSession.Schedule.Month())] + " " +
			strconv.Itoa(outpatientSession.Schedule.Year())

		res = append(res, dto.OutpatientSessionRes{
			ID:               outpatientSession.ID,
			CreatedAt:        outpatientSession.CreatedAt,
			UpdatedAt:        outpatientSession.UpdatedAt,
			DeletedAt:        outpatientSession.DeletedAt,
			DoctorId:         outpatientSession.DoctorId,
			PatientId:        outpatientSession.PatientId,
			Schedule:         outpatientSession.Schedule,
			Complaint:        outpatientSession.Complaint,
			IsApproved:       outpatientSession.IsApproved,
			IsFinish:         outpatientSession.IsFinish,
			FinishedAt:       outpatientSession.FinishedAt,
			ScheduleDate:     dateString,
			ScheduleDateIndo: dateIndoString,
			ScheduleTime:     timeString,
			Patient: dto.PatientRes{
				ID:            patient.ID,
				CreatedAt:     patient.CreatedAt,
				UpdatedAt:     patient.UpdatedAt,
				DeletedAt:     patient.DeletedAt,
				Nik:           patient.Nik,
				Name:          patient.Name,
				BirthDate:     patient.BirthDate,
				Gender:        patient.Gender,
				Address:       patient.Address,
				Phone:         patient.Phone,
				MaritalStatus: patient.MaritalStatus,
				ReligionID:    patient.ReligionID,
			},
			Doctor: dto.DoctorRes{
				ID:              doctor.ID,
				CreatedAt:       doctor.CreatedAt,
				UpdatedAt:       doctor.UpdatedAt,
				DeletedAt:       doctor.DeletedAt,
				Name:            user.Name,
				SpecialityId:    doctor.SpecialityId,
				LicenseNumber:   doctor.LicenseNumber,
				SpecialityName:  speciality.Name,
				DoctorSchedules: nil,
			},
		})
	}

	return res, nil
}
func (uc *outpatientSessionUseCase) GetById(id uint) (dto.OutpatientSessionRes, error) {
	var res dto.OutpatientSessionRes
	outpatientSession, err := uc.outptnRep.GetById(id)
	if err != nil {
		return res, err
	}

	doctor, err := uc.doctorRep.GetById(outpatientSession.DoctorId)
	if err != nil {
		return res, err
	}

	user, err := uc.userRep.GetById(doctor.UserId)
	if err != nil {
		return res, err
	}

	speciality, err := uc.specRep.GetById(doctor.SpecialityId)
	if err != nil {
		return res, err
	}

	patient, err := uc.patientRep.GetById(outpatientSession.PatientId)
	if err != nil {
		return res, err
	}

	dateString := strconv.Itoa(outpatientSession.Schedule.Year()) + "-" + strconv.Itoa(int(outpatientSession.Schedule.Month())) + "-" + fmt.Sprintf("%02d", outpatientSession.Schedule.Day())
	timeString := fmt.Sprintf("%02d", outpatientSession.Schedule.Hour()) + ":" + fmt.Sprintf("%02d", outpatientSession.Schedule.Minute())

	dateIndoString := fmt.Sprintf("%02d", outpatientSession.Schedule.Day()) + " " +
		constants.Bulan[int(outpatientSession.Schedule.Month())] + " " +
		strconv.Itoa(outpatientSession.Schedule.Year())

	res = dto.OutpatientSessionRes{
		ID:               outpatientSession.ID,
		CreatedAt:        outpatientSession.CreatedAt,
		UpdatedAt:        outpatientSession.UpdatedAt,
		DeletedAt:        outpatientSession.DeletedAt,
		DoctorId:         outpatientSession.DoctorId,
		PatientId:        outpatientSession.PatientId,
		Schedule:         outpatientSession.Schedule,
		Complaint:        outpatientSession.Complaint,
		IsApproved:       outpatientSession.IsApproved,
		IsFinish:         outpatientSession.IsFinish,
		FinishedAt:       outpatientSession.FinishedAt,
		ScheduleDate:     dateString,
		ScheduleDateIndo: dateIndoString,
		ScheduleTime:     timeString,
		Patient: dto.PatientRes{
			ID:            patient.ID,
			CreatedAt:     patient.CreatedAt,
			UpdatedAt:     patient.UpdatedAt,
			DeletedAt:     patient.DeletedAt,
			Nik:           patient.Nik,
			Name:          patient.Name,
			BirthDate:     patient.BirthDate,
			Gender:        patient.Gender,
			Address:       patient.Address,
			Phone:         patient.Phone,
			MaritalStatus: patient.MaritalStatus,
			ReligionID:    patient.ReligionID,
		},
		Doctor: dto.DoctorRes{
			ID:              doctor.ID,
			CreatedAt:       doctor.CreatedAt,
			UpdatedAt:       doctor.UpdatedAt,
			DeletedAt:       doctor.DeletedAt,
			Name:            user.Name,
			SpecialityId:    doctor.SpecialityId,
			LicenseNumber:   doctor.LicenseNumber,
			SpecialityName:  speciality.Name,
			DoctorSchedules: nil,
		},
	}

	return res, nil
}
func (uc *outpatientSessionUseCase) GetByDoctorId(doctorId uint) ([]dto.OutpatientSessionRes, error) {
	var res []dto.OutpatientSessionRes
	outpatientSessions, err := uc.outptnRep.GetByDoctorId(doctorId)
	if err != nil {
		return res, err
	}

	for _, outpatientSession := range outpatientSessions {
		doctor, err := uc.doctorRep.GetById(outpatientSession.DoctorId)
		if err != nil {
			return res, err
		}

		user, err := uc.userRep.GetById(doctor.UserId)
		if err != nil {
			return res, err
		}

		speciality, err := uc.specRep.GetById(doctor.SpecialityId)
		if err != nil {
			return res, err
		}

		patient, err := uc.patientRep.GetById(outpatientSession.PatientId)
		if err != nil {
			return res, err
		}

		dateString := strconv.Itoa(outpatientSession.Schedule.Year()) + "-" + strconv.Itoa(int(outpatientSession.Schedule.Month())) + "-" + fmt.Sprintf("%02d", outpatientSession.Schedule.Day())
		timeString := fmt.Sprintf("%02d", outpatientSession.Schedule.Hour()) + ":" + fmt.Sprintf("%02d", outpatientSession.Schedule.Minute())

		dateIndoString := fmt.Sprintf("%02d", outpatientSession.Schedule.Day()) + " " +
			constants.Bulan[int(outpatientSession.Schedule.Month())] + " " +
			strconv.Itoa(outpatientSession.Schedule.Year())

		res = append(res, dto.OutpatientSessionRes{
			ID:               outpatientSession.ID,
			CreatedAt:        outpatientSession.CreatedAt,
			UpdatedAt:        outpatientSession.UpdatedAt,
			DeletedAt:        outpatientSession.DeletedAt,
			DoctorId:         outpatientSession.DoctorId,
			PatientId:        outpatientSession.PatientId,
			Schedule:         outpatientSession.Schedule,
			Complaint:        outpatientSession.Complaint,
			IsApproved:       outpatientSession.IsApproved,
			IsFinish:         outpatientSession.IsFinish,
			FinishedAt:       outpatientSession.FinishedAt,
			ScheduleDate:     dateString,
			ScheduleDateIndo: dateIndoString,
			ScheduleTime:     timeString,
			Patient: dto.PatientRes{
				ID:            patient.ID,
				CreatedAt:     patient.CreatedAt,
				UpdatedAt:     patient.UpdatedAt,
				DeletedAt:     patient.DeletedAt,
				Nik:           patient.Nik,
				Name:          patient.Name,
				BirthDate:     patient.BirthDate,
				Gender:        patient.Gender,
				Address:       patient.Address,
				Phone:         patient.Phone,
				MaritalStatus: patient.MaritalStatus,
				ReligionID:    patient.ReligionID,
			},
			Doctor: dto.DoctorRes{
				ID:              doctor.ID,
				CreatedAt:       doctor.CreatedAt,
				UpdatedAt:       doctor.UpdatedAt,
				DeletedAt:       doctor.DeletedAt,
				Name:            user.Name,
				SpecialityId:    doctor.SpecialityId,
				LicenseNumber:   doctor.LicenseNumber,
				SpecialityName:  speciality.Name,
				DoctorSchedules: nil,
			},
		})
	}

	return res, nil
}
func (uc *outpatientSessionUseCase) GetByPatientId(patientId uint) ([]dto.OutpatientSessionRes, error) {
	var res []dto.OutpatientSessionRes
	outpatientSessions, err := uc.outptnRep.GetByPatientId(patientId)
	if err != nil {
		return res, err
	}

	for _, outpatientSession := range outpatientSessions {
		doctor, err := uc.doctorRep.GetById(outpatientSession.DoctorId)
		if err != nil {
			return res, err
		}

		user, err := uc.userRep.GetById(doctor.UserId)
		if err != nil {
			return res, err
		}

		speciality, err := uc.specRep.GetById(doctor.SpecialityId)
		if err != nil {
			return res, err
		}

		patient, err := uc.patientRep.GetById(outpatientSession.PatientId)
		if err != nil {
			return res, err
		}

		dateString := strconv.Itoa(outpatientSession.Schedule.Year()) + "-" + strconv.Itoa(int(outpatientSession.Schedule.Month())) + "-" + fmt.Sprintf("%02d", outpatientSession.Schedule.Day())
		timeString := fmt.Sprintf("%02d", outpatientSession.Schedule.Hour()) + ":" + fmt.Sprintf("%02d", outpatientSession.Schedule.Minute())

		dateIndoString := fmt.Sprintf("%02d", outpatientSession.Schedule.Day()) + " " +
			constants.Bulan[int(outpatientSession.Schedule.Month())] + " " +
			strconv.Itoa(outpatientSession.Schedule.Year())

		res = append(res, dto.OutpatientSessionRes{
			ID:               outpatientSession.ID,
			CreatedAt:        outpatientSession.CreatedAt,
			UpdatedAt:        outpatientSession.UpdatedAt,
			DeletedAt:        outpatientSession.DeletedAt,
			DoctorId:         outpatientSession.DoctorId,
			PatientId:        outpatientSession.PatientId,
			Schedule:         outpatientSession.Schedule,
			Complaint:        outpatientSession.Complaint,
			IsApproved:       outpatientSession.IsApproved,
			IsFinish:         outpatientSession.IsFinish,
			FinishedAt:       outpatientSession.FinishedAt,
			ScheduleDate:     dateString,
			ScheduleDateIndo: dateIndoString,
			ScheduleTime:     timeString,
			Patient: dto.PatientRes{
				ID:            patient.ID,
				CreatedAt:     patient.CreatedAt,
				UpdatedAt:     patient.UpdatedAt,
				DeletedAt:     patient.DeletedAt,
				Nik:           patient.Nik,
				Name:          patient.Name,
				BirthDate:     patient.BirthDate,
				Gender:        patient.Gender,
				Address:       patient.Address,
				Phone:         patient.Phone,
				MaritalStatus: patient.MaritalStatus,
				ReligionID:    patient.ReligionID,
			},
			Doctor: dto.DoctorRes{
				ID:              doctor.ID,
				CreatedAt:       doctor.CreatedAt,
				UpdatedAt:       doctor.UpdatedAt,
				DeletedAt:       doctor.DeletedAt,
				Name:            user.Name,
				SpecialityId:    doctor.SpecialityId,
				LicenseNumber:   doctor.LicenseNumber,
				SpecialityName:  speciality.Name,
				DoctorSchedules: nil,
			},
		})
	}

	return res, nil
}
func (uc *outpatientSessionUseCase) GetUnprocessedByDoctorId(doctorId uint) ([]dto.OutpatientSessionRes, error) {
	var res []dto.OutpatientSessionRes
	outpatientSessions, err := uc.outptnRep.GetUnprocessedByDoctorId(doctorId)
	if err != nil {
		return res, err
	}

	for _, outpatientSession := range outpatientSessions {
		doctor, err := uc.doctorRep.GetById(outpatientSession.DoctorId)
		if err != nil {
			return res, err
		}

		user, err := uc.userRep.GetById(doctor.UserId)
		if err != nil {
			return res, err
		}

		speciality, err := uc.specRep.GetById(doctor.SpecialityId)
		if err != nil {
			return res, err
		}

		patient, err := uc.patientRep.GetById(outpatientSession.PatientId)
		if err != nil {
			return res, err
		}

		dateString := strconv.Itoa(outpatientSession.Schedule.Year()) + "-" + strconv.Itoa(int(outpatientSession.Schedule.Month())) + "-" + fmt.Sprintf("%02d", outpatientSession.Schedule.Day())
		timeString := fmt.Sprintf("%02d", outpatientSession.Schedule.Hour()) + ":" + fmt.Sprintf("%02d", outpatientSession.Schedule.Minute())

		dateIndoString := fmt.Sprintf("%02d", outpatientSession.Schedule.Day()) + " " +
			constants.Bulan[int(outpatientSession.Schedule.Month())] + " " +
			strconv.Itoa(outpatientSession.Schedule.Year())

		res = append(res, dto.OutpatientSessionRes{
			ID:               outpatientSession.ID,
			CreatedAt:        outpatientSession.CreatedAt,
			UpdatedAt:        outpatientSession.UpdatedAt,
			DeletedAt:        outpatientSession.DeletedAt,
			DoctorId:         outpatientSession.DoctorId,
			PatientId:        outpatientSession.PatientId,
			Schedule:         outpatientSession.Schedule,
			Complaint:        outpatientSession.Complaint,
			IsApproved:       outpatientSession.IsApproved,
			IsFinish:         outpatientSession.IsFinish,
			FinishedAt:       outpatientSession.FinishedAt,
			ScheduleDate:     dateString,
			ScheduleDateIndo: dateIndoString,
			ScheduleTime:     timeString,
			Patient: dto.PatientRes{
				ID:            patient.ID,
				CreatedAt:     patient.CreatedAt,
				UpdatedAt:     patient.UpdatedAt,
				DeletedAt:     patient.DeletedAt,
				Nik:           patient.Nik,
				Name:          patient.Name,
				BirthDate:     patient.BirthDate,
				Gender:        patient.Gender,
				Address:       patient.Address,
				Phone:         patient.Phone,
				MaritalStatus: patient.MaritalStatus,
				ReligionID:    patient.ReligionID,
			},
			Doctor: dto.DoctorRes{
				ID:              doctor.ID,
				CreatedAt:       doctor.CreatedAt,
				UpdatedAt:       doctor.UpdatedAt,
				DeletedAt:       doctor.DeletedAt,
				Name:            user.Name,
				SpecialityId:    doctor.SpecialityId,
				LicenseNumber:   doctor.LicenseNumber,
				SpecialityName:  speciality.Name,
				DoctorSchedules: nil,
			},
		})
	}

	return res, nil
}
func (uc *outpatientSessionUseCase) GetProcessedByDoctorId(doctorId uint) ([]dto.OutpatientSessionRes, error) {
	var res []dto.OutpatientSessionRes
	outpatientSessions, err := uc.outptnRep.GetProcessedByDoctorId(doctorId)
	if err != nil {
		return res, err
	}

	for _, outpatientSession := range outpatientSessions {
		doctor, err := uc.doctorRep.GetById(outpatientSession.DoctorId)
		if err != nil {
			return res, err
		}

		user, err := uc.userRep.GetById(doctor.UserId)
		if err != nil {
			return res, err
		}

		speciality, err := uc.specRep.GetById(doctor.SpecialityId)
		if err != nil {
			return res, err
		}

		patient, err := uc.patientRep.GetById(outpatientSession.PatientId)
		if err != nil {
			return res, err
		}

		dateString := strconv.Itoa(outpatientSession.Schedule.Year()) + "-" + strconv.Itoa(int(outpatientSession.Schedule.Month())) + "-" + fmt.Sprintf("%02d", outpatientSession.Schedule.Day())
		timeString := fmt.Sprintf("%02d", outpatientSession.Schedule.Hour()) + ":" + fmt.Sprintf("%02d", outpatientSession.Schedule.Minute())

		dateIndoString := fmt.Sprintf("%02d", outpatientSession.Schedule.Day()) + " " +
			constants.Bulan[int(outpatientSession.Schedule.Month())] + " " +
			strconv.Itoa(outpatientSession.Schedule.Year())

		res = append(res, dto.OutpatientSessionRes{
			ID:               outpatientSession.ID,
			CreatedAt:        outpatientSession.CreatedAt,
			UpdatedAt:        outpatientSession.UpdatedAt,
			DeletedAt:        outpatientSession.DeletedAt,
			DoctorId:         outpatientSession.DoctorId,
			PatientId:        outpatientSession.PatientId,
			Schedule:         outpatientSession.Schedule,
			Complaint:        outpatientSession.Complaint,
			IsApproved:       outpatientSession.IsApproved,
			IsFinish:         outpatientSession.IsFinish,
			FinishedAt:       outpatientSession.FinishedAt,
			ScheduleDate:     dateString,
			ScheduleDateIndo: dateIndoString,
			ScheduleTime:     timeString,
			Patient: dto.PatientRes{
				ID:            patient.ID,
				CreatedAt:     patient.CreatedAt,
				UpdatedAt:     patient.UpdatedAt,
				DeletedAt:     patient.DeletedAt,
				Nik:           patient.Nik,
				Name:          patient.Name,
				BirthDate:     patient.BirthDate,
				Gender:        patient.Gender,
				Address:       patient.Address,
				Phone:         patient.Phone,
				MaritalStatus: patient.MaritalStatus,
				ReligionID:    patient.ReligionID,
			},
			Doctor: dto.DoctorRes{
				ID:              doctor.ID,
				CreatedAt:       doctor.CreatedAt,
				UpdatedAt:       doctor.UpdatedAt,
				DeletedAt:       doctor.DeletedAt,
				Name:            user.Name,
				SpecialityId:    doctor.SpecialityId,
				LicenseNumber:   doctor.LicenseNumber,
				SpecialityName:  speciality.Name,
				DoctorSchedules: nil,
			},
		})
	}

	return res, nil
}
func (uc *outpatientSessionUseCase) GetApprovedByDoctorId(doctorId uint) ([]dto.OutpatientSessionRes, error) {
	var res []dto.OutpatientSessionRes
	outpatientSessions, err := uc.outptnRep.GetApprovedByDoctorId(doctorId)
	if err != nil {
		return res, err
	}

	for _, outpatientSession := range outpatientSessions {
		doctor, err := uc.doctorRep.GetById(outpatientSession.DoctorId)
		if err != nil {
			return res, err
		}

		user, err := uc.userRep.GetById(doctor.UserId)
		if err != nil {
			return res, err
		}

		speciality, err := uc.specRep.GetById(doctor.SpecialityId)
		if err != nil {
			return res, err
		}

		patient, err := uc.patientRep.GetById(outpatientSession.PatientId)
		if err != nil {
			return res, err
		}

		dateString := strconv.Itoa(outpatientSession.Schedule.Year()) + "-" + strconv.Itoa(int(outpatientSession.Schedule.Month())) + "-" + fmt.Sprintf("%02d", outpatientSession.Schedule.Day())
		timeString := fmt.Sprintf("%02d", outpatientSession.Schedule.Hour()) + ":" + fmt.Sprintf("%02d", outpatientSession.Schedule.Minute())

		dateIndoString := fmt.Sprintf("%02d", outpatientSession.Schedule.Day()) + " " +
			constants.Bulan[int(outpatientSession.Schedule.Month())] + " " +
			strconv.Itoa(outpatientSession.Schedule.Year())

		res = append(res, dto.OutpatientSessionRes{
			ID:               outpatientSession.ID,
			CreatedAt:        outpatientSession.CreatedAt,
			UpdatedAt:        outpatientSession.UpdatedAt,
			DeletedAt:        outpatientSession.DeletedAt,
			DoctorId:         outpatientSession.DoctorId,
			PatientId:        outpatientSession.PatientId,
			Schedule:         outpatientSession.Schedule,
			Complaint:        outpatientSession.Complaint,
			IsApproved:       outpatientSession.IsApproved,
			IsFinish:         outpatientSession.IsFinish,
			FinishedAt:       outpatientSession.FinishedAt,
			ScheduleDate:     dateString,
			ScheduleDateIndo: dateIndoString,
			ScheduleTime:     timeString,
			Patient: dto.PatientRes{
				ID:            patient.ID,
				CreatedAt:     patient.CreatedAt,
				UpdatedAt:     patient.UpdatedAt,
				DeletedAt:     patient.DeletedAt,
				Nik:           patient.Nik,
				Name:          patient.Name,
				BirthDate:     patient.BirthDate,
				Gender:        patient.Gender,
				Address:       patient.Address,
				Phone:         patient.Phone,
				MaritalStatus: patient.MaritalStatus,
				ReligionID:    patient.ReligionID,
			},
			Doctor: dto.DoctorRes{
				ID:              doctor.ID,
				CreatedAt:       doctor.CreatedAt,
				UpdatedAt:       doctor.UpdatedAt,
				DeletedAt:       doctor.DeletedAt,
				Name:            user.Name,
				SpecialityId:    doctor.SpecialityId,
				LicenseNumber:   doctor.LicenseNumber,
				SpecialityName:  speciality.Name,
				DoctorSchedules: nil,
			},
		})
	}

	return res, nil
}
func (uc *outpatientSessionUseCase) GetRejectedByDoctorId(doctorId uint) ([]dto.OutpatientSessionRes, error) {
	var res []dto.OutpatientSessionRes
	outpatientSessions, err := uc.outptnRep.GetRejectedByDoctorId(doctorId)
	if err != nil {
		return res, err
	}

	for _, outpatientSession := range outpatientSessions {
		doctor, err := uc.doctorRep.GetById(outpatientSession.DoctorId)
		if err != nil {
			return res, err
		}

		user, err := uc.userRep.GetById(doctor.UserId)
		if err != nil {
			return res, err
		}

		speciality, err := uc.specRep.GetById(doctor.SpecialityId)
		if err != nil {
			return res, err
		}

		patient, err := uc.patientRep.GetById(outpatientSession.PatientId)
		if err != nil {
			return res, err
		}

		dateString := strconv.Itoa(outpatientSession.Schedule.Year()) + "-" + strconv.Itoa(int(outpatientSession.Schedule.Month())) + "-" + fmt.Sprintf("%02d", outpatientSession.Schedule.Day())
		timeString := fmt.Sprintf("%02d", outpatientSession.Schedule.Hour()) + ":" + fmt.Sprintf("%02d", outpatientSession.Schedule.Minute())

		dateIndoString := fmt.Sprintf("%02d", outpatientSession.Schedule.Day()) + " " +
			constants.Bulan[int(outpatientSession.Schedule.Month())] + " " +
			strconv.Itoa(outpatientSession.Schedule.Year())

		res = append(res, dto.OutpatientSessionRes{
			ID:               outpatientSession.ID,
			CreatedAt:        outpatientSession.CreatedAt,
			UpdatedAt:        outpatientSession.UpdatedAt,
			DeletedAt:        outpatientSession.DeletedAt,
			DoctorId:         outpatientSession.DoctorId,
			PatientId:        outpatientSession.PatientId,
			Schedule:         outpatientSession.Schedule,
			Complaint:        outpatientSession.Complaint,
			IsApproved:       outpatientSession.IsApproved,
			IsFinish:         outpatientSession.IsFinish,
			FinishedAt:       outpatientSession.FinishedAt,
			ScheduleDate:     dateString,
			ScheduleDateIndo: dateIndoString,
			ScheduleTime:     timeString,
			Patient: dto.PatientRes{
				ID:            patient.ID,
				CreatedAt:     patient.CreatedAt,
				UpdatedAt:     patient.UpdatedAt,
				DeletedAt:     patient.DeletedAt,
				Nik:           patient.Nik,
				Name:          patient.Name,
				BirthDate:     patient.BirthDate,
				Gender:        patient.Gender,
				Address:       patient.Address,
				Phone:         patient.Phone,
				MaritalStatus: patient.MaritalStatus,
				ReligionID:    patient.ReligionID,
			},
			Doctor: dto.DoctorRes{
				ID:              doctor.ID,
				CreatedAt:       doctor.CreatedAt,
				UpdatedAt:       doctor.UpdatedAt,
				DeletedAt:       doctor.DeletedAt,
				Name:            user.Name,
				SpecialityId:    doctor.SpecialityId,
				LicenseNumber:   doctor.LicenseNumber,
				SpecialityName:  speciality.Name,
				DoctorSchedules: nil,
			},
		})
	}

	return res, nil
}
func (uc *outpatientSessionUseCase) Create(payload dto.OutpatientSessionReq) (dto.OutpatientSessionRes, error) {
	var res dto.OutpatientSessionRes

	dateTimeString := payload.ScheduleDate + "T" + payload.ScheduleTime + "+07:00"
	dateTime, err := time.Parse(time.RFC3339, dateTimeString)
	if err != nil {
		return res, err
	}

	outpatient := models.OutpatientSession{
		Model:      gorm.Model{},
		DoctorId:   payload.DoctorId,
		PatientId:  payload.PatientId,
		Schedule:   dateTime,
		Complaint:  payload.Complaint,
		IsApproved: 0,
		IsFinish:   false,
		FinishedAt: time.Time{},
	}

	resUc, err := uc.outptnRep.Create(outpatient)
	if err != nil {
		return res, err
	}

	// TODO Preparing response create outpatient
	doctor, err := uc.doctorRep.GetById(resUc.DoctorId)
	if err != nil {
		return res, err
	}

	user, err := uc.userRep.GetById(doctor.UserId)
	if err != nil {
		return res, err
	}

	speciality, err := uc.specRep.GetById(doctor.SpecialityId)
	if err != nil {
		return res, err
	}

	patient, err := uc.patientRep.GetById(resUc.PatientId)
	if err != nil {
		return res, err
	}

	dateString := strconv.Itoa(resUc.Schedule.Year()) + "-" + strconv.Itoa(int(resUc.Schedule.Month())) + "-" + fmt.Sprintf("%02d", resUc.Schedule.Day())
	timeString := fmt.Sprintf("%02d", resUc.Schedule.Hour()) + ":" + fmt.Sprintf("%02d", resUc.Schedule.Minute())

	dateIndoString := fmt.Sprintf("%02d", resUc.Schedule.Day()) + " " +
		constants.Bulan[int(resUc.Schedule.Month())] + " " +
		strconv.Itoa(resUc.Schedule.Year())

	res = dto.OutpatientSessionRes{
		ID:               resUc.ID,
		CreatedAt:        resUc.CreatedAt,
		UpdatedAt:        resUc.UpdatedAt,
		DeletedAt:        resUc.DeletedAt,
		DoctorId:         resUc.DoctorId,
		PatientId:        resUc.PatientId,
		Schedule:         resUc.Schedule,
		Complaint:        resUc.Complaint,
		IsApproved:       resUc.IsApproved,
		IsFinish:         resUc.IsFinish,
		FinishedAt:       resUc.FinishedAt,
		ScheduleDate:     dateString,
		ScheduleDateIndo: dateIndoString,
		ScheduleTime:     timeString,
		Patient: dto.PatientRes{
			ID:            patient.ID,
			CreatedAt:     patient.CreatedAt,
			UpdatedAt:     patient.UpdatedAt,
			DeletedAt:     patient.DeletedAt,
			Nik:           patient.Nik,
			Name:          patient.Name,
			BirthDate:     patient.BirthDate,
			Gender:        patient.Gender,
			Address:       patient.Address,
			Phone:         patient.Phone,
			MaritalStatus: patient.MaritalStatus,
			ReligionID:    patient.ReligionID,
		},
		Doctor: dto.DoctorRes{
			ID:              doctor.ID,
			CreatedAt:       doctor.CreatedAt,
			UpdatedAt:       doctor.UpdatedAt,
			DeletedAt:       doctor.DeletedAt,
			Name:            user.Name,
			SpecialityId:    doctor.SpecialityId,
			LicenseNumber:   doctor.LicenseNumber,
			SpecialityName:  speciality.Name,
			DoctorSchedules: nil,
		},
	}

	return res, nil
}
func (uc *outpatientSessionUseCase) Update(id uint, payload dto.OutpatientSessionReq) (dto.OutpatientSessionRes, error) {
	var res dto.OutpatientSessionRes

	dateTimeString := payload.ScheduleDate + "T" + payload.ScheduleTime + "+07:00"
	dateTime, err := time.Parse(time.RFC3339, dateTimeString)
	if err != nil {
		return res, err
	}

	outpatient := models.OutpatientSession{
		Model:      gorm.Model{},
		DoctorId:   payload.DoctorId,
		PatientId:  payload.PatientId,
		Schedule:   dateTime,
		Complaint:  payload.Complaint,
		IsApproved: 0,
		IsFinish:   false,
		FinishedAt: time.Time{},
	}

	resUc, err := uc.outptnRep.Update(id, outpatient)
	if err != nil {
		return res, err
	}

	// TODO Preparing response create outpatient
	doctor, err := uc.doctorRep.GetById(resUc.DoctorId)
	if err != nil {
		return res, err
	}

	user, err := uc.userRep.GetById(doctor.UserId)
	if err != nil {
		return res, err
	}

	speciality, err := uc.specRep.GetById(doctor.SpecialityId)
	if err != nil {
		return res, err
	}

	patient, err := uc.patientRep.GetById(resUc.PatientId)
	if err != nil {
		return res, err
	}

	dateString := strconv.Itoa(resUc.Schedule.Year()) + "-" + strconv.Itoa(int(resUc.Schedule.Month())) + "-" + fmt.Sprintf("%02d", resUc.Schedule.Day())
	timeString := fmt.Sprintf("%02d", resUc.Schedule.Hour()) + ":" + fmt.Sprintf("%02d", resUc.Schedule.Minute())

	dateIndoString := fmt.Sprintf("%02d", resUc.Schedule.Day()) + " " +
		constants.Bulan[int(resUc.Schedule.Month())] + " " +
		strconv.Itoa(resUc.Schedule.Year())

	res = dto.OutpatientSessionRes{
		ID:               resUc.ID,
		CreatedAt:        resUc.CreatedAt,
		UpdatedAt:        resUc.UpdatedAt,
		DeletedAt:        resUc.DeletedAt,
		DoctorId:         resUc.DoctorId,
		PatientId:        resUc.PatientId,
		Schedule:         resUc.Schedule,
		Complaint:        resUc.Complaint,
		IsApproved:       resUc.IsApproved,
		IsFinish:         resUc.IsFinish,
		FinishedAt:       resUc.FinishedAt,
		ScheduleDate:     dateString,
		ScheduleDateIndo: dateIndoString,
		ScheduleTime:     timeString,
		Patient: dto.PatientRes{
			ID:            patient.ID,
			CreatedAt:     patient.CreatedAt,
			UpdatedAt:     patient.UpdatedAt,
			DeletedAt:     patient.DeletedAt,
			Nik:           patient.Nik,
			Name:          patient.Name,
			BirthDate:     patient.BirthDate,
			Gender:        patient.Gender,
			Address:       patient.Address,
			Phone:         patient.Phone,
			MaritalStatus: patient.MaritalStatus,
			ReligionID:    patient.ReligionID,
		},
		Doctor: dto.DoctorRes{
			ID:              doctor.ID,
			CreatedAt:       doctor.CreatedAt,
			UpdatedAt:       doctor.UpdatedAt,
			DeletedAt:       doctor.DeletedAt,
			Name:            user.Name,
			SpecialityId:    doctor.SpecialityId,
			LicenseNumber:   doctor.LicenseNumber,
			SpecialityName:  speciality.Name,
			DoctorSchedules: nil,
		},
	}

	return res, nil
}
func (uc *outpatientSessionUseCase) Approval(id uint, payload dto.ApprovalReq) (dto.OutpatientSessionRes, error) {
	var res dto.OutpatientSessionRes

	outpatientSession, err := uc.outptnRep.GetById(id)
	if err != nil {
		return res, err
	}

	resUc, err := uc.outptnRep.Approval(id, payload.IsApproved)
	if err != nil {
		return res, err
	}

	// TODO Preparing response create outpatient
	doctor, err := uc.doctorRep.GetById(resUc.DoctorId)
	if err != nil {
		return res, err
	}

	user, err := uc.userRep.GetById(doctor.UserId)
	if err != nil {
		return res, err
	}

	speciality, err := uc.specRep.GetById(doctor.SpecialityId)
	if err != nil {
		return res, err
	}

	patient, err := uc.patientRep.GetById(resUc.PatientId)
	if err != nil {
		return res, err
	}

	dateString := strconv.Itoa(resUc.Schedule.Year()) + "-" + strconv.Itoa(int(resUc.Schedule.Month())) + "-" + fmt.Sprintf("%02d", resUc.Schedule.Day())
	timeString := fmt.Sprintf("%02d", resUc.Schedule.Hour()) + ":" + fmt.Sprintf("%02d", resUc.Schedule.Minute())

	dateIndoString := fmt.Sprintf("%02d", resUc.Schedule.Day()) + " " +
		constants.Bulan[int(resUc.Schedule.Month())] + " " +
		strconv.Itoa(resUc.Schedule.Year())

	outpatientSession, err = uc.outptnRep.GetById(id)
	if err != nil {
		return res, err
	}

	res = dto.OutpatientSessionRes{
		ID:               outpatientSession.ID,
		CreatedAt:        outpatientSession.CreatedAt,
		UpdatedAt:        outpatientSession.UpdatedAt,
		DeletedAt:        outpatientSession.DeletedAt,
		DoctorId:         resUc.DoctorId,
		PatientId:        resUc.PatientId,
		Schedule:         resUc.Schedule,
		Complaint:        resUc.Complaint,
		IsApproved:       resUc.IsApproved,
		IsFinish:         resUc.IsFinish,
		FinishedAt:       resUc.FinishedAt,
		ScheduleDate:     dateString,
		ScheduleDateIndo: dateIndoString,
		ScheduleTime:     timeString,
		Patient: dto.PatientRes{
			ID:            patient.ID,
			CreatedAt:     patient.CreatedAt,
			UpdatedAt:     patient.UpdatedAt,
			DeletedAt:     patient.DeletedAt,
			Nik:           patient.Nik,
			Name:          patient.Name,
			BirthDate:     patient.BirthDate,
			Gender:        patient.Gender,
			Address:       patient.Address,
			Phone:         patient.Phone,
			MaritalStatus: patient.MaritalStatus,
			ReligionID:    patient.ReligionID,
		},
		Doctor: dto.DoctorRes{
			ID:              doctor.ID,
			CreatedAt:       doctor.CreatedAt,
			UpdatedAt:       doctor.UpdatedAt,
			DeletedAt:       doctor.DeletedAt,
			Name:            user.Name,
			SpecialityId:    doctor.SpecialityId,
			LicenseNumber:   doctor.LicenseNumber,
			SpecialityName:  speciality.Name,
			DoctorSchedules: nil,
		},
	}

	return res, nil
}
func (uc *outpatientSessionUseCase) Delete(id uint) error {

	err := uc.outptnRep.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
