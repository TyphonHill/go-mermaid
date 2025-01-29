```mermaid
---
title: Simple Login Flow
---

flowchart TB
	0("Start")
	1("Login Form")
	2("Validate")
	3("Success")
	4("Failure")
	5("End")
	0 --> 1
	1 --> 2
	2 -->|Valid| 3
	2 -->|Invalid| 4
	3 --> 5
	4 --> 1
```
