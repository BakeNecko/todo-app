package mdb

type Option func(db *Mongo)

// Some options for Mongo
// Example:
//func SomeOpt(value int) Option {
//	return func(c *Mongo) {
//		c.<field_name> = value
//	}
//}
