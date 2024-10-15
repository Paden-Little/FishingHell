**Fisher**  
The Fisher service is the service that facilitates catching of fish. Users begin a fish and the web app will notify the users whenever they have a pull. Upon a fish being hooked, the user has a short period of time where they can reel the fish in. When the fish is caught, it is added to the user's inventory.  
**Fisher database:**  
	Objects:

- Fish:  
  - ID  
  - Name – string  
  - Taxonomy – string  
  - Low weight – float  
  - High weight – float  
  - Peak – float nullable  
  - Stddev – float nullable  
- Pool:  
  - Name – unique string  

[//]: # (- Bait:  )

[//]: # (  - ID – unique string  )

[//]: # (  - Name – string  )

[//]: # (  - Cost – float  )

[//]: # (  - Description – string  )
- Attraction:  
  - ID – composite unique key between fishID and shop.baitID  
  - Value – float  
- CatchValue:  
  - ID – composite unique key between poolID and fishID   
  - Value – float  
- FishCaught:   
  - ID  
  - FishID – Foreign Key  
  - Weight – float (kg)  
  - 

**Fish Selection:**  
After the user casts a line, per **tick**, there is a chance for a fish to be drawn. The chance for a fish to be drawn is based on the **Pool** the user casted into modified by the type of **bait** the user is using. Upon selection the fish is generated.

- **Attraction**: Per each type of bait in existence, each type of fish must have an attraction value. The attraction value is a percentage, positive or negative, which will modify the chances of that fish being caught by adding that value to the fishes base catch value  
- **Pool:** Before casting a line the user selects which pool they would like to use. Pool influences what fish are able to be caught and at what chance per fish. Currently a static value but stretch goal is to have it fluctuate based on what has been caught previously (population simulation)  
- **Tick:** Tick is determined by what rod the user is using. Flat value per second.  
- **Bait:** Bait is a consumable item added to the rod, bait is required to fish. Bait modifies the **Catch Value**. Bait adds \+/– a percentage value to the catch value  
- **Catch Value:** Percentage chance of a fish being present in a given pool, will vary from pool to pool. Modified by **Bait**.

**Fish generation:**  
Upon a fish being caught, we pull the generation stats of the fish and generate the fish based on those stats.

- **Low weight:** Determines the lowest possible value the fish can be  
- **High weight:** Determines the high possible value the fish can be  
- **Peak** (Optional): Determines the most common weight value wanted  
- **Spread** (Optional): Determines the rarity of values distances from the peak

 Once a fish is generated, give the user a note of the fish's weight class as a percentile.

- A request from the Account service will run once an hour and cache the result of the current percentiles of the weights of all fish. Compare the weight of the new fish to that and present that to the user, but don’t store it. Percentile of fish weight is a derived value

**Fish Storage:**  
The user now has a choice to store the fish in the **ice chest**, **the freezer**, or immediately **sell it**. If the user selects to put it in the Ice chest or freezer, the fish object will be sent to the account service for storage. While if the user sells it, it will simply send the amount of money to add to the users account.

- **Freezer:**  
  - User can decide to freeze the fish, the fish will last forever and never spoil, but the fish will immediately take a cut in price (35%)  
- **Ice Chest:**  
  - Users can decide to put the ice in the ice chest where, each day it will lower in value before eventually spoiling after 5 days.  
  - Value lost per day is 7%  
- **Sell it:**  
  - Selling it will show the current market value per Kg and the total sell price based on the fishes weight


## Presentation Ideas
"The application can withstand X users so, if each and everyone of you were you recommend 3 people and then those 3 people reccomend 3 more, I would be able to handle that computation. But dont go anyfurther or we start breaking down"