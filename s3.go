package console

func (c UrlBuilder) S3Bucket(bucketName string) string {
	return c.baseUrl().Pathf("/s3/buckets/%s", bucketName).Query("bucketType", "general").String()
}

func (c UrlBuilder) S3Object(bucketName string, prefix string) string {
	return c.baseUrl().Pathf("/s3/object/%s", bucketName).
		Query("bucketType", "general").
		Query("prefix", prefix).
		String()
}
