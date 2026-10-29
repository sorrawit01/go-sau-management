// ทำ API

package main

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func studentRoute(app *fiber.App, db *gorm.DB) {
	// api post for create (Endpoint: /api/students)
	app.Post("/api/students", func(c fiber.Ctx) error {
		// ตัวแปรเก็บข้อมูลที่ส่งมาจาก client-user
		student := new(Student)

		//นำข้อมูลที่ส่งมาเก็บในตัวแปร พร้อมทั้งตรวจสอบ
		if error := c.Bind().Body(student); error != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid request body",
			})
		}
		// นำข้อมูลบันทึกลงในตาราง database พร้อมทั้งตรวจสอบ
		if error := createStudent(db, student); error != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Failed to create student",
			})
		}

		// กรณีบันทึกเรียบร้อยแล้ว
		return c.Status(201).JSON(fiber.Map{
			"message": "Student created successfully",
			"data":    student,
		})
	})

	// api get for read (Endpoint: /api/students)
	app.Get("/api/students", func(c fiber.Ctx) error {
		// ตัวแปรเก็บข้อมูลที่จะต้องส่งกลับไปยัง client-user
		var students []Student

		// -------------------------------
		// สร้างตัวแปรเก็บข้อมูลที่เป็นเงื่อนไขที่ส่งมาแบบ query string
		StuName := c.Query("StuName")
		StuLevel := c.Query("StuLevel")

		var error error

		if StuName != "" {
			// ดึงข้อมูลเงื่อนไข StuName
			students, error = getStudentsByName(db, StuName)
		} else if StuLevel != "" {
			// ดึงข้อมูลเงื่อนไข StuLevel
			StuLevelInt, err := strconv.Atoi(StuLevel)
			if err != nil {
				return c.Status(400).JSON(fiber.Map{
					"message": "Invalid request parameters",
				})
			}
			students, error = getStudentsByLevel(db, StuLevelInt)
		} else {
			// ดึงข้อมูลจากตาราง เก็บในตัวแปร
			students, error = getAllStudents(db)
		}
		// -------------------------------

		// ตรวจสอบว่ามีข้อมูลไหม
		if errors.Is(error, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"message": "Student not found",
			})
		}

		// ตรวจสอบว่ามีการทำงานผิดพลาดหรือไม่
		if error != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Failed to retrieve students",
			})
		}
		return c.Status(200).JSON(fiber.Map{
			"message": "Students retrieved successfully",
			"data":    students,
		})
	})

	// api get for read by id (Endpoint: /api/students/:id)
	// เงื่อนไขส่งมาแบบ path parameter ส่่งต่อท้าย Endpoint
	app.Get("/api/students/:id", func(c fiber.Ctx) error {
		// สร้างตัวแปรเก็บ PK คือ ID ที่ส่งมา
		id := c.Params("id")
		//  แปลง id เป็นตัวเลข
		idInt, error := strconv.Atoi(id)
		if error != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid request parameter",
			})
		}

		// ดึงข้อมูลเก็บในตัวแปร
		student, error := getStudentByID(db, uint(idInt))
		// ตรวจสอบว่ามีข้อมูลหรือไม่
		if errors.Is(error, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"message": "Student not found",
			})
		}
		// ตรวจสอบว่ามีอะไรผิดพลาดหรือไม่
		if error != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}
		return c.Status(200).JSON(fiber.Map{
			"message": "successfully",
			"data":    student,
		})
	})
	// api put for update (Endpoint:/api/studnet/:id)
	app.Put("/api/students/:id", func(c fiber.Ctx) error {
		// สร้างตัวแปรเก็บ PK คือ ID ที่ส่งมา
		id := c.Params("id")
		//  แปลง id เป็นตัวเลข
		idInt, error := strconv.Atoi(id)
		if error != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid request parameter",
			})
		}
		// ตัวแปรเก็บข้อมูลที่ส่งมาจาก client-user
		student := new(Student)
		// นำข้อมูลที่ส่งมาเก็บในตัวแปร พร้อมทั้งตรวจสอบ
		if error := c.Bind().Body(student); error != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid request body",
			})
		}
		// นำข้อมูลที่บันทึกแก้ไขลงตารางใน database พร้อมทั้งตรวจสอบ
		if error := updateStudent(db, idInt, student); error != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Failed to update student",
			})
		}
		// กรณีบันทึกข้อมูลสําเร็จ
		return c.Status(200).JSON(fiber.Map{
			"message": "Student updated successfully",
		})
	})
	// api delet for delete (Endpoint:/api/studnet/:id),
	app.Delete("/api/students/:id", func(c fiber.Ctx) error {
		// สร้างตัวแปรเก็บ PK คือ ID ที่ส่งมา
		id := c.Params("id")
		//  แปลง id เป็นตัวเลข
		idInt, error := strconv.Atoi(id)
		if error != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid request parameter",
			})
		}
		// นำข้อมูลที่บันทึกแก้ไขลงตารางใน database พร้อมทั้งตรวจสอบ
		if error := deleteStudent(db, idInt); error != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Failed to delete student",
			})
		}
		// กรณีบันทึกข้อมูลสําเร็จ
		return c.Status(200).JSON(fiber.Map{
			"message": "Student deleted successfully",
		})
	})
}
