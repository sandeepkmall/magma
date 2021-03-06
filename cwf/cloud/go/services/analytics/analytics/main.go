/*
 * Copyright 2020 The Magma Authors.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"flag"
	"magma/cwf/cloud/go/services/analytics/calculations"
	"strconv"
	"time"

	"magma/cwf/cloud/go/cwf"
	"magma/cwf/cloud/go/services/analytics"
	"magma/orc8r/cloud/go/orc8r"
	"magma/orc8r/cloud/go/service"
	"magma/orc8r/cloud/go/services/metricsd"
	"magma/orc8r/lib/go/metrics"
	"magma/orc8r/lib/go/service/config"

	"github.com/golang/glog"
	promAPI "github.com/prometheus/client_golang/api"
	"github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	ServiceName = "ANALYTICS"

	activeUsersMetricName           = "active_users_over_time"
	userThroughputMetricName        = "user_throughput"
	userConsumptionMetricName       = "user_consumption"
	userConsumptionHourlyMetricName = "user_consumption_hourly"
	apThroughputMetricName          = "throughput_per_ap"
	authenticationsMetricName       = "authentications_over_time"

	defaultAnalysisSchedule = "0 */12 * * *" // Every 12 hours
)

var (
	// Map from number of days to query to size the step should be to get best granularity
	// without causes prometheus to reject the query for having too many datapoints
	daysToQueryStepSize = map[int]time.Duration{1: 15 * time.Second, 7: time.Minute, 30: 5 * time.Minute}

	daysToCalculate = []int{1, 7, 30}
)

func main() {
	flag.Parse()

	// Create the service
	srv, err := service.NewOrchestratorService(cwf.ModuleName, ServiceName)
	if err != nil {
		glog.Fatalf("Error creating CWF Analytics service: %s", err)
	}

	analysisSchedule := defaultAnalysisSchedule
	providedSchedule, _ := srv.Config.GetString("analysisSchedule")
	if providedSchedule != "" {
		analysisSchedule = providedSchedule
	}

	calcs := getAnalyticsCalculations()
	promAPIClient := getPrometheusClient()
	exporter := getExporterIfRequired(srv)

	analyzer := analytics.NewPrometheusAnalyzer(promAPIClient, calcs, exporter)
	err = analyzer.Schedule(analysisSchedule)
	if err != nil {
		glog.Fatalf("Error scheduling analyzer: %s", err)
	}

	go analyzer.Run()

	// Run the service
	err = srv.Run()
	if err != nil {
		glog.Fatalf("Error running service: %s", err)
	}
}

func getExporterIfRequired(srv *service.OrchestratorService) analytics.Exporter {
	shouldExportData, _ := srv.Config.GetBool("exportMetrics")
	if shouldExportData {
		glog.Infof("Creating CWF Analytics Exporter")
		return analytics.NewWWWExporter(
			srv.Config.MustGetString("metricsPrefix"),
			srv.Config.MustGetString("appSecret"),
			srv.Config.MustGetString("appID"),
			srv.Config.MustGetString("metricExportURL"),
			srv.Config.MustGetString("categoryName"),
		)
	}
	return nil
}

var (
	xapLabels                   = []string{calculations.DaysLabel, metrics.NetworkLabelName}
	userThroughputLabels        = []string{calculations.DaysLabel, metrics.NetworkLabelName, calculations.DirectionLabel}
	userConsumptionLabels       = []string{calculations.DaysLabel, metrics.NetworkLabelName, calculations.DirectionLabel}
	hourlyUserConsumptionLabels = []string{"hours", metrics.NetworkLabelName, calculations.DirectionLabel}
	apThroughputLabels          = []string{calculations.DaysLabel, metrics.NetworkLabelName, calculations.DirectionLabel, calculations.APNLabel}
	authenticationsLabels       = []string{calculations.DaysLabel, metrics.NetworkLabelName, calculations.AuthCodeLabel}
)

func getAnalyticsCalculations() []calculations.Calculation {
	xapGauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: activeUsersMetricName}, xapLabels)
	userThroughputGauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: userThroughputMetricName}, userThroughputLabels)
	userConsumptionGauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: userConsumptionMetricName}, userConsumptionLabels)
	hourlyUserConsumptionGauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: userConsumptionHourlyMetricName}, hourlyUserConsumptionLabels)
	apThroughputGauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: apThroughputMetricName}, apThroughputLabels)
	authenticationsGauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: authenticationsMetricName}, authenticationsLabels)

	prometheus.MustRegister(xapGauge, userThroughputGauge, userConsumptionGauge,
		hourlyUserConsumptionGauge, apThroughputGauge, authenticationsGauge)

	allCalculations := make([]calculations.Calculation, 0)

	// MAP, WAP, DAP Calculations
	allCalculations = append(allCalculations, getXAPCalculations(daysToCalculate, xapGauge, activeUsersMetricName)...)

	// User Throughput Calculations
	allCalculations = append(allCalculations, getUserThroughputCalculations(daysToCalculate, userThroughputGauge, userThroughputMetricName)...)

	// AP Throughput Calculations
	allCalculations = append(allCalculations, getAPThroughputCalculations(daysToCalculate, apThroughputGauge, apThroughputMetricName)...)

	// User Consumption Calculations
	allCalculations = append(allCalculations, getUserConsumptionCalculations(daysToCalculate, userConsumptionGauge, userConsumptionMetricName)...)
	allCalculations = append(allCalculations, get1hourConsumptionCalculation(hourlyUserConsumptionGauge, userConsumptionHourlyMetricName)...)

	// Authentication Calculations
	allCalculations = append(allCalculations, getAuthenticationCalculations(daysToCalculate, authenticationsGauge, authenticationsMetricName)...)

	return allCalculations
}

