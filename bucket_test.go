package runscope

import "testing"

func TestCreateBucket(t *testing.T) {
	testPreCheck(t)
	client := clientConfigure()
	key, err := client.CreateBucket(Bucket{Name: "test", Team: Team{Id: teamId}})

	if err != nil {
		t.Error(err)
	}

	client.DeleteBucket(key)
}

func TestReadBucket(t *testing.T) {
	testPreCheck(t)
	client := clientConfigure()

	key, err := client.CreateBucket(Bucket{Name: "terraform-client.go-test", Team: Team{Id: teamId}})
	if err != nil {
		t.Error(err)
	}

	bucket, err := client.ReadBucket(key)
	if err != nil {
		t.Error(err)
	}

	if bucket.Data["key"] != key {
		t.Errorf("Bucket key expected %s was %s.", key, bucket.Data["key"])
	}

	client.DeleteBucket(key)
}