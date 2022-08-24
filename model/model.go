package model
import  "gorm.io/gorm"
type Employee struct {
 gorm.Model
 Employee_id  string
 Employee_name  string
 Employee_username    string
 Employee_password  string
 
}