func getXAPCalculations(daysList []int, gauge *prometheus.GaugeVec, metricName string) []calculations.Calculation {
	calcs := make([]calculations.Calculation, 0)
	for _, dayParam := range daysList {
		calcs = append(calcs, &calculations.XAPCalculation{
			CalculationParams: calculations.CalculationParams{
				Days:                dayParam,
				RegisteredGauge:     gauge,
				Labels:              prometheus.Labels{calculations.DaysLabel: strconv.Itoa(dayParam)},
				Name:                metricName,
				ExpectedGaugeLabels: xapLabels,
			},
		})
	}
	return calcs
}

func getUserThroughputCalculations(daysList []int, gauge *prometheus.GaugeVec, metricName string) []calculations.Calculation {
	calcs := make([]calculations.Calculation, 0)
	for _, dayParam := range daysList {
		for _, dir := range []calculations.ConsumptionDirection{calculations.ConsumptionIn, calculations.ConsumptionOut} {
			calcs = append(calcs, &calculations.UserThroughputCalculation{
				CalculationParams: calculations.CalculationParams{
					Days:                dayParam,
					RegisteredGauge:     gauge,
					Labels:              prometheus.Labels{calculations.DaysLabel: strconv.Itoa(dayParam)},
					Name:                metricName,
					ExpectedGaugeLabels: userThroughputLabels,
				},
				Direction:     dir,
				QueryStepSize: daysToQueryStepSize[dayParam],
			})
		}
	}
	return calcs
}

func getAPThroughputCalculations(daysList []int, gauge *prometheus.GaugeVec, metricName string) []calculations.Calculation {
	calcs := make([]calculations.Calculation, 0)
	for _, dayParam := range daysList {
		for _, dir := range []calculations.ConsumptionDirection{calculations.ConsumptionIn, calculations.ConsumptionOut} {
			calcs = append(calcs, &calculations.APThroughputCalculation{
				CalculationParams: calculations.CalculationParams{
					Days:                dayParam,
					RegisteredGauge:     gauge,
					Labels:              prometheus.Labels{calculations.DaysLabel: strconv.Itoa(dayParam)},
					Name:                metricName,
					ExpectedGaugeLabels: apThroughputLabels,
				},
				Direction:     dir,
				QueryStepSize: daysToQueryStepSize[dayParam],
			})
		}
	}
	return calcs
}

func getUserConsumptionCalculations(daysList []int, gauge *prometheus.GaugeVec, metricName string) []calculations.Calculation {
	calcs := make([]calculations.Calculation, 0)
	for _, dayParam := range daysList {
		for _, dir := range []calculations.ConsumptionDirection{calculations.ConsumptionIn, calculations.ConsumptionOut} {
			calcs = append(calcs, &calculations.UserConsumptionCalculation{
				CalculationParams: calculations.CalculationParams{
					Days:                dayParam,
					RegisteredGauge:     gauge,
					Labels:              prometheus.Labels{calculations.DaysLabel: strconv.Itoa(dayParam)},
					Name:                metricName,
					ExpectedGaugeLabels: userConsumptionLabels,
				},
				Direction: dir,
			})
		}
	}
	return calcs
}

func get1hourConsumptionCalculation(gauge *prometheus.GaugeVec, metricName string) []calculations.Calculation {
	calcs := make([]calculations.Calculation, 0)
	for _, dir := range []calculations.ConsumptionDirection{calculations.ConsumptionIn, calculations.ConsumptionOut} {
		calcs = append(calcs, &calculations.UserConsumptionCalculation{
			CalculationParams: calculations.CalculationParams{
				RegisteredGauge:     gauge,
				Labels:              prometheus.Labels{"hours": "1"},
				Name:                metricName,
				ExpectedGaugeLabels: hourlyUserConsumptionLabels,
			},
			Direction: dir,
			Hours:     1,
		})
	}
	return calcs
}

func getAuthenticationCalculations(daysList []int, gauge *prometheus.GaugeVec, metricName string) []calculations.Calculation {
	calcs := make([]calculations.Calculation, 0)
	for _, dayParam := range daysList {
		calcs = append(calcs, &calculations.AuthenticationsCalculation{
			CalculationParams: calculations.CalculationParams{
				Days:                dayParam,
				RegisteredGauge:     gauge,
				Labels:              prometheus.Labels{calculations.DaysLabel: strconv.Itoa(dayParam)},
				Name:                metricName,
				ExpectedGaugeLabels: authenticationsLabels,
			},
		})
	}
	return calcs
}

func getPrometheusClient() v1.API {
	metricsConfig, err := config.GetServiceConfig(orc8r.ModuleName, metricsd.ServiceName)
	if err != nil {
		glog.Fatalf("Could not retrieve metricsd configuration: %s", err)
	}
	promClient, err := promAPI.NewClient(promAPI.Config{Address: metricsConfig.MustGetString(metricsd.PrometheusQueryAddress)})
	if err != nil {
		glog.Fatalf("Error creating prometheus client: %s", promClient)
	}
	return v1.NewAPI(promClient)
}
