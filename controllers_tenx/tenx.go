package controllers_tenx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego"
)

type TenxController struct {
	beego.Controller
}

/* Use this tool to generate struct https://mholt.github.io/json-to-go/ */

type Health struct {
	Status string `json:"status"`
	Output struct {
		Timechecks struct {
			RoundStatus string `json:"round_status"`
			Epoch       int    `json:"epoch"`
			Round       int    `json:"round"`
		} `json:"timechecks"`
		Health struct {
			HealthServices []struct {
				Mons []struct {
					LastUpdated  string `json:"last_updated"`
					Name         string `json:"name"`
					AvailPercent int    `json:"avail_percent"`
					KbTotal      int    `json:"kb_total"`
					KbAvail      int    `json:"kb_avail"`
					Health       string `json:"health"`
					KbUsed       int    `json:"kb_used"`
					StoreStats   struct {
						BytesTotal  int    `json:"bytes_total"`
						BytesLog    int    `json:"bytes_log"`
						LastUpdated string `json:"last_updated"`
						BytesMisc   int    `json:"bytes_misc"`
						BytesSst    int    `json:"bytes_sst"`
					} `json:"store_stats"`
					HealthDetail string `json:"health_detail, omitempty"`
				} `json:"mons"`
			} `json:"health_services"`
		} `json:"health"`
		OverallStatus string `json:"overall_status"`
		Summary       []struct {
			Severity string `json:"severity"`
			Summary  string `json:"summary"`
		} `json:"summary"`
		Detail []string `json:"detail"`
	} `json:"output"`
}

//HealthItem cluster health data structure for UI
type HealthItem struct {
	OverallStatus string `json:"overall_status"`
	Summary       []struct {
		Severity string `json:"severity"`
		Summary  string `json:"summary"`
	} `json:"summary"`
	Detail []string `json:"detail"`
}

//Status status info, path: /status
type Status struct {
	Status string `json:"status"`
	Output struct {
		QuorumNames []string `json:"quorum_names"`
		Monmap      struct {
			Mons []struct {
				Name string `json:"name"`
				Addr string `json:"addr"`
			} `json:"mons, omitempty"`
		} `json:"monmap"`
		Pgmap struct {
			PgsByState []PgsByStateItem `json:"pgs_by_state"`
			NumPgs     int              `json:"num_pgs"`
		} `json:"pgmap"`
	} `json:"output"`
}

//OsdStat osd statistic info, path: /osd/stat
type OsdStat struct {
	Status string `json:"status"`
	Output struct {
		Epoch          int  `json:"epoch"`
		NumOsds        int  `json:"num_osds"`
		NumUpOsds      int  `json:"num_up_osds"`
		NumInOsds      int  `json:"num_in_osds"`
		Full           bool `json:"full"`
		Nearfull       bool `json:"nearfull"`
		NumRemappedPgs int  `json:"num_remapped_pgs"`
	} `json:"output"`
}

//OsdStatItem osd statistic info for UI
type OsdStatItem struct {
	NumOsds   int `json:"num_osds"`
	NumUpOsds int `json:"num_up_osds"`
	NumInOsds int `json:"num_in_osds"`
}

//MonStatItem mon statistic info for UI
type MonStatItem struct {
	MonNames    []string `json:"mon_names"`
	QuorumNames []string `json:"quorum_names"`
}

//OsdTree ceph-rest-api path: /osd/tree
type OsdTree struct {
	Status string `json:"status"`
	Output struct {
		Nodes []struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Type     string `json:"type"`
			TypeID   int    `json:"type_id"`
			Status   string `json:"status, omitempty"`
			Children []int  `json:"children, omitempty"`
		} `json:"nodes, omitempty"`
	} `json:"output"`
}

//MonStatus ceph-rest-api path: /mon_status
type MonStatus struct {
	Status string `json:"status"`
	Output struct {
		Monmap struct {
			Mons []struct {
				Name string `json:"name"`
				Addr string `json:"addr"`
			} `json:"mons, omitempty"`
		} `json:"monmap"`
	} `json:"output"`
}

