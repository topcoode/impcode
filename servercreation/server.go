package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type AppConfig struct {
	Port       int    `json:"Port"`
	PGConnStr  string `json:"PGConnStr"`
	PGPoolSize int    `json:"PGPoolSize"`
}

type PGConn struct {
	db                  *sql.DB
	isInUse             bool
	isConnected         bool
	connectionAttempted int
}

type WebConsole struct {
	DBConnPool     []PGConn
	DBConnPoolLock sync.Mutex
}

type PIParameter struct {
	Key string
	Val string
}

type PageInfo struct {
	CurrentPage int           `json:"CurrentPage"`
	PageSize    int           `json:"PageSize"`
	SortColumn  string        `json:"SortColumn"`
	SortOrder   string        `json:"SortOrder"`
	Parameter   []PIParameter `json:"Parameter"`
}

type FPageInfo struct {
	ID   string   `json:"insId"`
	Info PageInfo `json:"pageInfo"`
}

var appConfig AppConfig
var webConsole WebConsole

func (pgConn *PGConn) Connect() {

	pgConn.isInUse = false
	db, err := sql.Open("postgres", appConfig.PGConnStr)

	err = db.Ping()

	if err != nil {
		fmt.Println("connection failed to postgres: ", err)
		pgConn.isConnected = false
		pgConn.connectionAttempted++
	} else {
		pgConn.isConnected = true
		pgConn.connectionAttempted = 0
	}

	pgConn.db = db
}

func getPGConnection() *PGConn {

	webConsole.DBConnPoolLock.Lock()
	defer webConsole.DBConnPoolLock.Unlock()

	for i := 0; i < len(webConsole.DBConnPool); i++ {
		if webConsole.DBConnPool[i].isInUse == false {
			webConsole.DBConnPool[i].isInUse = true
			return &webConsole.DBConnPool[i]
		}
	}
	return nil
}

func putPGConnection(pgConn *PGConn) {
	if pgConn != nil {
		pgConn.isInUse = false
	}
}

// ------------------------------------------------------------------
// Table Name : fsubscriber
type Fsubscriber struct {
	Subscriberid  int
	Imsi          int `json:"Imsi,string"`
	Plmn          string
	Skey          string
	Seq           int
	Opc           string
	Amf           string
	Accountstatus int `json:"Accountstatus,string"`
	Defsst        int `json:"Defsst,string"`
	Defsd         string
	Uplink        int `json:"Uplink,string"`
	Downlink      int `json:"Downlink,string"`
	Dnnid1        string
	Dnn1Name      string
	Dnnid2        string
	Dnn2Name      string
	Dnnid3        string
	Dnn3Name      string
	Msisdn        string
	Createddate   string
	Createdby     int
	Updateddate   string
	Updatedby     int
	Isactive      int
	Rowversion    int
}

type Fsubscriber_list struct {
	TotalRecords int
	List         []Fsubscriber
}

// Table Name : dnn
type Dnn struct {
	Dnnid                      int
	Dnname                     string
	Supportedpdusessiontype    int `json:"Supportedpdusessiontype,string"`
	Allowedpdusessiontype      int `json:"Allowedpdusessiontype,string"`
	Supportedsscmode1          int
	Allowedsscmode1            int
	Supportedsscmode2          int
	Allowedsscmode2            int
	Supportedsscmode3          int
	Allowedsscmode3            int
	Qosprofileprioritylevel    int `json:"Qosprofileprioritylevel,string"`
	Qosprofile5qi              int `json:"Qosprofile5qi,string"`
	Qosprofilearpprioritylevel int `json:"Qosprofilearpprioritylevel,string"`
	Qosprofilearppreemptcap    int `json:"Qosprofilearppreemptcap,string"`
	Qosprofilearppreemptvuln   int `json:"Qosprofilearppreemptvuln,string"`
	Uplink                     int `json:"Uplink,string"`
	Downlink                   int `json:"Downlink,string"`
	Plmn                       string
	Isactive                   int
}

type Dnn_list struct {
	TotalRecords int
	List         []Dnn
}

// Table Name : accountstatus
type Accountstatus struct {
	Accountstatusid   int
	Accountstatusname string
	Isactive          int
}

type Accountstatus_list struct {
	TotalRecords int
	List         []Accountstatus
}

// Table Name : appuser
type Appuser struct {
	Appuserid   int
	Appusername string
	Password    string
	Plmn        string
	Isactive    int
	Isallowed   int
	Createddate string
	Createdby   int
	Updateddate string
	Updatedby   int
	Rowversion  int
}

type Appuser_list struct {
	TotalRecords int
	List         []Appuser
}

// ------------------------------------------------------------------
var loggedUser Appuser

func GetLoggedInUser() *Appuser {
	if loggedUser.Appuserid == 0 {
		loggedUser.Appuserid = 1
	}
	return &loggedUser
}

// ------------------------------------------------------------------

func Fsubscriber_Add(dbid int, ptr *Fsubscriber) bool {
	pgConn := getPGConnection()
	defer putPGConnection(pgConn)

	if pgConn != nil {
		insertStmt := `insert into fsubscriber( imsi, plmn, skey, seq, opc, amf, accountstatus, defsst, defsd, uplink, downlink, dnnid1, dnnid2, dnnid3, msisdn, createddate, createdby, updateddate, updatedby, isactive, rowversion) values( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21) RETURNING subscriberid`
		err := pgConn.db.QueryRow(insertStmt, ptr.Imsi, ptr.Plmn, ptr.Skey, ptr.Seq, ptr.Opc, ptr.Amf, ptr.Accountstatus, ptr.Defsst, ptr.Defsd, ptr.Uplink, ptr.Downlink, ptr.Dnnid1, ptr.Dnnid2, ptr.Dnnid3, ptr.Msisdn, ptr.Createddate, ptr.Createdby, ptr.Updateddate, ptr.Updatedby, ptr.Isactive, ptr.Rowversion).Scan(&ptr.Subscriberid)
		if err != nil {
			if strings.Index(err.Error(), "duplicate") >= 0 {
				ptr.Subscriberid = -222
			}
			fmt.Println(err)
			return false
		}

		return true
	}
	return false
}

