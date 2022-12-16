package patientConditionUseCase

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"hms-backend/constants"
	"hms-backend/dto"
	"hms-backend/helpers"
	"hms-backend/models"
	"hms-backend/repositories/doctorRepository"
	"hms-backend/repositories/doctorScheduleRepository"
	"hms-backend/repositories/historyRepository"
	"hms-backend/repositories/outpatientSessionRepository"
	"hms-backend/repositories/patientRepository"
	"hms-backend/repositories/specialityRepository"
	"hms-backend/repositories/treatmentRepository"
	"hms-backend/repositories/userRepository"
	"strconv"
	"time"
)

type PatientConditionUseCase interface {
	GetAll() ([]dto.PatientConditionRes, error)
	GetById(id uint) (dto.PatientConditionRes, error)
	GetByDoctorId(doctorId uint) ([]dto.PatientConditionRes, error)
	GetByPatientId(patientId uint) ([]dto.PatientConditionRes, error)
	Create(payload dto.InsertPatientCondition) (dto.InsertPatientConditionRes, error)
}

type patientConditionUseCase struct {
	treatmentRepo treatmentRepository.TreatmentRepository
	outptnRep     outpatientSessionRepository.OutpatientSessionRepository
	userRep       userRepository.UserRepository
	doctorRep     doctorRepository.DoctorRepository
	specRep       specialityRepository.SpecialityRepository
	schedRep      doctorScheduleRepository.DoctorScheduleRepository
	patientRep    patientRepository.PatientRepository
	historyRep    historyRepository.HistoryRepository
}

func New(
	treatmentRepo treatmentRepository.TreatmentRepository,
	outptRep outpatientSessionRepository.OutpatientSessionRepository,
	usrRep userRepository.UserRepository,
	dctRep doctorRepository.DoctorRepository,
	spcRep specialityRepository.SpecialityRepository,
	scdRep doctorScheduleRepository.DoctorScheduleRepository,
	ptRep patientRepository.PatientRepository,
	hstryRep historyRepository.HistoryRepository,
) *patientConditionUseCase {
	return &patientConditionUseCase{treatmentRepo, outptRep, usrRep, dctRep, spcRep, scdRep, ptRep, hstryRep}
}

