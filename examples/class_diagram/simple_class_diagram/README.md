```mermaid
---
title: Simple Class Diagram
---

classDiagram
	direction TB
	class User{
		+string id
		+string name
		+string email
		+login(password:string) bool
		+logout() 
	}
	class Account{
		+string id
		+float balance
		+deposit(amount:float) 
		+withdraw(amount:float) error
	}
	User -- Account : has

```
