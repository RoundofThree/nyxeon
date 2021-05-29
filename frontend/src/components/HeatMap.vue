<template>
<div>
<div id="heatmap"></div>

<div class="container mb-2 flex mx-auto w-full items-center justify-center">
<ul class="flex flex-box flex-col p-10 w-full">
  <li class="border-gray-400 flex flex-row mt-2" v-for="item in show" v-bind:key="item._id">
    <div class="select-none flex flex-1 items-center p-4 transition duration-500 ease-in-out transform hover:-translate-y-2 rounded-2xl border-2 p-6 hover:shadow-2xl">
      <div class="flex-1 pl-1 mr-16">
        <div class="font-medium">
          {{ item.content }}
        </div>
      </div>
      <div v-for="category in item.categories" v-bind:key="category" class="text-wrap text-center flex flex-col text-white text-bold rounded-md bg-black justify-center items-center mr-10 p-2">
        {{ category }}
      </div>
    </div>
  </li>
</ul>
</div>
</div>
</template>

<script>
import { calendarHeatmap } from '../plugins/heatmap'
import moment from 'moment'
import * as d3 from 'd3'

export default {
  data() {
    return {
      entries: this.entriesIn,
      colorRange: null, 
      tooltipEnabled: true,
      tooltipUnit: null, 
      locale: null,
      max: null, 
      map: null, 
      show: []
    }
  },
  props: ["entriesIn"],
  mounted() {
    var that = this 
    setTimeout(()=>{
      that.renderHeatMap()
    }, 2000)
  },
  methods: {
    renderHeatMap() {
      // preprocess data 
      this.entries = this.entries.reverse()
      let map = {}
      this.entries.forEach(e => {
        let date = new Date(e.date)
        date = new Date(date.toDateString())
        if (map[date]) {
          map[date].push(e)
        } else {
          map[date] = [e]
        }
      })
      this.map = map
      var now = moment().endOf('day').toDate()
      var yearAgo = moment().startOf('day').subtract(1, 'year').toDate() 
      var chartData = d3.time.days(yearAgo, now).map(e => {
        return {
          date: e,
          count: map[e]? map[e].length : 0
        }
      })
      var heatmap = calendarHeatmap()
        .data(chartData)
        .selector("#heatmap")
        .tooltipEnabled(true)
        .tooltipUnit("quest")
        .legendEnabled(false)
        .colorRange(["#f4f7f7", "#79a8a9"])
        .onClick(d => {
          // show the quests on that day 
          this.show = this.map[d.date]
        })
      heatmap()
    }
  }
}
</script>
<style>
text.month-name,
text.calendar-heatmap-legend-text,
text.day-initial {
  font-size: 10px;
  z-index: 100;
  fill: inherit;
  font-family: Helvetica, arial, 'Open Sans', sans-serif;
}
rect.day-cell:hover {
  stroke: #555555;
  stroke-width: 1px;
}
.day-cell-tooltip {
  position: absolute;
  z-index: 9999;
  padding: 5px 9px;
  color: #bbbbbb;
  font-size: 12px;
  background: rgba(0, 0, 0, 0.85);
  border-radius: 3px;
  text-align: center;
}
.day-cell-tooltip > span {
  font-family: Helvetica, arial, 'Open Sans', sans-serif
}
.calendar-heatmap {
  box-sizing: initial;
}

</style>
