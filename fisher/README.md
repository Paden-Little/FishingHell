Fish loose value per day after they are caught based on the fish. You can stop this by deciding to freeze the fish but the fish instantly looses a flat amount of total value because it has been frozen. There is a limited number of frozen "slots" in your inventory   


### Fish Weight

Fish weight is determined by multiple factors. 

#####

baseWeight is the lowest possible weight amount a fish can have
weightRarity is the rarity of weight disparity 

The algorithm functions as a larger the rarer categorization of fish where, depending on its weight disparity, the fish can theoreretically be infinitely large, however a programattic cap must be made for safety. It functions on rolls, where each rolls success is rarer than the last, upon failing a roll the weight of the fish is decided. Each roll also have dimmnishing weight returns where, each roll reduce the amount of weight added for a successful roll. So the rarirty of a successful roll, and a successful rolls weight addition are inversly proportional. Creating the ability for extreme outliers and world records, without flooding the market with multiple insanely large fish.

Upon receiving a fish the weight will be compared to all other weights of that fish and will be given a weight class based on that using percentile units. 

func DetermineWeight(baseWeight, weightRarity) -> weight{
	float weight;
	while(true){
		if(){
			
		}
		weight = weight + 
	}
}