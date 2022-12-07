package dashboardUseCase

import (
	"fmt"
	"hms-backend/constants"
	"hms-backend/dto"
	"hms-backend/helpers"
	"hms-backend/repositories/doctorRepository"
	"hms-backend/repositories/doctorScheduleRepository"
	"hms-backend/repositories/nurseRepository"
	"hms-backend/repositories/outpatientSessionRepository"
	"hms-backend/repositories/patientRepository"
	"hms-backend/repositories/religionRepository"
	"hms-backend/repositories/specialityRepository"
	"hms-backend/repositories/userRepository"
	"strconv"
	"time"
)

type DashboardUseCase interface {
	GetDataDashboardWeb() (dto.DashboardWeb, error)
}

type dashboardUseCase struct {
	outptnRep   outpatientSessionRepository.OutpatientSessionRepository
	userRep     userRepository.UserRepository
	doctorRep   doctorRepository.DoctorRepository
	specRep     specialityRepository.SpecialityRepository
	schedRep    doctorScheduleRepository.DoctorScheduleRepository
	nurseRep    nurseRepository.NurseRepository
	patientRep  patientRepository.PatientRepository
	religionRep religionRepository.ReligionRepository
}

func New(
	outptRep outpatientSessionRepository.OutpatientSessionRepository,
	usrRep userRepository.UserRepository,
	dctRep doctorRepository.DoctorRepository,
	spcRep specialityRepository.SpecialityRepository,
	scdRep doctorScheduleRepository.DoctorScheduleRepository,
	nrsRep nurseRepository.NurseRepository,
	ptRep patientRepository.PatientRepository,
	rlgRep religionRepository.ReligionRepository,
) *dashboardUseCase {
	return &dashboardUseCase{outptRep, usrRep, dctRep, spcRep, scdRep, nrsRep, ptRep, rlgRep}
}

func (uc *dashboardUseCase) GetDataDashboardWeb() (dto.DashboardWeb, error) {
	var res dto.DashboardWeb

	jakartaTime, _ := helpers.TimeIn(time.Now(), "Asia/Bangkok")
	today := jakartaTime.Weekday()

	countDoctors, err := uc.doctorRep.Count()
	if err != nil {
		return res, err
	}

	countNurses, err := uc.nurseRep.Count()
	if err != nil {
		return res, err
	}

	countPatients, err := uc.patientRep.Count()
	if err != nil {
		return res, err
	}

	var todayDoctors []dto.TodayDoctorRes

	todayScheds, err := uc.schedRep.GetByDay(int(today))
	if err != nil {
		return res, err
	}

	for _, todaySched := range todayScheds {
		doctor, err := uc.doctorRep.GetById(todaySched.DoctorId)
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

		todayDoctors = append(todayDoctors, dto.TodayDoctorRes{
			Name:           user.Name,
			LicenseNumber:  doctor.LicenseNumber,
			SpecialityName: speciality.Name,
			DayInt:         todaySched.Day,
			DayString:      constants.Hari[todaySched.Day],
			StartTime:      todaySched.StartTime,
			EndTime:        todaySched.EndTime,
		})
	}

	var todayOutPatientSessions []dto.OutpatientSessionDashboardRes
	outPatientSessions, err := uc.outptnRep.GetByDate(jakartaTime)
	if err != nil {
		return res, err
	}

	for _, outPatientSession := range outPatientSessions {
		patient, err := uc.patientRep.GetById(outPatientSession.PatientId)
		if err != nil {
			return res, err
		}

		religion, err := uc.religionRep.GetById(patient.ReligionID)
		if err != nil {
			return res, err
		}

		doctor, err := uc.doctorRep.GetById(outPatientSession.DoctorId)
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

		dateString := strconv.Itoa(outPatientSession.Schedule.Year()) + "-" + strconv.Itoa(int(outPatientSession.Schedule.Month())) + "-" + fmt.Sprintf("%02d", outPatientSession.Schedule.Day())
		timeString := fmt.Sprintf("%02d", outPatientSession.Schedule.Hour()) + ":" + fmt.Sprintf("%02d", outPatientSession.Schedule.Minute())

		todayOutPatientSessions = append(todayOutPatientSessions, dto.OutpatientSessionDashboardRes{
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
				Age:           helpers.Age(patient.BirthDate, jakartaTime),
				Phone:         patient.Phone,
				Address:       patient.Address,
				MaritalStatus: patient.MaritalStatus,
				ReligionName:  religion.Name,
			},
			Doctor: struct {
				Name           string `json:"name"`
				LicenseNumber  string `json:"license_number" form:"license_number"`
				SpecialityName string `json:"speciality_name" form:"speciality_name"`
			}{
				user.Name,
				doctor.LicenseNumber,
				speciality.Name,
			},
			Schedule:     outPatientSession.Schedule,
			Complaint:    outPatientSession.Complaint,
			IsApproved:   outPatientSession.IsApproved,
			IsFinish:     outPatientSession.IsFinish,
			FinishedAt:   outPatientSession.FinishedAt,
			ScheduleDate: dateString,
			ScheduleTime: timeString,
		})
	}

	var patients []dto.OutpatientSessionDashboardRes
	outPatientSessionDescs, err := uc.outptnRep.GetDesc(10)
	if err != nil {
		return res, err
	}

	for _, outPatientSessionDesc := range outPatientSessionDescs {
		patient, err := uc.patientRep.GetById(outPatientSessionDesc.PatientId)
		if err != nil {
			return res, err
		}

		religion, err := uc.religionRep.GetById(patient.ReligionID)
		if err != nil {
			return res, err
		}

		doctor, err := uc.doctorRep.GetById(outPatientSessionDesc.DoctorId)
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

		dateString := strconv.Itoa(outPatientSessionDesc.Schedule.Year()) + "-" +
			strconv.Itoa(int(outPatientSessionDesc.Schedule.Month())) + "-" +
			fmt.Sprintf("%02d", outPatientSessionDesc.Schedule.Day())

		timeString := fmt.Sprintf("%02d", outPatientSessionDesc.Schedule.Hour()) + ":" +
			fmt.Sprintf("%02d", outPatientSessionDesc.Schedule.Minute())

		patients = append(patients, dto.OutpatientSessionDashboardRes{
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
				Age:           helpers.Age(patient.BirthDate, jakartaTime),
				Phone:         patient.Phone,
				Address:       patient.Address,
				MaritalStatus: patient.MaritalStatus,
				ReligionName:  religion.Name,
			},
			Doctor: struct {
				Name           string `json:"name"`
				LicenseNumber  string `json:"license_number" form:"license_number"`
				SpecialityName string `json:"speciality_name" form:"speciality_name"`
			}{
				user.Name,
				doctor.LicenseNumber,
				speciality.Name,
			},
			Schedule:     outPatientSessionDesc.Schedule,
			Complaint:    outPatientSessionDesc.Complaint,
			IsApproved:   outPatientSessionDesc.IsApproved,
			IsFinish:     outPatientSessionDesc.IsFinish,
			FinishedAt:   outPatientSessionDesc.FinishedAt,
			ScheduleDate: dateString,
			ScheduleTime: timeString,
		})
	}

	res = dto.DashboardWeb{
		TotalDoctors:            countDoctors,
		TotalNurses:             countNurses,
		TotalPatients:           countPatients,
		TodayDoctors:            todayDoctors,
		TodayOutpatientSessions: todayOutPatientSessions,
		Patients:                patients,
	}

	return res, nil
}
