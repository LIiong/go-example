package main

import (
	"fmt"
	"math"
)
/*
狄克斯特拉算法
(1)找出“最便宜”的节点，即可在最短时间内到达的节点。
(2)更新该节点的邻居的开销，其含义将稍后介绍。
(3)重复这个过程，直到对图中的每个节点都这样做了。
(4)计算最终路径。
*/
func dijkstras() {
	graph := make(map[string]map[string]int)
	graph["start"] = map[string]int{}
	graph["start"]["a"] = 6
	graph["start"]["b"] = 2

	graph["a"] = map[string]int{}
	graph["a"]["finish"] = 1

	graph["b"] = map[string]int{}
	graph["b"]["a"] = 3
	graph["b"]["finish"] = 5

	graph["finish"] = map[string]int{}

	costs, parents := findShortestPath(graph, "start", "finish")
	fmt.Println(costs, parents)
}

// Finds shortest path using dijkstra algorithm
func findShortestPath(graph map[string]map[string]int, startNode string, finishNode string) (map[string]int, map[string]string)  {
	costs := make(map[string]int)
	costs[finishNode] = math.MaxInt32

	parents := make(map[string]string)
	parents[finishNode] = ""

	processed := make(map[string]bool)

	// Initialization of costs and parents
	for node, cost := range graph[startNode] {
		costs[node] = cost
		parents[node] = startNode
	}

	lowestCostNode := findLowestCostNode(costs, processed)
	for ; lowestCostNode != "" ; {
		// Calculation costs for neighbours
		for node, cost := range graph[lowestCostNode] {
			newCost := costs[lowestCostNode] + cost
			if newCost < costs[node] {
				// Set new cost for this node
				costs[node] = newCost
				parents[node] = lowestCostNode
			}
		}

		processed[lowestCostNode] = true
		lowestCostNode = findLowestCostNode(costs, processed)
	}

	return costs, parents
}

func findLowestCostNode(costs map[string]int, processed map[string]bool) string {
	lowestCost := math.MaxInt32
	lowestCostNode := ""
	for node, cost := range costs {
		if _, inProcessed := processed[node]; cost < lowestCost && !inProcessed {
			lowestCost = cost
			lowestCostNode = node
		}
	}

	return lowestCostNode
}