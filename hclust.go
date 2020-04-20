// Package hclust contains methods for performing agglomerative hierarchical clustering.
package hclust

import (
	"github.com/ivan-pindrop/hclust/cluster"
	"github.com/ivan-pindrop/hclust/dendrogram"
	"github.com/ivan-pindrop/hclust/distance"
	"github.com/ivan-pindrop/hclust/optimize"
	"github.com/ivan-pindrop/hclust/sort"
	"github.com/ivan-pindrop/hclust/tree"
	"github.com/ivan-pindrop/hclust/typedef"
)

// Cluster references the main cluster method in the cluster subpackage.
var Cluster = cluster.Cluster

// Dendrogram is an array of SubClusters.
type Dendrogram []SubCluster

// Distance references the main distance method in the distance subpackage.
var Distance = distance.Distance

// GetNodeHeights gets the height for each dendrogram node by summing child branch lengths.
var GetNodeHeight = dendrogram.GetNodeHeight

// Optimize references the main leaf optimization method in the optimize subpackage.
var Optimize = optimize.Optimize

// Sort references the main sort method in the sort subpackage
var Sort = sort.Sort

// SubCluster stores the node, distance and names of leafs for a subcluster.
type SubCluster = typedef.SubCluster

// TreeLayout contains a tree in newick format and the leaf order.
type TreeLayout = tree.Tree

// Tree references the main method for generating the newick tree in the tree subpackage.
var Tree = tree.Create