func Fsubscriber_Update(dbid int, ptr *Fsubscriber) bool {
	pgConn := getPGConnection()
	defer putPGConnection(pgConn)

	if pgConn != nil {
		updateStmt := `update fsubscriber set plmn=$1 , skey=$2 , opc=$3 , accountstatus=$4 , defsst=$5 , defsd=$6 , uplink=$7 , downlink=$8 , dnnid1=$9 , dnnid2=$10 , dnnid3=$11 , msisdn=$12 , updateddate=$13 , updatedby=$14 , isactive=$15 , rowversion=$16 where subscriberid=$17`
		_, err := pgConn.db.Exec(updateStmt, ptr.Plmn, ptr.Skey, ptr.Opc, ptr.Accountstatus, ptr.Defsst, ptr.Defsd, ptr.Uplink, ptr.Downlink, ptr.Dnnid1, ptr.Dnnid2, ptr.Dnnid3, ptr.Msisdn, ptr.Updateddate, ptr.Updatedby, ptr.Isactive, ptr.Rowversion, ptr.Subscriberid)

		if err != nil {
			fmt.Println(err)
			return false
		}

		return true
	}
	return false
}

func Fsubscriber_Get(dbid int, ptr *Fsubscriber) *Fsubscriber {
	pgConn := getPGConnection()
	defer putPGConnection(pgConn)

	if pgConn == nil {
		return nil
	}

	selectStmt := `select subscriberid, imsi, plmn, skey, seq, opc, amf, accountstatus, defsst, defsd, uplink, downlink, dnnid1, dnnid2, dnnid3, msisdn, createddate, createdby, updateddate, updatedby, isactive, rowversion from fsubscriber where subscriberid=$1 and isactive=$2`
	rows, err := pgConn.db.Query(selectStmt, ptr.Subscriberid, ptr.Isactive)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	var obj *Fsubscriber = new(Fsubscriber)

	for rows.Next() {
		err = rows.Scan(&obj.Subscriberid, &obj.Imsi, &obj.Plmn, &obj.Skey, &obj.Seq, &obj.Opc, &obj.Amf, &obj.Accountstatus, &obj.Defsst, &obj.Defsd, &obj.Uplink, &obj.Downlink, &obj.Dnnid1, &obj.Dnnid2, &obj.Dnnid3, &obj.Msisdn, &obj.Createddate, &obj.Createdby, &obj.Updateddate, &obj.Updatedby, &obj.Isactive, &obj.Rowversion)

		if err != nil {
			fmt.Println(err)
		}
	}
	rows.Close()

	return obj
}

func Fsubscriber_GetByPaging(dbid int, pageInfo FPageInfo, ptr *Fsubscriber) *Fsubscriber_list {
	pgConn := getPGConnection()
	defer putPGConnection(pgConn)

	if pgConn == nil {
		return nil
	}

	var slist *Fsubscriber_list = new(Fsubscriber_list)
	slist.TotalRecords = 0

	var selQuery string
	var err error
	var terr error
	var rows *sql.Rows
	var fQuery bool = false
	var Imsi string = ""

	if len(pageInfo.Info.Parameter) > 0 {
		if pageInfo.Info.Parameter[0].Key == "Imsi" {
			Imsi = pageInfo.Info.Parameter[0].Val
			if len(Imsi) == 15 {
				fQuery = true
			}
		}
	}

	if fQuery == false {
		selQuery = fmt.Sprintf("SELECT f.subscriberid, f.imsi, f.plmn, f.skey, f.seq, f.opc, f.amf, f.accountstatus, f.defsst, f.defsd, f.uplink, f.downlink, f.dnnid1, dnn1.dnname dnnname1, f.dnnid2, dnn2.dnname dnnname2, f.dnnid3, dnn3.dnname dnnname3, f.msisdn, f.createddate, f.createdby, f.updateddate, f.updatedby, f.isactive, f.rowversion FROM fsubscriber f INNER join dnn dnn1 ON dnn1.dnnid=f.dnnid1 INNER join dnn dnn2 ON dnn2.dnnid=f.dnnid2 INNER join dnn dnn3 ON dnn3.dnnid=f.dnnid3 where f.isactive=$1 ORDER BY subscriberid OFFSET %v LIMIT %v", ((pageInfo.Info.CurrentPage - 1) * pageInfo.Info.PageSize), pageInfo.Info.PageSize)
		rows, err = pgConn.db.Query(selQuery, ptr.Isactive)
	} else {
		//selQuery = fmt.Sprintf("SELECT f.subscriberid, f.imsi, f.plmn, f.skey, f.seq, f.opc, f.amf, f.accountstatus, f.defsst, f.defsd, f.uplink, f.downlink, f.dnnid1, dnn1.dnname dnnname1, f.dnnid2, dnn2.dnname dnnname2, f.dnnid3, dnn3.dnname dnnname3, f.msisdn, f.createddate, f.createdby, f.updateddate, f.updatedby, f.isactive, f.rowversion FROM fsubscriber f INNER join dnn dnn1 ON dnn1.dnnid=f.dnnid1 INNER join dnn dnn2 ON dnn2.dnnid=f.dnnid2 INNER join dnn dnn3 ON dnn3.dnnid=f.dnnid3 where f.isactive=$1 AND f.imsi=$2 ORDER BY subscriberid OFFSET %v LIMIT %v", ((pageInfo.Info.CurrentPage-1)*pageInfo.Info.PageSize), pageInfo.Info.PageSize)
		selQuery = fmt.Sprintf("SELECT f.subscriberid, f.imsi, f.plmn, f.skey, f.seq, f.opc, f.amf, f.accountstatus, f.defsst, f.defsd, f.uplink, f.downlink, f.dnnid1, dnn1.dnname dnnname1, f.dnnid2, dnn2.dnname dnnname2, f.dnnid3, dnn3.dnname dnnname3, f.msisdn, f.createddate, f.createdby, f.updateddate, f.updatedby, f.isactive, f.rowversion FROM fsubscriber f INNER join dnn dnn1 ON dnn1.dnnid=f.dnnid1 INNER join dnn dnn2 ON dnn2.dnnid=f.dnnid2 INNER join dnn dnn3 ON dnn3.dnnid=f.dnnid3 where f.isactive=$1 AND f.imsi=$2 ORDER BY subscriberid OFFSET %v LIMIT %v", 0, pageInfo.Info.PageSize)
		rows, err = pgConn.db.Query(selQuery, ptr.Isactive, Imsi)
	}

	if err == nil {
		for rows.Next() {
			var obj Fsubscriber
			if rerr := rows.Scan(&obj.Subscriberid, &obj.Imsi, &obj.Plmn, &obj.Skey, &obj.Seq, &obj.Opc, &obj.Amf, &obj.Accountstatus, &obj.Defsst, &obj.Defsd, &obj.Uplink, &obj.Downlink, &obj.Dnnid1, &obj.Dnn1Name, &obj.Dnnid2, &obj.Dnn2Name, &obj.Dnnid3, &obj.Dnn3Name, &obj.Msisdn, &obj.Createddate, &obj.Createdby, &obj.Updateddate, &obj.Updatedby, &obj.Isactive, &obj.Rowversion); rerr != nil {
				fmt.Println("error ", rerr)
			}
			slist.List = append(slist.List, obj)
		}
		rows.Close()

		if fQuery == false {
			terr = pgConn.db.QueryRow("SELECT count(1) id FROM fsubscriber where 1=1 and isactive=$1", ptr.Isactive).Scan(&slist.TotalRecords)
		} else {
			terr = pgConn.db.QueryRow("SELECT count(1) id FROM fsubscriber where 1=1 and isactive=$1 AND imsi=$2", ptr.Isactive, Imsi).Scan(&slist.TotalRecords)
		}

		if terr != nil {
			fmt.Println("terr ", terr)
		}
		if slist.TotalRecords == 0 {
			slist.List = []Fsubscriber{}
		}

		if slist.List == nil && slist.TotalRecords == 1 {
			slist.List = []Fsubscriber{}
			slist.TotalRecords = 0

			fmt.Println(selQuery, " - ", ptr.Isactive, " - ", Imsi)

		}

	} else {
		fmt.Println("err ", err)
	}

	return slist
}

