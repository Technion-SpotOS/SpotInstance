# SpotInstance
Spot-Instance Representation and Management Component of the Technion's SpotOS Project

This repository contains the API for SpotInstance CRD and the set of its managing controllers. 

### SpotInstance CR
A SpotIstance CR defines a spot instance configuration as outputted from the [CloudCostOptimizer](https://github.com/AdiY10/CloudCostOptimizer), and reflects its utilization in status when the spot-instance begins serving as a cluster node.

### SpotIstance Controllers
The controllers are responsible for:
- Ordering the specified spot instance
- Installing the acquired spot instance as a cluster node
    - Including tainting and labeling properly
- Updating CR status with assigned node name

TODO: enhance docs, add examples