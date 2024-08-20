package model

import (
	"database/sql/driver"
	"errors"
)

// Verification status.
// swagger:enum VerificationStatus
type VerificationStatus string

func (verificationStatus *VerificationStatus) Scan(value any) error {
	if value == nil {
		return nil
	}

	switch value.(int64) {
	case 1:
		*verificationStatus = VerificationStatusNotConfirmed
	case 2:
		*verificationStatus = VerificationStatusPendingApproval
	case 3:
		*verificationStatus = VerificationStatusConfirmed
	case 4:
		*verificationStatus = VerificationStatusFailed
	case 5:
		*verificationStatus = VerificationStatusRequired
	case 6:
		*verificationStatus = VerificationStatusRejected
	default:
		return errors.New("unknown verification status:")
	}

	return nil
}

func (verificationStatus VerificationStatus) Value() (driver.Value, error) {
	if verificationStatus == "" {
		return nil, nil
	}

	switch verificationStatus {
	case VerificationStatusNotConfirmed:
		return int64(1), nil
	case VerificationStatusPendingApproval:
		return int64(2), nil
	case VerificationStatusConfirmed:
		return int64(3), nil
	case VerificationStatusFailed:
		return int64(4), nil
	case VerificationStatusRequired:
		return int64(5), nil
	case VerificationStatusRejected:
		return int64(6), nil
	default:
		return nil, errors.New("unknown verification status: %v")
	}
}

var VerificationStatuses = []string{
	string(VerificationStatusNotConfirmed),
	string(VerificationStatusPendingApproval),
	string(VerificationStatusConfirmed),
	string(VerificationStatusFailed),
	string(VerificationStatusRequired),
	string(VerificationStatusRejected),
}

const (
	VerificationStatusNotConfirmed    VerificationStatus = "NotConfirmed"
	VerificationStatusPendingApproval VerificationStatus = "PendingApproval"
	VerificationStatusConfirmed       VerificationStatus = "Confirmed"
	VerificationStatusFailed          VerificationStatus = "Failed"
	VerificationStatusRequired        VerificationStatus = "Required"
	VerificationStatusRejected        VerificationStatus = "Rejected"
)
