package coreenum

type CTXEnumID string

const (
	CTXEnumIDUserID    CTXEnumID = "X-USER-ID"
	CTXEnumIDUserEmail CTXEnumID = "X-USER-EMAIL"
	CTXEnumIDUserName  CTXEnumID = "X-USER-NAME"
)

func (s *CTXEnumID) IsValid() bool {
	switch *s {
	case CTXEnumIDUserID, CTXEnumIDUserEmail, CTXEnumIDUserName:
		return true
	}
	return false
}

func (s *CTXEnumID) String() string {
	return string(*s)
}
