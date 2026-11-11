package service

import (
	"context"
	"fiber-app/dto"
	"fiber-app/repo"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2/log"
)

type productService struct {
	repo repo.IProductRepo
}

func NewProductService(repo repo.IProductRepo) IProductService {
	return &productService{repo: repo}
}

func (s *productService) CreateProduct(ctx context.Context, req dto.CreateProductReq) error {
    // Save main image
    mainImagePath, err := s.saveImage(req.ImageFile)
    if err != nil {
        log.Errorf("Failed to save main image: %v", err)
        return fmt.Errorf("Failed to save main image: %w", err)
    }

    // Save gallery images
    var galleryPaths []string
    for i, file := range req.ImageFiles {
        path, err := s.saveImage(file)
        if err != nil {
            log.Errorf("Failed to save gallery image %d: %v", i+1, err)
            s.cleanupImages(mainImagePath, galleryPaths)
            return fmt.Errorf("Failed to save gallery image %d: %w", i+1, err)
        }
        galleryPaths = append(galleryPaths, path)
    }

    // Build product model
    product := dto.NewProductDto{
        Name:        req.Name,
        Description: req.Description,
        ImageUrl:    mainImagePath,
        Mark:        req.Mark,
        Price:       req.Price,
        Discount:    req.Discount,
        SalesCount:  req.SalesCount,
        Count:       req.Count,
        Images:      galleryPaths,
        Features:    req.Features,
        Tags:        req.Tags,
    }

    // Save to database
    err = s.repo.Create(ctx, product)
    if err != nil {
        log.Errorf("Failed to create product in database: %v", err)
        s.cleanupImages(mainImagePath, galleryPaths)
        return err
    }

    return nil
}

func (s *productService) UpdateProduct(ctx context.Context, id uint64,
	req dto.CreateProductReq) error {

	// Get current images
	oldMainImage, oldGalleryImages, err := s.repo.GetProductImages(ctx, id)
	if err != nil {
		log.Errorf("Failed to get product images for ID %d: %v", id, err)
		return err
	}

	// Process main image
	var newMainImage string
	if req.ImageFile != nil {
		newMainImage, err = s.saveImage(req.ImageFile)
		if err != nil {
			log.Errorf("Failed to save new main image: %v", err)
			return fmt.Errorf("Failed to save main image: %w", err)
		}
	} else {
		newMainImage = oldMainImage
	}

	// Process gallery images
	var newGalleryImages []string
	if len(req.ImageFiles) > 0 {
		for _, file := range req.ImageFiles {
			path, err := s.saveImage(file)
			if err != nil {
				log.Errorf("Failed to save gallery image: %v", err)
				s.cleanupImages(newMainImage, newGalleryImages)
				return fmt.Errorf("Failed to save gallery image: %w", err)
			}
			newGalleryImages = append(newGalleryImages, path)
		}
	} else {
		newGalleryImages = oldGalleryImages
	}

	// Update database
	product := dto.NewProductDto{
		Name:        req.Name,
		Description: req.Description,
		ImageUrl:    newMainImage,
		Mark:        req.Mark,
		Price:       req.Price,
		Discount:    req.Discount,
		SalesCount:  req.SalesCount,
		Count:       req.Count,
		Images:      newGalleryImages,
		Features:    req.Features,
		Tags:        req.Tags,
	}

	err = s.repo.Update(ctx, id, product)
	if err != nil {
		log.Errorf("Failed to update product in database: %v", err)
		s.cleanupImages(newMainImage, newGalleryImages)
		return err
	}

	// Delete old images
	if req.ImageFile != nil && oldMainImage != "" {
		s.deleteImage(oldMainImage)
	}

	if len(req.ImageFiles) > 0 {
		// Delete all old gallery images
		for _, img := range oldGalleryImages {
			if img != "" {
				s.deleteImage(img)
			}
		}
	}

	return nil
}

func (s *productService) SoftDeleteProduct(ctx context.Context, id uint64) error {
	err := s.repo.SoftDelete(ctx, id)
	if err != nil {
		log.Errorf("Failed to soft delete product ID %d: %v", id, err)
		return err
	}

	return nil
}

func (s *productService) HardDeleteProduct(ctx context.Context, id uint64) error {
	// Get images before deletion
	mainImage, galleryImages, err := s.repo.GetProductImages(ctx, id)
	if err != nil {
		log.Errorf("Failed to get product images for ID %d: %v", id, err)
		return err
	}

	// Hard delete from database
	err = s.repo.HardDelete(ctx, id)
	if err != nil {
		log.Errorf("Failed to hard delete product ID %d: %v", id, err)
		return err
	}

	// Delete all images from storage
	if mainImage != "" {
		s.deleteImage(mainImage)
	}
	for _, img := range galleryImages {
		if img != "" {
			s.deleteImage(img)
		}
	}

	return nil
}

