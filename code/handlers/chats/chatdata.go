package chats

import (
	"app/db"
	"app/models"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"net/http"
	_ "os"
	"reflect"
	_ "reflect"
	_ "strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2"
)

func Createsss(c *gin.Context) {
	dbs := db.Dbcon
	fmt.Print("what is the task that you have to do? ")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := dbs.Collection("ttest").InsertOne(ctx, bson.D{
		{"task", "test4"},
		{"createdAt", "test"},
		{"modifiedAt", "test3"},
	})
	if err != nil {
		fmt.Println("", fmt.Errorf("createTask: task for to-do list couldn't be created: %v", err))
	}
	fmt.Println(res.InsertedID.(primitive.ObjectID).Hex())
	return
}

//chatdata record to Mongo
func Create(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	Reqdata := models.RequestId{}
	c.BindJSON(&Reqdata)
	if len(Reqdata.ReqId) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Id is Missing!", // cast it to string before showing
		})
		return
	}

	fmt.Println(reflect.TypeOf(Reqdata))
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	if len(Reqdata.ReqId) != 24 {
		fmt.Println("id", Reqdata.ReqId)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid Id", // cast it to string before showing
		})
		return
	}

	// url := "https://api.chatbot.com/archives/" + Reqdata.ReqId
	// // url := os.Getenv("CHAT_API_URL") + Reqdata.ReqId
	// req, err := http.NewRequest("GET", url, nil)
	// req.Header.Add("authorization", "Bearer XeHmiuiOERdfpYD_MIRPxhjwiOrv_oPD")
	// res, err := http.DefaultClient.Do(req)
	// defer res.Body.Close()
	// body, _ := ioutil.ReadAll(res.Body)
	// fmt.Println("res", res)
	// fmt.Println("body", string(body))
	// db := c.MustGet("db").(*mgo.Database)
	dbs := db.Dbcon

	data := `{
				"id": "5d0ac3ae96cf0ebd619fffff",
				"sessionId": "S1516088695.P2CNX1BJFY",
				"date": "2018-01-15T07:45:34.641Z",
				"metadata": {
					"activation_code":"ShiftPixy",
					"ssn":"785-852-852",
					"pp_dd1_name":"kishan",
					"pp_dd1_route":"qwe",
					"pp_dd1_account":"asd",
					"pp_dd1_check_or_save":"asdwf",
					"pp_dd1_allocation_type":"sdfd",
					"pp_dd2_name":"adw2",
					"pp_dd2_route":"sd2",
					"pp_dd2_account":"afwwfd2",
					"pp_dd2_check_or_save":"asdf2",
					"pp_dd2_allocation_type":"awd2",
					"pp_dd2_allocation":"afd2"
				},
				"lastMessage": {
					"resolvedQuery": "Hello"
				}
			}`
	_ = data

	// data := `{
	// 	"id": "5d0ac3ae96cf0ebd619f4444",
	// 	"sessionId": "S1516088695.P2CNX1BJFY",
	// 	"date": "2018-01-16T07:45:34.641Z",
	// 	"metadata": {
	// 		"activation_code":"qwerty",
	// 		"first_name":"785-852-852",
	// 		"last_name":"kishan",
	// 		"dob":"qwe",
	// 		"street_address":"asd",
	// 		"street_address_2":"asdwf",
	// 		"zip":"sdfd",
	// 		"city":"adw2",
	// 		"state":"sd2",
	// 		"ssn":"afwwfd2",
	// 		"w4_filing_status":"asdf2",
	// 		"w4_num_allowances":"awd2",
	// 		"w4_withhold_extra":"afd2",
	// 		"ca_filing_status":"sdd",
	// 		"ny_filing_status":"afwf",
	// 		"nj_filing_status":"wfQDF",
	// 		"ca_filing_status":"AQDQW",
	// 		"ca_allowances_num":"srgsef",
	// 		"state":"sge",
	// 		"ny_nyc_allowances_num":"gse",
	// 		"ny_yonkers_allowances_num":"sge",
	// 		"ny_nys_yonkers_allowances_num":"eesg",
	// 		"nj_allowances_num":"eg",
	// 		"ca_addl_withholdings_amount":"eg",
	// 		"nyc_addl_withholdings_amount":"eg",
	// 		"ny_yonkers_addl_withholdings_amount":"et",
	// 		"nj_addl_withholding_amount":"er",
	// 		"state":"eg"
	// 	},
	// 	"lastMessage": {
	// 		"resolvedQuery": "Hello"
	// 	}
	// }`
	// _=data

	// dataBytes := []byte(string(body))

	dataBytes := []byte(data)

	chat := models.Chats{}
	UnmarshalingErr := json.Unmarshal(dataBytes, &chat)

	if (models.Chats{}) == chat {

		fmt.Println("error", UnmarshalingErr)

		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid response from chatbot api or Invalid Id", // cast it to string before showing
		})
		return
	}

	if UnmarshalingErr != nil {

		fmt.Println("error", UnmarshalingErr)

		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "No Id or date found!", // cast it to string before showing
		})
		return
	}
	// colChats := &models.ChatStat{Id: chat.Id, Datetime: chat.Datetime}
	_, err := dbs.Collection(models.CollectionChats).InsertOne(ctx, bson.D{
		{"Id", chat.Id},
		{"Datetime", chat.Datetime},
	})
	// err := db.C(models.CollectionChats).Insert(colChats)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Failed to enter Chats data!", // cast it to string before showing
		})
		return
	}

	////conversion interface to map
	meta_data := chat.Metadata.(map[string]interface{})
	for key, val := range meta_data {
		// colChatDetails := &models.MetadataSet{Id: chat.Id, Key: key, Value: val}

		// errDet := db.C(models.CollectionChatDetails).Insert(colChatDetails)
		_, errDet := dbs.Collection(models.CollectionChatDetails).InsertOne(ctx, bson.D{
			{"Id", chat.Id},
			{"Key", key},
			{"Value", val},
		})
		if errDet != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "Failed to enter ChatDetails data!",
			})
			return
		}
	}
	c.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusOK,
		"message": "New record created successfully!",
	})
}

