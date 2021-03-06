package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccBigqueryReservationReservation_bigqueryReservationUpdate(t *testing.T) {
	t.Parallel()

	location := "asia-northeast1"
	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
		"location":      location,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBigqueryReservationReservationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccBigqueryReservationReservation_bigqueryReservationBasic(context),
			},
			{
				ResourceName:      "google_bigquery_reservation.reservation",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccBigqueryReservationReservation_bigqueryReservationUpdate(context),
			},
			{
				ResourceName:      "google_bigquery_reservation.reservation",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccBigqueryReservationReservation_bigqueryReservationBasic(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigquery_reservation" "reservation" {
	name           = "reservation%{random_suffix}"
	location       = "%{location}"
	// Set to 0 for testing purposes
	// In reality this would be larger than zero
	slot_capacity  = 0
	ignore_idle_slots = true
}
`, context)
}

func testAccBigqueryReservationReservation_bigqueryReservationUpdate(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigquery_reservation" "reservation" {
	name           = "reservation%{random_suffix}"
	location       = "%{location}"
	// Set to 0 for testing purposes
	// In reality this would be larger than zero
	slot_capacity  = 0
	ignore_idle_slots = false
}
`, context)
}
