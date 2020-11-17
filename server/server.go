package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/dannav/hhmmss"
	"github.com/golang/protobuf/ptypes"
	"github.com/v_task/db"
	"github.com/v_task/pb"
	"google.golang.org/grpc"
)

const grpcPort = ":8082"

var Dbx = db.ConnectSqlxDatabase()

type ActValid struct {
	Valid    string
	Duration string
}
type User struct {
	Name  string `db:"Name"`
	Phone string `db:"Phone"`
	Email string `db:"Email"`
}

type ActivityService struct{}

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute

	return fmt.Sprintf("%02d:%02d", h, m)
}
func main() {

	log.Println("Starting application")

	// start listening for grpc
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatal(err)
	}

	// Instanciate new gRPC server
	server := grpc.NewServer()

	//register service
	pb.RegisterActivityServiceServer(server, &ActivityService{})

	log.Println("Starting grpc server on port " + grpcPort)

	// Start the gRPC server in goroutine
	server.Serve(listen)

}

func (a *ActivityService) GetUserByName(ctx context.Context, req *pb.Name) (*pb.User, error) {

	var p User
	username := req.Name
	userDetails, err := Dbx.Queryx("SELECT * FROM user_table WHERE Name = ?", username)

	if err != nil {

		fmt.Println("Error in userDetails()", err.Error())
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovering from panic in userDetails() %v \n", r)
			}
		}()

	}

	for userDetails.Next() {
		err := userDetails.StructScan(&p)
		if err != nil {

			fmt.Println("Err in scanning userDetails()", err.Error())

			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("Recovering from panic in userDetails(): %v \n", r)
				}
			}()

		}
		//fmt.Println(p.Name, p.Phone, p.Email)

	}
	resp := &pb.User{
		Name:  p.Name,
		Phone: p.Phone,
		Email: p.Email,
	}

	return resp, err

}

func (a *ActivityService) UpdateUserDetails(ctx context.Context, req *pb.UpdateUser) (*pb.Response, error) {

	username := req.User.Name
	phone := req.User.Phone
	email := req.User.Email

	//fmt.Println(username, phone, email)

	updateUser, err := Dbx.Queryx("UPDATE user_table SET Phone = ?, Email = ? WHERE Name = ?", phone, email, username)

	if err != nil {

		fmt.Println("Error in  updateUser()", err.Error())
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovering from panic in updateUser(): %v \n", r)
			}
		}()

	}

	fmt.Println("User has been updated:", updateUser)

	resp := &pb.Response{
		Response: "The user " + username + " has been updated.",
	}
	return resp, err

}

func (a *ActivityService) RegisterUser(ctx context.Context, req *pb.User) (*pb.Response, error) {
	username := req.Name
	phone := req.Phone
	email := req.Email

	addUser, err := Dbx.Queryx("INSERT INTO user_table(Name, Phone, Email) VALUES(?,?,?)", username, phone, email)

	if err != nil {

		fmt.Println("Error in addUser(),", err.Error())
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovering from panic in addUser(): %v \n", r)
			}
		}()

	}

	fmt.Println("uer has been added:", addUser)

	resp := &pb.Response{
		Response: "The user " + username + " has been added into the db.",
	}

	return resp, err
}

func (a *ActivityService) AddUserActivity(ctx context.Context, req *pb.Activity) (*pb.Response, error) {

	activityType := req.Type
	timestamp := req.Timestamp // activity start time

	label := req.Label
	status := req.Status

	userTime, _ := ptypes.Timestamp(timestamp)
	//fmt.Println(activityType, userTime, label, status)
	//fmt.Printf("TYPE OF START TIME%T", userTime)
	var exists int
	ifActivityExists, err := Dbx.Queryx("SELECT type = ? FROM activity_table WHERE label = ?", activityType, label)
	if err != nil {

		fmt.Println("Error in ifActivityExists()", err.Error())
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovering from panic  in ifActivityExists(): %v \n", r)
			}
		}()

	}

	for ifActivityExists.Next() {

		err := ifActivityExists.Scan(&exists)
		if err != nil {

			fmt.Println("Error while scanning ifActivityExists()", err.Error())
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("Recovering from panic in ifActivityExists(): %v \n", r)
				}
			}()

		}

		//fmt.Println("exists", exists)

	}
	if exists == 0 {

		addActivity, err := Dbx.Queryx("INSERT into activity_table(Type, Label, Timestamp,Status) VALUES(?,?,?,?)", activityType, label, userTime, status)

		if err != nil {

			fmt.Println("Error in addActivity()", err.Error())
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("Recovering from panic in addActivity(): %v \n", r)
				}
			}()

		}

		fmt.Println("activity has been added", addActivity)
	} else {
		updateActivity, err := Dbx.Queryx("UPDATE activity_table SET Timestamp =?,Status=? WHERE Type =? AND Label=?", userTime, status, activityType, label)
		if err != nil {

			fmt.Println("Error in updateActivity()", err.Error())
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("Recovering from panic in updateActivity(): %v \n", r)
				}
			}()

		}
		fmt.Println("activity has been updated", updateActivity)
	}

	resp := &pb.Response{
		Response: "The activity " + activityType + " has been added into the db.",
	}

	return resp, err

}

