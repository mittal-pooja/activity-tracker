package main

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/v_task/pb"
	"google.golang.org/grpc"
)

func main() {

	grpcServer := "localhost:8082"

	conn, err := grpc.Dial(grpcServer, grpc.WithInsecure())
	if err != nil {
		fmt.Println("error connecting grpc: ", err)
	}
	defer conn.Close()

	client := pb.NewActivityServiceClient(conn)

	//GetUserDetailsByName
	name := &pb.Name{

		Name: "pooja",
	}
	userDetails, err := client.GetUserByName(context.Background(), name)

	if err != nil {

		fmt.Println("Error in calling users using grpc client", err.Error())

	}
	fmt.Println("user details are:", userDetails)

	//Update UserDetails
	details := &pb.UpdateUser{
		User: &pb.User{
			Name:  "Pooja",
			Phone: "1234567894",
			Email: "12345689@example.com",
		},
	}

	updateUser, err := client.UpdateUserDetails(context.Background(), details)

	if err != nil {

		fmt.Println("Error in updating  users using grpc client", err.Error())

	}
	fmt.Println(updateUser)

	//Add User
	addUserDetails := &pb.User{

		Name:  "Pranjal",
		Phone: "1456123789",
		Email: "agarwalpranjal@yahoo.com",
	}

	addUser, err := client.RegisterUser(context.Background(), addUserDetails)
	if err != nil {

		fmt.Println("Error in updating  users using grpc client", err.Error())

	}
	fmt.Println(addUser)

	//Add User's Activity
	//Here "Label" will always refer to the "username", because it maps the activity to its user.
	//Also, inside MySQL database the "Label" column of the "activity_table" acts as Foreign Key Constraint to "Name" (Primary Key) column of "user_table".

	startTime, _ := ptypes.TimestampProto(time.Now())

	userActivity := &pb.Activity{

		Type:      "Eat",
		Label:     "Pooja",
		Timestamp: startTime,
		Status:    "NOT DONE",
	}

	addUserActivity, err := client.AddUserActivity(context.Background(), userActivity)
	if err != nil {

		fmt.Println("Error in updating  users using grpc client", err.Error())

	}
	fmt.Println(addUserActivity)

	//Stop User's Activity
	//time.Sleep(10000 * time.Millisecond)
	stopTime, _ := ptypes.TimestampProto(time.Now())
	//fmt.Println(time.Now())
	stopActivity := &pb.Activity{
		//ptypes.DurationProto
		Type:      "Eat",
		Timestamp: stopTime,
		Label:     "Pooja",
		Status:    "DONE",
	}

	stopUserActivity, err := client.StopUserActivity(context.Background(), stopActivity)
	if err != nil {

		fmt.Println("Error in updating  users using grpc client", err.Error())

	}
	fmt.Println(stopUserActivity)

	//is Done
	activityDone := &pb.MapUserToActivity{

		Label: "Pooja",
		Type:  "Eat",
	}

	checkIsDone, err := client.ActivityIsDone(context.Background(), activityDone)
	if err != nil {

		fmt.Println("Error in updating  users using grpc client", err.Error())

	}
	fmt.Println("Activity status(isDone)", checkIsDone)

	//is Valid

	activityValid := &pb.MapUserToActivity{

		Label: "Pooja",
		Type:  "Eat",
	}

	checkIsValid, err := client.ActivityIsValid(context.Background(), activityValid)
	if err != nil {

		fmt.Println("Error in updating  users using grpc client", err.Error())

	}
	fmt.Println(checkIsValid)

}
