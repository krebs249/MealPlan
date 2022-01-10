package models

type RecipesAPI struct {
	Recipes []struct {
		Vegetarian               bool    `json:"vegetarian"`
		Vegan                    bool    `json:"vegan"`
		GlutenFree               bool    `json:"glutenFree"`
		DairyFree                bool    `json:"dairyFree"`
		VeryHealthy              bool    `json:"veryHealthy"`
		Cheap                    bool    `json:"cheap"`
		VeryPopular              bool    `json:"veryPopular"`
		Sustainable              bool    `json:"sustainable"`
		WeightWatcherSmartPoints int     `json:"weightWatcherSmartPoints"`
		Gaps                     string  `json:"gaps"`
		LowFodmap                bool    `json:"lowFodmap"`
		PreparationMinutes       int     `json:"preparationMinutes,omitempty"`
		CookingMinutes           int     `json:"cookingMinutes,omitempty"`
		AggregateLikes           int     `json:"aggregateLikes"`
		SpoonacularScore         float64 `json:"spoonacularScore"`
		HealthScore              float64 `json:"healthScore"`
		CreditsText              string  `json:"creditsText"`
		SourceName               string  `json:"sourceName"`
		PricePerServing          float64 `json:"pricePerServing"`
		ExtendedIngredients      []struct {
			ID              int           `json:"id"`
			Aisle           string        `json:"aisle"`
			Image           string        `json:"image"`
			Consistency     string        `json:"consistency"`
			Name            string        `json:"name"`
			NameClean       string        `json:"nameClean"`
			Original        string        `json:"original"`
			OriginalString  string        `json:"originalString"`
			OriginalName    string        `json:"originalName"`
			Amount          float64       `json:"amount"`
			Unit            string        `json:"unit"`
			Meta            []interface{} `json:"meta"`
			MetaInformation []interface{} `json:"metaInformation"`
			Measures        struct {
				Us struct {
					Amount    float64 `json:"amount"`
					UnitShort string  `json:"unitShort"`
					UnitLong  string  `json:"unitLong"`
				} `json:"us"`
				Metric struct {
					Amount    float64 `json:"amount"`
					UnitShort string  `json:"unitShort"`
					UnitLong  string  `json:"unitLong"`
				} `json:"metric"`
			} `json:"measures"`
		} `json:"extendedIngredients"`
		ID                   int           `json:"id"`
		Title                string        `json:"title"`
		ReadyInMinutes       int           `json:"readyInMinutes"`
		Servings             int           `json:"servings"`
		SourceURL            string        `json:"sourceUrl"`
		Image                string        `json:"image,omitempty"`
		ImageType            string        `json:"imageType,omitempty"`
		Summary              string        `json:"summary"`
		Cuisines             []string      `json:"cuisines"`
		DishTypes            []string      `json:"dishTypes"`
		Diets                []interface{} `json:"diets"`
		Occasions            []interface{} `json:"occasions"`
		Instructions         string        `json:"instructions"`
		AnalyzedInstructions []struct {
			Name  string `json:"name"`
			Steps []struct {
				Number      int    `json:"number"`
				Step        string `json:"step"`
				Ingredients []struct {
					ID            int    `json:"id"`
					Name          string `json:"name"`
					LocalizedName string `json:"localizedName"`
					Image         string `json:"image"`
				} `json:"ingredients"`
				Equipment []struct {
					ID            int    `json:"id"`
					Name          string `json:"name"`
					LocalizedName string `json:"localizedName"`
					Image         string `json:"image"`
				} `json:"equipment"`
				Length struct {
					Number int    `json:"number"`
					Unit   string `json:"unit"`
				} `json:"length"`
			} `json:"steps"`
		} `json:"analyzedInstructions"`
		OriginalID           interface{} `json:"originalId"`
		SpoonacularSourceURL string      `json:"spoonacularSourceUrl"`
		License              string      `json:"license,omitempty"`
	} `json:"recipes"`
}