func Fsubscriber_GetAll(dbid int, ptr *Fsubscriber) *Fsubscriber_list {
	pgConn := getPGConnection()
	defer putPGConnection(pgConn)

	if pgConn == nil {
		return nil
	}

	var slist *Fsubscriber_list = new(Fsubscriber_list)
	slist.TotalRecords = 0

	selQuery := fmt.Sprintf("SELECT subscriberid, imsi, plmn, skey, seq, opc, amf, accountstatus, defsst, defsd, uplink, downlink, dnnid1, dnnid2, dnnid3, msisdn, createddate, createdby, updateddate, updatedby, isactive, rowversion FROM fsubscriber where isactive=$1 ORDER BY subscriberid")
	rows, err := pgConn.db.Query(selQuery, ptr.Isactive)

	if err == nil {
		for rows.Next() {
			var obj Fsubscriber
			if rerr := rows.Scan(&obj.Subscriberid, &obj.Imsi, &obj.Plmn, &obj.Skey, &obj.Seq, &obj.Opc, &obj.Amf, &obj.Accountstatus, &obj.Defsst, &obj.Defsd, &obj.Uplink, &obj.Downlink, &obj.Dnnid1, &obj.Dnnid2, &obj.Dnnid3, &obj.Msisdn, &obj.Createddate, &obj.Createdby, &obj.Updateddate, &obj.Updatedby, &obj.Isactive, &obj.Rowversion); rerr != nil {
				fmt.Println("error ", rerr)
			}
			slist.List = append(slist.List, obj)
		}
		rows.Close()
		if slist.List != nil {
			slist.TotalRecords = len(slist.List)
		}
	} else {
		fmt.Println("err ", err)
	}

	return slist
}

func Fsubscriber_SelectByIMSI(dbid int, ptr *Fsubscriber) *Fsubscriber {
	pgConn := getPGConnection()
	defer putPGConnection(pgConn)

	if pgConn == nil {
		return nil
	}

	selectStmt := `select subscriberid, imsi, plmn, skey, seq, opc, amf, accountstatus, defsst, defsd, uplink, downlink, dnnid1, dnnid2, dnnid3, msisdn, createddate, createdby, updateddate, updatedby, isactive, rowversion from "fsubscriber" where imsi=$1 and isactive=$2`
	rows, err := pgConn.db.Query(selectStmt, ptr.Imsi, ptr.Isactive)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	var obj *Fsubscriber = new(Fsubscriber)

	for rows.Next() {
		err = rows.Scan(&obj.Subscriberid, &obj.Imsi, &obj.Plmn, &obj.Skey, &obj.Seq, &obj.Opc, &obj.Amf, &obj.Accountstatus, &obj.Defsst, &obj.Defsd, &obj.Uplink, &obj.Downlink, &obj.Dnnid1, &obj.Dnnid2, &obj.Dnnid3, &obj.Msisdn, &obj.Createddate, &obj.Createdby, &obj.Updateddate, &obj.Updatedby, &obj.Isactive, &obj.Rowversion)

		if err != nil {
			fmt.Println(err)
		}
	}
	rows.Close()

	return obj
}