func (uc *patientConditionUseCase) GetAll() ([]dto.PatientConditionRes, error) {
	var res []dto.PatientConditionRes

	jakartaTime, _ := helpers.TimeIn(time.Now(), "Asia/Bangkok")

	treatments, err := uc.treatmentRepo.GetAll()
	if err != nil {
		return res, err
	}

	for _, treatment := range treatments {
		outpatientSession, err := uc.outptnRep.GetById(treatment.SessionId)
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

		finishedAtIndoString := fmt.Sprintf("%02d", outpatientSession.FinishedAt.Day()) + " " +
			constants.Bulan[int(outpatientSession.FinishedAt.Month())] + " " +
			strconv.Itoa(outpatientSession.FinishedAt.Year())

		patientAge := helpers.Age(patient.BirthDate, jakartaTime)

		res = append(res, dto.PatientConditionRes{
			ID: treatment.ID,
			Patient: struct {
				NIK           string    `json:"nik"`
				Name          string    `json:"name"`
				BirthDate     time.Time `json:"birth_date"`
				Gender        int       `json:"gender"`
				Age           int       `json:"age"`
				Phone         string    `json:"phone"`
				Address       string    `json:"address"`
				MaritalStatus bool      `json:"marital_status"`
				ReligionName  string    `json:"religion_name"`
			}{
				NIK:           patient.Nik,
				Name:          patient.Name,
				BirthDate:     patient.BirthDate,
				Gender:        patient.Gender,
				Age:           patientAge,
				Phone:         patient.Phone,
				Address:       patient.Address,
				MaritalStatus: patient.MaritalStatus,
				ReligionName:  "",
			},
			Doctor: struct {
				Name           string `json:"name"`
				LicenseNumber  string `json:"license_number" form:"license_number"`
				SpecialityName string `json:"speciality_name" form:"speciality_name"`
			}{
				Name:           user.Name,
				LicenseNumber:  doctor.LicenseNumber,
				SpecialityName: speciality.Name,
			},
			Schedule:         outpatientSession.Schedule,
			ScheduleDate:     dateString,
			ScheduleDateIndo: dateIndoString,
			ScheduleTime:     timeString,
			Complaint:        outpatientSession.Complaint,
			IsApproved:       outpatientSession.IsApproved,
			Description:      treatment.Description,
			Medicine:         treatment.Medicine,
			Allergy:          treatment.Allergy,
			IsFinish:         outpatientSession.IsFinish,
			FinishedAt:       outpatientSession.FinishedAt,
			FinishedAtIndo:   finishedAtIndoString,
		})
	}

	return res, nil
}
func (uc *patientConditionUseCase) GetById(id uint) (dto.PatientConditionRes, error) {
	var res dto.PatientConditionRes

	jakartaTime, _ := helpers.TimeIn(time.Now(), "Asia/Bangkok")

	treatment, err := uc.treatmentRepo.GetById(id)
	if err != nil {
		return res, err
	}
	outpatientSession, err := uc.outptnRep.GetById(treatment.SessionId)
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

	finishedAtIndoString := fmt.Sprintf("%02d", outpatientSession.FinishedAt.Day()) + " " +
		constants.Bulan[int(outpatientSession.FinishedAt.Month())] + " " +
		strconv.Itoa(outpatientSession.FinishedAt.Year())

	patientAge := helpers.Age(patient.BirthDate, jakartaTime)

	res = dto.PatientConditionRes{
		ID: treatment.ID,
		Patient: struct {
			NIK           string    `json:"nik"`
			Name          string    `json:"name"`
			BirthDate     time.Time `json:"birth_date"`
			Gender        int       `json:"gender"`
			Age           int       `json:"age"`
			Phone         string    `json:"phone"`
			Address       string    `json:"address"`
			MaritalStatus bool      `json:"marital_status"`
			ReligionName  string    `json:"religion_name"`
		}{
			NIK:           patient.Nik,
			Name:          patient.Name,
			BirthDate:     patient.BirthDate,
			Gender:        patient.Gender,
			Age:           patientAge,
			Phone:         patient.Phone,
			Address:       patient.Address,
			MaritalStatus: patient.MaritalStatus,
			ReligionName:  finishedAtIndoString,
		},
		Doctor: struct {
			Name           string `json:"name"`
			LicenseNumber  string `json:"license_number" form:"license_number"`
			SpecialityName string `json:"speciality_name" form:"speciality_name"`
		}{
			Name:           user.Name,
			LicenseNumber:  doctor.LicenseNumber,
			SpecialityName: speciality.Name,
		},
		Schedule:         outpatientSession.Schedule,
		ScheduleDate:     dateString,
		ScheduleDateIndo: dateIndoString,
		ScheduleTime:     timeString,
		Complaint:        outpatientSession.Complaint,
		IsApproved:       outpatientSession.IsApproved,
		Description:      treatment.Description,
		Medicine:         treatment.Medicine,
		Allergy:          treatment.Allergy,
		IsFinish:         outpatientSession.IsFinish,
		FinishedAt:       outpatientSession.FinishedAt,
		FinishedAtIndo:   "",
	}

	return res, nil
}
func (uc *patientConditionUseCase) GetByDoctorId(doctorId uint) ([]dto.PatientConditionRes, error) {
	var res []dto.PatientConditionRes

	jakartaTime, _ := helpers.TimeIn(time.Now(), "Asia/Bangkok")

	outpatientSessions, err := uc.outptnRep.GetByDoctorId(doctorId)
	if err != nil {
		return res, err
	}

	for _, outpatientSession := range outpatientSessions {

		treatment, err := uc.treatmentRepo.GetByOutpatientSessionId(outpatientSession.ID)
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

		finishedAtIndoString := fmt.Sprintf("%02d", outpatientSession.FinishedAt.Day()) + " " +
			constants.Bulan[int(outpatientSession.FinishedAt.Month())] + " " +
			strconv.Itoa(outpatientSession.FinishedAt.Year())

		patientAge := helpers.Age(patient.BirthDate, jakartaTime)

		res = append(res, dto.PatientConditionRes{
			ID: treatment.ID,
			Patient: struct {
				NIK           string    `json:"nik"`
				Name          string    `json:"name"`
				BirthDate     time.Time `json:"birth_date"`
				Gender        int       `json:"gender"`
				Age           int       `json:"age"`
				Phone         string    `json:"phone"`
				Address       string    `json:"address"`
				MaritalStatus bool      `json:"marital_status"`
				ReligionName  string    `json:"religion_name"`
			}{
				NIK:           patient.Nik,
				Name:          patient.Name,
				BirthDate:     patient.BirthDate,
				Gender:        patient.Gender,
				Age:           patientAge,
				Phone:         patient.Phone,
				Address:       patient.Address,
				MaritalStatus: patient.MaritalStatus,
				ReligionName:  finishedAtIndoString,
			},
			Doctor: struct {
				Name           string `json:"name"`
				LicenseNumber  string `json:"license_number" form:"license_number"`
				SpecialityName string `json:"speciality_name" form:"speciality_name"`
			}{
				Name:           user.Name,
				LicenseNumber:  doctor.LicenseNumber,
				SpecialityName: speciality.Name,
			},
			Schedule:         outpatientSession.Schedule,
			ScheduleDate:     dateString,
			ScheduleDateIndo: dateIndoString,
			ScheduleTime:     timeString,
			Complaint:        outpatientSession.Complaint,
			IsApproved:       outpatientSession.IsApproved,
			Description:      treatment.Description,
			Medicine:         treatment.Medicine,
			Allergy:          treatment.Allergy,
			IsFinish:         outpatientSession.IsFinish,
			FinishedAt:       outpatientSession.FinishedAt,
			FinishedAtIndo:   "",
		})
	}

	return res, nil
}
func (uc *patientConditionUseCase) GetByPatientId(patientId uint) ([]dto.PatientConditionRes, error) {
	var res []dto.PatientConditionRes

	jakartaTime, _ := helpers.TimeIn(time.Now(), "Asia/Bangkok")

	outpatientSessions, err := uc.outptnRep.GetFinishedByPatientIdDesc(patientId)
	if err != nil {
		return res, err
	}

	for _, outpatientSession := range outpatientSessions {

		treatment, err := uc.treatmentRepo.GetByOutpatientSessionId(outpatientSession.ID)
		if err != nil {
			continue
		}

		doctor, err := uc.doctorRep.GetById(outpatientSession.DoctorId)
		if err != nil {
			continue
		}

		user, err := uc.userRep.GetById(doctor.UserId)
		if err != nil {
			continue
		}

		speciality, err := uc.specRep.GetById(doctor.SpecialityId)
		if err != nil {
			continue
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

		finishedAtIndoString := fmt.Sprintf("%02d", outpatientSession.FinishedAt.Day()) + " " +
			constants.Bulan[int(outpatientSession.FinishedAt.Month())] + " " +
			strconv.Itoa(outpatientSession.FinishedAt.Year())

		patientAge := helpers.Age(patient.BirthDate, jakartaTime)

		res = append(res, dto.PatientConditionRes{
			ID: treatment.ID,
			Patient: struct {
				NIK           string    `json:"nik"`
				Name          string    `json:"name"`
				BirthDate     time.Time `json:"birth_date"`
				Gender        int       `json:"gender"`
				Age           int       `json:"age"`
				Phone         string    `json:"phone"`
				Address       string    `json:"address"`
				MaritalStatus bool      `json:"marital_status"`
				ReligionName  string    `json:"religion_name"`
			}{
				NIK:           patient.Nik,
				Name:          patient.Name,
				BirthDate:     patient.BirthDate,
				Gender:        patient.Gender,
				Age:           patientAge,
				Phone:         patient.Phone,
				Address:       patient.Address,
				MaritalStatus: patient.MaritalStatus,
				ReligionName:  finishedAtIndoString,
			},
			Doctor: struct {
				Name           string `json:"name"`
				LicenseNumber  string `json:"license_number" form:"license_number"`
				SpecialityName string `json:"speciality_name" form:"speciality_name"`
			}{
				Name:           user.Name,
				LicenseNumber:  doctor.LicenseNumber,
				SpecialityName: speciality.Name,
			},
			Schedule:         outpatientSession.Schedule,
			ScheduleDate:     dateString,
			ScheduleDateIndo: dateIndoString,
			ScheduleTime:     timeString,
			Complaint:        outpatientSession.Complaint,
			IsApproved:       outpatientSession.IsApproved,
			Description:      treatment.Description,
			Medicine:         treatment.Medicine,
			Allergy:          treatment.Allergy,
			IsFinish:         outpatientSession.IsFinish,
			FinishedAt:       outpatientSession.FinishedAt,
			FinishedAtIndo:   "",
		})
	}

	return res, nil
}
func (uc *patientConditionUseCase) Create(payload dto.InsertPatientCondition) (dto.InsertPatientConditionRes, error) {
	var res dto.InsertPatientConditionRes

	jakartaTime, _ := helpers.TimeIn(time.Now(), "Asia/Bangkok")

	// TODO Check if treatment exist in this out patient session
	existedTreatment, _ := uc.treatmentRepo.GetByOutpatientSessionId(payload.OutpatientSessionId)
	if existedTreatment.ID > 0 {
		return res, errors.New("this outpatient session is finished")
	}

	outpatientSession, err := uc.outptnRep.GetById(payload.OutpatientSessionId)
	if err != nil {
		return res, err
	}

	// TODO check if this outpatient session is approved by doctor
	if outpatientSession.IsApproved != 1 {
		return res, errors.New("this outpatient session is not approved")
	}

	// TODO Preparing to insert into treatment

	treatment := models.Treatment{
		Model:       gorm.Model{},
		SessionId:   payload.OutpatientSessionId,
		Diagnose:    "",
		Description: payload.Description,
		Medicine:    payload.Medicine,
		Allergy:     payload.Allergy,
	}

	outpatientSession.IsFinish = true
	outpatientSession.FinishedAt = jakartaTime

	_, err = uc.outptnRep.Update(outpatientSession.ID, outpatientSession)
	if err != nil {
		return res, err
	}

	createTreatment, err := uc.treatmentRepo.Create(treatment)
	if err != nil {
		return res, err
	}

	// TODO insert into history

	history := models.History{
		Model:       gorm.Model{},
		DoctorId:    outpatientSession.DoctorId,
		PatientId:   outpatientSession.PatientId,
		Diagnose:    "",
		Description: payload.Description,
		Medicine:    payload.Medicine,
		Allergy:     payload.Medicine,
		FinishedAt:  outpatientSession.FinishedAt,
	}

	_, err = uc.historyRep.Create(history)
	if err != nil {
		return res, err
	}

	finishedAtIndo := fmt.Sprintf("%02d", outpatientSession.FinishedAt.Day()) + " " +
		constants.Bulan[int(outpatientSession.FinishedAt.Month())] + " " +
		strconv.Itoa(outpatientSession.FinishedAt.Year())

	res = dto.InsertPatientConditionRes{
		ID:                  createTreatment.ID,
		OutpatientSessionId: payload.OutpatientSessionId,
		Description:         payload.Description,
		Medicine:            payload.Medicine,
		Allergy:             payload.Allergy,
		IsFinish:            outpatientSession.IsFinish,
		FinishedAt:          outpatientSession.FinishedAt,
		FinishedAtIndo:      finishedAtIndo,
	}

	return res, nil
}
