package controller

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/DangNguyen146/bloggolangreactjs/tree/main/blogbe/database"
	"github.com/DangNguyen146/bloggolangreactjs/tree/main/blogbe/models"
	"github.com/DangNguyen146/bloggolangreactjs/tree/main/blogbe/util"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreatePost(c *fiber.Ctx) error {
	var blogPost models.Blog
	if err := c.BodyParser(&blogPost); err != nil {
		fmt.Println("Unable to parse body")
	}

	if err := database.DB.Create(&blogPost).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid payload",
			"err":     err,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Your port is live, image not store - please contact admin to backup",
	})
}

func AllPost(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 5
	offset := (page - 1) * limit
	var total int64
	var blogPosts []models.Blog

	database.DB.Preload(("User")).Offset(offset).Limit(limit).Find(&blogPosts)
	database.DB.Model(&models.Blog{}).Count(&total)

	return c.JSON(fiber.Map{
		"data": blogPosts,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})
}

func DetailPost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var blogpost models.Blog
	database.DB.Where("id=?", id).Preload("User").First(&blogpost)
	return c.JSON(fiber.Map{
		"data": blogpost,
	})
}

func UpdatePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blog := models.Blog{
		Id: uint(id),
	}

	if err := c.BodyParser(&blog); err != nil {
		fmt.Println("Unable to parse body")
	}

	database.DB.Model(&blog).Updates(blog)
	return c.JSON(fiber.Map{
		"data": blog,
	})
}

func UniquePost(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, _ := util.ParseJwt(cookie)
	var blog []models.Blog
	database.DB.Model(&blog).Where("user_id=?", id).Preload("User").Find(&blog)
	return c.JSON(blog)
}

func DeletePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	blog := models.Blog{
		Id: uint(id),
	}
	deleteQuyery := database.DB.Delete(&blog)
	if errors.Is(deleteQuyery.Error, gorm.ErrRecordNotFound) {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "No blog with that Id",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Your post is deleted",
	})
}