func Fsubscriber_UpdateByIMSI(dbid int, ptr *Fsubscriber) bool {
	pgConn := getPGConnection()
	defer putPGConnection(pgConn)

	if pgConn != nil {
		updateStmt := `update fsubscriber set plmn=$1 , skey=$2 , opc=$3 , accountstatus=$4 , defsst=$5 , defsd=$6 , uplink=$7 , downlink=$8 , dnnid1=$9 , dnnid2=$10 , dnnid3=$11 , msisdn=$12 , updateddate=$13 , updatedby=$14 , rowversion=$15 where imsi=$16 and isactive=$17`
		_, err := pgConn.db.Exec(updateStmt, ptr.Plmn, ptr.Skey, ptr.Opc, ptr.Accountstatus, ptr.Defsst, ptr.Defsd, ptr.Uplink, ptr.Downlink, ptr.Dnnid1, ptr.Dnnid2, ptr.Dnnid3, ptr.Msisdn, ptr.Updateddate, ptr.Updatedby, ptr.Rowversion, ptr.Imsi, ptr.Isactive)

		if err != nil {
			fmt.Println(err)
			return false
		}

		return true
	}
	return false
}

func Dnn_Add(dbid int, ptr *Dnn) bool {
	pgConn := getPGConnection()
	defer putPGConnection(pgConn)

	if pgConn != nil {
		insertStmt := `insert into dnn( dnname, supportedpdusessiontype, allowedpdusessiontype, supportedsscmode1, allowedsscmode1, supportedsscmode2, allowedsscmode2, supportedsscmode3, allowedsscmode3, qosprofileprioritylevel, qosprofile5qi, qosprofilearpprioritylevel, qosprofilearppreemptcap, qosprofilearppreemptvuln, uplink, downlink, plmn, isactive) values( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18) RETURNING dnnid`
		err := pgConn.db.QueryRow(insertStmt, ptr.Dnname, ptr.Supportedpdusessiontype, ptr.Allowedpdusessiontype, ptr.Supportedsscmode1, ptr.Allowedsscmode1, ptr.Supportedsscmode2, ptr.Allowedsscmode2, ptr.Supportedsscmode3, ptr.Allowedsscmode3, ptr.Qosprofileprioritylevel, ptr.Qosprofile5qi, ptr.Qosprofilearpprioritylevel, ptr.Qosprofilearppreemptcap, ptr.Qosprofilearppreemptvuln, ptr.Uplink, ptr.Downlink, ptr.Plmn, ptr.Isactive).Scan(&ptr.Dnnid)
		if err != nil {
			fmt.Println(err)
			return false
		}

		return true
	}
	return false
}

func Dnn_Update(dbid int, ptr *Dnn) bool {
	pgConn := getPGConnection()
	defer putPGConnection(pgConn)

	if pgConn != nil {
		updateStmt := `update dnn set dnname=$1 , supportedpdusessiontype=$2 , allowedpdusessiontype=$3 , supportedsscmode1=$4 , allowedsscmode1=$5 , supportedsscmode2=$6 , allowedsscmode2=$7 , supportedsscmode3=$8 , allowedsscmode3=$9 , qosprofileprioritylevel=$10 , qosprofile5qi=$11 , qosprofilearpprioritylevel=$12 , qosprofilearppreemptcap=$13 , qosprofilearppreemptvuln=$14 , uplink=$15 , downlink=$16 , plmn=$17 , isactive=$18 where dnnid=$19`
		_, err := pgConn.db.Exec(updateStmt, ptr.Dnname, ptr.Supportedpdusessiontype, ptr.Allowedpdusessiontype, ptr.Supportedsscmode1, ptr.Allowedsscmode1, ptr.Supportedsscmode2, ptr.Allowedsscmode2, ptr.Supportedsscmode3, ptr.Allowedsscmode3, ptr.Qosprofileprioritylevel, ptr.Qosprofile5qi, ptr.Qosprofilearpprioritylevel, ptr.Qosprofilearppreemptcap, ptr.Qosprofilearppreemptvuln, ptr.Uplink, ptr.Downlink, ptr.Plmn, ptr.Isactive, ptr.Dnnid)

		if err != nil {
			fmt.Println(err)
			return false
		}

		return true
	}
	return false
}

func Dnn_Get(dbid int, ptr *Dnn) *Dnn {
	pgConn := getPGConnection()
	defer putPGConnection(pgConn)

	if pgConn == nil {
		return nil
	}

	selectStmt := `select dnnid, dnname, supportedpdusessiontype, allowedpdusessiontype, supportedsscmode1, allowedsscmode1, supportedsscmode2, allowedsscmode2, supportedsscmode3, allowedsscmode3, qosprofileprioritylevel, qosprofile5qi, qosprofilearpprioritylevel, qosprofilearppreemptcap, qosprofilearppreemptvuln, uplink, downlink, plmn, isactive from dnn where dnnid=$1 and isactive=$2`
	rows, err := pgConn.db.Query(selectStmt, ptr.Dnnid, ptr.Isactive)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	var obj *Dnn = new(Dnn)

	for rows.Next() {
		err = rows.Scan(&obj.Dnnid, &obj.Dnname, &obj.Supportedpdusessiontype, &obj.Allowedpdusessiontype, &obj.Supportedsscmode1, &obj.Allowedsscmode1, &obj.Supportedsscmode2, &obj.Allowedsscmode2, &obj.Supportedsscmode3, &obj.Allowedsscmode3, &obj.Qosprofileprioritylevel, &obj.Qosprofile5qi, &obj.Qosprofilearpprioritylevel, &obj.Qosprofilearppreemptcap, &obj.Qosprofilearppreemptvuln, &obj.Uplink, &obj.Downlink, &obj.Plmn, &obj.Isactive)

		if err != nil {
			fmt.Println(err)
		}
	}
	rows.Close()

	return obj
}

