package pokeapi

type LocationDetails struct {
	EcounterMethodRates []EncounterMethod      `json:"encounter_method_rates"`
	GameIndex           int                    `json:"game_index"`
	ID                  int                    `json:"id"`
	Location            []CoreItem             `json:"location"`
	Name                string                 `json:"name"`
	Names               []t_Name_Localized     `json:"names"`
	PokemonEncounters   []t_Pokemon_Encounters `json:"pokemon_encounters"`
}

type EncounterMethod struct {
	Name           string             `json:"name"`
	URL            string             `json:"url"`
	VersionDetails []t_VersionDetails `json:"version_details"`
}

type t_Name_Localized struct {
	Language CoreItem `json:"language"`
	Name     string   `json:"name"`
}

type t_Pokemon_Encounters struct {
	Pokemon        CoreItem              `json:"pokemon"`
	VersionDetails []t_Encounter_Details `json:"version_details"`
	MaxChance      int                   `json:"max_chance"`
	Version        CoreItem              `json:"version"`
}

type t_Encounter_Details struct {
	EncounterDetails struct {
		Chance          int      `json:"chance"`
		ConditionValues []any    `json:"condition_values"`
		MaxLevel        int      `json:"max_level"`
		Method          CoreItem `json:"method"`
		MinLevel        int      `json:"min_level"`
	} `json:"encounter_details"`
}

type t_VersionDetails struct {
	Rate    int      `json:"rate"`
	Version CoreItem `json:"version"`
}

type CoreItem struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