//PoolNames ceph-rest-api path: /osd/pool/ls
type PoolNames struct {
	Status string   `json:"status"`
	Output []string `json:"output"`
}

//PoolSize ceph-rest-api path: /osd/pool/get?pool={pool-name}&var=size
type PoolSize struct {
	Status string `json:"status"`
	Output struct {
		Size int `json:"size"`
	} `json:"output"`
}

//PoolPgNum ceph-rest-api path: /osd/pool/get?pool={pool-name}&var=pg_num
type PoolPgNum struct {
	Status string `json:"status"`
	Output struct {
		PgNum int `json:"pg_num"`
	} `json:"output"`
}

const (
	HOST_TYPE string = "host"
	OSD_TYPE  string = "osd"
)

//DiskUsageItem Disk usage data structure for UI
type DiskUsageItem struct {
	TotalKb      int `json:"total_kb"`
	TotalKbAvail int `json:"total_kb_avail"`
	TotalKbUsed  int `json:"total_kb_used"`
}

//PgsByStateItem PgsByState data structure for UI
type PgsByStateItem struct {
	StateName string `json:"state_name"`
	Count     int    `json:"count"`
}

//PgItem PG data structure for UI
type PgStatItem struct {
	PgsByState []PgsByStateItem `json:"pgs_by_state"`
	NumPgs     int              `json:"num_pgs"`
}

//PoolItem Pool data structure for UI
type PoolItem struct {
	Name  string `json:"name"`
	Size  int    `json:"size"`
	PgNum int    `json:"pg_num"`
}

//MonItem monitor node data structure for UI
type MonItem struct {
	Name         string `json:"name"`
	Addr         string `json:"addr"`
	Health       string `json:"health"`
	HealthDetail string `json:"health_detail"`
}

//OsdItem OSD node data structure for UI
type OsdItem struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Kb      int    `json:"kb"`
	KbUsed  int    `json:"kb_used"`
	KbAvail int    `json:"kb_avail"`
}

//HostItem host data structure for UI
type HostItem struct {
	Name string    `json:"name"`
	Mon  MonItem   `json:"mons"`
	Osds []OsdItem `json:"osds"`
}

type OsdDf struct {
	Status string `json:"status"`
	Output struct {
		Nodes []struct {
			Kb          int     `json:"kb"`
			Name        string  `json:"name"`
			TypeID      int     `json:"type_id"`
			Reweight    float64 `json:"reweight"`
			CrushWeight float64 `json:"crush_weight"`
			Utilization float64 `json:"utilization"`
			Depth       int     `json:"depth"`
			KbAvail     int     `json:"kb_avail"`
			KbUsed      int     `json:"kb_used"`
			Var         float64 `json:"var"`
			Type        string  `json:"type"`
			ID          int     `json:"id"`
		} `json:"nodes"`
		Stray   []interface{} `json:"stray"`
		Summary struct {
			TotalKb            int     `json:"total_kb"`
			Dev                float64 `json:"dev"`
			MaxVar             float64 `json:"max_var"`
			TotalKbAvail       int     `json:"total_kb_avail"`
			MinVar             float64 `json:"min_var"`
			AverageUtilization float64 `json:"average_utilization"`
			TotalKbUsed        int     `json:"total_kb_used"`
		} `json:"summary"`
	} `json:"output"`
}

func getPgStatInfo(status Status) PgStatItem {
	pgStatItem := PgStatItem{}

	pgStatItem.NumPgs = status.Output.Pgmap.NumPgs

	pgStatItem.PgsByState = status.Output.Pgmap.PgsByState

	return pgStatItem
}

func getMonStatInfo(status Status) MonStatItem {
	monStatItem := MonStatItem{}

	monStatItem.QuorumNames = status.Output.QuorumNames

	monStatItem.MonNames = make([]string, 0)

	for _, mon := range status.Output.Monmap.Mons {
		monStatItem.MonNames = append(monStatItem.MonNames, mon.Name)
	}

	return monStatItem
}

func getOsdStatInfo(osdStat OsdStat) OsdStatItem {
	osdStatItem := OsdStatItem{}

	osdStatItem.NumOsds = osdStat.Output.NumOsds
	osdStatItem.NumUpOsds = osdStat.Output.NumUpOsds
	osdStatItem.NumInOsds = osdStat.Output.NumInOsds

	return osdStatItem
}