func (a *ActivityService) StopUserActivity(ctx context.Context, req *pb.Activity) (*pb.Response, error) {

	activityType := req.Type
	stopTime := req.Timestamp
	activityStopTime, _ := ptypes.Timestamp(stopTime)
	//fmt.Printf("activityStopTime %T", activityStopTime)
	//fmt.Println("stop time %T \n", stopTime)

	label := req.Label
	status := req.Status

	//Calculate Duration and insert it
	sTime, err := Dbx.Query("SELECT Timestamp FROM activity_table WHERE Type = ? AND Label = ?", activityType, label)
	if err != nil {

		fmt.Println("Error in sTime()", err.Error())
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovering from panic in sTime(): %v \n", r)
			}
		}()

	}
	var activityStartTime string
	for sTime.Next() {

		err := sTime.Scan(&activityStartTime)

		if err != nil {

			fmt.Println("Error while scanning sTime()", err.Error())
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("Recovering from panic in sTime(): %v \n", r)
				}
			}()
			//fmt.Println(activityStartTime)

		}
	}

	layOut := "2006-01-02 15:04:05"
	startdateStamp, err := time.Parse(layOut, activityStartTime)

	if err != nil {
		fmt.Println(err)

	}

	duration := activityStopTime.Sub(startdateStamp)

	durStr := fmtDuration(duration)

	updateDuration, err := Dbx.Queryx("UPDATE activity_table SET Duration=?, Status=? WHERE Type = ? AND Label = ?", durStr, status, activityType, label)

	if err != nil {

		fmt.Println("Error in updateDuration()", err.Error())
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovering from panic in updateDuration(): %v \n", r)
			}
		}()

	}

	fmt.Println("Duration has been updated", updateDuration)

	resp := &pb.Response{
		Response: "The activity " + activityType + " of user" + label + " has been stopped.",
	}

	return resp, err
}

func (a *ActivityService) ActivityIsDone(ctx context.Context, req *pb.MapUserToActivity) (*pb.ActivityStatus, error) {

	label := req.Label
	activityType := req.Type

	isDone, err := Dbx.Queryx("SELECT Status FROM activity_table WHERE Label = ? AND Type = ? ", label, activityType)

	if err != nil {

		fmt.Println("Error in isDone()", err.Error())
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovering from panic in isDone(): %v \n", r)
			}
		}()

	}

	var stat string
	var activityDone bool
	for isDone.Next() {

		err := isDone.Scan(&stat)

		if err != nil {

			fmt.Println("Error while scanning for isDone()", err.Error())
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("Recovering from panic in isDone(): %v \n", r)
				}
			}()
		}
		//fmt.Println("stat", stat)

		if stat == "DONE" {
			activityDone = true

		} else {
			activityDone = false

		}

	}
	if activityDone {
		fmt.Println("Activity is done")
	} else {
		fmt.Println("Activity is not done")
	}
	resp := &pb.ActivityStatus{
		Status: activityDone,
	}
	return resp, err
}
func (a *ActivityService) ActivityIsValid(ctx context.Context, req *pb.MapUserToActivity) (*pb.Valid, error) {

	label := req.Label
	activityType := req.Type

	getDuration, err := Dbx.Queryx("SELECT Duration FROM activity_table WHERE Label = ? AND Type = ? ", label, activityType)

	if err != nil {

		fmt.Println("Error in getDuration()", err.Error())
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovering from panic in getDuration(): %v \n", r)
			}
		}()

	}
	var dur string

	for getDuration.Next() {

		err := getDuration.Scan(&dur)

		if err != nil {

			fmt.Println("Error in getDuration()", err.Error())
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("Recovering from panic in getDuration(): %v \n", r)
				}
			}()

		}

	}
	var val string

	durs, _ := hhmmss.Parse(dur)

	fmt.Println("durs", durs)

	if durs.Hours() >= 6.0 && durs.Hours() <= 8.0 {

		//fmt.Println("Activity is Valid")
		val = "VALID"

	} else {

		//fmt.Println("Activity is InValid")
		val = "INVALID"
	}
	updateValid, err := Dbx.Queryx("UPDATE activity_table SET Valid=? WHERE Type =? AND Label=?", val, activityType, label)
	if err != nil {

		fmt.Println("Error in updateValid()", err.Error())
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovering from panic in updateValid(): %v \n", r)
			}
		}()

	}
	fmt.Println("Validity of an activity has been updated", updateValid)

	resp := &pb.Valid{
		Valid: val,
	}

	return resp, err
}
