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
				c.log.Error(fmt.Sprintf("unable to update scholarship to registration status %#v \n", err))
			}
		}
	})

	// review status checker
	c.scheduler.Every(d).Seconds().Do(func(){
		ids , err := c.scholarshipSvc.ReviewStatusScheduler()
		if err != nil {
			c.log.Error(fmt.Sprintf("unable to get review scholarship status %#v \n", err))
		}
		for _, id := range ids {
			if err := c.scholarshipSvc.UpdateScholarshipStatus(int(util.REVIEW), id); err != nil {
				c.log.Error(fmt.Sprintf("unable to update scholarship to review status %#v \n", err))
			}
		}
	})

	// announcement status checker
	c.scheduler.Every(d).Seconds().Do(func(){
		ids , err := c.scholarshipSvc.AnnouncementStatusScheduler()
		if err != nil {
			c.log.Error(fmt.Sprintf("unable to get announcement scholarship status %#v \n", err))
		}
		for _, id := range ids {
			if err := c.scholarshipSvc.UpdateScholarshipStatus(int(util.ANNOUNCEMENT), id); err != nil {
				c.log.Error(fmt.Sprintf("unable to update scholarship to announcement status %#v \n", err))
			}
		}
	})

	// funding status checker
	c.scheduler.Every(d).Seconds().Do(func(){
		ids , err := c.scholarshipSvc.FundingStatusScheduler()
		if err != nil {
			c.log.Error(fmt.Sprintf("unable to get funding scholarship status %#v \n", err))
		}
		for _, id := range ids {
			if err := c.scholarshipSvc.UpdateScholarshipStatus(int(util.FUNDING), id); err != nil {
				c.log.Error(fmt.Sprintf("unable to update scholarship to funding status %#v \n", err))
			}
		}
	})
	// finish status checker
	c.scheduler.Every(d).Seconds().Do(func(){
		ids , err := c.scholarshipSvc.FinishStatusScheduler()
		if err != nil {
			c.log.Error(fmt.Sprintf("unable to get finish scholarship status %#v \n", err))
		}
		for _, id := range ids {
			if err := c.scholarshipSvc.UpdateScholarshipStatus(int(util.FINISH), id); err != nil {
				c.log.Error(fmt.Sprintf("unable to update scholarship to finish status %#v \n", err))
			}
		}
	})

	c.scheduler.StartAsync()
	return nil
}
