package schedulers

import (
	dbmodels "billing-engine/internal/models/db"
	"billing-engine/internal/repositories"
	"billing-engine/internal/utils"
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	loanRepo    *repositories.LoanRepository
	billingRepo *repositories.BillingRepository
}

func NewScheduler(
	loanRepo *repositories.LoanRepository,
	billingRepo *repositories.BillingRepository,
) *Scheduler {
	return &Scheduler{
		loanRepo:    loanRepo,
		billingRepo: billingRepo,
	}
}

func (s *Scheduler) StartScheduler() {
	c := cron.New()

	c.AddFunc("@daily", func() {
		loans, err := s.loanRepo.GetAllLoans()
		if err != nil {
			fmt.Println("Error fetching loans", err.Error())
			return
		}

		for _, loan := range loans {
			fmt.Printf("Processing loan %d\n", loan.ID)

			billings, err := s.billingRepo.GetBillingsByLoadID(loan.ID)
			if err != nil {
				fmt.Println("Error fetching billings", err.Error())
				continue
			}

			missedPayments := 0
			for _, billing := range billings {
				if !billing.IsPaid {
					missedPayments++
				}
			}

			if missedPayments > 3 {
				loan.IsDelinquent = true
				s.loanRepo.UpdateLoan(&loan)
			}
		}
	})

	c.AddFunc("@daily", func() {

		loans, err := s.loanRepo.GetAllLoans()
		if err != nil {
			fmt.Println("Error fetching loans", err.Error())
			return
		}

		for _, loan := range loans {
			fmt.Printf("Processing loan %d\n", loan.ID)

			billings, err := s.billingRepo.GetBillingsByLoadID(loan.ID)
			if err != nil {
				fmt.Println("Error fetching billings", err.Error())
				continue
			}

			today := time.Now()
			if len(billings) == 0 {
				difference := utils.CalculateDifferenceWeeksFloat(loan.StartDate, today)
				if difference == 1.0 {
					billing := dbmodels.Billing{
						LoanID:    loan.ID,
						Amount:    loan.WeeklyPayment,
						DueDate:   today.AddDate(0, 0, 7),
						Week:      1,
						CreatedAt: today,
					}
					s.billingRepo.CreateBilling(&billing)
				}
			} else {
				lastBilling := billings[len(billings)-1]
				aWeekLater := lastBilling.DueDate.AddDate(0, 0, 7)
				difference := utils.CalculateDifferenceWeeksFloat(aWeekLater, today)
				estimated := today.Sub(loan.EndDate)
				if difference == 1.0 && estimated > 0 {
					billing := dbmodels.Billing{
						LoanID:    loan.ID,
						Amount:    loan.WeeklyPayment,
						DueDate:   today.AddDate(0, 0, 7),
						CreatedAt: today,
						Week:      lastBilling.Week + 1,
					}
					s.billingRepo.CreateBilling(&billing)
				}
			}
		}
	})
	c.Start()
}
