package runscope

import "fmt"

// Schedule determines how often a test is executed. See https://api.blazemeter.com/api-monitoring/#schedules
type Schedule struct {
	ID            string `json:"id,omitempty"`
	EnvironmentID string `json:"environment_id,omitempty"`
	Interval      string `json:"interval,omitempty"`
	Note          string `json:"note,omitempty"`
}

// NewSchedule creates a new schedule struct
func NewSchedule() *Schedule {
	return &Schedule{}
}

// CreateSchedule creates a new test schedule. See https://api.blazemeter.com/api-monitoring/#schedule-details
func (client *Client) CreateSchedule(schedule *Schedule, bucketKey string, testID string) (*Schedule, error) {
	newResource, error := client.createResource(schedule, "schedule", schedule.Note,
		fmt.Sprintf("/buckets/%s/tests/%s/schedules", bucketKey, testID))
	if error != nil {
		return nil, error
	}

	newSchedule, error := getScheduleFromResponse(newResource.Data)
	if error != nil {
		return nil, error
	}

	return newSchedule, nil
}

// ReadSchedule list details about an existing test schedule. See https://api.blazemeter.com/api-monitoring/#schedule-details
func (client *Client) ReadSchedule(schedule *Schedule, bucketKey string, testID string) (*Schedule, error) {
	resource, error := client.readResource("schedule", schedule.ID,
		fmt.Sprintf("/buckets/%s/tests/%s/schedules/%s", bucketKey, testID, schedule.ID))
	if error != nil {
		return nil, error
	}

	readSchedule, error := getScheduleFromResponse(resource.Data)
	if error != nil {
		return nil, error
	}

	return readSchedule, nil
}

// ListSchedules list all the schedules for a given test. See https://api.blazemeter.com/api-monitoring/#test-schedule-list
func (client *Client) ListSchedules(bucketKey string, testID string) ([]*Schedule, error) {
	resource, error := client.readResource("[]schedule", testID,
		fmt.Sprintf("/buckets/%s/tests/%s/schedules", bucketKey, testID))
	if error != nil {
		return nil, error
	}

	readSchedules, error := getSchedulesFromResponse(resource.Data)
	if error != nil {
		return nil, error
	}

	return readSchedules, nil
}

// UpdateSchedule updates an existing test schedule. See https://api.blazemeter.com/api-monitoring/#modify-schedule
func (client *Client) UpdateSchedule(schedule *Schedule, bucketKey string, testID string) (*Schedule, error) {
	resource, error := client.updateResource(schedule, "schedule", schedule.ID,
		fmt.Sprintf("/buckets/%s/tests/%s/schedules/%s", bucketKey, testID, schedule.ID))
	if error != nil {
		return nil, error
	}

	readSchedule, error := getScheduleFromResponse(resource.Data)
	if error != nil {
		return nil, error
	}

	return readSchedule, nil
}

// DeleteSchedule delete an existing test schedule. See https://api.blazemeter.com/api-monitoring/#delete-test-schedule
func (client *Client) DeleteSchedule(schedule *Schedule, bucketKey string, testID string) error {
	return client.deleteResource("schedule", schedule.ID,
		fmt.Sprintf("/buckets/%s/tests/%s/schedules/%s", bucketKey, testID, schedule.ID))
}

func getScheduleFromResponse(response interface{}) (*Schedule, error) {
	schedule := new(Schedule)
	err := decode(schedule, response)
	return schedule, err
}

func getSchedulesFromResponse(response interface{}) ([]*Schedule, error) {
	var schedules []*Schedule
	err := decode(&schedules, response)
	return schedules, err
}
