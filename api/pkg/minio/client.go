package minio

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/DemoLiang/hridoc/api/internal/config"
)

type Client struct {
	client *minio.Client
	bucket string
	expiry time.Duration
}

func NewClient(cfg config.MinIO) (*Client, error) {
	mc, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
		Secure: cfg.UseSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("minio new client: %w", err)
	}

	ctx := context.Background()
	exists, err := mc.BucketExists(ctx, cfg.Bucket)
	if err != nil {
		return nil, fmt.Errorf("minio check bucket: %w", err)
	}
	if !exists {
		err = mc.MakeBucket(ctx, cfg.Bucket, minio.MakeBucketOptions{})
		if err != nil {
			return nil, fmt.Errorf("minio make bucket: %w", err)
		}
	}

	return &Client{
		client: mc,
		bucket: cfg.Bucket,
		expiry: time.Duration(cfg.PresignedExpiry) * time.Second,
	}, nil
}

func (c *Client) PresignedPutURL(ctx context.Context, objectName string) (string, error) {
	u, err := c.client.PresignedPutObject(ctx, c.bucket, objectName, c.expiry)
	if err != nil {
		return "", fmt.Errorf("minio presigned put: %w", err)
	}
	return u.String(), nil
}

func (c *Client) PresignedGetURL(ctx context.Context, objectName string, expiry time.Duration) (string, error) {
	u, err := c.client.PresignedGetObject(ctx, c.bucket, objectName, expiry, nil)
	if err != nil {
		return "", fmt.Errorf("minio presigned get: %w", err)
	}
	return u.String(), nil
}

func (c *Client) PutObject(ctx context.Context, objectName string, reader io.Reader, size int64, contentType string) error {
	_, err := c.client.PutObject(ctx, c.bucket, objectName, reader, size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return fmt.Errorf("minio put object: %w", err)
	}
	return nil
}

func (c *Client) GetObject(ctx context.Context, objectName string) (io.ReadCloser, error) {
	obj, err := c.client.GetObject(ctx, c.bucket, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("minio get object: %w", err)
	}
	return obj, nil
}

func (c *Client) RemoveObject(ctx context.Context, objectName string) error {
	err := c.client.RemoveObject(ctx, c.bucket, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("minio remove object: %w", err)
	}
	return nil
}

func (c *Client) ObjectURL(objectName string) string {
	return fmt.Sprintf("%s/%s/%s", c.client.EndpointURL(), c.bucket, objectName)
}
