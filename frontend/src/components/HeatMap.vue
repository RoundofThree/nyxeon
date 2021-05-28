<template>
  <div class="heatmap">I am a heatmap</div>
</template>

<script>
// import * as d3 from 'd3'

export default {
  data() {
    return {
      entries: [],
      colorRange: null, 
      tooltipEnabled: true,
      tooltipUnit: null, 
      locale: null,
      max: null, 
      onClick: null,
    }
  },
  created() {
    this.renderHeatMap()
  },
  watch: {
    entries() {
      this.renderHeatMap() 
    }
  },
  methods: {
    renderHeatMap() {
      // main logic 
    }
  }
}
</script>