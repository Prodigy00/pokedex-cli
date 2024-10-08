package api

type LocationArea struct {
	EncounterMethodRates `json:"encounter_method_rates"`
	GameIndex            int               `json:"game_index"`
	ID                   int               `json:"id"`
	Name                 string            `json:"name"`
	Names                []AreaDetails     `json:"names"`
	PokemonEncounters    PokemonEncounters `json:"pokemon_encounters"`
}
type AreaDetails struct {
	Language `json:"language"`
	Name     string `json:"name"`
}

type Language struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type EncounterMethodRates []EncounterMethodRate

type EncounterMethodRate struct {
	EncounterMethod   `json:"encounter_method"`
	VersionDetailsEMR `json:"version_details"`
}

type EncounterMethod struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type VersionDetailsEMR []VersionDetailEMR

type VersionDetailEMR struct {
	Rate    int     `json:"rate"`
	Version Version `json:"version"`
}

type PokemonEncounters []PokemonEncounter

type PokemonEncounter struct {
	Pokemon        `json:"pokemon"`
	VersionDetails `json:"version_details"`
}

type Pokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type VersionDetails []VersionDetail

type VersionDetail struct {
	EncounterDetails `json:"encounter_details"`
	MaxChance        int     `json:"max_chance"`
	Version          Version `json:"version"`
}

type EncounterDetails []EncounterDetail

type EncounterDetail struct {
	Chance          int      `json:"chance"`
	ConditionValues []Result `json:"condition_values"`
	MaxLevel        int      `json:"max_level"`
	Method          Result   `json:"method"`
	MinLevel        int      `json:"min_level"`
}

type Version struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type CatchPokemonResult struct {
	Abilities      []AbilityData `json:"abilities"`
	BaseExperience int           `json:"base_experience"`
	Cries          struct {
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
	Forms                  []Form          `json:"forms"`
	GameIndices            []GameIndexData `json:"game_indices"`
	Height                 int             `json:"height"`
	HeldItems              []ItemsData     `json:"held_items"`
	ID                     int             `json:"id"`
	IsDefault              bool            `json:"is_default"`
	LocationAreaEncounters string          `json:"location_area_encounters"`
	Moves                  []MoveData      `json:"moves"`
	Name                   string          `json:"name"`
	Order                  int             `json:"order"`
	PastAbilities          []any           `json:"past_abilities"`
	PastTypes              []any           `json:"past_types"`
	Species                struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"species"`
	Sprites any        `json:"sprites"`
	Stats   []PokeStat `json:"stats"`
	Types   []TypeData `json:"types"`
	Weight  int        `json:"weight"`
}

type AbilityData struct {
	Ability struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"ability"`
	IsHidden bool `json:"is_hidden"`
	Slot     int  `json:"slot"`
}

type Form struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type GameIndexData struct {
	GameIndex int     `json:"game_index"`
	Version   Version `json:"version"`
}

type ItemsData struct {
	Item struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"item"`

	VersionDetails []struct {
		Rarity  int `json:"rarity"`
		Version struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		}
	} `json:"version_details"`
}

type MoveData struct {
	Move struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	VersionGroupDetails []struct {
		LevelLearnedAt  int `json:"level_learned_at"`
		MoveLearnMethod struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		}
		VersionGroup struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		}
	}
}

type PokeStat struct {
	BaseStat int `json:"base_stat"`
	Effort   int `json:"effort"`
	Stat     struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
}

type TypeData struct {
	Slot int `json:"slot"`
	Type struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
}
