```mermaid
journey
	title E-Commerce User Journey
	section Discovery
		See Social Media Ad: 3: Customer,Marketing
		Visit Website: 5: Customer
		Browse Categories: 4: Customer
		Read Product Reviews: 4: Customer,Previous Buyers
	section Account
		Create Account: 2: Customer,System
		Verify Email: 3: Customer,System
		Complete Profile: 3: Customer
		Save Payment Info: 4: Customer,Payment Provider
	section Shopping
		Search Products: 4: Customer
		Apply Filters: 3: Customer
		Compare Items: 4: Customer
		Add to Cart: 5: Customer,System
		Adjust Quantities: 4: Customer,System
	section Checkout
		Review Cart: 5: Customer
		Apply Coupon: 4: Customer,System
		Select Shipping: 3: Customer,Shipping Provider
		Complete Payment: 4: Customer,Payment Provider
		Receive Confirmation: 5: Customer,System
	section Post-Purchase
		Track Order: 5: Customer,Shipping Provider
		Receive Updates: 4: Customer,System,Shipping Provider
		Get Delivery: 5: Customer,Delivery Driver
		Rate Purchase: 3: Customer
		Contact Support: 2: Customer,Support Team
	section Support
		Find FAQ: 3: Customer
		Chat with Bot: 2: Customer,Chat Bot
		Email Support: 3: Customer,Support Team
		Live Chat: 4: Customer,Support Agent
		Resolution: 4: Customer,Support Agent

```
