package scheduler

import (
	"fmt"
	"github.com/Nusantara-Muda/scholarship-api/src/business/util"
)

func (c *sch) ScholarshipStatusChecker(d int) error {
	// registration status checker
	c.scheduler.Every(d).Seconds().Do(func(){
		ids , err := c.scholarshipSvc.RegistrationStatusScheduler()
		if err != nil {
			c.log.Error(fmt.Sprintf("unable to get registration scholarship status %#v \n", err))
		}
		for _, id := range ids {
			if err := c.scholarshipSvc.UpdateScholarshipStatus(int(util.REGISTRATION), id); err != nil {
				c.log.Error(fmt.Sprintf("unable to update scholarship status to registration %#v \n", err))
			}
		}
	})

	// TODO: review status checker
	// TODO: announcement status checker
	// TODO: funding status checker
	// TODO: finish status checker

	c.scheduler.StartAsync()
	return nil
}
