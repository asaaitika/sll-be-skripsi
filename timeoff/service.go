package timeoff

type Service interface {
	CreateRequestTimeOff(input CreateRequestTimeOffInput, file string) (TimeOff, error)
	ListTimeOff(input SearchTimeOffInput, id int) ([]TimeOff, error)
	ListRequestTimeOff(input SearchRequestTimeOffInput) ([]TimeOff, error)
	GetTimeOffById(input GetTimeOffDetailInput) (TimeOff, error)
	UpdateRequestTimeOff(inputId GetTimeOffDetailInput, inputData UpdateStatusTimeOffInput) (TimeOff, error)
	DeleteTimeOff(input GetTimeOffDetailInput) (TimeOff, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateRequestTimeOff(input CreateRequestTimeOffInput, file string) (TimeOff, error) {
	timeoff := TimeOff{}

	timeoff.TimeoffType = input.TimeOffType
	timeoff.TimeoffSaldo = input.TimeOffSaldo
	timeoff.StartDate = input.StartDate
	timeoff.EndDate = input.EndDate
	timeoff.RequestType = input.RequestType
	timeoff.Reason = input.Reason
	timeoff.File = file
	timeoff.StatusTimeoff = "N"
	timeoff.EmployeeId = input.EmployeeId

	newTimeOff, err := s.repository.Save(timeoff)
	if err != nil {
		return newTimeOff, err
	}

	return newTimeOff, nil
}

func (s *service) ListTimeOff(input SearchTimeOffInput, id int) ([]TimeOff, error) {
	if input.Status != "" || input.Month != "" || input.Year != "" {
		timeoff, err := s.repository.SearchTimeOff(input, id)
		if err != nil {
			return timeoff, err
		}

		return timeoff, nil
	}

	timeoff, err := s.repository.FindAllByEmployeeId(id)
	if err != nil {
		return timeoff, err
	}

	return timeoff, nil
}

func (s *service) ListRequestTimeOff(input SearchRequestTimeOffInput) ([]TimeOff, error) {
	if input.Keyword != "" || input.RequestDate != "" || input.Status != "" {
		timeoff, err := s.repository.SearchRequestTimeOff(input)
		if err != nil {
			return timeoff, err
		}

		return timeoff, nil
	}

	timeoff, err := s.repository.FindAll()
	if err != nil {
		return timeoff, err
	}

	return timeoff, nil
}

func (s *service) GetTimeOffById(input GetTimeOffDetailInput) (TimeOff, error) {
	timeoffs, err := s.repository.FindById(input.Id)

	if err != nil {
		return timeoffs, err
	}

	return timeoffs, nil
}

func (s *service) UpdateRequestTimeOff(inputId GetTimeOffDetailInput, inputData UpdateStatusTimeOffInput) (TimeOff, error) {
	timeoff, err := s.repository.FindById(inputId.Id)
	if err != nil {
		return timeoff, err
	}

	timeoff.TimeoffId = inputId.Id
	timeoff.Remark = inputData.Remark
	timeoff.StatusTimeoff = inputData.StatusTimeoff

	updateTimeOff, err := s.repository.Update(timeoff)
	if err != nil {
		return updateTimeOff, err
	}

	return updateTimeOff, nil
}

func (s *service) DeleteTimeOff(input GetTimeOffDetailInput) (TimeOff, error) {
	timeoffs, err := s.repository.FindById(input.Id)

	if err != nil {
		return timeoffs, err
	}

	timeoff, err := s.repository.Delete(input.Id)
	if err != nil {
		return timeoff, err
	}

	return timeoff, nil
}
