package main

import (
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	lr "github.com/tlkamp/litter-api"
)

type LitterRobotCollector struct {
	cycles      *prometheus.Desc
	drawerFull  *prometheus.Desc
	capacity    *prometheus.Desc
	nightlight  *prometheus.Desc
	sleepMode   *prometheus.Desc
	unitStatus  *prometheus.Desc
	panelLocked *prometheus.Desc
	lrClient    *lr.Client
}

func NewCollector(email, password, apiKey, clientSecret, clientId, endpoint, authEndpoint string) *LitterRobotCollector {
	client, err := lr.NewClient(&lr.Config{
		ApiKey:       apiKey,
		Email:        email,
		Password:     password,
		ClientSecret: clientSecret,
		ClientId:     clientId,
		ApiUrl:       endpoint,
		AuthUrl:      authEndpoint,
	})
	if err != nil {
		log.Fatal(err)
	}

	labels := []string{
		"serial",
		"id",
		"name",
	}
	return &LitterRobotCollector{
		cycles: prometheus.NewDesc(
			"litterrobot_cycles", "The number of cycles completed", labels, nil,
		),
		drawerFull: prometheus.NewDesc(
			"litterrobot_drawer_full", "Drawer full", labels, nil,
		),
		capacity: prometheus.NewDesc(
			"litterrobot_capacity", "Total number of cycles until full", labels, nil,
		),
		nightlight: prometheus.NewDesc(
			"litterrobot_nightlight", "the status of the nightlight", labels, nil,
		),
		sleepMode: prometheus.NewDesc(
			"litterrobot_sleepmode", "sleepmode status", labels, nil,
		),
		unitStatus: prometheus.NewDesc(
			"litterrobot_unit_status", "unit status", labels, nil,
		),
		panelLocked: prometheus.NewDesc(
			"litterrobot_panel_locked", "panel locked", labels, nil,
		),
		lrClient: client,
	}
}

func (lrc *LitterRobotCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- lrc.cycles
	ch <- lrc.drawerFull
	ch <- lrc.capacity
	ch <- lrc.nightlight
	ch <- lrc.sleepMode
	ch <- lrc.unitStatus
	ch <- lrc.panelLocked
}

func (lrc *LitterRobotCollector) Collect(ch chan<- prometheus.Metric) {
	states, _ := lrc.lrClient.States()
	for _, s := range states {
		labels := []string{s.LitterRobotSerial, s.LitterRobotID, s.Name}

		ch <- prometheus.MustNewConstMetric(lrc.capacity, prometheus.CounterValue, s.CycleCapacity, labels...)
		ch <- prometheus.MustNewConstMetric(lrc.cycles, prometheus.CounterValue, s.CycleCount, labels...)
		ch <- prometheus.MustNewConstMetric(lrc.drawerFull, prometheus.GaugeValue, bool2float(s.DFITriggered), labels...)
		ch <- prometheus.MustNewConstMetric(lrc.nightlight, prometheus.GaugeValue, bool2float(s.NightLightActive), labels...)
		ch <- prometheus.MustNewConstMetric(lrc.sleepMode, prometheus.GaugeValue, bool2float(s.SleepModeActive), labels...)
		ch <- prometheus.MustNewConstMetric(lrc.unitStatus, prometheus.GaugeValue, s.UnitStatus, labels...)
		ch <- prometheus.MustNewConstMetric(lrc.panelLocked, prometheus.GaugeValue, bool2float(s.PanelLockActive), labels...)
	}
}

func bool2float(b bool) float64 {
	if b {
		return 1
	}
	return 0
}
