package billing

import (
	"time"

	"hcc/piano/lib/logger"

	errors "innogrid.com/hcloud-classic/hcc_errors"
)

var DriverBilling = &Billing{
	lastUpdate:  time.Now(),
	updateTimer: nil,
	StopTimer:   nil,
}

func reserveRegisterUpdateTimer() {
	defer DriverBilling.UpdateBillingInfo()
	logger.Logger.Println("Register billing info update timer")

	now := time.Now().Add(1 * time.Hour)
	<-time.After(time.Until(time.Date(now.Year(), now.Month(), now.Day(),
		now.Hour(), 0, 0, 0, now.Location())))

	DriverBilling.RunUpdateTimer()
}

func Init() *errors.HccError {
	logger.Logger.Println("Update billing info in boot up time")
	DriverBilling.UpdateBillingInfo()

	go reserveRegisterUpdateTimer()

	return nil
}

func End() {
	if DriverBilling.StopTimer != nil {
		DriverBilling.StopTimer()
	}
}
