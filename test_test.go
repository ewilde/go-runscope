package runscope

import "testing"

func TestCreateTest(t *testing.T) {
	testPreCheck(t)
	client := clientConfigure()
	key, err := client.CreateBucket(Bucket{Name: "test", Team: Team{Id: teamId}})
	defer client.DeleteBucket(key)

	if err != nil {
		t.Error(err)
	}

	test := Test{Name: "tf_test", Description: "This is a tf test", BucketId: key}
	test, err = client.CreateTest(test)
	defer client.DeleteTest(test)

	if err != nil {
		t.Error(err)
	}

	if len(test.Id) == 0 {
		t.Error("Test id should not be empty")
	}
}

func TestReadTest(t *testing.T) {
	testPreCheck(t)
	client := clientConfigure()
	key, err := client.CreateBucket(Bucket{Name: "test", Team: Team{Id: teamId}})
	defer client.DeleteBucket(key)

	if err != nil {
		t.Error(err)
	}

	test := Test{Name: "tf_test", Description: "This is a tf test", BucketId: key}
	test, err = client.CreateTest(test)
	defer client.DeleteTest(test)

	if err != nil {
		t.Error(err)
	}

	resource, err := client.ReadTest(test)
	if err != nil {
		t.Error(err)
	}

	if resource.Data["name"] != test.Name {
		t.Errorf("Expected name %s, actual %s", test.Name, resource.Data["name"])
	}
}

func TestUpdateTest(t *testing.T) {
	testPreCheck(t)
	client := clientConfigure()
	key, err := client.CreateBucket(Bucket{Name: "test", Team: Team{Id: teamId}})
	defer client.DeleteBucket(key)

	if err != nil {
		t.Error(err)
	}

	test := Test{Name: "tf_test", Description: "This is a tf test", BucketId: key}
	test, err = client.CreateTest(test)
	defer client.DeleteTest(test)

	if err != nil {
		t.Error(err)
	}

	test.Description = "New description"
	resource, err := client.UpdateTest(test)
	if err != nil {
		t.Error(err)
	}

	if resource.Data["description"] != test.Description {
		t.Errorf("Expected description %s, actual %s", test.Description, resource.Data["description"])
	}
}