func getClusterHealthInfo(health Health) HealthItem {
	healthItem := HealthItem{}

	healthItem.OverallStatus = health.Output.OverallStatus
	healthItem.Summary = health.Output.Summary
	healthItem.Detail = health.Output.Detail

	return healthItem
}

func getMonStatusInfo(monStatus MonStatus) map[string]MonItem {
	var monMap = make(map[string]MonItem)

	for _, node := range monStatus.Output.Monmap.Mons {
		mon := MonItem{}
		mon.Name = node.Name
		mon.Addr = node.Addr
		monMap[mon.Name] = mon
	}

	return monMap
}

func getMonHealthInfo(health Health) map[string]MonItem {
	var monMap = make(map[string]MonItem)

	for _, service := range health.Output.Health.HealthServices {
		for _, node := range service.Mons {
			mon := MonItem{}
			mon.Name = node.Name
			mon.Health = node.Health
			mon.HealthDetail = node.HealthDetail
			monMap[mon.Name] = mon
		}
	}

	return monMap
}

func getDiskUsageInfo(osdDf OsdDf) DiskUsageItem {
	var diskUsageItem DiskUsageItem

	diskUsageItem.TotalKb = osdDf.Output.Summary.TotalKb
	diskUsageItem.TotalKbAvail = osdDf.Output.Summary.TotalKbAvail
	diskUsageItem.TotalKbUsed = osdDf.Output.Summary.TotalKbUsed

	return diskUsageItem
}

func getOsdStorageInfo(osdDf OsdDf) map[int]OsdItem {
	var osdItemMap = make(map[int]OsdItem)

	for _, node := range osdDf.Output.Nodes {
		if node.Type == OSD_TYPE {
			osd := OsdItem{}
			osd.Name = node.Name
			osd.Kb = node.Kb
			osd.KbAvail = node.KbAvail
			osd.KbUsed = node.KbUsed
			osdItemMap[node.ID] = osd
		}
	}

	return osdItemMap
}

func convertOsdTree2HostItem(osdTree OsdTree, monStatusMap map[string]MonItem, monHealthMap map[string]MonItem,
	osdStoragemap map[int]OsdItem) []HostItem {
	var hosts = make([]HostItem, 0)
	var osdMap = make(map[int]OsdItem)

	for _, node := range osdTree.Output.Nodes {
		if node.Type == OSD_TYPE {
			osd := OsdItem{}
			osd.Name = node.Name
			osd.Status = node.Status
			osdMap[node.ID] = osd
		}
	}

	for _, node := range osdTree.Output.Nodes {
		if node.Type == HOST_TYPE {
			host := HostItem{}
			host.Name = node.Name
			if _, ok := monStatusMap[host.Name]; ok {
				host.Mon = monStatusMap[host.Name]
				host.Mon.Health = monHealthMap[host.Name].Health
				host.Mon.HealthDetail = monHealthMap[host.Name].HealthDetail
			}
			host.Osds = make([]OsdItem, 0)
			for _, id := range node.Children {
				osd := osdStoragemap[id]
				osd.Status = osdMap[id].Status
				host.Osds = append(host.Osds, osd)
			}
			hosts = append(hosts, host)
		}
	}

	return hosts
}