//ChatData List
func ChatDataList(c *gin.Context) {

	db := c.MustGet("db").(*mgo.Database)
	// results := make([]models.Chats, 0, 10)\
	//hellocommit
	pipeline := []bson.M{
		bson.M{
			"$group": bson.M{
				"_id":      "$_id",
				"datetime": bson.M{"$last": "$datetime"},
			},
		},
	}

	// pipeline := []bson.M{sort, group}

	result := []bson.M{}

	err := db.C(models.CollectionChats).Pipe(pipeline).All(&result)

	// resp := []bson.M{}
	// iter := pipe.Iter()
	// err := iter.All(&resp)

	// err := db.C(models.CollectionChats).Find(nil).Sort("-datetime").All(&results)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

	// Marshal provided interface into JSON structure
	mResult, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}

	_ = mResult

	c.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusOK,
		"message": "List ChatData successfully",
		"data":    result,
	})
}

//For DirectData Api
func loop(c *gin.Context, id string) []models.MetadataSet {

	db := c.MustGet("db").(*mgo.Database)

	resultsLp := make([]models.MetadataSet, 0, 100)

	errLp := db.C(models.CollectionChatDetails).Find(bson.M{"id": id}).All(&resultsLp)

	if errLp != nil {
		fmt.Println(errLp)
	}

	testData := false

	_ = testData

	fmt.Println(resultsLp)

	for _, res := range resultsLp {

		switch res.Key {
		case "activation_code":
			if res.Value == "ShiftPixy" {
				testData = true
			}
			break
		case "ssn":
			str := fmt.Sprint(res.Value)
			ssn := formatSSN(str)
			res.Value = ssn
			break
		case "pp_dd1_name":
			pp_dd1_name := res.Value
			res.Value = pp_dd1_name
			break
		case "pp_dd1_route":
			pp_dd1_route := res.Value
			res.Value = pp_dd1_route
			break
		case "pp_dd1_account":
			pp_dd1_account := res.Value
			res.Value = pp_dd1_account
			break
		case "pp_dd1_check_or_save":
			pp_dd1_check_or_save := res.Value
			res.Value = pp_dd1_check_or_save
			break
		case "pp_dd1_allocation_type":
			pp_dd1_allocation_type := res.Value
			res.Value = pp_dd1_allocation_type
		case "pp_dd2_name":
			pp_dd2_name := res.Value
			res.Value = pp_dd2_name
			break
		case "pp_dd2_route":
			pp_dd2_route := res.Value
			res.Value = pp_dd2_route
			break
		case "pp_dd2_account":
			pp_dd2_account := res.Value
			res.Value = pp_dd2_account
			break
		case "pp_dd2_check_or_save":
			pp_dd2_check_or_save := res.Value
			res.Value = pp_dd2_check_or_save
			break
		case "pp_dd2_allocation_type":
			pp_dd2_allocation_type := res.Value
			res.Value = pp_dd2_allocation_type
			break
		case "pp_dd2_allocation":
			pp_dd2_allocation := res.Value
			res.Value = pp_dd2_allocation
			break
		default:
			break
		}
	}

	if testData == true {
		return resultsLp
	} else {
		return nil
	}
}

