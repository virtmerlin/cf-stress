package main

import (
	"bytes"
	"log"
	"os"
	"strconv"
	"time"
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
	loop_full, err := strconv.Atoi(os.Getenv("LOOPCOUNT"))
	if err != nil {
		log.Fatalln(err)
	}
	loop_gets, err := strconv.Atoi(os.Getenv("LOOPCOUNTGETS"))
	if err != nil {
		log.Fatalln(err)
	}

	// Create File for uploads
	fn_gen_file("", f_size_min, f_size_max)

	// gen UUIDs for the instance
	f_uuids := fn_gen_uuid_bank()
	slice_uuids := f_uuids[0:len(f_uuids)]

	for y := 0; y < loop_full; y++ {
		// Sleep to randomize Writes
		s_seconds := fn_random(15, 30)
		log.Println("Sleeping for ", strconv.Itoa(s_seconds), " Seconds ...")
		time.Sleep(time.Second * time.Duration(s_seconds))

		// Upload 5 x File w/ unique Object IDs
		for index, element := range slice_uuids {
			log.Printf("Uploading file %v to bucket as %v", strconv.Itoa(index), element)
			fn_s3_put(s3_key, s3_secret, s3_bucket, element, "upload.dat", s3_content_type)
		}

		// Read the files Down
		for x := 0; x < loop_gets; x++ {
			for index, element := range slice_uuids {
				var file_name bytes.Buffer
				file_name.WriteString(element)
				file_name.WriteString(".dat")
				log.Printf("Reading object %v from bucket as %v", strconv.Itoa(index), file_name.String())
				fn_s3_get(s3_key, s3_secret, s3_bucket, element, file_name.String())
			}
		}

		// Clean up the S3 Bucket
		for index, element := range slice_uuids {
			log.Printf("Removing object %v from bucket named %v", strconv.Itoa(index), element)
			fn_s3_rm(s3_key, s3_secret, s3_bucket, element)
		}
	}
}