func RequestJson(url string) []byte {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

type MdsStat struct {
	Status string `json:"status"`
	Output struct {
		MdsmapFirstCommitted int `json:"mdsmap_first_committed"`
		Mdsmap               struct {
			SessionAutoclose int `json:"session_autoclose"`
			Up               struct {
				Mds0 int `json:"mds_0"`
			} `json:"up"`
			LastFailureOsdEpoch int           `json:"last_failure_osd_epoch"`
			In                  []int         `json:"in"`
			LastFailure         int           `json:"last_failure"`
			MaxFileSize         int64         `json:"max_file_size"`
			Tableserver         int           `json:"tableserver"`
			MetadataPool        int           `json:"metadata_pool"`
			Failed              []interface{} `json:"failed"`
			Epoch               int           `json:"epoch"`
			Flags               int           `json:"flags"`
			MaxMds              int           `json:"max_mds"`
			Compat              struct {
				Compat struct {
				} `json:"compat"`
				RoCompat struct {
				} `json:"ro_compat"`
				Incompat struct {
					Feature8 string `json:"feature_8"`
					Feature2 string `json:"feature_2"`
					Feature3 string `json:"feature_3"`
					Feature1 string `json:"feature_1"`
					Feature6 string `json:"feature_6"`
					Feature4 string `json:"feature_4"`
					Feature5 string `json:"feature_5"`
				} `json:"incompat"`
			} `json:"compat"`
			DataPools []int `json:"data_pools"`
			Info      struct {
				Gid4106 struct {
					StandbyForRank int           `json:"standby_for_rank"`
					ExportTargets  []interface{} `json:"export_targets"`
					Name           string        `json:"name"`
					Incarnation    int           `json:"incarnation"`
					StateSeq       int           `json:"state_seq"`
					State          string        `json:"state"`
					Gid            int           `json:"gid"`
					Rank           int           `json:"rank"`
					StandbyForName string        `json:"standby_for_name"`
					Addr           string        `json:"addr"`
				} `json:"gid_4106"`
			} `json:"info"`
			FsName         string        `json:"fs_name"`
			Created        string        `json:"created"`
			Enabled        bool          `json:"enabled"`
			Modified       string        `json:"modified"`
			SessionTimeout int           `json:"session_timeout"`
			Stopped        []interface{} `json:"stopped"`
			Root           int           `json:"root"`
		} `json:"mdsmap"`
		MdsmapLastCommitted int `json:"mdsmap_last_committed"`
	} `json:"output"`
}

type OsdCrushDump struct {
	Status string `json:"status"`
	Output struct {
		Rules []struct {
			MinSize  int    `json:"min_size"`
			RuleName string `json:"rule_name"`
			Steps    []struct {
				ItemName string `json:"item_name"`
				Item     int    `json:"item"`
				Op       string `json:"op"`
			} `json:"steps"`
			Ruleset int `json:"ruleset"`
			Type    int `json:"type"`
			RuleID  int `json:"rule_id"`
			MaxSize int `json:"max_size"`
		} `json:"rules"`
		Tunables struct {
			Profile                  string `json:"profile"`
			HasV3Rules               int    `json:"has_v3_rules"`
			HasV4Buckets             int    `json:"has_v4_buckets"`
			ChooseTotalTries         int    `json:"choose_total_tries"`
			RequireFeatureTunables3  int    `json:"require_feature_tunables3"`
			LegacyTunables           int    `json:"legacy_tunables"`
			ChooseleafDescendOnce    int    `json:"chooseleaf_descend_once"`
			AllowedBucketAlgs        int    `json:"allowed_bucket_algs"`
			ChooseLocalFallbackTries int    `json:"choose_local_fallback_tries"`
			HasV2Rules               int    `json:"has_v2_rules"`
			StrawCalcVersion         int    `json:"straw_calc_version"`
			RequireFeatureTunables2  int    `json:"require_feature_tunables2"`
			OptimalTunables          int    `json:"optimal_tunables"`
			ChooseLocalTries         int    `json:"choose_local_tries"`
			ChooseleafVaryR          int    `json:"chooseleaf_vary_r"`
			RequireFeatureTunables   int    `json:"require_feature_tunables"`
		} `json:"tunables"`
		Buckets []struct {
			Hash     string `json:"hash"`
			Name     string `json:"name"`
			Weight   int    `json:"weight"`
			TypeID   int    `json:"type_id"`
			Alg      string `json:"alg"`
			TypeName string `json:"type_name"`
			Items    []struct {
				ID     int `json:"id"`
				Weight int `json:"weight"`
				Pos    int `json:"pos"`
			} `json:"items"`
			ID int `json:"id"`
		} `json:"buckets"`
		Types []struct {
			Name   string `json:"name"`
			TypeID int    `json:"type_id"`
		} `json:"types"`
		Devices []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"devices"`
	} `json:"output"`
}

type PgStat struct {
	Status string `json:"status"`
	Output struct {
		NumPgByState []struct {
			Num  int    `json:"num"`
			Name string `json:"name"`
		} `json:"num_pg_by_state"`
		NumPgs        int   `json:"num_pgs"`
		RawBytes      int64 `json:"raw_bytes"`
		NumBytes      int   `json:"num_bytes"`
		Version       int   `json:"version"`
		RawBytesUsed  int64 `json:"raw_bytes_used"`
		RawBytesAvail int64 `json:"raw_bytes_avail"`
	} `json:"output"`
}

func (c *TenxController) Get() {

	baseUrl := "http://192.168.1.87:5000/api/v0.1"
	//baseUrl := "http://192.168.99.100:5000/api/v0.1"

	// Request osd df
	body := RequestJson(baseUrl + "/mon_status")
	var monStatus MonStatus
	err := json.Unmarshal(body, &monStatus)
	if err != nil {
		panic(err)
	}

	monStatusMap := getMonStatusInfo(monStatus)

	// Request health
	body = RequestJson(baseUrl + "/health")
	var health Health
	err = json.Unmarshal(body, &health)
	if err != nil {
		panic(err)
	}

	healthItem := getClusterHealthInfo(health)

	c.Data["health"] = healthItem

	monHealthMap := getMonHealthInfo(health)

	// Request osd df
	body = RequestJson(baseUrl + "/osd/df")
	var osdDf OsdDf
	err = json.Unmarshal(body, &osdDf)
	if err != nil {
		panic(err)
	}
	c.Data["osdDf"] = osdDf

	osdStorageMap := getOsdStorageInfo(osdDf)
	diskUsageItem := getDiskUsageInfo(osdDf)
	c.Data["diskUsageItem"] = diskUsageItem

	// Request osd tree
	body = RequestJson(baseUrl + "/osd/tree")
	var osdTree OsdTree
	err = json.Unmarshal(body, &osdTree)
	if err != nil {
		panic(err)
	}
	fmt.Printf("body: %v", string(body))
	hosts := convertOsdTree2HostItem(osdTree, monStatusMap, monHealthMap, osdStorageMap)

	hosts_json, _ := json.Marshal(hosts)
	c.Data["hosts"] = string(hosts_json)

	// Request osd statistic info
	body = RequestJson(baseUrl + "/osd/stat")
	var osdStat OsdStat
	err = json.Unmarshal(body, &osdStat)
	if err != nil {
		panic(err)
	}
	osdStatItem := getOsdStatInfo(osdStat)
	c.Data["osdStatItem"] = osdStatItem

	// Request mon statistic info
	body = RequestJson(baseUrl + "/status")
	var status Status
	err = json.Unmarshal(body, &status)
	if err != nil {
		panic(err)
	}
	monStatItem := getMonStatInfo(status)
	pgStatItem := getPgStatInfo(status)
	c.Data["monStatItem"] = monStatItem
	c.Data["pgStatItem"] = pgStatItem

	// Request pool info
	body = RequestJson(baseUrl + "/osd/pool/ls")
	var poolNames PoolNames
	err = json.Unmarshal(body, &poolNames)
	if err != nil {
		panic(err)
	}

	poolItems := make([]PoolItem, 0)
	for _, name := range poolNames.Output {
		var poolItem PoolItem
		poolItem.Name = name

		body = RequestJson(baseUrl + "/osd/pool/get?pool=" + name + "&var=size")
		var poolSize PoolSize
		err = json.Unmarshal(body, &poolSize)
		if err != nil {
			panic(err)
		}
		poolItem.Size = poolSize.Output.Size

		body = RequestJson(baseUrl + "/osd/pool/get?pool=" + name + "&var=pg_num")
		var poolPgNum PoolPgNum
		err = json.Unmarshal(body, &poolPgNum)
		if err != nil {
			panic(err)
		}
		poolItem.PgNum = poolPgNum.Output.PgNum

		poolItems = append(poolItems, poolItem)
	}

	c.Data["poolItems"] = poolItems

	c.TplName = "tenx.tpl"
}