//For DirectData Api
func loopempolyee(c *gin.Context, id string) []models.MetadataSet {

	db := c.MustGet("db").(*mgo.Database)

	resultsLp := make([]models.MetadataSet, 0, 100)

	errLp := db.C(models.CollectionChatDetails).Find(bson.M{"id": id}).All(&resultsLp)

	if errLp != nil {
		fmt.Println(errLp)
	}

	testData := false

	_ = testData

	fmt.Println(resultsLp)

	for _, res := range resultsLp {

		switch res.Key {
		case "activation_code":
			if res.Value == "ShiftPixy" {
				testData = true
			}
			break
		case "first_name":
			first_name := res.Value
			res.Value = first_name
			break
		case "last_name":
			last_name := res.Value
			res.Value = last_name
			break
		case "dob":
			dob := res.Value
			res.Value = dob
			break
		case "street_address":
			street_address := res.Value
			res.Value = street_address
			break
		case "street_address_2":
			str := fmt.Sprint(res.Value)
			street_address_2 := cleanData(str)
			res.Value = street_address_2
			break
		case "zip":
			zip := res.Value
			res.Value = zip
			break
		case "city":
			city := res.Value
			res.Value = city
			break
		case "state":
			state := res.Value
			res.Value = state
			break
		case "ssn":
			str := fmt.Sprint(res.Value)
			ssn := formatSSN(str)
			res.Value = ssn
			break
		case "w4_filing_status":
			str := fmt.Sprint(res.Value)
			w4_filing_statuser := setFilingStatus(str)
			res.Value = w4_filing_statuser
			break
		case "w4_num_allowances":
			str := fmt.Sprint(res.Value)
			w4_num_allowances := cleanData(str)
			res.Value = w4_num_allowances
			break
		case "w4_withhold_extra":
			str := fmt.Sprint(res.Value)
			w4_withhold_extra := cleanData(str)
			res.Value = w4_withhold_extra
			break
		case "ca_filing_status":
			str := fmt.Sprint(res.Value)
			state_filing_status := setCaliforniaFilingStatus(str)
			res.Value = state_filing_status
			break
		case "ny_filing_status":
			str := fmt.Sprint(res.Value)
			state_filing_status := setNYFilingStatus(str)
			res.Value = state_filing_status
			break
		case "nj_filing_status":
			str := fmt.Sprint(res.Value)
			state_filing_status := setNJFilingStatus(str)
			res.Value = state_filing_status
			break
		case "ca_allowances_num":
			str := fmt.Sprint(res.Value)
			state_allowances_num := cleanData(str)
			res.Value = state_allowances_num
			break
		case "ny_nyc_allowances_num":
			str := fmt.Sprint(res.Value)
			state_allowances_num := cleanData(str)
			res.Value = state_allowances_num
			break
		case "ny_yonkers_allowances_num":
			str := fmt.Sprint(res.Value)
			state_allowances_num := cleanData(str)
			res.Value = state_allowances_num
			break
		case "ny_nys_yonkers_allowances_num":
			str := fmt.Sprint(res.Value)
			state_allowances_num := cleanData(str)
			res.Value = state_allowances_num
			break
		case "nj_allowances_num":
			str := fmt.Sprint(res.Value)
			state_allowances_num := cleanData(str)
			res.Value = state_allowances_num
			break
		case "ca_addl_withholdings_amount":
			str := fmt.Sprint(res.Value)
			addl_withholdings_amount := cleanData(str)
			res.Value = addl_withholdings_amount
			break
		case "nyc_addl_withholdings_amount":
			str := fmt.Sprint(res.Value)
			addl_withholdings_amount := cleanData(str)
			res.Value = addl_withholdings_amount
			break
		case "ny_yonkers_addl_withholdings_amount":
			str := fmt.Sprint(res.Value)
			addl_withholdings_amount := cleanData(str)
			res.Value = addl_withholdings_amount
			break
		case "nj_addl_withholding_amount":
			str := fmt.Sprint(res.Value)
			addl_withholdings_amount := cleanData(str)
			res.Value = addl_withholdings_amount
			break
		default:
			break
		}
	}

	if testData == true {
		return resultsLp
	} else {
		return nil
	}
}