func Dnn_GetByPaging(dbid int, pageInfo FPageInfo, ptr *Dnn) *Dnn_list {
	pgConn := getPGConnection()
	defer putPGConnection(pgConn)

	if pgConn == nil {
		return nil
	}

	var slist *Dnn_list = new(Dnn_list)
	slist.TotalRecords = 0

	selQuery := fmt.Sprintf("SELECT dnnid, dnname, supportedpdusessiontype, allowedpdusessiontype, supportedsscmode1, allowedsscmode1, supportedsscmode2, allowedsscmode2, supportedsscmode3, allowedsscmode3, qosprofileprioritylevel, qosprofile5qi, qosprofilearpprioritylevel, qosprofilearppreemptcap, qosprofilearppreemptvuln, uplink, downlink, plmn, isactive FROM dnn where isactive IN ($1) ORDER BY dnnid OFFSET %v LIMIT %v", ((pageInfo.Info.CurrentPage - 1) * pageInfo.Info.PageSize), pageInfo.Info.PageSize)
	rows, err := pgConn.db.Query(selQuery, ptr.Isactive)

	if err == nil {
		for rows.Next() {
			var obj Dnn
			if rerr := rows.Scan(&obj.Dnnid, &obj.Dnname, &obj.Supportedpdusessiontype, &obj.Allowedpdusessiontype, &obj.Supportedsscmode1, &obj.Allowedsscmode1, &obj.Supportedsscmode2, &obj.Allowedsscmode2, &obj.Supportedsscmode3, &obj.Allowedsscmode3, &obj.Qosprofileprioritylevel, &obj.Qosprofile5qi, &obj.Qosprofilearpprioritylevel, &obj.Qosprofilearppreemptcap, &obj.Qosprofilearppreemptvuln, &obj.Uplink, &obj.Downlink, &obj.Plmn, &obj.Isactive); rerr != nil {
				fmt.Println("error ", rerr)
			}
			slist.List = append(slist.List, obj)
		}
		rows.Close()
		terr := pgConn.db.QueryRow("SELECT count(1) id FROM dnn where 1=1 and isactive IN ($1)", ptr.Isactive).Scan(&slist.TotalRecords)
		if terr != nil {
			fmt.Println("terr ", terr)
		}
		if slist.TotalRecords == 0 {
			slist.List = []Dnn{}
		}
	} else {
		fmt.Println("err ", err)
	}

	return slist
}

func Dnn_GetAll(dbid int, ptr *Dnn) *Dnn_list {
	pgConn := getPGConnection()
	defer putPGConnection(pgConn)

	if pgConn == nil {
		return nil
	}

	var slist *Dnn_list = new(Dnn_list)
	slist.TotalRecords = 0

	selQuery := fmt.Sprintf("SELECT dnnid, dnname, supportedpdusessiontype, allowedpdusessiontype, supportedsscmode1, allowedsscmode1, supportedsscmode2, allowedsscmode2, supportedsscmode3, allowedsscmode3, qosprofileprioritylevel, qosprofile5qi, qosprofilearpprioritylevel, qosprofilearppreemptcap, qosprofilearppreemptvuln, uplink, downlink, plmn, isactive FROM dnn where isactive IN ($1,2) ORDER BY dnnid")
	rows, err := pgConn.db.Query(selQuery, ptr.Isactive)

	if err == nil {
		for rows.Next() {
			var obj Dnn
			if rerr := rows.Scan(&obj.Dnnid, &obj.Dnname, &obj.Supportedpdusessiontype, &obj.Allowedpdusessiontype, &obj.Supportedsscmode1, &obj.Allowedsscmode1, &obj.Supportedsscmode2, &obj.Allowedsscmode2, &obj.Supportedsscmode3, &obj.Allowedsscmode3, &obj.Qosprofileprioritylevel, &obj.Qosprofile5qi, &obj.Qosprofilearpprioritylevel, &obj.Qosprofilearppreemptcap, &obj.Qosprofilearppreemptvuln, &obj.Uplink, &obj.Downlink, &obj.Plmn, &obj.Isactive); rerr != nil {
				fmt.Println("error ", rerr)
			}
			slist.List = append(slist.List, obj)
		}
		rows.Close()
		if slist.List != nil {
			slist.TotalRecords = len(slist.List)
		}
	} else {
		fmt.Println("err ", err)
	}

	return slist
}

func Accountstatus_Add(dbid int, ptr *Accountstatus) bool {
	pgConn := getPGConnection()
	defer putPGConnection(pgConn)

	if pgConn != nil {
		insertStmt := `insert into accountstatus( accountstatusname, isactive) values( $1, $2) RETURNING accountstatusid`
		err := pgConn.db.QueryRow(insertStmt, ptr.Accountstatusname, ptr.Isactive).Scan(&ptr.Accountstatusid)
		if err != nil {
			fmt.Println(err)
			return false
		}

		return true
	}
	return false
}

func Accountstatus_Update(dbid int, ptr *Accountstatus) bool {
	pgConn := getPGConnection()
	defer putPGConnection(pgConn)

	if pgConn != nil {
		updateStmt := `update accountstatus set accountstatusname=$1 , isactive=$2 where accountstatusid=$3`
		_, err := pgConn.db.Exec(updateStmt, ptr.Accountstatusname, ptr.Isactive, ptr.Accountstatusid)

		if err != nil {
			fmt.Println(err)
			return false
		}

		return true
	}
	return false
}

func Accountstatus_Get(dbid int, ptr *Accountstatus) *Accountstatus {
	pgConn := getPGConnection()
	defer putPGConnection(pgConn)

	if pgConn == nil {
		return nil
	}

	selectStmt := `select accountstatusid, accountstatusname, isactive from accountstatus where accountstatusid=$1 and isactive=$2`
	rows, err := pgConn.db.Query(selectStmt, ptr.Accountstatusid, ptr.Isactive)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	var obj *Accountstatus = new(Accountstatus)

	for rows.Next() {
		err = rows.Scan(&obj.Accountstatusid, &obj.Accountstatusname, &obj.Isactive)

		if err != nil {
			fmt.Println(err)
		}
	}
	rows.Close()

	return obj
}

