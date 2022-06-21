/*
 * File: /error.go                                                             *
 * Project: errors                                                             *
 * Created At: Monday, 2022/06/20 , 11:11:43                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/06/21 , 11:23:41                               *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package errors

type MyError struct {
	// Code defines the business error code.
	Code int `json:"code"`

	// Message contains the detail of this message.
	// This message is suitable to be exposed to external
	Message string `json:"message"`

	Specific string `json:"specific,omitempty"`

	Details []MyError `json:"details,omitempty"`
}

func (me MyError) Error() string {
	return me.Message
}

func ToMyError(e error) MyError {
	myError := MyError{}
	if e, ok := e.(*withCode); ok {
		myError.Code = e.code
		myError.Message = e.Error()
		myError.Specific = e.err.Error()
		// if len(myError.Message) > 0 {
		// 	myError.Message += ": "
		// }
		// myError.Message += e.Error()

		// if msg := e.err.Error(); len(msg) != 0 {
		// 	myError.Message = msg
		// } else {
		// 	myError.Message = e.Error()
		// }
		if _, ok := e.details.(*withCode); ok {
			myError.Details = parseErrorList(e.details)
		}
	}
	return myError
}

func parseErrorList(errs ...error) []MyError {
	var myErrors []MyError
	for _, e := range errs {
		myError := MyError{}

		if e, ok := e.(*withCode); ok {
			myError.Code = e.code
			myError.Message = e.Error()
			myError.Specific = e.err.Error()
			// if len(myError.Message) > 0 {
			// 	myError.Message += ": "
			// }
			// myError.Message += e.Error()

			// if msg := e.err.Error(); len(msg) != 0 {
			// 	myError.Message = msg
			// } else {
			// 	myError.Message = e.Error()
			// }
			myError.Details = parseErrorList(e.details)
			myErrors = append(myErrors, myError)
		}
	}

	return myErrors
}