//For SSN Format
func formatSSN(val string) string {
	//Strip spaces
	// $ssn = str_replace(" ", "",$ssn);
	strings.ReplaceAll(val, " ", "")

	//Strip dashes
	// $ssn = str_replace("-", "",$ssn);
	strings.ReplaceAll(val, "-", "")

	//Add dashes
	if len(val) == 9 {
		newArr := strings.Split(val, "")
		// $ssn = substr_replace($ssn, "-", 3, 0);
		// $ssn = substr_replace($ssn, "-", 6, 0);
		val = newArr[0] + newArr[1] + newArr[2] + "-" + newArr[3] + newArr[4] + newArr[5] + "-" + newArr[6] + newArr[7] + newArr[8]
	}
	return val
}

//DirectDtata Api
func GetChatDeatails(c *gin.Context) {

	type Reqdata struct {
		FromDate string `json:"fromDate"`
		ToDate   string `json:"toDate"`
	}

	var Req Reqdata

	c.BindJSON(&Req)

	fmt.Printf("easfwefawf %+v\n", Req)

	db := c.MustGet("db").(*mgo.Database)

	results := make([]models.Chats, 0, 1000)

	err := db.C(models.CollectionChats).Find(bson.M{
		"datetime": bson.M{
			"$gt": Req.FromDate,
			"$lt": Req.ToDate,
		},
	}).All(&results)

	fmt.Println(results)

	if err != nil {
		fmt.Println(err)
	}

	var count int
	count = len(results)
	resultsMS := make([][]models.MetadataSet, count)

	for ii, res := range results {
		fmt.Println("iisiis", ii)
		resultsMS[ii] = loop(c, res.Id)
	}

	fmt.Println(resultsMS)

	c.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusOK,
		"message": "List Direct-Deposite Data successfully",
		"data":    resultsMS,
	})

}

//EmployeeDtata Api
func GetEmployeeChatDeatails(c *gin.Context) {

	type Reqdata struct {
		FromDate string `json:"fromDate"`
		ToDate   string `json:"toDate"`
	}

	var Req Reqdata

	c.BindJSON(&Req)

	fmt.Printf("easfwefawf %+v\n", Req)

	db := c.MustGet("db").(*mgo.Database)

	results := make([]models.Chats, 0, 10)

	err := db.C(models.CollectionChats).Find(bson.M{
		"datetime": bson.M{
			"$gt": Req.FromDate,
			"$lt": Req.ToDate,
		},
	}).All(&results)

	fmt.Println(results)

	if err != nil {
		fmt.Println(err)
	}

	var count int
	count = len(results)
	resultsMS := make([][]models.MetadataSet, count)

	// var resultsMS []CollchatsGroup

	for ii, res := range results {
		fmt.Println("iisiis", ii)
		resultsMS[ii] = loopempolyee(c, res.Id)
	}

	fmt.Println(resultsMS)

	c.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusOK,
		"message": "List Direct-Deposite Data successfully",
		"data":    resultsMS,
	})

}

func setFilingStatus(input string) string {

	var s string

	switch input {
	case "Single":
		s = "S"
		break
	case "Married":
		s = "M"
		break
	case "Married, withhold at single rate":
		s = "S"
		break
	case "Married but withhold":
		s = "M"
		break
	default:
		s = input
		break

	}

	return s
}

func setCaliforniaFilingStatus(input string) string {

	var s string

	switch input {
	case "SINGLE or MARRIED (with two or more incomes)":
		s = "MT"
		break
	case "MARRIED (one income)":
		s = "M"
		break
	case "Single":
		s = "S"
		break
	case "Head of household":
		s = "H"
		break
	default:
		s = input
		break
	}
	return s
}

func setNYFilingStatus(input string) string {
	var s string
	switch input {
	case "Yes, I'm exempt":
		s = "N/A"
		break
	case "Single or Head of household":
		s = "S"
		break
	case "Married":
		s = "M"
		break
	case "Married, but withhold at single rate":
		s = "MH" //Change to S if there is a problem
		break
	default:
		s = input
		break
	}
	return s
}

func setNJFilingStatus(input string) string {
	var s string
	switch input {
	case "Head of Household":
		s = "H"
		break
	case "Married/Civil Union Couple Joint":
		s = "M"
		break
	case "Married/Civil Union Couple Separate":
		s = "MS"
		break
	case "Yes, I'm exempt":
		s = "NA"
		break
	case "Single":
		s = "S"
		break
	case "Qualify Widow(er)/Surviving Civil Union Partner":
		s = "W"
		break
	default:
		s = input
		break
	}
	return s
}

func cleanData(input string) string {

	return input

}
