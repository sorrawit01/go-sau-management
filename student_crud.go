// ทำงานกับ DB-CRUD

package main

import "gorm.io/gorm"

// ฟังก์ชั่นเพิ่มข้อมูลลงตาราง Student
func createStudent(db *gorm.DB, student *Student) error {
	return db.Create(student).Error
}

// ฟังก์ชั่นดึงข้อมูลทั้งหมดจากตาราง
func getAllStudents(db *gorm.DB) ([]Student, error) {
	// ส้ร้างตัวแปรเก็บข้อมูลที่ได้จากดึง
	var students []Student
	result := db.Find(&students)
	return students, result.Error
}

func getStudentByID(db *gorm.DB, id uint) (*Student, error) {
	// สร้างตัวแปรเก็บข้อมูลที่ได้จากการดึง
	var student Student

	result := db.First(&student, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return &student, nil
}

// ฟังก์ชั่นดึงข้อมูลจากตาราง เงื่อนไขไม่ใช่ PK เช่น ชื่อ ตรงกับที่ส่งมาบางส่วน
func getStudentsByName(db *gorm.DB, StuName string) ([]Student, error) {
	// สร้างตัวแปรเก็บข้อมูลที่ได้จากการดึง
	var students []Student

	result := db.Where("stu_name LIKE ?", "%"+StuName+"%").Find(&students)

	return students, result.Error
}

// ฟังก์ชั่นดึงข้อมูลจากตาราง เงื่อนไขไม่ใช่ PK เช่น ชั้นปีตรงกับที่ส่งมา
func getStudentsByLevel(db *gorm.DB, StuLevel int) ([]Student, error) {
	// สร้างตัวแปรเก็บข้อมูลที่ได้จากการดึง
	var students []Student

	result := db.Where("stu_level = ?", StuLevel).Find(&students)

	return students, result.Error
}

// ฟังก์ชั่น แก้ไขข้อมูลในตาราง เงื่อนไข PK คือ ID
func updateStudent(db *gorm.DB, id int, student *Student) error {
	return db.Model(&student).Where("id = ?", id).Updates(student).Error
}

// ฟังก์ชั่น ลบข้อมูลในตาราง เงื่อนไข PK คือ ID
func deleteStudent(db *gorm.DB, id int) error {
	return db.Delete(&Student{}, id).Error
}