func Accountstatus_GetByPaging(dbid int, pageInfo FPageInfo, ptr *Accountstatus) *Accountstatus_list {
	pgConn := getPGConnection()
	defer putPGConnection(pgConn)

	if pgConn == nil {
		return nil
	}

	var slist *Accountstatus_list = new(Accountstatus_list)
	slist.TotalRecords = 0

	selQuery := fmt.Sprintf("SELECT accountstatusid, accountstatusname, isactive FROM accountstatus where isactive=$1 ORDER BY accountstatusid OFFSET %v LIMIT %v", ((pageInfo.Info.CurrentPage - 1) * pageInfo.Info.PageSize), pageInfo.Info.PageSize)
	rows, err := pgConn.db.Query(selQuery, ptr.Isactive)

	if err == nil {
		for rows.Next() {
			var obj Accountstatus
			if rerr := rows.Scan(&obj.Accountstatusid, &obj.Accountstatusname, &obj.Isactive); rerr != nil {
				fmt.Println("error ", rerr)
			}
			slist.List = append(slist.List, obj)
		}
		rows.Close()
		terr := pgConn.db.QueryRow("SELECT count(1) id FROM accountstatus where 1=1 and isactive=$1", ptr.Isactive).Scan(&slist.TotalRecords)
		if terr != nil {
			fmt.Println("terr ", terr)
		}
		if slist.TotalRecords == 0 {
			slist.List = []Accountstatus{}
		}
	} else {
		fmt.Println("err ", err)
	}

	return slist
}

func Accountstatus_GetAll(dbid int, ptr *Accountstatus) *Accountstatus_list {
	pgConn := getPGConnection()
	defer putPGConnection(pgConn)

	if pgConn == nil {
		return nil
	}

	var slist *Accountstatus_list = new(Accountstatus_list)
	slist.TotalRecords = 0

	selQuery := fmt.Sprintf("SELECT accountstatusid, accountstatusname, isactive FROM accountstatus where isactive=$1 ORDER BY accountstatusid")
	rows, err := pgConn.db.Query(selQuery, ptr.Isactive)

	if err == nil {
		for rows.Next() {
			var obj Accountstatus
			if rerr := rows.Scan(&obj.Accountstatusid, &obj.Accountstatusname, &obj.Isactive); rerr != nil {
				fmt.Println("error ", rerr)
			}
			slist.List = append(slist.List, obj)
		}
		rows.Close()
		if slist.List != nil {
			slist.TotalRecords = len(slist.List)
		}
	} else {
		fmt.Println("err ", err)
	}

	return slist
}

func Appuser_Add(dbid int, ptr *Appuser) bool {
	pgConn := getPGConnection()
	defer putPGConnection(pgConn)

	if pgConn != nil {
		insertStmt := `insert into appuser( appusername, password, plmn, isactive, isallowed, createddate, createdby, updateddate, updatedby, rowversion) values( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING appuserid`
		err := pgConn.db.QueryRow(insertStmt, ptr.Appusername, ptr.Password, ptr.Plmn, ptr.Isactive, ptr.Isallowed, ptr.Createddate, ptr.Createdby, ptr.Updateddate, ptr.Updatedby, ptr.Rowversion).Scan(&ptr.Appuserid)
		if err != nil {
			fmt.Println(err)
			return false
		}

		return true
	}
	return false
}

func Appuser_Update(dbid int, ptr *Appuser) bool {
	pgConn := getPGConnection()
	defer putPGConnection(pgConn)

	if pgConn != nil {
		updateStmt := `update appuser set appusername=$1 , password=$2 , plmn=$3 , isactive=$4 , isallowed=$5 , updateddate=$6 , updatedby=$7 , rowversion=$8 where appuserid=$9`
		_, err := pgConn.db.Exec(updateStmt, ptr.Appusername, ptr.Password, ptr.Plmn, ptr.Isactive, ptr.Isallowed, ptr.Updateddate, ptr.Updatedby, ptr.Rowversion, ptr.Appuserid)

		if err != nil {
			fmt.Println(err)
			return false
		}

		return true
	}
	return false
}

func Appuser_Get(dbid int, ptr *Appuser) *Appuser {
	pgConn := getPGConnection()
	defer putPGConnection(pgConn)

	if pgConn == nil {
		return nil
	}

	selectStmt := `select appuserid, appusername, password, plmn, isactive, isallowed, createddate, createdby, updateddate, updatedby, rowversion from appuser where appuserid=$1 and isactive=$2`
	rows, err := pgConn.db.Query(selectStmt, ptr.Appuserid, ptr.Isactive)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	var obj *Appuser = new(Appuser)

	for rows.Next() {
		err = rows.Scan(&obj.Appuserid, &obj.Appusername, &obj.Password, &obj.Plmn, &obj.Isactive, &obj.Isallowed, &obj.Createddate, &obj.Createdby, &obj.Updateddate, &obj.Updatedby, &obj.Rowversion)

		if err != nil {
			fmt.Println(err)
		}
	}
	rows.Close()

	return obj
}

func Appuser_GetByPaging(dbid int, pageInfo FPageInfo, ptr *Appuser) *Appuser_list {
	pgConn := getPGConnection()
	defer putPGConnection(pgConn)

	if pgConn == nil {
		return nil
	}

	var slist *Appuser_list = new(Appuser_list)
	slist.TotalRecords = 0

	selQuery := fmt.Sprintf("SELECT appuserid, appusername, password, plmn, isactive, isallowed, createddate, createdby, updateddate, updatedby, rowversion FROM appuser where isactive=$1 ORDER BY appuserid OFFSET %v LIMIT %v", ((pageInfo.Info.CurrentPage - 1) * pageInfo.Info.PageSize), pageInfo.Info.PageSize)
	rows, err := pgConn.db.Query(selQuery, ptr.Isactive)

	if err == nil {
		for rows.Next() {
			var obj Appuser
			if rerr := rows.Scan(&obj.Appuserid, &obj.Appusername, &obj.Password, &obj.Plmn, &obj.Isactive, &obj.Isallowed, &obj.Createddate, &obj.Createdby, &obj.Updateddate, &obj.Updatedby, &obj.Rowversion); rerr != nil {
				fmt.Println("error ", rerr)
			}
			slist.List = append(slist.List, obj)
		}
		rows.Close()
		terr := pgConn.db.QueryRow("SELECT count(1) id FROM appuser where 1=1 and isactive=$1", ptr.Isactive).Scan(&slist.TotalRecords)
		if terr != nil {
			fmt.Println("terr ", terr)
		}
		if slist.TotalRecords == 0 {
			slist.List = []Appuser{}
		}
	} else {
		fmt.Println("err ", err)
	}

	return slist
}

