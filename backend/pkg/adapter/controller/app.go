package controller

type AppController struct {
	User     interface{ User }
	Merchant interface{ Merchant }
}
