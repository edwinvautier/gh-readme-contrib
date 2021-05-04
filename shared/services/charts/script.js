AmCharts.makeChart( "chart1", {
  "type": "serial",
  "theme": "light",
  "autoMargins": false,
  "marginLeft": 50,
  "marginRight": 8,
  "marginTop": 30,
  "marginBottom": 26,
  "dataProvider": [ {
    "country": "USA",
    "visits": 2025
  }, {
    "country": "China",
    "visits": 1882
  }, {
    "country": "Japan",
    "visits": 1809
  }, {
    "country": "Germany",
    "visits": 1322
  }, {
    "country": "UK",
    "visits": 1122
  }, {
    "country": "France",
    "visits": 1114
  }, {
    "country": "India",
    "visits": 984
  }, {
    "country": "Spain",
    "visits": 711
  } ],
  "valueAxes": [ {
    "gridColor": "#FFFFFF",
    "gridAlpha": 0.2,
    "axisAlpha": 1,
    "dashLength": 0
  } ],
  "startDuration": 1,
  "graphs": [ {
    "balloonText": "[[category]]: <b>[[value]]</b>",
    "fillAlphas": 0.8,
    "lineAlpha": 0.2,
    "type": "column",
    "valueField": "visits"
  } ],
  "chartCursor": {
    "categoryBalloonEnabled": false,
    "cursorAlpha": 0,
    "zoomable": false
  },
  "categoryField": "country",
  "categoryAxis": {
    "gridPosition": "start",
    "gridAlpha": 0,
    "axisAlpha": 1,
    "tickPosition": "start",
    "tickLength": 20
  }

} );

/**
 * Second chart
 */

AmCharts.makeChart( "chart2", {
  "type": "serial",
  "addClassNames": true,
  "theme": "light",
  "autoMargins": false,
  "marginLeft": 50,
  "marginRight": 8,
  "marginTop": 30,
  "marginBottom": 26,
  "balloon": {
    "adjustBorderColor": false,
    "horizontalPadding": 10,
    "verticalPadding": 8,
    "color": "#ffffff"
  },

  "dataProvider": [ {
    "year": 2009,
    "income": 23.5,
    "expenses": 21.1
  }, {
    "year": 2010,
    "income": 26.2,
    "expenses": 30.5
  }, {
    "year": 2011,
    "income": 30.1,
    "expenses": 34.9
  }, {
    "year": 2012,
    "income": 29.5,
    "expenses": 31.1
  }, {
    "year": 2013,
    "income": 30.6,
    "expenses": 28.2,
  }, {
    "year": 2014,
    "income": 34.1,
    "expenses": 32.9,
    "dashLengthColumn": 5,
    "alpha": 0.2,
    "additional": "(projection)"
  } ],
  "valueAxes": [ {
    "gridColor": "#FFFFFF",
    "gridAlpha": 0.2,
    "axisAlpha": 1,
    "dashLength": 0
  } ],
  "startDuration": 1,
  "graphs": [ {
    "alphaField": "alpha",
    "balloonText": "<span style='font-size:12px;'>[[title]] in [[category]]:<br><span style='font-size:20px;'>[[value]]</span> [[additional]]</span>",
    "fillAlphas": 1,
    "title": "Income",
    "type": "column",
    "valueField": "income",
    "dashLengthField": "dashLengthColumn"
  }, {
    "id": "graph2",
    "balloonText": "<span style='font-size:12px;'>[[title]] in [[category]]:<br><span style='font-size:20px;'>[[value]]</span> [[additional]]</span>",
    "bullet": "round",
    "lineThickness": 3,
    "bulletSize": 7,
    "bulletBorderAlpha": 1,
    "bulletColor": "#FFFFFF",
    "useLineColorForBulletBorder": true,
    "bulletBorderThickness": 3,
    "fillAlphas": 0,
    "lineAlpha": 1,
    "title": "Expenses",
    "valueField": "expenses"
  } ],
  "categoryField": "year",
  "categoryAxis": {
    "gridPosition": "start",
    "gridAlpha": 0,
    "axisAlpha": 1,
    "tickPosition": "start",
    "tickLength": 20
  }
} );

/**
 * Third chart
 */

AmCharts.makeChart( "chart3", {
  "type": "serial",
  "theme": "light",
  "autoMargins": false,
  "marginLeft": 50,
  "marginRight": 8,
  "marginTop": 30,
  "marginBottom": 26,
  "dataProvider": [ {
    "date": "2012-03-01",
    "price": 20
  }, {
    "date": "2012-03-02",
    "price": 75
  }, {
    "date": "2012-03-03",
    "price": 15
  }, {
    "date": "2012-03-04",
    "price": 75
  }, {
    "date": "2012-03-05",
    "price": 158
  }, {
    "date": "2012-03-06",
    "price": 57
  }, {
    "date": "2012-03-07",
    "price": 107
  }, {
    "date": "2012-03-08",
    "price": 89
  }, {
    "date": "2012-03-09",
    "price": 75
  }, {
    "date": "2012-03-10",
    "price": 132
  }, {
    "date": "2012-03-11",
    "price": 158
  }, {
    "date": "2012-03-12",
    "price": 56
  }, {
    "date": "2012-03-13",
    "price": 169
  }, {
    "date": "2012-03-14",
    "price": 24
  }, {
    "date": "2012-03-15",
    "price": 147
  } ],
  "valueAxes": [ {
    "logarithmic": true,
    "dashLength": 1,
    "axisAlpha": 1,
    "guides": [ {
      "dashLength": 6,
      "inside": true,
      "label": "average",
      "lineAlpha": 1,
      "value": 90.4
    } ],
    "position": "left"
  } ],
  "graphs": [ {
    "bullet": "round",
    "id": "g1",
    "bulletBorderAlpha": 1,
    "bulletColor": "#FFFFFF",
    "bulletSize": 7,
    "lineThickness": 2,
    "lineAlpha": 1,
    "title": "Price",
    "useLineColorForBulletBorder": true,
    "valueField": "price"
  } ],
  "chartCursor": {
    "valueLineEnabled": true,
    "valueLineBalloonEnabled": true,
    "valueLineAlpha": 0.5,
    "fullWidth": true,
    "cursorAlpha": 0.05
  },
  "dataDateFormat": "YYYY-MM-DD",
  "categoryField": "date",
  "categoryAxis": {
    "parseDates": true,
    "gridPosition": "start",
    "gridAlpha": 0,
    "axisAlpha": 1,
    "tickPosition": "start",
    "tickLength": 20
  }
} );