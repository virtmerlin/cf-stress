package main

import (
	"log"
	"os"
	"strconv"
)

func main() {

	//Get ENv Vars

	f_size_min, err := strconv.Atoi(os.Getenv("FILESIZEMIN"))
	if err != nil {
		log.Fatalln(err)
	}
	f_size_max, err := strconv.Atoi(os.Getenv("FILESIZEMAX"))
	if err != nil {
		log.Fatalln(err)
	}
	s3_key := os.Getenv("S3KEY")
	s3_secret := os.Getenv("S3SECRET")
	s3_bucket := os.Getenv("S3BUCKET")
	s3_content_type := "binary/octet-stream"
	//	loop_full, err := strconv.Atoi(os.Getenv("LOOPCOUNT"))
	if err != nil {
		log.Fatalln(err)
	}
	loop_gets, err := strconv.Atoi(os.Getenv("LOOPCOUNTGETS"))
	if err != nil {
		log.Fatalln(err)
	}

	// Create File for uploads
	fn_gen_file("", f_size_min, f_size_max)

	// Upload 5 x File w/ unique Object IDs

	f_uuids := fn_gen_uuid_bank()
	slice_uuids := f_uuids[0:len(f_uuids)]
	for index, element := range slice_uuids {
		log.Printf("Uploading file %v to bucket as %v", strconv.Itoa(index), element)
		fn_s3_put(s3_key, s3_secret, s3_bucket, element, "upload.dat", s3_content_type)
	}

	// Read the files Down
	for x := 0; x < loop_gets; x++ {
		for index, element := range slice_uuids {
			log.Printf("Reading file %v from bucket as %v", strconv.Itoa(index), element)
			fn_s3_get(s3_key, s3_secret, s3_bucket, element, element)
		}
	}

	// Clean up the S3 Bucket
	for index, element := range slice_uuids {
		log.Printf("Removing file %v from bucket as %v", strconv.Itoa(index), element)
		fn_s3_rm(s3_key, s3_secret, s3_bucket, element)
	}

	// Clean up container
	for index, element := range slice_uuids {
		log.Printf("Removing file %v from container as %v", strconv.Itoa(index), element)
		fn_delete_file(element)
	}

}
