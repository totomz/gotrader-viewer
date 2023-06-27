package main

import (
	"fmt"
	"strings"
)

const template = `
      ,{"datasource": {
        "type": "redis-datasource",
        "uid": "redislocal"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisCenteredZero": false,
            "axisColorMode": "text",
            "axisLabel": "",
            "axisPlacement": "auto",
            "barAlignment": 0,
            "drawStyle": "line",
            "fillOpacity": 0,
            "gradientMode": "none",
            "hideFrom": {
              "legend": false,
              "tooltip": false,
              "viz": false
            },
            "lineInterpolation": "linear",
            "lineWidth": 1,
            "pointSize": 5,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "auto",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          }
        },
        "overrides": []
      },
      "gridPos": {
        "h": 7,
        "w": 12,
        "x": 12,
        "y": 30
      },
      "id": %v,
      "options": {
        "legend": {
          "calcs": [],
          "displayMode": "list",
          "placement": "bottom",
          "showLegend": true
        },
        "tooltip": {
          "mode": "single",
          "sort": "none"
        }
      },
      "targets": [
        {
          "command": "ts.range",
          "datasource": {
            "type": "redis-datasource",
            "uid": "redislocal"
          },
          "keyName": "gotrader.$symbol.%s",
          "query": "",
          "refId": "A",
          "type": "timeSeries"
        }
      ],
      "title": "%s",
      "type": "timeseries"
    }`

func main() {

	dio := strings.ReplaceAll(template, "\n", "")
	names := []string{
		// "trade_pl",
		"trade_loss_count",
		"trade_profit_count",
		"psar",
		"lr_candle_midprice_6",
		// "position_size",
		// "position_price",
		"c_psar_strength",
		"psar_trend",
		"psar_strength",
		"psar_inversion",
		"candle_color",
		"candle_psar_const",
		"profit_loss",
		// "daily_pl",
		"entry_psar_strength",
		"entry_psar_cons",
		// "zero",
		// "enter_signal",
		// "enter_signal_threshold",
		"order_x",
		"take_profit_threshold",
		"stop_loss_threshold",
		"exit_part_stoploss",
		"exit_part_takeprofit",
		"exit_part_psarcandlepos",
		"order_leg",
		"exit_signal_contrario",
		// "exit_signal",
		// "exit_signal_threshold",
	}

	for i, n := range names {
		j := i + 12
		println(fmt.Sprintf(dio, j, n, n))
	}
}
