/*
 * File: /error.go                                                             *
 * Project: errors                                                             *
 * Created At: Monday, 2022/06/20 , 11:11:43                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/06/20 , 15:35:13                                *
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

	Details error `json:"details,omitempty"`
}

func (me MyError) Error() string {
	return me.Message
}

func ToMyError(e error) MyError {
	myError := MyError{}

	if e, ok := e.(*ErrorWithCode); ok {
		myError.Code = e.Code
		myError.Message = e.Error()
		if _, ok := e.Details.(*ErrorWithCode); ok {
			myError.Details = ToMyError(e.Details)
		}
	}
	return myError
}
