package main

type Bucket struct {
	Seed string
}

func createBucket(seed string) Bucket {
	bucket := Bucket{seed}
	return bucket
}