func Appuser_GetAll(dbid int, ptr *Appuser) *Appuser_list {
	pgConn := getPGConnection()
	defer putPGConnection(pgConn)

	if pgConn == nil {
		return nil
	}

	var slist *Appuser_list = new(Appuser_list)
	slist.TotalRecords = 0

	selQuery := fmt.Sprintf("SELECT appuserid, appusername, password, plmn, isactive, isallowed, createddate, createdby, updateddate, updatedby, rowversion FROM appuser where isactive=$1 ORDER BY appuserid")
	rows, err := pgConn.db.Query(selQuery, ptr.Isactive)

	if err == nil {
		for rows.Next() {
			var obj Appuser
			if rerr := rows.Scan(&obj.Appuserid, &obj.Appusername, &obj.Password, &obj.Plmn, &obj.Isactive, &obj.Isallowed, &obj.Createddate, &obj.Createdby, &obj.Updateddate, &obj.Updatedby, &obj.Rowversion); rerr != nil {
				fmt.Println("error ", rerr)
			}
			slist.List = append(slist.List, obj)
		}
		rows.Close()
		if slist.List != nil {
			slist.TotalRecords = len(slist.List)
		}
	} else {
		fmt.Println("err ", err)
	}

	return slist
}

// Table Name : fsubscriber

func addFsubscriber(c *gin.Context) {
	var new_fsubscriber Fsubscriber

	if err := c.BindJSON(&new_fsubscriber); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if new_fsubscriber.Dnnid2 == "" {
		new_fsubscriber.Dnnid2 = "1"
	}

	if new_fsubscriber.Dnnid3 == "" {
		new_fsubscriber.Dnnid3 = "1"
	}

	//set custom feild values
	t := time.Now()
	user := GetLoggedInUser()
	new_fsubscriber.Amf = "8000"
	new_fsubscriber.Createddate = t.Format("2006-01-02 15:04:05")
	new_fsubscriber.Createdby = user.Appuserid
	new_fsubscriber.Updateddate = t.Format("2006-01-02 15:04:05")
	new_fsubscriber.Updatedby = user.Appuserid
	new_fsubscriber.Isactive = 1
	new_fsubscriber.Rowversion = 1

	Fsubscriber_Add(0, &new_fsubscriber)
	c.IndentedJSON(http.StatusOK, new_fsubscriber.Subscriberid)
}

func updateFsubscriber(c *gin.Context) {
	var obj_fsubscriber Fsubscriber

	if err := c.BindJSON(&obj_fsubscriber); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if obj_fsubscriber.Dnnid2 == "" {
		obj_fsubscriber.Dnnid2 = "1"
	}

	if obj_fsubscriber.Dnnid3 == "" {
		obj_fsubscriber.Dnnid3 = "1"
	}

	//set custom feild values
	t := time.Now()
	user := GetLoggedInUser()
	obj_fsubscriber.Updateddate = t.Format("2006-01-02 15:04:05")
	obj_fsubscriber.Updatedby = user.Appuserid
	obj_fsubscriber.Isactive = 1
	obj_fsubscriber.Rowversion = 1
	obj_fsubscriber.Amf = "8000"

	sts := Fsubscriber_Update(0, &obj_fsubscriber)
	c.IndentedJSON(http.StatusOK, sts)
}

func getFsubscriberList(c *gin.Context) {
	var pageInfo FPageInfo

	if perr := c.BindJSON(&pageInfo); perr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": perr.Error()})
		return
	}

	var obj Fsubscriber
	obj.Isactive = 1

	list := Fsubscriber_GetByPaging(0, pageInfo, &obj)
	c.IndentedJSON(http.StatusOK, list)
}

// Table Name : dnn