func (s *productService) RestoreProduct(ctx context.Context, id uint64) error {
	err := s.repo.Restore(ctx, id)
	if err != nil {
		log.Errorf("Failed to restore product ID %d: %v", id, err)
		return err
	}

	return nil
}

func (s *productService) GetProductByID(ctx context.Context, id uint64) (*dto.ProductDto, error) {
	product, err := s.repo.FindByID(ctx, id)
	if err != nil {
		log.Errorf("Failed to get product ID %d: %v", id, err)
		return nil, err
	}

	return product, nil
}

func (s *productService) AddProductImages(ctx context.Context,
	id uint64, files []*multipart.FileHeader) error {

	product, err := s.repo.FindByID(ctx, id)
	if err != nil {
		log.Errorf("Product not found ID %d: %v", id, err)
		return err
	}

	// Save new images
	var newImages []string
	for i, file := range files {
		path, err := s.saveImage(file)
		if err != nil {
			log.Errorf("Failed to save image %d: %v", i+1, err)
			s.cleanupImages("", newImages)
			return fmt.Errorf("Failed to save image %d: %w", i+1, err)
		}
		newImages = append(newImages, path)
	}

	// Combine old and new gallery
	updatedGallery := append(product.Images, newImages...)

	// Update in database (only gallery)
	err = s.repo.UpdateGallery(ctx, id, updatedGallery)
	if err != nil {
		log.Errorf("Failed to update gallery for product %d: %v", id, err)
		s.cleanupImages("", newImages)
		return err
	}

	return nil
}


func (s *productService) RemoveProductImages(ctx context.Context, id uint64,
	imagePaths []string) error {

	// Get product
	product, err := s.repo.FindByID(ctx, id)
	if err != nil {
		log.Errorf("Product not found ID %d: %v", id, err)
		return err
	}

	// Create map for quick lookup
	removeMap := make(map[string]bool)
	for _, path := range imagePaths {
		removeMap[path] = true
	}

	// Filter images
	var updatedGallery []string
	for _, img := range product.Images {
		if !removeMap[img] {
			updatedGallery = append(updatedGallery, img)
		}
	}

	// Update in database
	err = s.repo.UpdateGallery(ctx, id, updatedGallery)
	if err != nil {
		log.Errorf("Failed to update gallery for product %d: %v", id, err)
		return err
	}

	// Delete images from storage
	for _, path := range imagePaths {
		if path != "" {
			s.deleteImage(path)
		}
	}

	return nil
}

// ==================== Helpers =====================

func (s *productService) saveImage(file *multipart.FileHeader) (string, error) {
	// Generate unique filename by appending nanosecond timestamp to original name
	ext := filepath.Ext(file.Filename)
	baseName := strings.TrimSuffix(file.Filename, ext)
	newFileName := fmt.Sprintf("%d_%s%s", time.Now().UnixNano(), baseName, ext)
	
	// Storage path
	savePath := filepath.Join("../storage/images", newFileName)
	
	// Open the uploaded file for reading
	src, err := file.Open()
	if err != nil {
		log.Errorf("Failed to open uploaded file %s: %v", file.Filename, err)
		return "", err
	}
	defer src.Close()

	// Create the destination file on the server filesystem
	dst, err := os.Create(savePath)
	if err != nil {
		log.Errorf("Failed to create destination file %s: %v", savePath, err)
		return "", err
	}
	defer dst.Close()

	// Copy the uploaded file content to the destination
	bytesWritten, err := dst.ReadFrom(src)
	if err != nil {
		log.Errorf("Failed to copy file content to %s: %v", savePath, err)
		// Clean up partially written file
		os.Remove(savePath)
		return "", err
	}

	log.Infof("Image saved successfully: %s (size: %d bytes, original: %s)", 
		newFileName, bytesWritten, file.Filename)

	// Return the relative path for database storage
	return "/images/" + newFileName, nil
}

func (s *productService) cleanupImages(mainPath string, galleryPaths []string) {
	paths := append([]string{mainPath}, galleryPaths...)
	for _, path := range paths {
		if path == "" {
			continue
		}
		// Convert relative URL path to physical filesystem path
		filename := filepath.Base(path)
		physicalPath := filepath.Join("../storage/images", filename)
		
		if err := os.Remove(physicalPath); err != nil && !os.IsNotExist(err) {
			log.Warnf("Failed to cleanup image: %s, error: %v", physicalPath, err)
		}
	}
}

func (s *productService) deleteImage(imagePath string) {
	if imagePath == "" {
		return
	}
	filename := filepath.Base(imagePath)
	physicalPath := filepath.Join("../storage/images", filename)

	if err := os.Remove(physicalPath); err != nil && !os.IsNotExist(err) {
		log.Warnf("Failed to delete image: %s, error: %v", physicalPath, err)
	}
}