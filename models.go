package main

type Building struct {
	BUILDING_ID                  string `json:"building_id"`
	BLDG_DESC                    string `json:"bldg_desc"`
	BLDG_LOCATION                string `json:"bldg_location"`
	BLDG_TYPE                    string `json:"bldg_type"`
	BLDG_LONG_DESC               string `json:"bldg_long_desc"`
	BLDG_CITY                    string `json:"bldg_city"`
	BLDG_STATE                   string `json:"bldg_state"`
	BLDG_ZIP                     string `json:"bldg_zip"`
	BLDG_SECTOR                  string `json:"bldg_sector"`
	BLDG_LATITUDE                string `json:"bldg_latitude"`
	BLDG_LONGITUDE               string `json:"bldg_longitude"`
	BUILDINGS_ADD_DATE           string `json:"bldg_add_date"`
	BUILDINGS_CHGDATE            string `json:"bldg_chgdate"`
	BLDG_TYPE_REPRESENTATION     string `json:"bldg_representation"`
	BLDG_LOCATION_REPRESENTATION string `json:"bldg_location_representation"`
	BLDG_SECTOR_REPRESENTATION   string `json:"bldg_sector_representation"`
}

// slice of pointer to Building struct to define its properties
// choose mcewen as example
//   (N'AVPS', N'Gray Pavilion - Pol. Science', N'EL', N'IH', N'Academic Village-Political Science', 
// N'Elon', N'NC', N'27244', N'AVIL', '36.102012', '-79.502129', '2004-06-24 00:00:00.000', 
// '2022-08-19 00:00:00.000', N'Inactive Housing', N'Elon University Campus', N'Lambert Academic Village'),
var buildings = []*Building{
	{
		BUILDING_ID: 					"AVPS",
		BLDG_DESC: 						"Gray Pavilion - Pol. Science",                              
		BLDG_LOCATION: 					"EL",             
		BLDG_TYPE:     					"IH",         
		BLDG_LONG_DESC: 				"Academic Village-Political Science",          
		BLDG_CITY:    					"Elon",              
		BLDG_STATE:   					"NC",               
		BLDG_ZIP:     					"27244",           
		BLDG_SECTOR:  					"AVIL",  
		BLDG_LATITUDE: 					"36.102012",        
		BLDG_LONGITUDE:  				"-79.502129",  
		BUILDINGS_ADD_DATE: 			"2004-06-24 00:00:00.000",          
		BUILDINGS_CHGDATE:  			"2022-08-19 00:00:00.000",           
		BLDG_TYPE_REPRESENTATION: 		"Inactive Housing",  
		BLDG_LOCATION_REPRESENTATION: 	"Elon University Campus",
		BLDG_SECTOR_REPRESENTATION: 	"Lambert Academic Village"
	},
}


// point to building struct and return buildings...emulating a database
func listBuildings() []*Building {
	return buildings
}

// return building with the same id as input id
func getBuilding(id string) *Building{
	for _, building := range buildings {
		if building.BUILDING_ID == id {
			return building
		}
	}
	return nil
}

// adds buildings to the collection
func storeBuilding(building Building){
	buildings = append(building, &building)
}

// removes building that matches with the same id by slicing
func deleteBuilding(id string) *Building{
	for i, building := range buildings{
		if building.BUILDING_ID == id {
			buildings = append(buildings[:i], (buildings)[i+1:]...)
		}
	}
	return nil
}

// building with id as input, and replace that building with buildingUpdate
func updateBuilding(id string, buildingUpdate Building) *Building{
	for i, building := range buildings{
		if building.BUILDING_ID == id {
			buildings[i] = &buildingUpdate
			return building
		}
	}
	return nil
}