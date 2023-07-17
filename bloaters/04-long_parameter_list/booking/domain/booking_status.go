package domain

type BookingStatus string

const (
	BookingStatusActive   BookingStatus = "active"
	BookingStatusFinished BookingStatus = "finished"
	BookingStatusNotStart BookingStatus = "not_start"
)
