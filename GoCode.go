//this is somemore work from the Documentation
package GoCode//pretty sure this is the file name

import "fmt"

//hello become an obj and it returns the message that is inside
func Hello(name string) string{//yea u see it
//returns a gretting that has "name" in it
  message:= fmt.Sprintf("Hi, %v. Welcome!", name)//this is like C %v specifacion for strings
  return message
}
//just realized I can make my own packages