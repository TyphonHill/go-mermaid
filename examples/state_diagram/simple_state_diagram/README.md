```mermaid
---
title: Simple State Machine
---

stateDiagram-v2
	state "Idle State" as idle
	note left of idle: System waiting for new orders
	state "Processing" as processing
	state "Error" as error
	idle --> processing: Start Process
	processing --> idle: Complete
	processing --> error: Error Occurred
	error --> idle: Reset

```
