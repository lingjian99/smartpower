package utils

import (
	"github.com/gofrs/uuid"
)

//var v4 = uuid.Must(uuid.NewV4())

func UUID() string {
	v4, err := uuid.NewV4()
	if err != nil {
		return ""
	}

	return v4.String()
}
