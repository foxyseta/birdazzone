<script lang="ts" setup>
import "leaflet/dist/leaflet.css"
import { LMap, LTileLayer } from '@vue-leaflet/vue-leaflet'

const ROME = [41.53, 12.30]

const center = ROME
const url = 'https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png'
const attribution = '&copy; <a target="_blank" href="http://osm.org/copyright">OpenStreetMap</a> contributors'
     
</script>
<template>
  <div class="bg-foreground rounded-xl w-100 p-3">
    <LMap style="height:40rem; width: 40rem;" :zoom="5" :center="center">
     <LTileLayer :url="url" :attribution="attribution"></LTileLayer>
    </LMap>
  </div>
</template>
