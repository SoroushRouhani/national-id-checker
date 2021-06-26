package model

import "strconv"

var placevalue = [10]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

const residualConst = 11
const controlDigitIndex = 9

type NationalID struct {
	
	stringFormat	string
	digitFormat		[10]int
	IsValid			bool
	cityCode		string

}

func NewNationalID(NewID string) *NationalID {

	nid := &NationalID{stringFormat: NewID}
	nid.IsValid = nid.validation()
	return nid

}

func (id *NationalID) GetCity() string {
	return getCityName(id.cityCode)

}
func (id *NationalID) validation() bool {

	if len(id.stringFormat) == 10 {
		for index, val := range id.stringFormat {
			digit, _ := strconv.Atoi(string(val))

			id.digitFormat[index] = digit
		}

		ret := id.isCheckSumValid()
		if ret == false {
			return false
		} else {
			id.IsValid = true
			id.cityCode = id.stringFormat[0:3]
			return true
		}
	}

	return false
}

func (id *NationalID) isCheckSumValid() bool {

	var (

		sum				int
		residual        int
		controlDigit 	int
	)

	for index := 0; index < controlDigitIndex; index++ {

		sum = sum + id.digitFormat[index]*placevalue[index]

	}

	residual = sum % residualConst

	controlDigit = id.digitFormat[controlDigitIndex]

	if residual < 2 {

		if residual == controlDigit {
			return true
		}

	} else {

		if controlDigit == (residualConst - residual) {
			return true
		}
	}

	return false
	
}