func addDnn(c *gin.Context) {
	var new_Dnn Dnn

	if err := c.BindJSON(&new_Dnn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//set custom feild values
	new_Dnn.Isactive = 1

	Dnn_Add(0, &new_Dnn)
	c.IndentedJSON(http.StatusOK, new_Dnn.Dnnid)
}

func updateDnn(c *gin.Context) {
	var obj_Dnn Dnn

	if err := c.BindJSON(&obj_Dnn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//set custom feild values
	obj_Dnn.Isactive = 1

	sts := Dnn_Update(0, &obj_Dnn)
	c.IndentedJSON(http.StatusOK, sts)
}

func delDnn(c *gin.Context) {
	var obj_Dnn Dnn

	if err := c.BindJSON(&obj_Dnn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//set custom feild values
	obj_Dnn.Isactive = 0

	sts := Dnn_Update(0, &obj_Dnn)
	c.IndentedJSON(http.StatusOK, sts)
}

func getDnnList(c *gin.Context) {
	var pageInfo FPageInfo

	if perr := c.BindJSON(&pageInfo); perr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": perr.Error()})
		return
	}

	var obj Dnn
	obj.Isactive = 1

	list := Dnn_GetByPaging(0, pageInfo, &obj)
	c.IndentedJSON(http.StatusOK, list)
}

func getDnnListAll(c *gin.Context) {
	var pageInfo FPageInfo

	if perr := c.BindJSON(&pageInfo); perr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": perr.Error()})
		return
	}

	var obj Dnn
	obj.Isactive = 1

	list := Dnn_GetAll(0, &obj)
	c.IndentedJSON(http.StatusOK, list)
}

// Table Name : accountstatus

func addAccountstatus(c *gin.Context) {
	var new_Accountstatus Accountstatus

	if err := c.BindJSON(&new_Accountstatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//set custom feild values

	Accountstatus_Add(0, &new_Accountstatus)
	c.IndentedJSON(http.StatusOK, new_Accountstatus.Accountstatusid)
}

func updateAccountstatus(c *gin.Context) {
	var obj_Accountstatus Accountstatus

	if err := c.BindJSON(&obj_Accountstatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//set custom feild values

	sts := Accountstatus_Update(0, &obj_Accountstatus)
	c.IndentedJSON(http.StatusOK, sts)
}

func getAccountstatusList(c *gin.Context) {
	var pageInfo FPageInfo

	if perr := c.BindJSON(&pageInfo); perr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": perr.Error()})
		return
	}

	var obj Accountstatus
	list := Accountstatus_GetByPaging(0, pageInfo, &obj)
	c.IndentedJSON(http.StatusOK, list)
}

// Table Name : appuser

func addAppuser(c *gin.Context) {
	var new_appuser Appuser

	if err := c.BindJSON(&new_appuser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//set custom feild values
	t := time.Now()
	user := GetLoggedInUser()
	new_appuser.Isactive = 1
	new_appuser.Createddate = t.Format("2006-01-02 15:04:05")
	new_appuser.Createdby = user.Appuserid
	new_appuser.Updateddate = t.Format("2006-01-02 15:04:05")
	new_appuser.Updatedby = user.Appuserid
	new_appuser.Rowversion = 1

	Appuser_Add(0, &new_appuser)
	c.IndentedJSON(http.StatusOK, new_appuser.Appuserid)
}

func updateAppuser(c *gin.Context) {
	var obj_appuser Appuser

	if err := c.BindJSON(&obj_appuser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//set custom feild values
	t := time.Now()
	user := GetLoggedInUser()
	obj_appuser.Isactive = 1
	obj_appuser.Updateddate = t.Format("2006-01-02 15:04:05")
	obj_appuser.Updatedby = user.Appuserid
	obj_appuser.Rowversion = 1

	sts := Appuser_Update(0, &obj_appuser)
	c.IndentedJSON(http.StatusOK, sts)
}

func delAppuser(c *gin.Context) {
	var obj_appuser Appuser

	if err := c.BindJSON(&obj_appuser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//set custom feild values
	t := time.Now()
	user := GetLoggedInUser()
	obj_appuser.Isactive = 0
	obj_appuser.Updateddate = t.Format("2006-01-02 15:04:05")
	obj_appuser.Updatedby = user.Appuserid
	obj_appuser.Rowversion = 1

	sts := Appuser_Update(0, &obj_appuser)
	c.IndentedJSON(http.StatusOK, sts)
}

func getAppuserList(c *gin.Context) {
	var pageInfo FPageInfo

	if perr := c.BindJSON(&pageInfo); perr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": perr.Error()})
		return
	}

	var obj Appuser
	obj.Isactive = 1

	list := Appuser_GetByPaging(0, pageInfo, &obj)
	c.IndentedJSON(http.StatusOK, list)
}

func createTestSubscriber(imsi int) {
	var ss Fsubscriber

	ss.Imsi = imsi
	ss.Plmn = "0"
	ss.Skey = "12121212121212121212121212121212"
	ss.Seq = 0
	ss.Opc = "12121212121212121212121212121212"
	ss.Amf = "8000"
	ss.Accountstatus = 1
	ss.Defsst = 1
	ss.Defsd = ""
	ss.Uplink = 5000
	ss.Downlink = 5000
	ss.Dnnid1 = "2"
	ss.Dnnid2 = "1"
	ss.Dnnid3 = "1"
	ss.Msisdn = ""
	ss.Createddate = "2022-08-30 7:41"
	ss.Createdby = 1
	ss.Updateddate = "2022-08-30 7:41"
	ss.Updatedby = 1
	ss.Isactive = 1
	ss.Rowversion = 1

	Fsubscriber_Add(0, &ss)
}

// ALTER TABLE fsubscriber ADD CONSTRAINT fsubscriber_imsi UNIQUE (imsi);

func createTestSubscribers() {
	var startImsi = 200000000000010
	var total = 200000
	var i = 0

	t := time.Now()
	fmt.Println("Started Inserting ", t)

	for i = 1; i <= total; i++ {
		createTestSubscriber((startImsi + i))
	}

	t = time.Now()
	fmt.Println("End Inserting ", t)
}

func main() {

	connStr := "user=postgres password=123456 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")
	router := gin.Default()

	router.Static("/web", "./html")

	// Table Name : fsubscriber
	router.POST("/fsubscriber/add", addFsubscriber)
	router.POST("/fsubscriber/upd", updateFsubscriber)
	router.POST("/fsubscriber/get", getFsubscriberList)

	// Table Name : dnn
	router.POST("/dnn/add", addDnn)
	router.POST("/dnn/upd", updateDnn)
	router.POST("/dnn/get", getDnnList)
	router.POST("/dnn/del", delDnn)
	router.POST("/dnn/geta", getDnnListAll)

	// // Table Name : accountstatus
	// router.POST("/accountstatus/add", addAccountstatus)
	// router.POST("/accountstatus/upd", updateAccountstatus)
	// router.POST("/accountstatus/get", getAccountstatusList)

	// Table Name : appuser
	router.POST("/appuser/add", addAppuser)
	router.POST("/appuser/upd", updateAppuser)
	router.POST("/appuser/get", getAppuserList)
	router.POST("/appuser/del", delAppuser)

	t := time.Now()
	fmt.Println("Started Server ", t)
	//router.Run(fmt.Sprintf(":%v", appConfig.Port))
	router.Run(":8080")
}
