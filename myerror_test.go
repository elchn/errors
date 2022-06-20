/*
 * File: /error.go                                                             *
 * Project: errors                                                             *
 * Created At: Monday, 2022/06/20 , 11:11:43                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Monday, 2022/06/20 , 15:36:20                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package errors

import (
	"encoding/json"
	"fmt"
)

func ExampleExported() {
	err := loadConfig()
	myError := ToMyError(err)
	bytes, _ := json.MarshalIndent(myError, "", " ")
	fmt.Println(string(bytes))
	// Output:
	// {
	//  "code": 1000,
	//  "message": "ConfigurationNotValid error",
	//  "details": {
	//   "code": 1001,
	//   "message": "Data is not valid JSON",
	//   "details": {
	//    "code": 1002,
	//    "message": "End of input"
	//   }
	//  }
	// }
}
