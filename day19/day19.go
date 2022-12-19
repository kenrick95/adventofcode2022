package main

import (
	"fmt"
	"regexp"
	"strconv"
	"utils"
)

type Blueprint struct {
	id               int
	producers        map[ResourceType]ResourceProducer
	initialResources map[ResourceType]int
}
type ResourceProducer struct {
	resourceType ResourceType
	quantity     int
	costs        map[ResourceType]int
}
type ResourceType int

const (
	Ore = iota + 1
	Clay
	Obsidian
	Geode
)

type State struct {
	oreQty     int
	oreProdQty int

	clayQty     int
	clayProdQty int

	obsQty     int
	obsProdQty int

	geoQty     int
	geoProdQty int

	timeLeft int
}

// TODO: Not solved yet... seems like wrong approach

func main() {
	lines := utils.ReadFileToLines("day19.in")
	re, _ := regexp.Compile(`Blueprint (\d+): Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian.`)
	blueprintMap := map[int]Blueprint{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		matches := re.FindStringSubmatch(line)
		blueprintIndex, _ := strconv.Atoi(matches[1])
		oreRobotCostOre, _ := strconv.Atoi(matches[2])
		clayRobotCostOre, _ := strconv.Atoi(matches[3])
		obsidianRobotCostOre, _ := strconv.Atoi(matches[4])
		obsidianRobotCostClay, _ := strconv.Atoi(matches[5])
		geodeRobotCostOre, _ := strconv.Atoi(matches[6])
		geodeRobotCostObsidian, _ := strconv.Atoi(matches[7])
		blueprintMap[blueprintIndex] = Blueprint{
			id: blueprintIndex,
			producers: map[ResourceType]ResourceProducer{
				Ore: {
					resourceType: Ore,
					quantity:     1,
					costs: map[ResourceType]int{
						Ore:      oreRobotCostOre,
						Clay:     0,
						Obsidian: 0,
						Geode:    0,
					},
				},
				Clay: {
					resourceType: Clay,
					quantity:     0,
					costs: map[ResourceType]int{
						Ore:      clayRobotCostOre,
						Clay:     0,
						Obsidian: 0,
						Geode:    0,
					},
				},
				Obsidian: {
					resourceType: Obsidian,
					quantity:     0,
					costs: map[ResourceType]int{
						Ore:      obsidianRobotCostOre,
						Clay:     obsidianRobotCostClay,
						Obsidian: 0,
						Geode:    0,
					},
				},
				Geode: {
					resourceType: Geode,
					quantity:     0,
					costs: map[ResourceType]int{
						Ore:      geodeRobotCostOre,
						Clay:     0,
						Obsidian: 0,
						Geode:    geodeRobotCostObsidian,
					},
				},
			},
			initialResources: map[ResourceType]int{
				Ore:      0,
				Clay:     0,
				Obsidian: 0,
				Geode:    0,
			},
		}
	}

	getMaxGeode := func(blueprint Blueprint) int {
		memo := map[State]int{}
		producerTypes := []ResourceType{
			Ore,
			Clay,
			Obsidian,
			Geode,
		}
		println("blueprint", blueprint.id)
		for _, producerType := range producerTypes {
			println(producerType, "Ore", blueprint.producers[producerType].costs[Ore], "Clay", blueprint.producers[producerType].costs[Clay], "Obs", blueprint.producers[producerType].costs[Obsidian], "Geo", blueprint.producers[producerType].costs[Geode])
		}

		canAfford := func(currentResources map[ResourceType]int, producerType ResourceType) bool {
			for resource, resourceQty := range currentResources {
				if resourceQty < blueprint.producers[producerType].costs[resource] {
					return false
				}
			}
			return true
		}

		cloneMap := func(prevMap map[ResourceType]int) map[ResourceType]int {
			newMap := map[ResourceType]int{
				Ore:      prevMap[Ore],
				Clay:     prevMap[Clay],
				Obsidian: prevMap[Obsidian],
				Geode:    prevMap[Geode],
			}
			return newMap
		}

		deductResources := func(currentResources map[ResourceType]int, producerType ResourceType) map[ResourceType]int {
			newResources := cloneMap(currentResources)
			for resource, resourceQty := range currentResources {
				newResources[resource] = resourceQty - blueprint.producers[producerType].costs[resource]
			}
			return newResources
		}

		var check func(resourceQuantity map[ResourceType]int, producerQuantity map[ResourceType]int, timeLeft int) int
		check = func(resourceQuantity map[ResourceType]int, producerQuantity map[ResourceType]int, timeLeft int) int {
			fmt.Println(timeLeft, "Ore", resourceQuantity[Ore], producerQuantity[Ore], "Clay", resourceQuantity[Clay], producerQuantity[Clay], "Obs", resourceQuantity[Obsidian], producerQuantity[Obsidian], "Geo", resourceQuantity[Geode], producerQuantity[Geode])
			if timeLeft == 0 {
				return resourceQuantity[Geode]
			}

			currentState := State{
				oreQty:      resourceQuantity[Ore],
				oreProdQty:  producerQuantity[Ore],
				clayQty:     resourceQuantity[Clay],
				clayProdQty: producerQuantity[Clay],
				obsQty:      resourceQuantity[Obsidian],
				obsProdQty:  producerQuantity[Obsidian],
				geoQty:      resourceQuantity[Geode],
				geoProdQty:  producerQuantity[Geode],
				timeLeft:    timeLeft,
			}
			memoValue, memoExist := memo[currentState]
			if memoExist {
				// fmt.Printf("currentState: %v %v\n", currentState, memoValue)
				return memoValue
			}

			var res = 0

			nextResourceQuantity := map[ResourceType]int{
				Ore:      resourceQuantity[Ore] + producerQuantity[Ore],
				Clay:     resourceQuantity[Clay] + producerQuantity[Clay],
				Obsidian: resourceQuantity[Obsidian] + producerQuantity[Obsidian],
				Geode:    resourceQuantity[Geode] + producerQuantity[Geode],
			}

			// Do nothing
			res = utils.Max(res, check(nextResourceQuantity, producerQuantity, timeLeft-1))

			for _, producerType := range producerTypes {
				if canAfford(resourceQuantity, producerType) {
					// fmt.Println(timeLeft, "producing", producerType, "Ore", resourceQuantity[Ore], producerQuantity[Ore], "Clay", resourceQuantity[Clay], producerQuantity[Clay], "Obs", resourceQuantity[Obsidian], producerQuantity[Obsidian], "Geo", resourceQuantity[Geode], producerQuantity[Geode])
					nextProducerQuantity := cloneMap(producerQuantity)
					nextProducerQuantity[producerType] += 1
					nextResourceQuantityAfterProducing := deductResources(nextResourceQuantity, producerType)

					res = utils.Max(res, check(nextResourceQuantityAfterProducing, nextProducerQuantity, timeLeft-1))
				}
			}
			memo[currentState] = res

			return res
		}

		return check(map[ResourceType]int{
			Ore:      0,
			Clay:     0,
			Obsidian: 0,
			Geode:    0,
		}, map[ResourceType]int{
			Ore:      1,
			Clay:     0,
			Obsidian: 0,
			Geode:    0,
		}, 24)

	}

	{
		blueprint := blueprintMap[1]
		res := getMaxGeode(blueprint)
		println("res", blueprint.id, res)
	}

	// ansPart1 := 0
	// for idx, blueprint := range blueprintMap {
	// 	res := getMaxGeode(blueprint)
	// 	println("res", blueprint.id, res)
	// 	ansPart1 += idx * res
	// }
	// println("ansPart1", ansPart1)

}
