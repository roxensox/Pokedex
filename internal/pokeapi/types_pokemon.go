package pokeapi

type t_Pokemon struct {
	Abilities []struct {
		Ability  CoreItem `json:"ability"`
		IsHidden bool     `json:"is_hidden"`
		Slot     int      `json:"slot"`
	} `json:"abilities"`
	BaseExperience int `json:"base_experience"`
	Cries          struct {
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
	GameIndices []struct {
		GameIndex int        `json:"game_index"`
		Version   []CoreItem `json:"version"`
	} `json:"game_indices"`
	Height    int `json:"height"`
	HeldItems []struct {
		Item           CoreItem `json:"item"`
		VersionDetails struct {
			Rarity  int      `json:"rarity"`
			Version CoreItem `json:"version"`
		} `json:"version_details"`
	} `json:"held_items"`
	ID                     int    `json:"id"`
	IsDefault              bool   `json:"is_default"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move                CoreItem `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int      `json:"level_learned_at"`
			MoveLearnMethod CoreItem `json:"move_learn_method"`
			Order           any      `json:"order"`
			VersionGroup    CoreItem `json:"version_group"`
		} `json:"version_group_details"`
	}
	Name   string `json:"name"`
	Order  int    `json:"order"`
	Weight int    `json:"weight"`
	Stats  []struct {
		BaseStat int      `json:"base_stat"`
		Effort   int      `json:"effort"`
		Stat     CoreItem `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int      `json:"slot"`
		Type CoreItem `json:"type"`
	} `json:"types"`
}
